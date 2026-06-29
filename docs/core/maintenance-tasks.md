# Maintenance Tasks

Use this page for recurring generated-file tasks that are easy to forget.

## Regenerate Sitemap

Run this after changing routes, locales, API catalog entries, examples, systems,
or case-study pages.

```bash
docker exec requiem-dev-dashboard-1 bin/rails sitemap:refresh
```

This updates:

- `apps/dashboard/public/sitemap.xml`
- `apps/dashboard/public/core-sitemap.xml`

The sitemap is generated from `apps/dashboard/config/sitemap.rb`. Do not edit
the generated XML files by hand.

## Regenerate OpenAPI Spec

Run this after changing dashboard API docs or the API catalog.

```bash
pnpm i && pnpm lint:fix && pnpm format
```

The generated OpenAPI file is ignored by the shared worker Biome config, so the
temporary config above applies the same formatter settings without changing the
repo-wide ignore rules.

This reads:

- `apps/dashboard/config/api_docs/*.yml`
- `apps/dashboard/config/api_catalog.yml`

This updates:

- `apps/workers/auth-gateway/src/generated/openapi.ts`

The Auth Gateway serves this generated spec as `/openapi.json`. Do not edit the
generated TypeScript file by hand.

## Fill Missing Translations (rori18n)

Run after adding new EN locale keys (new tool pages, new UI text, etc.) to
fill matching ES and FR files automatically.

```bash
cd /path/to/rori18n   # repo: github.com/bobadilla-tech/rori18n
GOOGLE_APPLICATION_CREDENTIALS=google.json go run . translate \
  --root ../requiems-api/apps/dashboard \
  --from en --to es,fr \
  --protect-file ../requiems-api/apps/dashboard/.translate-dictionary.txt
```

Use `--dry-run` to preview without writing. rori18n only fills empty or
`TODO: ...` placeholder values — never overwrites a real translation.

The `.translate-dictionary.txt` file at `apps/dashboard/` lists brand names
that must not be translated (Requiems API, NeverBounce, IPstack, etc.).

To check what's missing before translating:

```bash
go run . audit --root ../requiems-api/apps/dashboard
```

## Generated File Checklist

Before committing generated-file updates:

1. Regenerate the relevant files.
2. Review the generated diffs.
3. Commit generated outputs with the source docs, catalog, or route changes.
4. Avoid hand-editing generated outputs.
