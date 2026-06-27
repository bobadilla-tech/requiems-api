# locale-sync

Go CLI for automating Rails i18n YAML management. Extracts hardcoded strings, validates `t()` calls, translates missing keys, and prunes dead translations.

## Requirements

- Go 1.21+
- Rails app with `config/locales/{lang}/` layout (e.g. `config/locales/en/home.en.yml`)
- Google Cloud Translation API credentials (only for `translate`)

## Build

```sh
cd scripts/locale-sync
go build -o locale-sync .
# or run without installing:
go run . <command> [flags]
```

---

## Recommended workflow

Commands have a natural order. Follow this sequence the first time, then use individual commands as needed.

### 1. See what needs i18n

```sh
locale-sync report --root ../../apps/dashboard
```

Reports hardcoded strings in ERB/Ruby. No files changed. Use `--fail-on-found` in CI to prevent regressions.

### 2. Extract hardcoded strings into YAML

```sh
# Preview first
locale-sync generate --root ../../apps/dashboard --fix --dry-run

# Apply
locale-sync generate --root ../../apps/dashboard --fix
```

Finds duplicates, writes them into `shared.{lang}.yml`, replaces inline strings with `t()` calls. Works on ERB and Ruby files. Use `--erb-only` to skip `.rb` files.

To extract only from files changed in a PR:

```sh
git diff --name-only origin/main | \
  locale-sync generate --root ../../apps/dashboard --fix --changed-files -
```

### 3. Validate all t() calls resolve

```sh
locale-sync lint --root ../../apps/dashboard
```

Exits 0 if every `t('key')` call has a matching YAML entry. Exits 1 with `file:line: error: missing key "..."` output. Run this in CI — it catches keys that will silently fail or show `translation missing:` errors in production.

If lint reports genuinely missing keys, add them:

```sh
locale-sync add-key \
  --root ../../apps/dashboard \
  --key dashboard.settings.title \
  --value "Settings"
```

Repeat `lint` until it exits 0.

### 4. Remove dead translations

```sh
# Always preview first
locale-sync prune --root ../../apps/dashboard --dry-run

# Remove orphaned keys
locale-sync prune --root ../../apps/dashboard
```

Removes YAML keys that no `t()` call references. Safe to run after `generate --fix` or any refactor. The prune command understands pluralization (`t('foo', count: n)` keeps `foo.one`/`foo.other`) and array/hash access patterns.

### 5. Fill translations for other languages

```sh
# Set credentials
export GOOGLE_APPLICATION_CREDENTIALS=/path/to/service-account.json

# Preview
locale-sync translate --root ../../apps/dashboard --to=es,fr --dry-run

# Translate
locale-sync translate --root ../../apps/dashboard --to=es,fr

# Translate all known languages
locale-sync translate --root ../../apps/dashboard --to=all
```

Only fills keys that are missing or placeholder values — never overwrites real translations. Use `--protect-words` for brand names that must not be translated:

```sh
locale-sync translate \
  --root ../../apps/dashboard \
  --to=es \
  --protect-words="Requiems API,AbstractAPI,IPstack"
```

Or maintain a dictionary file:

```sh
# .translate-dictionary.txt
# lines starting with # are comments
Requiems API
AbstractAPI
NeverBounce

locale-sync translate --root ../../apps/dashboard --to=es --protect-file=.translate-dictionary.txt
```

---

## All commands

| Command        | What it does                                                                       |
| -------------- | ---------------------------------------------------------------------------------- |
| `report`       | List hardcoded strings in source (no changes)                                      |
| `generate`     | Extract duplicates to shared YAML; optionally replace hardcoded strings with `t()` |
| `lint`         | Exit 1 if any `t()` call references an undefined key (CI gate)                     |
| `audit`        | Cross-reference YAML keys vs source calls; show orphaned, missing, empty           |
| `add-key`      | Add a single key-value pair to the correct YAML file                               |
| `prune`        | Delete YAML keys that are never called in source                                   |
| `translate`    | Fill missing translations via Google Cloud Translation API                         |
| `analyze`      | Find duplicate key names or identical values across YAML files                     |
| `consolidate`  | Deduplicate keys and rewrite all callers in one shot                               |
| `refactor-key` | Rename a key in YAML and all `t()` callers                                         |

---

## Command reference

### `report`

```sh
locale-sync report --root <path>

# CI: fail the build if hardcoded strings exist
locale-sync report --root ../../apps/dashboard --fail-on-found

# Only check ERB/Haml
locale-sync report --root ../../apps/dashboard --erb-only

# Only changed files
git diff --name-only origin/main | \
  locale-sync report --root ../../apps/dashboard --changed-files -
```

### `generate`

```sh
# Dry run — see what would change
locale-sync generate --root ../../apps/dashboard --fix --dry-run

# Extract strings and generate ES/FR skeletons
locale-sync generate --root ../../apps/dashboard --fix --languages es,fr

# Only safe replacements (reuse existing keys, no new keys)
locale-sync generate --root ../../apps/dashboard --fix --safe-only

# Skip shared YAML consolidation
locale-sync generate --root ../../apps/dashboard --fix --no-shared
```

### `lint`

```sh
locale-sync lint --root ../../apps/dashboard

# Check a specific language
locale-sync lint --root ../../apps/dashboard --lang fr
```

Exit codes: `0` = all resolved, `1` = missing keys found.

### `audit`

```sh
# Show orphaned keys (defined but never called)
locale-sync audit --root ../../apps/dashboard --orphaned

# Show missing keys (called but not defined)
locale-sync audit --root ../../apps/dashboard --missing

# Show both
locale-sync audit --root ../../apps/dashboard --all

# Show keys with empty values
locale-sync audit --root ../../apps/dashboard --empty-values

# Compare EN against FR coverage
locale-sync audit --root ../../apps/dashboard --compare-locale fr
```

### `add-key`

```sh
locale-sync add-key \
  --root ../../apps/dashboard \
  --key shared.buttons.save \
  --value "Save changes"

# Preview without writing
locale-sync add-key \
  --root ../../apps/dashboard \
  --key shared.buttons.save \
  --value "Save changes" \
  --dry-run

# Add to a specific language
locale-sync add-key \
  --root ../../apps/dashboard \
  --lang es \
  --key shared.buttons.save \
  --value "Guardar cambios"
```

Key is placed in the YAML file that owns that namespace (e.g. `shared.*` → `shared.en.yml`, `dashboard.*` → `dashboard.en.yml`). File and intermediate keys are created if needed.

### `prune`

```sh
# Always preview first
locale-sync prune --root ../../apps/dashboard --dry-run

# Prune a specific language
locale-sync prune --root ../../apps/dashboard --lang fr

# Limit to a specific namespace
locale-sync prune --root ../../apps/dashboard --pattern 'shared\.common\.'
```

Pluralization forms (`.one`, `.other`, etc.) are never pruned when their base key is called in source with `count:`.

### `translate`

```sh
export GOOGLE_APPLICATION_CREDENTIALS=/path/to/service-account.json

# Preview
locale-sync translate --root ../../apps/dashboard --to=es,fr --dry-run

# Translate with cache (default)
locale-sync translate --root ../../apps/dashboard --to=es

# Force API call, bypass cache
locale-sync translate --root ../../apps/dashboard --to=es --no-cache

# Save a JSON report
locale-sync translate --root ../../apps/dashboard --to=es \
  --report-file reports/translate.json
```

### `analyze`

```sh
# Find duplicate key names and identical values
locale-sync analyze --root ../../apps/dashboard

# Include keys with different values under the same name
locale-sync analyze --root ../../apps/dashboard --all

# Also scan source for hardcoded strings
locale-sync analyze --root ../../apps/dashboard --source
```

### `consolidate`

One-shot deduplication: finds duplicate keys → writes to shared YAML → rewrites all `t()` callers → deletes old keys.

```sh
# Always dry-run first
locale-sync consolidate --root ../../apps/dashboard --dry-run

# Run it
locale-sync consolidate --root ../../apps/dashboard

# Rewrite callers but skip deleting old keys (review with prune later)
locale-sync consolidate --root ../../apps/dashboard --no-prune
```

### `refactor-key`

```sh
# Rename a key (dry-run first)
locale-sync refactor-key \
  --root ../../apps/dashboard \
  --old shared.common.copy_btn \
  --new shared.buttons.copy \
  --dry-run

# Apply (old key stays — run prune after verifying)
locale-sync refactor-key \
  --root ../../apps/dashboard \
  --old shared.common.copy_btn \
  --new shared.buttons.copy

# After verifying app works:
locale-sync prune --root ../../apps/dashboard
```

---

## CI integration

Minimal gate — fails if any `t()` call is undefined:

```yaml
# GitHub Actions example
- name: Lint i18n keys
  run: |
    cd scripts/locale-sync
    go run . lint --root ../../apps/dashboard
```

Full pipeline on PRs:

```yaml
- name: Check for hardcoded strings
  run: |
    cd scripts/locale-sync
    git diff --name-only origin/main | \
      go run . report --root ../../apps/dashboard --fail-on-found --changed-files -

- name: Lint i18n keys
  run: |
    cd scripts/locale-sync
    go run . lint --root ../../apps/dashboard
```

---

## YAML file layout

The tool expects locale files under `config/locales/{lang}/`:

```
config/locales/
  en/
    home.en.yml          # en.home.*
    dashboard.en.yml     # en.dashboard.*
    shared.en.yml        # en.shared.*
    auth.en.yml          # en.auth.*
  es/
    home.es.yml
    ...
```

Keys in `add-key` and `generate` are routed to the file whose name matches the top-level namespace (e.g. `dashboard.foo.bar` → `dashboard.en.yml`). Keys with an unknown namespace go into `shared.{lang}.yml`.
