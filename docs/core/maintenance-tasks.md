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

## Generated File Checklist

Before committing generated-file updates:

1. Regenerate the relevant files.
2. Review the generated diffs.
3. Commit generated outputs with the source docs, catalog, or route changes.
4. Avoid hand-editing generated outputs.
