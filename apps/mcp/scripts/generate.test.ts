import { describe, expect, test } from "bun:test";
import {
    deriveToolName,
    extractOperations,
    generateToolFileSource,
    normalizeName,
    shouldInclude,
    zodFieldsForRequestBody,
    zodForParamSchema,
} from "./generate";

describe("shouldInclude", () => {
    test("accepts /v1/* paths", () => {
        expect(shouldInclude("/v1/technology/convert")).toBe(true);
    });

    test("rejects non-v1 paths", () => {
        expect(shouldInclude("/v2/technology/convert")).toBe(false);
        expect(shouldInclude("/technology/convert")).toBe(false);
    });

    test("rejects batch endpoints, case-insensitively", () => {
        expect(shouldInclude("/v1/technology/batch")).toBe(false);
        expect(shouldInclude("/v1/technology/Batch/convert")).toBe(false);
    });
});

describe("normalizeName", () => {
    test("lowercases, strips non-alphanumerics, trims underscores", () => {
        expect(normalizeName("Get--User__Info!!")).toBe("get_user_info");
        expect(normalizeName("__leading_and_trailing__")).toBe("leading_and_trailing");
    });
});

describe("deriveToolName", () => {
    test("prefers operationId when present", () => {
        expect(deriveToolName("/v1/anything", "get", "Custom Operation ID")).toBe(
            "custom_operation_id",
        );
    });

    test("derives from path when operationId absent", () => {
        expect(deriveToolName("/v1/technology/convert", "get")).toBe("technology_convert");
    });

    test("replaces path param braces with the param name", () => {
        expect(deriveToolName("/v1/finance/bin/{bin}", "get")).toBe("finance_bin_bin");
    });

    test("suffixes non-GET methods to disambiguate from GET on the same path", () => {
        expect(deriveToolName("/v1/technology/counter/{namespace}", "post")).toBe(
            "technology_counter_namespace_post",
        );
    });

    test("does not suffix GET (the common case stays clean)", () => {
        expect(deriveToolName("/v1/technology/counter/{namespace}", "get")).toBe(
            "technology_counter_namespace",
        );
    });
});

describe("zodForParamSchema", () => {
    test("maps primitive types", () => {
        expect(zodForParamSchema({ type: "integer" })).toBe("z.number().int()");
        expect(zodForParamSchema({ type: "number" })).toBe("z.number()");
        expect(zodForParamSchema({ type: "boolean" })).toBe("z.boolean()");
        expect(zodForParamSchema({ type: "string" })).toBe("z.string()");
    });

    test("defaults to z.string() when schema or type is missing", () => {
        expect(zodForParamSchema(undefined)).toBe("z.string()");
        expect(zodForParamSchema({})).toBe("z.string()");
    });

    test("maps string enums to z.enum", () => {
        expect(zodForParamSchema({ type: "string", enum: ["a", "b"] })).toBe(
            'z.enum(["a", "b"])',
        );
    });

    test("falls back to the base type when enum contains non-strings", () => {
        expect(zodForParamSchema({ type: "integer", enum: [1, 2] })).toBe("z.number().int()");
    });
});

describe("zodFieldsForRequestBody", () => {
    test("marks required fields as-is and optional fields with .optional()", () => {
        const fields = zodFieldsForRequestBody({
            type: "object",
            required: ["from"],
            properties: {
                from: { type: "string" },
                to: { type: "string" },
            },
        });
        const byName = Object.fromEntries(fields.map((f) => [f.name, f]));
        expect(byName.from?.optional).toBe(false);
        expect(byName.from?.zodExpr).toBe("z.string()");
        expect(byName.to?.optional).toBe(true);
        expect(byName.to?.zodExpr).toBe("z.string().optional()");
    });

    test("wraps array items", () => {
        const fields = zodFieldsForRequestBody({
            type: "object",
            required: ["tags"],
            properties: { tags: { type: "array", items: { type: "string" } } },
        });
        expect(fields[0]?.zodExpr).toBe("z.array(z.string())");
    });

    test("defers nested objects to z.record(z.any())", () => {
        const fields = zodFieldsForRequestBody({
            type: "object",
            required: ["meta"],
            properties: { meta: { type: "object" } },
        });
        expect(fields[0]?.zodExpr).toBe("z.record(z.any())");
    });

    test("defers oneOf/anyOf to z.any()", () => {
        const fields = zodFieldsForRequestBody({
            type: "object",
            required: ["value"],
            properties: { value: { oneOf: [{ type: "string" }, { type: "integer" }] } },
        });
        expect(fields[0]?.zodExpr).toBe("z.any()");
    });

    test("returns no fields when schema has no properties", () => {
        expect(zodFieldsForRequestBody({ type: "object" })).toEqual([]);
    });
});

describe("extractOperations", () => {
    test("extracts, filters to /v1/*, and sorts deterministically", () => {
        const ops = extractOperations({
            paths: {
                "/v1/b/thing": { get: { summary: "b" } },
                "/v1/a/thing": { get: { summary: "a" } },
                "/v2/ignored": { get: { summary: "ignored" } },
            },
        });
        expect(ops.map((o) => o.toolName)).toEqual(["a_thing", "b_thing"]);
    });

    test("throws on a genuine tool name collision", () => {
        expect(() =>
            extractOperations({
                paths: {
                    "/v1/foo/{bar}": { get: { summary: "one" } },
                    "/v1/foo/baz": {
                        get: { operationId: "foo_bar", summary: "two" },
                    },
                },
            }),
        ).toThrow(/Tool name collision/);
    });

    test("does not throw when the same path+method is seen twice (not a real collision)", () => {
        expect(() =>
            extractOperations({
                paths: {
                    "/v1/foo/{bar}": { get: { summary: "one" } },
                },
            }),
        ).not.toThrow();
    });
});

describe("generateToolFileSource", () => {
    test("emits a tool file importing zod + runtime, with the handler calling requiemsRequest", () => {
        const [op] = extractOperations({
            paths: { "/v1/entertainment/chuck-norris": { get: { summary: "Chuck Norris fact" } } },
        });
        const source = generateToolFileSource(op!);

        expect(source).toContain('import { z } from "zod"');
        expect(source).toContain('import { requiemsRequest } from "../runtime"');
        expect(source).toContain("export const entertainment_chuck_norris");
        expect(source).toContain('path: "/v1/entertainment/chuck-norris"');
        expect(source).toContain("await requiemsRequest(");
    });
});
