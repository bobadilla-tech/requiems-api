<p align="center">
  <p align="center">
    <a href="https://requiems.xyz/?utm_source=github&utm_medium=logo" target="_blank">
      <img src="https://raw.githubusercontent.com/bobadilla-tech/requiems-api/refs/heads/main/apps/dashboard/app/assets/images/logo.png" alt="Requiems API" width="280" />
    </a>
  </p>
  <p align="center">
    All-in-one backend for SaaS products.
  </p>
  <p align="center">
    <i>A product by <a href="https://bobadilla.tech">Bobadilla Technologies</a></i>
  </p>
</p>

# Requiems API Skills

Installable AI skills that give your coding agent direct knowledge of
[Requiems API](https://requiems.xyz), authentication, fraud detection, payments
intelligence, and global data, all through one unified API.

Drop a skill into your agent and start calling Requiems API endpoints without
copying docs or writing boilerplate.

## Install in Claude Code

Register this repository as a Claude Code plugin marketplace:

```
/plugin marketplace add bobadilla-tech/requiems-api-skills
```

Then browse and install:

1. Select **Browse and install plugins**
2. Select **requiems-api-skills**
3. Select **Install now**

Or install directly:

```
/plugin install requiems-api-skills@requiems-api-skills
```

Once installed, mention the skill naturally: _"Use the identity skill to protect
this signup endpoint."_

## Install in OpenCode

OpenCode discovers skills from `.opencode/skills/<name>/SKILL.md` (project) or
`~/.config/opencode/skills/<name>/SKILL.md` (global). Copy the folder in:

```bash
npm install @requiems/api-skills
mkdir -p ~/.config/opencode/skills/requiems-identity
cp node_modules/@requiems/api-skills/skills/identity-risk-post-protect/SKILL.md ~/.config/opencode/skills/requiems-identity/SKILL.md
```

Already have the Claude Code plugin installed? OpenCode also reads
`.claude/skills/` and `~/.claude/skills/` directly, so it picks up the same
skills with no extra step.

## Install in GitHub Copilot

Copilot's agent skills use the same `SKILL.md` convention, one folder per skill
under `.github/skills/<name>/SKILL.md` (repo) or
`~/.copilot/skills/<name>/SKILL.md` (personal, cross-repo). This is separate
from the older single-file `.github/copilot-instructions.md`.

```bash
npm install @requiems/api-skills
mkdir -p .github/skills/requiems-identity
cp node_modules/@requiems/api-skills/skills/identity-risk-post-protect/SKILL.md .github/skills/requiems-identity/SKILL.md
```

Already have the Claude Code plugin installed? Copilot also reads
`.claude/skills/` directly, so it picks up the same skills with no extra step.

## Install Manually (any agent)

Skills are plain Markdown files. Works anywhere that accepts custom instructions
or a `SKILL.md` convention:

```bash
npm install @requiems/api-skills
# Then copy the .md file from node_modules/@requiems/api-skills/skills/ wherever your agent expects it
```

## What's Inside

Each skill gives your agent:

- Full endpoint reference (paths, params, response shapes)
- Copy-paste request examples
- Error codes and handling patterns
- Engine composition guides (e.g. `/v1/signup/protect`)
