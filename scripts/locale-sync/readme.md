# locale-sync

Go CLI for Rails i18n — a full replacement for `i18n-tasks` with capabilities it lacks.

**What i18n-tasks can't do that locale-sync can:**
- Extract hardcoded strings from ERB/Ruby and inject `t()` calls automatically
- Merge sentence fragments split across ERB interpolations into a single `t()` call
- Translate missing keys via Google Cloud Translation with brand-name protection
- Rename a key across YAML and all source callers in one command
- Deduplicate identical strings across files and consolidate them to `shared.*`

**What i18n-tasks does that locale-sync also does:**
- Lint: exit 1 if any `t('key')` call references an undefined key
- Audit: report missing, orphaned, and empty keys
- Prune: delete YAML keys that no source file calls

---

## Requirements

- Go 1.21+
- Rails app with `config/locales/{lang}/` layout
- Google Cloud Translation API service account JSON (only for `translate`)

## Build

```sh
cd scripts/locale-sync
go build -o locale-sync .
# or without building:
go run . <command> [flags]
```

---

## Workflow

Commands have a natural order. Run the full sequence to bootstrap i18n from scratch,
then use individual commands as needed.

### Step 1 — See what needs extracting

```sh
locale-sync report --root ../../apps/dashboard
```

Scans all `.erb`, `.rb`, `.haml`, `.slim` files under `app/` and lists every hardcoded
user-visible string with file and line number. No files are changed.

```sh
# CI gate: fail if hardcoded strings exist in changed files
git diff --name-only origin/main | \
  locale-sync report --root ../../apps/dashboard --fail-on-found --changed-files -
```

### Step 2 — Extract strings into YAML

```sh
# Preview — see what keys would be created and which source lines would change
locale-sync generate --root ../../apps/dashboard --fix --dry-run

# Apply — writes YAML keys and replaces hardcoded strings with t() calls
locale-sync generate --root ../../apps/dashboard --fix
```

This does three things in one pass:
1. Finds strings used in more than one place → writes them to `shared.{lang}.yml`
2. Finds unique strings → writes them to the correct topic file (e.g. `home.en.yml`)
3. Rewrites source files replacing the string with the `t('key')` call

After this, the app works identically — only the source and YAML have changed.

```sh
# Only process files changed in the current branch
git diff --name-only origin/main | \
  locale-sync generate --root ../../apps/dashboard --fix --changed-files -

# Generate ES/FR skeleton files at the same time
locale-sync generate --root ../../apps/dashboard --fix --languages es,fr
```

### Step 3 — Merge ERB fragments

For sentence fragments split across ERB interpolations:

```erb
<!-- Before -->
<p>Hello <%= current_user.name %>, you have <%= count %> messages.</p>

<!-- After -->
<p><%= t('.greeting', name: current_user.name, count: count) %></p>
```

```sh
# Preview complex cases (requires manual attention)
locale-sync merge-fragments --root ../../apps/dashboard --dry-run

# Apply simple cases automatically
locale-sync merge-fragments --root ../../apps/dashboard --fix
```

Complex cases (boolean operators, nested HTML, multiple arguments) are flagged for
manual review and never auto-patched.

### Step 4 — Lint

```sh
locale-sync lint --root ../../apps/dashboard
```

Exits 0 if every `t('key')` call resolves to a defined YAML entry. Exits 1 with
`file:line: error: missing key "..."` output. Add to CI.

```sh
# Check a specific language
locale-sync lint --root ../../apps/dashboard --lang fr
```

If lint reports missing keys, add them:

```sh
locale-sync add-key \
  --root ../../apps/dashboard \
  --key dashboard.settings.title \
  --value "Settings"
```

### Step 5 — Translate

```sh
export GOOGLE_APPLICATION_CREDENTIALS=scripts/locale-sync/google.json

# Preview what would be translated
locale-sync translate --root ../../apps/dashboard --to es,fr --dry-run

# Translate — only fills empty and placeholder values, never overwrites real translations
locale-sync translate --root ../../apps/dashboard --to es,fr
```

Protect brand names from being translated:

```sh
locale-sync translate \
  --root ../../apps/dashboard \
  --to es,fr \
  --protect-words "Requiems API,AbstractAPI"

# Or from a file (one word/phrase per line, # = comment)
locale-sync translate \
  --root ../../apps/dashboard \
  --to es,fr \
  --protect-file .translate-dictionary.txt
```

Translation results are cached in `reports/translate-cache.json` — re-running the
same source string doesn't hit the API again. Use `--no-cache` to force a fresh call.

**Write safety:** `translate` only overwrites values that are empty or match
`^TODO:|^FIXME:` (case-sensitive). Any value written by a human translator is preserved.

### Step 6 — Prune dead keys

```sh
# Always preview first
locale-sync prune --root ../../apps/dashboard --dry-run

# Delete orphaned keys
locale-sync prune --root ../../apps/dashboard

# Prune only a specific namespace
locale-sync prune --root ../../apps/dashboard --pattern 'shared\.common\.'
```

Understands pluralization — `foo.one` and `foo.other` are kept when source calls
`t('foo', count: n)`.

---

## All commands

| Command | What it does |
|---|---|
| `report` | List hardcoded strings in source (read-only, CI-safe) |
| `generate` | Extract strings to YAML; optionally replace them with `t()` calls |
| `merge-fragments` | Merge ERB sentence fragments into single `t()` calls |
| `lint` | Exit 1 if any `t()` call references an undefined key |
| `audit` | Report missing, orphaned, empty keys |
| `add-key` | Add one key-value pair to the correct YAML file |
| `prune` | Delete YAML keys never referenced in source |
| `translate` | Fill missing keys via Google Cloud Translation |
| `analyze` | Find duplicate key names or identical values across files |
| `consolidate` | Deduplicate keys and rewrite all callers in one shot |
| `refactor-key` | Rename a key in YAML and all `t()` callers |

---

## Command reference

### `report`

```sh
locale-sync report --root <path>

# Fail CI on any hardcoded string
locale-sync report --root ../../apps/dashboard --fail-on-found

# Limit to changed files (reads newline-separated paths from stdin)
git diff --name-only origin/main | \
  locale-sync report --root ../../apps/dashboard --changed-files -

# Only ERB/Haml (skip Ruby files)
locale-sync report --root ../../apps/dashboard --erb-only
```

### `generate`

```sh
# Dry run
locale-sync generate --root ../../apps/dashboard --fix --dry-run

# Extract and write
locale-sync generate --root ../../apps/dashboard --fix

# Also create ES/FR skeleton files
locale-sync generate --root ../../apps/dashboard --fix --languages es,fr

# Only reuse existing keys, don't create new ones
locale-sync generate --root ../../apps/dashboard --fix --safe-only

# Skip shared.yml consolidation (write all keys to topic files)
locale-sync generate --root ../../apps/dashboard --fix --no-shared

# Changed files only
git diff --name-only origin/main | \
  locale-sync generate --root ../../apps/dashboard --fix --changed-files -
```

### `merge-fragments`

```sh
# Preview all fragments (simple + complex)
locale-sync merge-fragments --root ../../apps/dashboard --dry-run

# Auto-patch simple cases
locale-sync merge-fragments --root ../../apps/dashboard --fix

# Only report, don't patch anything
locale-sync merge-fragments --root ../../apps/dashboard
```

Complex cases are printed but never auto-patched. They require assigning ERB
expressions to local variables first (eliminates boolean operators and side effects).

### `lint`

```sh
locale-sync lint --root ../../apps/dashboard

# Specific language
locale-sync lint --root ../../apps/dashboard --lang fr
```

Exit codes: `0` = all resolved, `1` = missing keys found.

### `audit`

```sh
# Show orphaned keys (in YAML, never called in source)
locale-sync audit --root ../../apps/dashboard --orphaned

# Show missing keys (called in source, not in YAML)
locale-sync audit --root ../../apps/dashboard --missing

# Show both
locale-sync audit --root ../../apps/dashboard --all

# Show keys with empty values
locale-sync audit --root ../../apps/dashboard --empty-values

# Compare EN coverage against FR
locale-sync audit --root ../../apps/dashboard --compare-locale fr
```

### `add-key`

```sh
locale-sync add-key \
  --root ../../apps/dashboard \
  --key shared.buttons.save \
  --value "Save changes"

# Dry run
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

Key is routed to the YAML file matching its top-level namespace
(`shared.*` → `shared.en.yml`, `dashboard.*` → `dashboard.en.yml`).
File and intermediate keys are created if they don't exist.

### `prune`

```sh
locale-sync prune --root ../../apps/dashboard --dry-run
locale-sync prune --root ../../apps/dashboard

# Specific language
locale-sync prune --root ../../apps/dashboard --lang fr

# Limit to a namespace
locale-sync prune --root ../../apps/dashboard --pattern 'shared\.common\.'
```

### `translate`

```sh
export GOOGLE_APPLICATION_CREDENTIALS=/path/to/service-account.json

# Preview
locale-sync translate --root ../../apps/dashboard --to es,fr --dry-run

# Translate (uses cache)
locale-sync translate --root ../../apps/dashboard --to es

# Force re-translate (bypass cache)
locale-sync translate --root ../../apps/dashboard --to es --no-cache

# Save JSON report
locale-sync translate --root ../../apps/dashboard --to es \
  --report-file reports/translate.json

# Protect brand names
locale-sync translate --root ../../apps/dashboard --to es \
  --protect-words "Requiems API,AbstractAPI"
```

### `analyze`

```sh
# Find duplicate key names and identical values
locale-sync analyze --root ../../apps/dashboard

# Include keys with same name but different values
locale-sync analyze --root ../../apps/dashboard --all

# Also scan source for hardcoded strings
locale-sync analyze --root ../../apps/dashboard --source
```

### `consolidate`

One-shot deduplication: finds duplicate keys → writes to `shared.*` → rewrites
all callers → deletes old keys.

```sh
locale-sync consolidate --root ../../apps/dashboard --dry-run
locale-sync consolidate --root ../../apps/dashboard

# Rewrite callers but skip deletion (review with prune later)
locale-sync consolidate --root ../../apps/dashboard --no-prune
```

### `refactor-key`

Renames a key across all locale files (all languages by default) and all source callers.
Old key is deleted immediately — no separate prune step needed.

```sh
# Dry run first
locale-sync refactor-key \
  --root ../../apps/dashboard \
  --old shared.common.copy_btn \
  --new shared.buttons.copy \
  --dry-run

# Apply
locale-sync refactor-key \
  --root ../../apps/dashboard \
  --old shared.common.copy_btn \
  --new shared.buttons.copy

# Only rename in EN (skip FR/ES)
locale-sync refactor-key \
  --root ../../apps/dashboard \
  --old shared.common.copy_btn \
  --new shared.buttons.copy \
  --all-lang=false
```

---

## CI integration

Minimal gate — fails if any `t()` call is undefined:

```yaml
- name: Lint i18n keys
  run: |
    cd scripts/locale-sync
    go run . lint --root ../../apps/dashboard
```

Full pipeline — also blocks hardcoded strings in changed files:

```yaml
- name: Check hardcoded strings
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

```
config/locales/
  en/
    home.en.yml          # en.home.*
    dashboard.en.yml     # en.dashboard.*
    shared.en.yml        # en.shared.*  (deduplicated strings)
    auth.en.yml          # en.auth.*
  es/
    home.es.yml
    shared.es.yml
    ...
  fr/
    home.fr.yml
    shared.fr.yml
    ...
```

Keys are routed to the file whose name matches the top-level namespace
(`dashboard.foo.bar` → `dashboard.{lang}.yml`). Unknown namespaces go
to `shared.{lang}.yml`.

---

## vs i18n-tasks

| Feature | locale-sync | i18n-tasks |
|---|---|---|
| Extract hardcoded strings from ERB | ✅ `report`, `generate` | ❌ |
| Inject `t()` calls into source | ✅ `generate --fix` | ❌ |
| Merge ERB sentence fragments | ✅ `merge-fragments` | ❌ |
| Translate missing keys | ✅ `translate` (Google API) | ✅ (multiple backends) |
| Lint undefined `t()` calls | ✅ `lint` | ✅ `health` / `missing` |
| Prune orphaned keys | ✅ `prune` | ✅ `remove-unused` |
| Rename keys + update callers | ✅ `refactor-key` | ✅ `mv` |
| Deduplicate shared values | ✅ `consolidate` | ❌ |
| Dynamic key inference | `--pattern` exclusion | `strict: false` wildcard |
| Rails integration | standalone Go binary | Ruby gem in Gemfile |

i18n-tasks is still useful as a passive health checker (`bundle exec i18n-tasks health`)
because its Prism-based scanner handles some Rails-specific patterns (before_actions,
model translations) that locale-sync's static scanner does not. Run locale-sync for all
write operations; use i18n-tasks health-only in CI.
