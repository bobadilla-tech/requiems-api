/**
 * Backup-export for the auth-gateway Worker's KV namespace and D1 database,
 * written ahead of the Phase 3 items 6-7 Worker-retirement session
 * (docs/plans/2026-08-22-go-auth-foundation-phase-6-usage-multiplier-and-loose-ends.md)
 * so that session isn't writing this from scratch with live Cloudflare
 * access already in hand.
 *
 * Run from the auth-gateway working directory so wrangler.toml bindings
 * resolve correctly (same convention as ./scripts/seed-dev.ts).
 *
 * Usage:
 *   node ./scripts/export-backup.ts --local [--out <dir>]
 *   node ./scripts/export-backup.ts --remote [--out <dir>]
 *
 * --local  targets the local wrangler dev persistence (Miniflare's SQLite/KV
 *          simulation, e.g. WRANGLER_PERSIST_TO) — safe to run anytime, no
 *          Cloudflare credentials required. Used to dry-run this script.
 * --remote targets the real, live production KV namespace and D1 database
 *          (wrangler.toml's [[env.production.kv_namespaces]]/[[env.production.d1_databases]]
 *          ids) — requires `wrangler login` with account access. Only run
 *          this once Phase 3 items 6-7 are actually being executed.
 *
 * Output: <out>/d1-requiem-usage-<timestamp>.sql (full D1 dump, schema + data)
 *         <out>/kv-dump-<timestamp>.json (array of {key, value} pairs, KV
 *         values are stored as JSON strings so each `value` here is the raw
 *         string as stored — parse it again if you need the object)
 */

import { execSync } from "node:child_process";
import { mkdirSync, writeFileSync } from "node:fs";
import { join } from "node:path";

const WRANGLER = "pnpm exec wrangler";
// Mirrors ./scripts/seed-dev.ts's withPersist(): --local without an
// explicit --persist-to falls back to wrangler's default
// (.wrangler/state relative to cwd), not the shared dev persistence this
// repo's docker-compose.dev.yml actually seeds into
// (WRANGLER_PERSIST_TO=/workers/.wrangler/state) — omitting it silently
// exports an empty, unrelated local D1/KV instead of the seeded dev data.
const WRANGLER_PERSIST_TO = process.env.WRANGLER_PERSIST_TO?.trim();

interface KvListEntry {
  name: string;
}

function usageAndExit(message: string): never {
  console.error(message);
  console.error("Usage: node ./scripts/export-backup.ts --local|--remote [--out <dir>]");
  process.exit(1);
  throw new Error("unreachable");
}

function parseArgs(argv: string[]): { target: "local" | "remote"; outDir: string } {
  const hasLocal = argv.includes("--local");
  const hasRemote = argv.includes("--remote");

  if (hasLocal === hasRemote) {
    usageAndExit(
      hasLocal
        ? "Pass exactly one of --local or --remote, not both."
        : "Pass exactly one of --local or --remote.",
    );
  }

  const outIndex = argv.indexOf("--out");
  const outDir = outIndex !== -1 && argv[outIndex + 1] ? argv[outIndex + 1] : "./backup-export";

  return { target: hasLocal ? "local" : "remote", outDir };
}

function run(cmd: string): string {
  return execSync(cmd, { encoding: "utf8" });
}

function withPersist(cmd: string, targetFlag: string): string {
  if (targetFlag !== "--local" || !WRANGLER_PERSIST_TO) {
    return cmd;
  }
  return `${cmd} --persist-to=${WRANGLER_PERSIST_TO}`;
}

function timestamp(): string {
  return new Date().toISOString().replace(/[:.]/g, "-");
}

function exportD1(targetFlag: string, outDir: string, ts: string): string {
  const outFile = join(outDir, `d1-requiem-usage-${ts}.sql`);

  console.log(`Exporting D1 database "requiem-usage" (${targetFlag})...`);
  // Unlike `d1 execute`/`kv key`, `d1 export` has no --persist-to flag (as
  // of wrangler 4.x) — it always resolves local state relative to cwd
  // (./.wrangler/state). If this repo's shared dev persistence
  // (WRANGLER_PERSIST_TO, see docker-compose.dev.yml) lives elsewhere, a
  // --local export will silently produce an empty file instead of erroring;
  // symlink ./.wrangler to the real persist-to directory before running
  // --local from a shell where that mismatch applies.
  run(`${WRANGLER} d1 export requiem-usage ${targetFlag} --skip-confirmation --output="${outFile}"`);
  console.log(`  -> ${outFile}`);

  return outFile;
}

function exportKv(targetFlag: string, outDir: string, ts: string): string {
  console.log(`Listing KV keys (${targetFlag})...`);
  const listRaw = run(withPersist(`${WRANGLER} kv key list --binding=KV ${targetFlag}`, targetFlag));
  const keys: KvListEntry[] = JSON.parse(listRaw);
  console.log(`  ${keys.length} key(s) found.`);

  const entries: Array<{ key: string; value: string }> = [];

  for (const [i, { name }] of keys.entries()) {
    if (i > 0 && i % 50 === 0) {
      console.log(`  ...${i}/${keys.length}`);
    }

    // --text decodes the value as utf8; every value this Worker writes to KV
    // (see seed-dev.ts / apps/workers/auth-gateway/src/*) is a JSON string,
    // never binary, so this is lossless.
    const value = run(
      withPersist(`${WRANGLER} kv key get '${name}' --binding=KV ${targetFlag} --text`, targetFlag),
    ).replace(/\n$/, "");
    entries.push({ key: name, value });
  }

  const outFile = join(outDir, `kv-dump-${ts}.json`);
  writeFileSync(outFile, JSON.stringify(entries, null, 2));
  console.log(`  -> ${outFile} (${entries.length} entries)`);

  return outFile;
}

const { target, outDir } = parseArgs(process.argv.slice(2));
const targetFlag = `--${target}`;
const ts = timestamp();

mkdirSync(outDir, { recursive: true });

console.log(`Backup export starting (target: ${target}, out: ${outDir})\n`);

const d1File = exportD1(targetFlag, outDir, ts);
const kvFile = exportKv(targetFlag, outDir, ts);

console.log("\nDone.");
console.log(`  D1: ${d1File}`);
console.log(`  KV: ${kvFile}`);
