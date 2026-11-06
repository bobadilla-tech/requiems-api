# Continuous Integration — Requiems API Clients

## Context

`example-apps/` contains minimal test apps that validate auto-generated API
clients compile correctly and can communicate with the Requiems API.

## Current State

- Franzaina implemented local testing across all languages. Everything tested
  locally and working.
- Tests previously ran without an API key, treating HTTP 401 as the success
  signal (proves network + serialization work).
- Next step: run these tests in CI with a real API key so tests validate actual
  API responses.

## Proposed CI Workflow (GitHub Actions)

**Trigger:** any PR that modifies files under `clients/`

**Steps:**

1. Checkout repo
2. Build TypeScript client
   (`cd clients/typescript && npm install && npm run build`)
3. Set up all language runtimes (Node.js, .NET 8, Python 3, Ruby, Go)
4. Run each example app with `REQUIEMS_API_KEY` injected as an env var
5. Fail the pipeline if any app exits non-zero

**Secret required:** added `REQUIEMS_API_KEY` to the GitHub repository secrets.

## Test Success Criteria

| Scenario     | Expected result                                 |
| ------------ | ----------------------------------------------- |
| With API key | HTTP 2xx — confirms full end-to-end integration |

## Example Apps Quick Reference

| Language    | Setup                                                  | Run                    |
| ----------- | ------------------------------------------------------ | ---------------------- |
| TypeScript  | `npm install`                                          | `npx ts-node index.ts` |
| C# (.NET 8) | —                                                      | `dotnet run`           |
| Python 3    | `pip install urllib3 pydantic python-dateutil aiohttp` | `python3 main.py`      |
| Ruby        | `gem install typhoeus`                                 | `ruby main.rb`         |
| Go          | `go mod tidy`                                          | `go run main.go`       |

> **Note:** TypeScript client must be built before running the TypeScript
> example app. Python: no venv needed in CI — install deps directly into the
> runner environment.

## Open Items

- [ ] Create `.github/workflows/ci.yml` implementing the workflow above
- [ ] Add `REQUIEMS_API_KEY` to GitHub repository secrets
- [ ] Update example apps to exit non-zero on unexpected responses (not just log
      them)
