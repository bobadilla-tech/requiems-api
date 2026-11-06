# Requiems API Tutorials — Claude Context

This repo contains self-contained example projects demonstrating how to use the
[Requiems API](https://requiems.xyz). Each project lives in its own directory,
uses real API calls, and is intended to be runnable from scratch in under five
minutes.

The full API source code lives in the same repo: `../api/` (Go), with endpoint docs in `../../docs/`. This directory is `apps/tutorials` of the `requiems-api` monorepo.

---

## Before You Build — Read This First

| Task                            | Read first                                           |
| ------------------------------- | ---------------------------------------------------- |
| Create any new example          | [contributing.md](./contributing.md)                 |
| Understand the API architecture | `../../claude.md`                          |
| Look up a specific endpoint     | `../../docs/apis/{category}/{endpoint}.md` |

---

## API Reference

| Detail           | Value                                    |
| ---------------- | ---------------------------------------- |
| Base URL         | `https://api.requiems.xyz`               |
| Auth header      | `requiems-api-key: YOUR_API_KEY`         |
| Content-Type     | `application/json`                       |
| API key location | `.env` in `apps/tutorials/` → `REQUIEMS_API=...` |
| Online docs      | https://requiems.xyz/en/apis             |
| Local docs       | `../../docs/apis/`             |

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

- Every project is a **self-contained directory** directly inside `apps/tutorials/`.
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
   `../../docs/apis/{category}/{endpoint}.md` — understand the request
   shape, response fields, and error cases.

2. **Create the project directory**: `mkdir {framework}-{feature}` in `apps/tutorials/`
   root.

3. **Init the project** with the correct tooling (pnpm / uv — see above).

4. **Implement the demo** — keep it focused on one endpoint. Avoid boilerplate
   beyond what's needed to run the example.

5. **Write the project README** covering: what it does, prerequisites, setup
   steps, and example output.

6. **Create `.env.example`** with `REQUIEMS_API=your_api_key_here`.

7. **Update `apps/tutorials/readme.md`** — add a row to the Projects table.

---

## API Categories Quick Reference

| Category              | Local docs path                                  | Endpoint prefix       |
| --------------------- | ------------------------------------------------ | --------------------- |
| AI & Computer Vision  | `../../docs/apis/ai-computer-vision/`  | `/v1/ai/`             |
| Animals               | `../../docs/apis/animals/`             | `/v1/animals/`        |
| Email                 | `../../docs/apis/email/`               | `/v1/email/`          |
| Entertainment         | `../../docs/apis/entertainment/`       | `/v1/entertainment/`  |
| Finance               | `../../docs/apis/finance/`             | `/v1/finance/`        |
| Health & Wellness     | `../../docs/apis/health-wellness/`     | `/v1/health/`         |
| Internet & Technology | `../../docs/apis/internet-technology/` | `/v1/tech/`           |
| Miscellaneous         | `../../docs/apis/miscellaneous/`       | `/v1/misc/`           |
| Places                | `../../docs/apis/places/`              | `/v1/places/`         |
| Text                  | `../../docs/apis/text/`                | `/v1/text/`           |
| Transportation        | `../../docs/apis/transportation/`      | `/v1/transportation/` |
| Weather               | `../../docs/apis/weather/`             | `/v1/weather/`        |

---

## Current Projects

| Directory                  | Status   | API Used                  |
| -------------------------- | -------- | ------------------------- |
| `nextjs-email-validation/` | Complete | `POST /v1/email/validate` |
