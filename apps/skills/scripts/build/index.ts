/**
 * Reads API documentation YAML files from requiems-api and transforms them
 * into Markdown skill files consumable by AI agents.
 */

import { load } from "js-yaml";
import { join } from "node:path";
import { mkdir, readdir, readFile, writeFile } from "node:fs/promises";
import { parseArgs } from "node:util";
import type { ApiDoc, Endpoint } from "./types.ts";

// ---------------------------------------------------------------------------
// Markdown builder
// ---------------------------------------------------------------------------

// Claude Code skill names: lowercase letters, digits and hyphens only. Path
// params like `{bin}` and ids with underscores are normalized.
function skillSlug(api: ApiDoc, endpoint: Endpoint, segments = 1): string {
  return `${api.api_id}-${endpoint.method.toLowerCase()}-${
    endpoint.path.split("/").slice(-segments).join("-")
  }`
    .toLowerCase()
    .replace(/[^a-z0-9]+/g, "-")
    .replace(/^-+|-+$/g, "");
}

function buildSkillMarkdown(
  api: ApiDoc,
  endpoint: Endpoint,
  slug: string,
): string {
  const lines: string[] = [];

  // Front-matter
  lines.push("---");
  lines.push(`name: ${slug}`);
  lines.push(`api: ${JSON.stringify(api.api_name)}`);
  lines.push(`method: ${endpoint.method}`);
  lines.push(`path: ${endpoint.path}`);
  lines.push(`base_url: ${api.base_url}`);
  // JSON strings are valid YAML scalars, so colons/quotes in descriptions are safe.
  lines.push(
    `description: ${
      JSON.stringify(endpoint.description.replace(/\s+/g, " ").trim())
    }`,
  );
  lines.push("---");
  lines.push("");

  // Endpoint URL — shown first so agents never have to infer it
  const fullUrl = `${api.base_url}${endpoint.path}`;
  lines.push("## Endpoint");
  lines.push("");
  lines.push(`**${endpoint.method} ${fullUrl}**`);
  lines.push("");

  // Description
  lines.push(`## ${endpoint.name}`);
  lines.push("");
  lines.push(endpoint.description.trim());
  lines.push("");

  // Parameters
  if (endpoint.parameters && endpoint.parameters.length > 0) {
    lines.push("## Parameters");
    lines.push("");
    lines.push("| Name | Type | Required | Location | Description |");
    lines.push("| ---- | ---- | -------- | -------- | ----------- |");
    for (const p of endpoint.parameters) {
      const req = p.required ? "yes" : "no";
      lines.push(
        `| \`${p.name}\` | ${p.type} | ${req} | ${p.location} | ${p.description} |`,
      );
    }
    lines.push("");
  }

  // Request example
  if (endpoint.request_example) {
    lines.push("## Request Example");
    lines.push("");
    lines.push("```json");
    lines.push(endpoint.request_example.trim());
    lines.push("```");
    lines.push("");
  }

  // Response example
  if (endpoint.response_example) {
    lines.push("## Response Example");
    lines.push("");
    lines.push("```json");
    lines.push(endpoint.response_example.trim());
    lines.push("```");
    lines.push("");
  }

  // Response fields
  if (endpoint.response_fields && endpoint.response_fields.length > 0) {
    lines.push("## Response Fields");
    lines.push("");
    lines.push("| Field | Type | Description |");
    lines.push("| ----- | ---- | ----------- |");
    for (const f of endpoint.response_fields) {
      lines.push(`| \`${f.name}\` | ${f.type} | ${f.description} |`);
    }
    lines.push("");
  }

  // Errors
  if (endpoint.errors && endpoint.errors.length > 0) {
    lines.push("## Errors");
    lines.push("");
    for (const e of endpoint.errors) {
      const codeLabel = e.code ? ` **${e.code}**` : "";
      lines.push(`- \`${e.status}\`${codeLabel} — ${e.description}`);
    }
    lines.push("");
  }

  return lines.join("\n");
}

// ---------------------------------------------------------------------------
// Main
// ---------------------------------------------------------------------------

async function main() {
  // Parse CLI arguments
  const { values: args } = parseArgs({
    options: {
      source: { type: "string" },
      output: { type: "string", default: "./skills" },
    },
  });

  const sourcePath = args.source;
  const outputPath = args.output as string;

  if (!sourcePath) {
    console.error("Error: --source is required.");
    console.error(
      "Usage: node scripts/build/index.ts --source <path> [--output <path>]",
    );

    process.exit(1);
  }

  await mkdir(outputPath, { recursive: true });

  const seen = new Set<string>();
  let processed = 0;
  let skipped = 0;

  const entries = await readdir(sourcePath, { withFileTypes: true });

  for (const entry of entries) {
    if (!entry.isFile() || !entry.name.endsWith(".yml")) {
      skipped++;
      continue;
    }

    const filePath = join(sourcePath, entry.name);
    const raw = await readFile(filePath, "utf-8");

    let api: ApiDoc;
    try {
      api = load(raw) as ApiDoc;
    } catch (err) {
      console.warn(`Skipping ${entry.name}: YAML parse error — ${err}`);
      skipped++;
      continue;
    }

    if (!api.endpoints || api.endpoints.length === 0) {
      console.warn(`Skipping ${entry.name}: no endpoints found.`);
      skipped++;
      continue;
    }

    for (const endpoint of api.endpoints) {
      // Different paths can share a last segment (e.g. .../encode/batch and
      // .../decode/batch); widen every colliding name instead of overwriting.
      const base = skillSlug(api, endpoint);
      const collides = api.endpoints.filter((e) => skillSlug(api, e) === base)
        .length > 1;
      const slug = collides ? skillSlug(api, endpoint, 2) : base;
      if (seen.has(slug)) {
        throw new Error(`Duplicate skill name "${slug}" (${entry.name})`);
      }
      seen.add(slug);
      const skillDir = join(outputPath, slug);
      await mkdir(skillDir, { recursive: true });
      const outFile = join(skillDir, "SKILL.md");
      const content = buildSkillMarkdown(api, endpoint, slug);
      await writeFile(outFile, content, "utf-8");
      console.log(`✓ ${slug}.md`);
      processed++;
    }
  }

  console.log(
    `\nDone. ${processed} skills generated, ${skipped} files skipped.`,
  );
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});
