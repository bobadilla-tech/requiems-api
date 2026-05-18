/**
 * Seeds local wrangler KV and D1 for the full-stack dev environment.
 * Run from the auth-gateway working directory so wrangler.toml bindings
 * and schema.sql resolve correctly.
 *
 * Usage: node ./scripts/seed-dev.ts
 */

import { execSync } from "node:child_process";

interface DevKey {
  apiKey: string;
  userId: string;
  plan: string;
}

const WRANGLER = "pnpm exec wrangler";
const WRANGLER_PERSIST_TO = process.env.WRANGLER_PERSIST_TO?.trim();

// Keys use unique 12-char prefixes so D1 UNIQUE(key_prefix) constraint is satisfied.
// Format: rq_{plan-abbr}_{6-digit-id}  →  prefix = first 12 chars
const DEV_KEYS: DevKey[] = [
  { apiKey: "rq_free_000001", userId: "dev_user_free", plan: "free" },
  { apiKey: "rq_devl_000001", userId: "dev_user_developer", plan: "developer" },
  { apiKey: "rq_bizz_000001", userId: "dev_user_business", plan: "business" },
  {
    apiKey: "rq_prof_000001",
    userId: "dev_user_professional",
    plan: "professional",
  },
];

const CREATED_AT = "2025-01-01T00:00:00.000Z";

function run(cmd: string): void {
  execSync(cmd, { stdio: "inherit" });
}

function withPersist(cmd: string): string {
  if (!WRANGLER_PERSIST_TO) {
    return cmd;
  }
  return `${cmd} --persist-to=${WRANGLER_PERSIST_TO}`;
}

console.log("Applying D1 schema...");
// Local dev uses shared D1 persistence; reset tables first so schema changes are applied deterministically.
run(
  withPersist(
    `${WRANGLER} d1 execute requiem-usage --local --yes --command="DROP TABLE IF EXISTS credit_usage; DROP TABLE IF EXISTS api_keys;"`,
  ),
);
run(withPersist(`${WRANGLER} d1 execute requiem-usage --local --yes --file=./schema.sql`));

console.log("Applying D1 migrations...");
run(
  withPersist(
    `${WRANGLER} d1 execute requiem-usage --local --yes --file=./migrations/001_remove_credit_usage_unique_constraint.sql`,
  ),
);

console.log("\nSeeding KV with dev test keys...");
for (const { apiKey, userId, plan } of DEV_KEYS) {
  const kvKey = `key:${apiKey}`;
  const value = JSON.stringify({ userId, plan, createdAt: CREATED_AT });
  run(withPersist(`${WRANGLER} kv key put '${kvKey}' '${value}' --binding=KV --local`));
}

console.log("\nSeeding D1 api_keys with dev test keys...");
for (const { apiKey, userId, plan } of DEV_KEYS) {
  const keyPrefix = apiKey.substring(0, 12);
  const sql =
    `INSERT OR REPLACE INTO api_keys (key_prefix, user_id, plan, active, created_at, billing_cycle_start)` +
    ` VALUES ('${keyPrefix}', '${userId}', '${plan}', 1, '${CREATED_AT}', '${CREATED_AT}')`;
  run(withPersist(`${WRANGLER} d1 execute requiem-usage --local --yes --command="${sql}"`));
}

console.log("\nDev test keys seeded (header: requiems-api-key):");
for (const { apiKey, plan } of DEV_KEYS) {
  console.log(`  ${plan.padEnd(12)} -> ${apiKey}`);
}
console.log(
  `\nExample: curl -H 'requiems-api-key: ${
    DEV_KEYS[0].apiKey
  }' http://localhost:4455/v1/text/advice\n`,
);
