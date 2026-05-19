# Skill: High-Quality Plane Ticket Writing

## Purpose

Create self-contained, implementation-ready Plane tickets with enough context
that a developer can start work immediately without chat follow-up.

## Mandatory Rules

1. Never write dates in the ticket title or description.
2. Never write vague summaries like "fix bug" without explaining:
   - what failed
   - who/what was impacted
   - why it happened (if known)
   - how it was fixed
3. If context is incomplete, investigate first. Do not publish low-context
   tickets.
4. Always include business impact, not only technical details.
5. Tickets must be self-contained handoff docs, not short reminders.
6. Never leave implementation tickets with minimal one-paragraph descriptions.

## Required Investigation Before Writing

Collect context in this order when available:

1. PR title and description
2. PR/commit diff (files changed, meaningful code paths)
3. Commit titles and commit messages
4. Related docs/design notes
5. Validation evidence (tests/checks/runbooks)

If a source is missing, say that explicitly and continue with remaining
evidence.

## Ticket Structure (Use Exactly)

### Title

- Specific outcome, not process wording.
- Include domain and affected area.
- Avoid generic prefixes unless the team standard requires them.

### Original Problem

- Describe the real failure mode.
- Include user/system symptom and scope.

### Scope

- Define what is in scope for this ticket.
- Define explicit non-goals if adjacent work exists.

### Technical Context

- Name key services/files/routes/tables that the assignee must inspect.
- Include current-state constraints the assignee must account for.

### Previous State

- Explain how the system behaved before the change.
- Mention constraints, gaps, or risks in prior implementation.

### Work Done

- Summarize the implemented changes by component.
- Mention key files/services touched when useful.
- Link behavior changes to the problem they resolve.

### Acceptance Criteria

- Provide concrete completion checks.
- Include behavior expectations, not just "code added."

### Dependencies

- List prerequisite tickets/decisions.
- If none, explicitly write "None."

### Result

- Describe post-change behavior in production terms.
- Include business/user outcome and risk reduction.

### Business Impact

- Explicitly state impact on reliability, support load, developer velocity,
  cost, security, or conversion/retention.

### Evidence

- List commit subjects and PR references used to derive the ticket.
- For planning/spec tickets, cite stakeholder/product inputs and docs analyzed.

## Quality Bar (Must Pass)

1. A reader can understand the issue without opening code.
2. Problem statement is concrete and falsifiable.
3. Result section states measurable or observable effect.
4. Business impact is present and non-generic.
5. No dates appear anywhere in the ticket body/title.
6. No placeholder text remains.
7. Ticket has enough implementation context (files/flows/data model) to start
   work immediately.
8. Acceptance criteria are testable and specific.

## Guardrails

- Do not invent root causes. If unknown, write: "Root cause under investigation;
  observed behavior and mitigation documented."
- Do not claim business metrics you cannot support.
- Prefer concise, dense writing over long narrative.

## Output Template

Title: <specific outcome>

Original Problem:
<concrete failure and symptoms>

Scope:

- <in scope item>
- <in scope item>
- Non-goals: <what is out of scope>

Technical Context:

- <service/file/route/table context 1>
- <service/file/route/table context 2>
- <important current-state constraint>

Previous State:
<how things worked before and why that was insufficient>

Work Done:

- <change 1>
- <change 2>
- <change 3>

Acceptance Criteria:

- <behavior check 1>
- <behavior check 2>
- <behavior check 3>

Dependencies:

- <ticket/decision dependency or None>

Result:
<post-change system behavior>

Business Impact:

- <impact 1>
- <impact 2>

Evidence:

- PR: <link or id>
- Commits:
  - <sha> <subject>
  - <sha> <subject>
- Additional Context:
  - <stakeholder requirement or design doc reference>

## Suggested Invocation

"Use skill plane-ticket-quality to draft Plane issues from these commits/PRs. If
context is weak, inspect diffs and PR descriptions first, then write tickets
with business impact and no dates."

---

## Plane API (tasks.bobadilla.tech): identifiers and `curl` examples

Use this to **call the Plane REST API** from a shell. Do **not** paste real PATs
into skills, commits, or chat; use `export PLANE_API_KEY='plane_api_…'` locally
or a private `.env` that stays untracked.

### Identifiers (this workspace)

| What                                      | Value                                  |
| ----------------------------------------- | -------------------------------------- |
| Host (UI + API)                           | `https://tasks.bobadilla.tech`         |
| API prefix                                | `/api/v1/`                             |
| Workspace **slug**                        | `core`                                 |
| Project id (**UUID**)                     | `5da3caa1-cbba-4e0d-b9a9-0bcad39bba09` |
| Example backlog **state** id (new issues) | `57725193-924f-4477-b9bc-b305e553efb9` |

Re-check **state** and **project** ids in the Plane UI if workflows were
reconfigured. Assignees and labels in JSON bodies are **member UUIDs** and
**label UUIDs** from the workspace (Members / Labels screens or API responses).

**Base path for issues in this project:**

```text
https://tasks.bobadilla.tech/api/v1/workspaces/core/projects/5da3caa1-cbba-4e0d-b9a9-0bcad39bba09/issues/
```

### Headers that work through Cloudflare

- **`X-API-Key`**: personal access token (Plane); not `Authorization: Bearer …`
  for PAT-based calls.
- **`Content-Type: application/json`** and **`Accept: application/json`**
- **`User-Agent`**: use a normal browser string. A generic script UA can trigger
  Cloudflare **error 1010**; a desktop Chrome string is a safe default.

Define once:

```bash
export PLANE_API_KEY='plane_api_REPLACE_ME'
export PLANE_UA='Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36'
```

### `curl`: create an issue (POST)

```bash
curl -sS -w '\nHTTP %{http_code}\n' -X POST \
  'https://tasks.bobadilla.tech/api/v1/workspaces/core/projects/5da3caa1-cbba-4e0d-b9a9-0bcad39bba09/issues/' \
  -H "X-API-Key: ${PLANE_API_KEY}" \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json' \
  -H "User-Agent: ${PLANE_UA}" \
  -d '{
    "name": "Example: created via curl",
    "description_html": "<p>Short HTML body. Priorities: urgent, high, medium, low.</p>",
    "priority": "medium",
    "state": "57725193-924f-4477-b9bc-b305e553efb9"
  }'
```

Save the returned JSON `id` (issue UUID) for the examples below.

### `curl`: list issues (GET, first page)

Query params vary by Plane version; `cursor` / `per_page` are common. If the
server ignores unknown params, you still get a page of results.

```bash
curl -sS -w '\nHTTP %{http_code}\n' -G \
  'https://tasks.bobadilla.tech/api/v1/workspaces/core/projects/5da3caa1-cbba-4e0d-b9a9-0bcad39bba09/issues/' \
  -H "X-API-Key: ${PLANE_API_KEY}" \
  -H 'Accept: application/json' \
  -H "User-Agent: ${PLANE_UA}" \
  --data-urlencode 'per_page=20'
```

### `curl`: fetch one issue (GET)

Replace `ISSUE_UUID` with the work item id from create or list.

```bash
ISSUE_UUID='00000000-0000-0000-0000-000000000000'

curl -sS -w '\nHTTP %{http_code}\n' \
  "https://tasks.bobadilla.tech/api/v1/workspaces/core/projects/5da3caa1-cbba-4e0d-b9a9-0bcad39bba09/issues/${ISSUE_UUID}/" \
  -H "X-API-Key: ${PLANE_API_KEY}" \
  -H 'Accept: application/json' \
  -H "User-Agent: ${PLANE_UA}"
```

Use a **trailing slash** on URLs; some setups return redirects without it.

### `curl`: partial update (PATCH)

Many installs accept **`/issues/{uuid}/`**. Current public Plane docs also show
**`/work-items/{uuid}/`** — if PATCH on `issues` fails with 404/301, try the
same JSON against `work-items` instead.

```bash
ISSUE_UUID='00000000-0000-0000-0000-000000000000'

curl -sS -w '\nHTTP %{http_code}\n' -X PATCH \
  "https://tasks.bobadilla.tech/api/v1/workspaces/core/projects/5da3caa1-cbba-4e0d-b9a9-0bcad39bba09/issues/${ISSUE_UUID}/" \
  -H "X-API-Key: ${PLANE_API_KEY}" \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json' \
  -H "User-Agent: ${PLANE_UA}" \
  -d '{
    "name": "Example: renamed via curl",
    "priority": "high",
    "description_html": "<p>Updated description.</p>"
  }'
```

Example with **assignees** (array of **user UUIDs** from Members):

```bash
curl -sS -w '\nHTTP %{http_code}\n' -X PATCH \
  "https://tasks.bobadilla.tech/api/v1/workspaces/core/projects/5da3caa1-cbba-4e0d-b9a9-0bcad39bba09/issues/${ISSUE_UUID}/" \
  -H "X-API-Key: ${PLANE_API_KEY}" \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json' \
  -H "User-Agent: ${PLANE_UA}" \
  -d '{
    "assignees": ["MEMBER_UUID_1", "MEMBER_UUID_2"]
  }'
```

### Rate limits and errors

Roughly **60 requests/minute** per key is a safe cadence for bulk creates. On
**HTTP 429**, sleep and retry. **403** with **error 1010** is usually Cloudflare
blocking the client fingerprint—fix **User-Agent** / network path first, not the
JSON body.
