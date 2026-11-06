// Validates the YAML frontmatter of every skills/*/SKILL.md so a skill that
// Claude Code cannot load fails CI instead of shipping.
import { readdirSync, readFileSync, existsSync } from "node:fs";
import { join } from "node:path";
import { load } from "js-yaml";

const root = process.argv[2] ?? "skills";
const NAME_RE = /^[a-z0-9]+(-[a-z0-9]+)*$/;
const errors: string[] = [];
let count = 0;

for (const dir of readdirSync(root, { withFileTypes: true })) {
  if (!dir.isDirectory()) continue;
  const file = join(root, dir.name, "SKILL.md");
  if (!existsSync(file)) {
    errors.push(`${dir.name}: missing SKILL.md`);
    continue;
  }
  count++;
  const text = readFileSync(file, "utf8");
  const m = text.match(/^---\n([\s\S]*?)\n---\n/);
  if (!m) {
    errors.push(`${dir.name}: no YAML frontmatter`);
    continue;
  }
  let fm: Record<string, unknown>;
  try {
    fm = load(m[1]) as Record<string, unknown>;
  } catch (e) {
    errors.push(`${dir.name}: invalid YAML frontmatter (${(e as Error).message.split("\n")[0]})`);
    continue;
  }
  const { name, description } = fm ?? {};
  if (typeof name !== "string" || !NAME_RE.test(name) || name.length > 64) {
    errors.push(`${dir.name}: name "${name}" must be lowercase letters, digits and hyphens (max 64)`);
  } else if (name !== dir.name) {
    errors.push(`${dir.name}: name "${name}" does not match directory`);
  }
  if (typeof description !== "string" || !description.trim() || description.length > 1024) {
    errors.push(`${dir.name}: description must be a non-empty string (max 1024 chars)`);
  }
}

// package.json and the Claude Code plugin manifest must carry the same version.
const pkg = JSON.parse(readFileSync("package.json", "utf8"));
const plugin = JSON.parse(readFileSync(".claude-plugin/plugin.json", "utf8"));
if (pkg.version !== plugin.version) {
  errors.push(
    `version mismatch: package.json ${pkg.version} vs .claude-plugin/plugin.json ${plugin.version}`,
  );
}

if (errors.length) {
  console.error(`${errors.length} problem(s) in ${count} skills:`);
  for (const e of errors.slice(0, 50)) console.error(`  - ${e}`);
  process.exit(1);
}
console.log(`OK: ${count} skills have valid frontmatter`);
