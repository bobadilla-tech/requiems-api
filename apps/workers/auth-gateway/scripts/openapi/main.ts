/**
 * Generates an OpenAPI 3.0 spec from api_docs YAML files in the dashboard.
 * Run: pnpm generate:openapi
 * Auto-runs before dev and deploy via predev/predeploy hooks.
 */

import { mkdirSync, writeFileSync } from "node:fs";
import { join } from "node:path";

import {
  buildOperation,
  buildSpec,
  loadAPIDoc,
  loadAPIDocs,
  loadCatalog,
} from "./helpers";

import type { CatalogEntry } from "./types";

const outputDir = join(import.meta.dirname, "../../src/generated");
const outputPath = join(outputDir, "openapi.ts");

async function main() {
  const catalog = await loadCatalog();

  const catalogMap = new Map<string, CatalogEntry>();

  for (const entry of catalog.apis ?? []) {
    catalogMap.set(entry.id, entry);
  }

  const docFiles = await loadAPIDocs();

  const paths: Record<string, Record<string, unknown>> = {};
  const tags: { name: string; description: string }[] = [];

  for (const file of docFiles) {
    const doc = await loadAPIDoc(file);

    if (!doc?.api_id || !doc.endpoints?.length) continue;

    const catalogEntry = catalogMap.get(doc.api_id);

    tags.push({
      name: doc.api_id,
      description: catalogEntry?.description ?? doc.description ?? doc.api_name,
    });

    for (const endpoint of doc.endpoints) {
      const pathKey = endpoint.path;
      const method = endpoint.method.toLowerCase();

      if (!paths[pathKey]) paths[pathKey] = {};

      paths[pathKey][method] = buildOperation(endpoint, doc.api_id);
    }
  }

  const spec = buildSpec(tags, paths);

  try {
    mkdirSync(outputDir, { recursive: true });

    writeFileSync(
      outputPath,
      `// AUTO-GENERATED — do not edit. Run \`pnpm generate:openapi\` to regenerate.\n` +
        `export const openApiSpec = ${JSON.stringify(spec, null, 2)};\n`,
    );
  } catch (err) {
    console.error(`❌ Failed to write output to ${outputPath}:`, err);
    process.exit(1);
  }

  console.log(
    `✅ OpenAPI spec generated: ${
      Object.keys(paths).length
    } paths across ${docFiles.length} APIs`,
  );
}

await main();
