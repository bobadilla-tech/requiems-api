# Requiems API Tutorials — Claude Context

This repo contains self-contained example projects demonstrating how to use the
[Requiems API](https://requiems.xyz). Each project lives in its own directory,
uses real API calls, and is intended to be runnable from scratch in under five
minutes.

The full API source code is at `../requiems-api/` (sibling directory).

---

## Before You Build — Read This First

| Task                            | Read first                                           |
| ------------------------------- | ---------------------------------------------------- |
| Create any new example          | [contributing.md](./contributing.md)                 |
| Understand the API architecture | `../requiems-api/claude.md`                          |
| Look up a specific endpoint     | `../requiems-api/docs/apis/{category}/{endpoint}.md` |

---

## API Reference

| Detail           | Value                                    |
| ---------------- | ---------------------------------------- |
| Base URL         | `https://api.requiems.xyz`               |
| Auth header      | `requiems-api-key: YOUR_API_KEY`         |
| Content-Type     | `application/json`                       |
| API key location | `.env` at repo root → `REQUIEMS_API=...` |
| Online docs      | https://requiems.xyz/en/apis             |
| Local docs       | `../requiems-api/docs/apis/`             |

**Minimal request example:**

```bash
curl -X POST https://api.requiems.xyz/v1/email/validate \
  -H "requiems-api-key: $REQUIEMS_API" \
  -H "Content-Type: application/json" \
  -d '{"email": "user@gmial.com"}'
```

**Response envelope:**

```json
{
  "data": { ... },
  "metadata": { "timestamp": "2026-01-01T00:00:00Z" }
}
```

All payload is inside `data`. Always check `data`, not the root object.

---

## Repo Conventions

- Every project is a **self-contained directory** at the repo root.
- Directory naming: `{framework}-{api-feature}` — lowercase, hyphenated.
  - Examples: `nextjs-email-validation`, `fastapi-qr-code`,
    `express-spell-check`
- The root `.env` holds the shared API key (`REQUIEMS_API`). Projects either
  read `../.env` directly or copy it locally. Never commit keys.
- Each project must have its own `README.md` and `.env.example`.

---

## Tooling Rules (Non-Negotiable)

### Node.js / TypeScript

- Package manager: **pnpm** — never npm, yarn, or bun
- Init: `pnpm create next-app` / `pnpm init`
- Install: `pnpm install` / `pnpm add`
- Run: `pnpm dev` / `pnpm build`

### Python

- Project manager: **uv** — never pip, poetry, or pipenv
- Linter/formatter: **ruff** — never flake8, black, isort, or pylint
- Init: `uv init {project-name}`
- Add deps: `uv add requests`
- Run: `uv run python main.py`
- Lint: `ruff check .`
- Format: `ruff format .`

### Other languages

- Go: standard Go modules (`go mod init`, `go run`)
- Ruby: Bundler (`bundle init`, `bundle exec`)

---

## How to Build a New Example

1. **Read the endpoint doc** at
   `../requiems-api/docs/apis/{category}/{endpoint}.md` — understand the request
   shape, response fields, and error cases.

2. **Create the project directory**: `mkdir {framework}-{feature}` at the repo
   root.

3. **Init the project** with the correct tooling (pnpm / uv — see above).

4. **Implement the demo** — keep it focused on one endpoint. Avoid boilerplate
   beyond what's needed to run the example.

5. **Write the project README** covering: what it does, prerequisites, setup
   steps, and example output.

6. **Create `.env.example`** with `REQUIEMS_API=your_api_key_here`.

7. **Update the root `readme.md`** — add a row to the Projects table.

---

## API Categories Quick Reference

| Category              | Local docs path                                  | Endpoint prefix       |
| --------------------- | ------------------------------------------------ | --------------------- |
| AI & Computer Vision  | `../requiems-api/docs/apis/ai-computer-vision/`  | `/v1/ai/`             |
| Animals               | `../requiems-api/docs/apis/animals/`             | `/v1/animals/`        |
| Email                 | `../requiems-api/docs/apis/email/`               | `/v1/email/`          |
| Entertainment         | `../requiems-api/docs/apis/entertainment/`       | `/v1/entertainment/`  |
| Finance               | `../requiems-api/docs/apis/finance/`             | `/v1/finance/`        |
| Health & Wellness     | `../requiems-api/docs/apis/health-wellness/`     | `/v1/health/`         |
| Internet & Technology | `../requiems-api/docs/apis/internet-technology/` | `/v1/tech/`           |
| Miscellaneous         | `../requiems-api/docs/apis/miscellaneous/`       | `/v1/misc/`           |
| Places                | `../requiems-api/docs/apis/places/`              | `/v1/places/`         |
| Text                  | `../requiems-api/docs/apis/text/`                | `/v1/text/`           |
| Transportation        | `../requiems-api/docs/apis/transportation/`      | `/v1/transportation/` |
| Weather               | `../requiems-api/docs/apis/weather/`             | `/v1/weather/`        |

---

## Current Projects

| Directory                  | Status   | API Used                  |
| -------------------------- | -------- | ------------------------- |
| `nextjs-email-validation/` | Complete | `POST /v1/email/validate` |
