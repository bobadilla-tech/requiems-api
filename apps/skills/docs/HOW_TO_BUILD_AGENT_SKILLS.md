# How to Auto-Generate AI Agent Skills From Your API Docs

This is a write-up of the pattern behind this repo, generalized. If you already
have structured API documentation somewhere, you can turn it into installable
AI-agent "skills" with a small, boring pipeline instead of hand-maintaining a
separate context file. This doc is meant to double as source material for a blog
post — it explains _why_ each piece exists, not just what it does. For the
Requiems-specific implementation details, see [DESIGN.md](DESIGN.md).

## The core idea

Treat an agent skill as a **build artifact**, not hand-written content. If you
already maintain structured docs for humans (OpenAPI, YAML endpoint specs,
whatever), you have a source of truth. A "skill" is just that source of truth,
transformed into the format an agent expects, regenerated on a schedule, and
shipped like any other package.

That reframing gets you two things for free that a hand-maintained `AGENT.md`
never has: it can't drift from reality (it's derived, not typed by hand), and
updating it is a CI job, not a task on someone's todo list.

## What a "skill" actually is

Across Claude Code, OpenCode, and GitHub Copilot's agent-skills support, the
convention has converged on the same shape: a `SKILL.md` file — plain Markdown,
with a small YAML front-matter header for machine-readable fields (name, and
whatever else your agent's loader cares about) and a natural-language body
underneath. No proprietary schema, no SDK. That's a deliberate design point
worth keeping in your own implementation: the lowest common denominator across
agent tools is "a markdown file in a folder," so that's the compile target, not
any one vendor's format.

## The pipeline, in five stages

This repo's version of the pattern is `requiems-api` (structured API docs) →
`requiems-api-skills` (this repo, published as `@requiems/api-skills` on npm).
Mapped onto the general pattern:

**1. A structured source of truth.** Your API docs need to already be data, not
prose — YAML, JSON, an OpenAPI spec, a database table. If the source is a Word
doc or a wiki page, write that structured layer first; there's no transforming
free text reliably. Here, that's `apps/dashboard/config/api_docs/*.yml`.

**2. A pure transform function.** One function: takes the parsed doc, returns a
markdown string. No file I/O, no CLI parsing, no side effects — see
`buildSkillMarkdown()` in [`scripts/build/index.ts`](../scripts/build/index.ts).
Keeping this pure paid off directly: when this repo migrated its build script
off Deno onto Node, the entire markdown-building function needed _zero_ changes.
Only the I/O shell around it (reading files, writing files, parsing CLI args)
had runtime-specific code to port. If you take one thing from this doc, take
this: the part of your transform that shapes output should never touch a
filesystem or a runtime API directly.

**3. A CLI wrapper around the transform.** Something you can run locally with
`node script.ts --source <dir> --output <dir>` and get the same result CI gets.
If you can't run your generator on your own laptop without spinning up CI, you
can't debug it either.

**4. A scheduled job that runs the transform and diffs the output.** Don't try
to track "what changed upstream" — that's stateful and fragile. Instead,
regenerate everything from scratch every time, and let `git status
--porcelain`
tell you what's different. It's simpler, self-healing (a missed run doesn't lose
anything, the next run just produces a bigger diff), and it's the one thing git
is already good at.

**5. A human gate before anything ships.** Regeneration produces a pull request,
not a direct commit, and merging that PR is _not_ the same event as publishing a
new package version. Two separate, deliberate actions: "this reflects current
docs" (merge) and "this is now public" (a maintainer pushes a release tag). See
[Current Architecture](DESIGN.md#current-architecture) for the concrete version
of this.

## Design decisions worth stealing (and their tradeoffs)

### One generated file per endpoint, not one per API

A single `skills/sentiment.md` covering ten endpoints means every regen touches
that one file, and a PR reviewer can't tell which endpoint actually changed
without reading the whole diff. `skills/<api>-<method>-<path>/SKILL.md` per
endpoint keeps diffs scoped: if only one endpoint's docs changed upstream, only
one file shows up in the PR.

### Gate the version bump on the actual diff, not the cron firing

A weekly job that always bumps the version — even when nothing changed —
produces a version-only PR every single week, forever. Diff `skills/` first;
only bump if there's a real change. No-op weeks produce no PR at all.

### Decouple "merged" from "published"

The tempting version of this pipeline auto-publishes whenever `main` changes. We
deliberately didn't do that: publishing requires a maintainer to push a `vX.Y.Z`
tag by hand. The honest tradeoff is that it's an extra step someone can forget —
but it means nothing reaches the public registry that a human didn't explicitly
decide to ship, which matters a lot more once real users depend on the package.

### Guard the publish step by cross-checking the tag against the package version

Before `npm publish` runs, verify the pushed tag's version actually matches
`package.json`. Costs three lines of bash; the alternative is discovering the
mismatch as a cryptic npm 403, or worse, successfully publishing the wrong
version.

### Keep the toolchain boring

This repo's transform script briefly ran on Deno while the rest of the repo
(`package.json`, publishing) was Node — motivated by Deno's nicer
`--allow-read`/`--allow-write` permission model for a script that reads and
writes files. In practice the script never touched anything untrusted, so that
sandboxing bought nothing, and running two runtimes meant two lockfiles and an
extra `setup-deno` CI step just to run one file. Consolidating onto Node (using
its native TypeScript type-stripping — no `ts-node`, no build step) removed a
whole toolchain for free. Don't pay for isolation you don't need.

### A reviewer ping beats a broken review gate

GitHub's "request review from a team" API needs org-scope token permissions that
a workflow's default `GITHUB_TOKEN` usually doesn't have — it fails with
`Resource not accessible by integration` unless you provision a PAT or GitHub
App just for this. A plain `@org/team` mention written directly into the PR body
needs no special permissions and still notifies everyone on the team; it just
isn't an enforced approval gate. Whether that tradeoff is acceptable depends on
how much you trust "someone will look before merging" versus needing GitHub to
block the merge outright.

## Applying this to your own API

If you want to do this for your own API:

1. Make sure your API docs are structured data somewhere, not just prose.
2. Write one pure function: doc → markdown string, matching the `SKILL.md`
   convention (front-matter + body).
3. Wrap it in a CLI you can run locally.
4. Add a scheduled CI job: run the CLI, diff the output, open a PR only if
   something changed.
5. Keep publishing a separate, explicit step from merging — don't let a docs
   update silently become a public release.
6. Pick one runtime and stay boring about it.

This repo currently generates 60+ `SKILL.md` files from ~60 source YAML docs,
regenerated weekly with zero manual editing, published as a single npm package.
The pipeline that does it is under 200 lines of TypeScript plus three GitHub
Actions steps.
