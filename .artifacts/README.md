# Work artifacts

This directory stores historical artifacts created while completing a unit of
work: plans, audits, investigation notes, design decisions, handoff prompts,
implementation logs, review notes, and similar Markdown files.

## Layout

Each work unit has a dated folder under its year and month:

```
.artifacts/YYYY/MM/YYYYMMDD--<ticket-id>--<short-slug>/
```

Omit the ticket segment when there is no ticket. Keep the slug short and
lowercase. A work-unit folder can contain multiple related files; use clear
names such as `plan.md`, `audit.md`, `handoff-prompt.md`, and
`implementation-log.md`.

## How to use these records

Artifacts are snapshots, not living documentation or a backlog. Save plans and
working notes here when they are useful to preserve, alongside the work they
describe. Once a work unit is complete, leave its folder as a historical record.
Verify every claim against the current code and system state before relying on
it. Product documentation, runbooks, and maintained references belong in their
normal locations under `docs/` or alongside the relevant code.

A plan starts as research: understand the current state related to the
problem, improvement, or bug before deciding anything. Only after that does it
become an approach: what we will build, fix, test, or benchmark, and how,
described in prose rather than code where possible. Finish the write-up with a
"Final notes" section filled in after the work lands (what shipped, deviations,
follow-ups).

## Migration note

Before 2026-11 these records lived in `docs/plans/` (`YYYY-MM-DD-topic.md`),
`docs/audits/`, and `apps/clients/design-plans/`. They were moved here
unchanged: `docs/plans/2026-08-21-go-auth-foundation-phase-2.md` is now
`.artifacts/2026/08/20260821--go-auth-foundation-phase-2/plan.md`, audits are
`audit.md`, and the client plans use a `clients-` slug prefix. Old path
references inside historical records were intentionally left as written.
