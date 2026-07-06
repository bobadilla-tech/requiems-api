// Codegen CLI for the Requiems API MCP server.
//
// Reads the Requiems OpenAPI spec (v1 only), generates one MCP tool file
// per operation under generated/tools/, plus a generated/index.ts that
// registers them all. Batch endpoints (paths containing "/batch") are
// excluded — see "Non-goals" in the POC spec.
//
// Usage:
//   bun run scripts/generate.ts \
//     --input https://api.requiems.xyz/openapi.json \
//     --output generated
//
//   bun run scripts/generate.ts --input ./openapi.json --output generated

import { mkdir, readFile, rm, writeFile } from "node:fs/promises";
import path from "node:path";

// ---------------------------------------------------------------------------
// Minimal OpenAPI 3.0 types (only what we use)
// ---------------------------------------------------------------------------

interface OpenApiSchema {
    type?: string;
    format?: string;
    description?: string;
    example?: unknown;
    properties?: Record<string, OpenApiSchema>;
    items?: OpenApiSchema;
    required?: string[];
    enum?: unknown[];
    oneOf?: OpenApiSchema[];
    anyOf?: OpenApiSchema[];
    nullable?: boolean;
}

interface OpenApiParameter {
    name: string;
    in: "path" | "query" | "header" | "cookie";
    required?: boolean;
    description?: string;
    schema?: OpenApiSchema;
}

interface OpenApiRequestBody {
    required?: boolean;
    content?: Record<string, { schema?: OpenApiSchema }>;
}

interface OpenApiOperation {
    operationId?: string;
    summary?: string;
    description?: string;
    tags?: string[];
    parameters?: OpenApiParameter[];
    requestBody?: OpenApiRequestBody;
    responses?: Record<string, unknown>;
}

type OpenApiMethod = "get" | "post" | "put" | "patch" | "delete";

type OpenApiPathItem = Partial<Record<OpenApiMethod, OpenApiOperation>>;

interface OpenApiSpec {
    openapi?: string;
    info?: { title?: string; version?: string };
    servers?: { url: string }[];
    paths: Record<string, OpenApiPathItem>;
}

// ---------------------------------------------------------------------------
// CLI args
// ---------------------------------------------------------------------------

interface CliArgs {
    input: string;
    output: string;
    baseUrl?: string;
}

function parseArgs(argv: string[]): CliArgs {
    const args: Partial<CliArgs> = {};
    for (let i = 0; i < argv.length; i++) {
        const arg = argv[i];
        if (arg === "--input") args.input = argv[++i];
        else if (arg === "--output") args.output = argv[++i];
        else if (arg === "--base-url") args.baseUrl = argv[++i];
    }
    if (!args.input) {
        throw new Error("--input <path-or-url> is required");
    }
    return {
        input: args.input,
        output: args.output ?? "generated",
        baseUrl: args.baseUrl,
    };
}

// ---------------------------------------------------------------------------
// Spec loading + validation
// ---------------------------------------------------------------------------

async function loadSpec(input: string): Promise<OpenApiSpec> {
    let raw: string;
    if (/^https?:\/\//.test(input)) {
        const res = await fetch(input, { signal: AbortSignal.timeout(10_000) });
        if (!res.ok) {
            throw new Error(`Failed to fetch spec from ${input}: ${res.status}`);
        }
        raw = await res.text();
    } else {
        raw = await readFile(input, "utf-8");
    }

    let spec: OpenApiSpec;
    try {
        spec = JSON.parse(raw);
    } catch (err) {
        throw new Error(`Spec at ${input} is not valid JSON: ${(err as Error).message}`);
    }

    validateSpec(spec);
    return spec;
}

function validateSpec(spec: OpenApiSpec): void {
    if (!spec || typeof spec !== "object") {
        throw new Error("Spec did not parse to an object");
    }
    if (!spec.paths || typeof spec.paths !== "object") {
        throw new Error("Spec is missing a top-level `paths` object");
    }
    if (Object.keys(spec.paths).length === 0) {
        throw new Error("Spec `paths` is empty — nothing to generate");
    }
}

// ---------------------------------------------------------------------------
// Operation extraction
// ---------------------------------------------------------------------------

interface ExtractedOperation {
    toolName: string;
    method: OpenApiMethod;
    /** Raw OpenAPI path, e.g. /v1/technology/convert or /v1/finance/bin/{bin} */
    rawPath: string;
    summary: string;
    description: string;
    pathParams: OpenApiParameter[];
    queryParams: OpenApiParameter[];
    requestBodySchema: OpenApiSchema | null;
    requestBodyRequired: boolean;
}

const HTTP_METHODS: OpenApiMethod[] = ["get", "post", "put", "patch", "delete"];

/** Only operate on /v1/* paths, and skip any batch variant. */
function shouldInclude(rawPath: string): boolean {
    if (!rawPath.startsWith("/v1/")) return false;
    if (rawPath.toLowerCase().includes("/batch")) return false;
    return true;
}

/**
 * Derive a stable tool name from operationId (preferred) or the path,
 * per the naming rules in the POC spec section 6.
 */
function deriveToolName(
    rawPath: string,
    method: OpenApiMethod,
    operationId?: string,
): string {
    if (operationId && operationId.trim().length > 0) {
        return normalizeName(operationId);
    }

    let normalized = rawPath
        .replace(/^\/v1\//, "")
        .replace(/\{([^}]+)\}/g, "$1")
        .replace(/\//g, "_")
        .toLowerCase();

    normalized = normalizeName(normalized);

    // Disambiguate GET/POST pairs on the same path (e.g. counter get+increment)
    // by suffixing the method, but only for non-GET so the common case
    // ("technology_convert") stays clean.
    if (method !== "get") {
        normalized = `${normalized}_${method}`;
    }
    return normalized;
}

function normalizeName(raw: string): string {
    return raw
        .toLowerCase()
        .replace(/[^a-z0-9]+/g, "_")
        .replace(/^_+|_+$/g, "");
}

function extractOperations(spec: OpenApiSpec): ExtractedOperation[] {
    const operations: ExtractedOperation[] = [];
    const seenNames = new Map<string, string>(); // toolName -> rawPath+method, for collision detection

    for (const [rawPath, pathItem] of Object.entries(spec.paths)) {
        if (!shouldInclude(rawPath)) continue;

        for (const method of HTTP_METHODS) {
            const op = pathItem[method];
            if (!op) continue;

            const toolName = deriveToolName(rawPath, method, op.operationId);

            const collisionKey = `${method.toUpperCase()} ${rawPath}`;
            const existing = seenNames.get(toolName);
            if (existing && existing !== collisionKey) {
                throw new Error(
                    `Tool name collision: "${toolName}" derived from both "${existing}" and "${collisionKey}". ` +
                    `Add an explicit operationId in the spec to disambiguate.`,
                );
            }
            seenNames.set(toolName, collisionKey);

            const parameters = op.parameters ?? [];
            const pathParams = parameters.filter((p) => p.in === "path");
            const queryParams = parameters.filter((p) => p.in === "query");

            const jsonBody = op.requestBody?.content?.["application/json"];

            operations.push({
                toolName,
                method,
                rawPath,
                summary: op.summary ?? toolName,
                description: op.description ?? op.summary ?? "",
                pathParams,
                queryParams,
                requestBodySchema: jsonBody?.schema ?? null,
                requestBodyRequired: op.requestBody?.required ?? false,
            });
        }
    }

    // Stable, deterministic ordering (section 12: same spec -> same files).
    operations.sort((a, b) => a.toolName.localeCompare(b.toolName));

    return operations;
}

// ---------------------------------------------------------------------------
// OpenAPI schema -> Zod codegen
// ---------------------------------------------------------------------------

/**
 * Convert a single OpenAPI parameter schema into a Zod expression string,
 * e.g. `z.string()`, `z.number().int()`, `z.boolean()`.
 *
 * POC scope (section 11): string, number, boolean, integer. Arrays/objects
 * in path/query params are rare in this API and fall back to z.string().
 */
function zodForParamSchema(schema: OpenApiSchema | undefined): string {
    if (!schema) return "z.string()";

    if (schema.enum && schema.enum.length > 0) {
        const literals = schema.enum
            .filter((v): v is string => typeof v === "string")
            .map((v) => JSON.stringify(v));
        if (literals.length === schema.enum.length && literals.length > 0) {
            return `z.enum([${literals.join(", ")}])`;
        }
    }

    switch (schema.type) {
        case "integer":
            return "z.number().int()";
        case "number":
            return "z.number()";
        case "boolean":
            return "z.boolean()";
        case "string":
        default:
            return "z.string()";
    }
}

/**
 * Convert an OpenAPI requestBody object schema (flat, single-depth — see
 * section 11) into a record of field name -> Zod expression string.
 *
 * Deferred per POC scope: nested objects, oneOf/anyOf. Both fall back to
 * `z.any()` with a comment rather than failing generation, so codegen
 * stays resilient to the long tail of the spec.
 */
function zodFieldsForRequestBody(
    schema: OpenApiSchema,
): { name: string; zodExpr: string; optional: boolean; description?: string }[] {
    if (!schema.properties) return [];

    const required = new Set(schema.required ?? []);
    const fields: { name: string; zodExpr: string; optional: boolean; description?: string }[] = [];

    for (const [propName, propSchema] of Object.entries(schema.properties)) {
        const isRequired = required.has(propName);
        let zodExpr: string;

        if (propSchema.oneOf || propSchema.anyOf) {
            // Deferred per POC scope (section 11) — keep generation moving.
            zodExpr = "z.any()";
        } else if (propSchema.type === "array") {
            const itemExpr = propSchema.items
                ? zodForParamSchema(propSchema.items)
                : "z.any()";
            zodExpr = `z.array(${itemExpr})`;
        } else if (propSchema.type === "object") {
            // Deferred per POC scope (section 11): nested objects beyond single
            // depth. Accept any plain object so valid requests aren't blocked.
            zodExpr = "z.record(z.any())";
        } else {
            zodExpr = zodForParamSchema(propSchema);
        }

        if (!isRequired) {
            zodExpr += ".optional()";
        }

        fields.push({
            name: propName,
            zodExpr,
            optional: !isRequired,
            description: propSchema.description,
        });
    }

    return fields;
}

// ---------------------------------------------------------------------------
// Tool file generation
// ---------------------------------------------------------------------------

function jsDocLines(text: string, indent = " * "): string {
    if (!text) return "";
    return text
        .split("\n")
        .map((line) => `${indent}${line}`)
        .join("\n");
}

function relativeImportToRuntime(): string {
    return "../runtime";
}

function buildInputSchemaSource(op: ExtractedOperation): {
    schemaFieldsSource: string;
    fieldNames: string[];
} {
    const lines: string[] = [];
    const fieldNames: string[] = [];

    for (const p of op.pathParams) {
        const zodExpr = zodForParamSchema(p.schema);
        const withDesc = p.description
            ? `${zodExpr}.describe(${JSON.stringify(p.description)})`
            : zodExpr;
        // Path params are always required by definition.
        lines.push(`  ${p.name}: ${withDesc},`);
        fieldNames.push(p.name);
    }

    for (const p of op.queryParams) {
        let zodExpr = zodForParamSchema(p.schema);
        if (!p.required) zodExpr += ".optional()";
        const withDesc = p.description
            ? `${zodExpr}.describe(${JSON.stringify(p.description)})`
            : zodExpr;
        lines.push(`  ${p.name}: ${withDesc},`);
        fieldNames.push(p.name);
    }

    if (op.requestBodySchema) {
        const fields = zodFieldsForRequestBody(op.requestBodySchema);
        for (const f of fields) {
            const withDesc = f.description
                ? `${f.zodExpr}.describe(${JSON.stringify(f.description)})`
                : f.zodExpr;
            lines.push(`  ${f.name}: ${withDesc},`);
            fieldNames.push(f.name);
        }
    }

    return { schemaFieldsSource: lines.join("\n"), fieldNames };
}

function buildHandlerSource(op: ExtractedOperation): string {
    const pathParamsObj =
        op.pathParams.length > 0
            ? `{ ${op.pathParams.map((p) => `${p.name}: args.${p.name}`).join(", ")} }`
            : "undefined";

    const queryParamsObj =
        op.queryParams.length > 0
            ? `{ ${op.queryParams.map((p) => `${p.name}: args.${p.name}`).join(", ")} }`
            : "undefined";

    let bodyExpr = "undefined";
    if (op.requestBodySchema) {
        const fields = zodFieldsForRequestBody(op.requestBodySchema);
        if (fields.length > 0) {
            bodyExpr = `{ ${fields.map((f) => `${f.name}: args.${f.name}`).join(", ")} }`;
        }
    }

    return `requiemsRequest({
      method: "${op.method.toUpperCase()}",
      path: "${op.rawPath}",
      pathParams: ${pathParamsObj},
      query: ${queryParamsObj},
      body: ${bodyExpr},
    })`;
}

function generateToolFileSource(op: ExtractedOperation): string {
    const { schemaFieldsSource, fieldNames } = buildInputSchemaSource(op);
    const handlerBody = buildHandlerSource(op);

    const hasFields = fieldNames.length > 0;
    const schemaBlock = hasFields ? `{\n${schemaFieldsSource}\n}` : "{}";

    const description = op.description || op.summary;

    return `// AUTO-GENERATED by scripts/generate.ts — do not edit by hand.
// Source operation: ${op.method.toUpperCase()} ${op.rawPath}
// Regenerate with: bun run scripts/generate.ts --input <spec> --output generated

import { z } from "zod";
import { requiemsRequest } from "${relativeImportToRuntime()}";

/**
${jsDocLines(description)}
 */
export const inputSchema = z.object(${schemaBlock});

export type ${toPascalCase(op.toolName)}Input = z.infer<typeof inputSchema>;

export const ${op.toolName} = {
  name: "${op.toolName}",
  description: ${JSON.stringify(description)},
  inputSchema,

  handler: async (rawArgs: Record<string, unknown>) => {
    const args = inputSchema.parse(rawArgs);
    const result = await ${handlerBody};
    return result;
  },
};
`;
}

function toPascalCase(snake: string): string {
    return snake
        .split("_")
        .filter(Boolean)
        .map((part) => part.charAt(0).toUpperCase() + part.slice(1))
        .join("");
}

// ---------------------------------------------------------------------------
// Index generation
// ---------------------------------------------------------------------------

function generateIndexSource(operations: ExtractedOperation[]): string {
    const imports = operations
        .map((op) => `import { ${op.toolName} } from "./tools/${op.toolName}";`)
        .join("\n");

    const list = operations.map((op) => `  ${op.toolName},`).join("\n");

    return `// AUTO-GENERATED by scripts/generate.ts — do not edit by hand.
// Regenerate with: bun run scripts/generate.ts --input <spec> --output generated

${imports}

export const tools = [
${list}
];
`;
}

// ---------------------------------------------------------------------------
// Main
// ---------------------------------------------------------------------------

async function main() {
    const args = parseArgs(process.argv.slice(2));

    console.log(`Loading spec from ${args.input} ...`);
    const spec = await loadSpec(args.input);

    const title = spec.info?.title ?? "Requiems API";
    const version = spec.info?.version ?? "unknown";
    console.log(`Loaded "${title}" v${version}`);

    const operations = extractOperations(spec);
    console.log(
        `Extracted ${operations.length} v1 operations (batch endpoints excluded)`,
    );

    const outputDir = path.resolve(process.cwd(), args.output);
    const toolsDir = path.join(outputDir, "tools");

    // Idempotent regeneration: wipe and rewrite generated/tools + index.
    // We do NOT touch runtime.ts here since it's also generated but stable;
    // re-run with --force-runtime if you need to regenerate it too.
    await rm(toolsDir, { recursive: true, force: true });
    await mkdir(toolsDir, { recursive: true });

    for (const op of operations) {
        const source = generateToolFileSource(op);
        const filePath = path.join(toolsDir, `${op.toolName}.ts`);
        await writeFile(filePath, source, "utf-8");
    }

    const indexSource = generateIndexSource(operations);
    await writeFile(path.join(outputDir, "index.ts"), indexSource, "utf-8");

    console.log(`Wrote ${operations.length} tool files to ${toolsDir}`);
    console.log(`Wrote index to ${path.join(outputDir, "index.ts")}`);
}

main().catch((err) => {
    console.error("Codegen failed:", err);
    process.exit(1);
});