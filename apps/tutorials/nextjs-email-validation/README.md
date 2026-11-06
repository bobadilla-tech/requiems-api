# Next.js Email Validation with Requiems API

A tutorial project demonstrating real-time email validation in a
Next.js 16 signup flow using the [Requiems API](https://requiems.xyz).

## What this shows

| Feature                                     | File                                                                             |
| ------------------------------------------- | -------------------------------------------------------------------------------- |
| Requiems API client (server-only)           | [src/lib/requiems.ts](src/lib/requiems.ts)                                       |
| Server-side proxy route (keeps key secret)  | [src/app/api/validate-email/route.ts](src/app/api/validate-email/route.ts)       |
| Better Auth config + email validation hooks | [src/lib/auth.ts](src/lib/auth.ts)                                               |
| Debounced real-time validation form         | [src/components/SignupForm.tsx](src/components/SignupForm.tsx)                   |
| All 8 API response fields displayed         | [src/components/EmailValidationCard.tsx](src/components/EmailValidationCard.tsx) |
| Protected dashboard (session required)      | [src/app/dashboard/page.tsx](src/app/dashboard/page.tsx)                         |
| Next.js 16 proxy (route protection)         | [proxy.ts](proxy.ts)                                                             |
| Typed env validation                        | [src/env.ts](src/env.ts)                                                         |

## Architecture

```
Browser → /api/validate-email (Next.js proxy) → api.requiems.xyz/v1/email/validate
                  ↑
         REQUIEMS_API lives here only (no NEXT_PUBLIC_ prefix)
```

The client never calls the Requiems API directly. A Next.js route handler adds
the API key server-side. Better Auth's `hooks.before` re-validates on every
sign-up and sign-in, blocking invalid emails even if someone bypasses the UI.

## Setup

```bash
git clone <this-repo>
cd nextjs-email-validation
pnpm install
```

Copy the example env file and fill in your values:

```bash
cp .env.example .env.local
```

Required env vars:

| Variable             | Description                                                             |
| -------------------- | ----------------------------------------------------------------------- |
| `REQUIEMS_API`       | Your Requiems API key — get one at [requiems.xyz](https://requiems.xyz) |
| `BETTER_AUTH_SECRET` | Random secret ≥ 32 chars (`openssl rand -base64 32`)                    |
| `BETTER_AUTH_URL`    | Base URL of your app (`http://localhost:3000` for local dev)            |

Run the dev server:

```bash
pnpm dev
```

Open [http://localhost:3000](http://localhost:3000) — it redirects to `/signup`.

## Test emails for the YouTube demo

| Email                                   | What it triggers                                  |
| --------------------------------------- | ------------------------------------------------- |
| `demo@gmail.com`                        | Happy path — all 8 fields green                   |
| `user@gmial.com`                        | Typo suggestion → click to correct to `gmail.com` |
| `not-an-email`                          | `syntax_valid: false` — format invalid            |
| `test@mailinator.com`                   | `disposable: true` — blocked with alert           |
| `user@nonexistent-xyz-domain-12345.com` | `mx_valid: false` — no DNS records                |
| `User+Tag@Googlemail.COM`               | Normalization demo → `user@gmail.com`             |
| `test@yaho.com`                         | Another typo → suggestion `yahoo.com`             |

## API response fields

All 8 fields from `POST /v1/email/validate` are displayed in the
`EmailValidationCard` component:

```json
{
  "data": {
    "email": "User+Tag@Googlemail.COM",
    "valid": true,
    "syntax_valid": true,
    "mx_valid": true,
    "disposable": false,
    "normalized": "user@gmail.com",
    "domain": "googlemail.com",
    "suggestion": null
  }
}
```

> **Note:** `suggestion` is a **domain** suggestion (e.g. `"gmail.com"`), not a
> full email. The UI reconstructs the full address as `localPart@suggestion`.

## Auth header

The Requiems API uses `requiems-api-key` header (not `Authorization: Bearer`).
See [src/lib/requiems.ts](src/lib/requiems.ts) for the exact fetch call.

## Tech stack

- **Next.js 16** (App Router, TypeScript, Tailwind CSS v4)
- **shadcn/ui** (latest)
- **Better Auth** — `emailAndPassword` + `hooks.before`
- **better-sqlite3** — local SQLite for demo (no external DB needed)
- **@t3-oss/env-nextjs** — typed, validated env vars
- **Requiems API** — `POST /v1/email/validate`
