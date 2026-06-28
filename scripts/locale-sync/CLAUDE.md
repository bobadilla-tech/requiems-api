# locale-sync — Claude Code context

## What this tool is

A Go CLI that fully replaces `i18n-tasks` for a Rails app and adds capabilities i18n-tasks lacks:
automatic hardcoded-string extraction from ERB, automatic `t()` call injection, ERB fragment merging,
Google Translate integration, deduplication, and key refactoring.

It is the **primary i18n toolchain** for `apps/dashboard`. Do not use `i18n-tasks` write commands
(add-missing, remove-unused, normalize) — use locale-sync instead. i18n-tasks is kept only as a
passive health-check tool (`bundle exec i18n-tasks health`).

## Project layout

```
scripts/locale-sync/
  main.go                     entry point (delegates to cmd.Execute())
  cmd/                        one file per command
  internal/
    locale/                   YAML read/write, key dedup, value linting
    source/                   ERB/Ruby scanning, t() injection, fragment merging
    translate/                Google Cloud Translation API client
```

See [ARCHITECTURE.md](ARCHITECTURE.md) for full package internals.

## Key invariants

**Write safety:** Every write command uses `setYAMLPath()` with a `nil` overwrite callback — it
skips any key that already has a non-empty value. `translate` is the only exception: it overwrites
values that match `IsPlaceholder()` (empty string or `^TODO:|^FIXME:` regex). Real translations
written by a human are never overwritten by any command.

**Key routing:** Keys are routed to YAML files by their top-level namespace.
`dashboard.foo.bar` → `config/locales/en/dashboard.en.yml`. Unknown namespaces → `shared.en.yml`.
This is determined in `locale.UpsertTopicFile()`.

**All-lang by default:** `refactor-key` has `--all-lang` (default true) — it renames the key in
every language directory. `discoverLangs()` reads subdirectories of `config/locales/` to find them.

**Mailer keys:** Mailer views use relative keys (`t '.subject'`) that static scanners can't resolve.
`*_mailer.*` is in `ignore_unused` in `i18n-tasks.yml` for this reason. locale-sync's `lint` and
`audit` commands have a parallel limitation — mailer relative keys are treated as used by matching
against the view-path resolved key in `source.ResolveRelativeKey()`.

**Dynamic keys:** ERB that interpolates full key segments (`t("#{i18n_scope}.heading")`) cannot be
statically resolved. These are excluded from prune via `--pattern` flag or pre-emptively excluded
in `i18n-tasks.yml` with `ignore_unused` patterns.

**YAML 1.1 boolean key bug:** Keys named `yes`, `no`, `true`, `false` must be quoted in YAML
(`"yes":`) or Ruby's Psych parser treats them as booleans and lookups return `nil`.
locale-sync does not currently detect this — it's caught at runtime or by manual YAML review.

## Adding a new command

1. Create `cmd/<name>.go` following the pattern in any existing cmd file.
2. Register with `rootCmd.AddCommand(...)` in an `init()` function.
3. Use `locale.Scan(root, lang)` to read entries, `locale.UpsertTopicFile()` to write.
4. Add `--dry-run` and `--root` flags to every write command.
5. No tests required in `cmd/` — test the logic in `internal/` packages instead.

## Running

```sh
cd scripts/locale-sync

# build once
go build -o locale-sync .

# or run directly
go run . <command> [flags]

# tests
go test ./...
```

## Google Translate credentials

`google.json` in this directory is the service account key. Never commit it.
Pass via environment variable:

```sh
GOOGLE_APPLICATION_CREDENTIALS=google.json go run . translate \
  --root ../../apps/dashboard --from en --to es
```

## Common tasks

**Add a missing key across all locales:**
```sh
go run . add-key --root ../../apps/dashboard \
  --key shared.buttons.save --value "Save changes"
# then translate it:
GOOGLE_APPLICATION_CREDENTIALS=google.json go run . translate \
  --root ../../apps/dashboard --to es,fr
```

**Rename a key:**
```sh
go run . refactor-key --root ../../apps/dashboard \
  --old shared.common.copy_btn --new shared.buttons.copy --dry-run
go run . refactor-key --root ../../apps/dashboard \
  --old shared.common.copy_btn --new shared.buttons.copy
```

**Extract hardcoded strings from a new view file:**
```sh
go run . report --root ../../apps/dashboard   # see what's there
go run . generate --root ../../apps/dashboard --fix --dry-run
go run . generate --root ../../apps/dashboard --fix
GOOGLE_APPLICATION_CREDENTIALS=google.json go run . translate \
  --root ../../apps/dashboard --to es,fr
```

**After a big refactor (check for dead keys):**
```sh
go run . prune --root ../../apps/dashboard --dry-run
go run . prune --root ../../apps/dashboard
```
