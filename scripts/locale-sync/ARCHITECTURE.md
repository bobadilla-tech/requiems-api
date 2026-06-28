# Architecture

## Overview

```
main.go
  └── cmd.Execute()           cobra root command
        ├── cmd/report.go
        ├── cmd/generate.go
        ├── cmd/lint.go
        ├── cmd/audit.go
        ├── cmd/add_key.go
        ├── cmd/prune.go
        ├── cmd/translate.go
        ├── cmd/analyze.go
        ├── cmd/consolidate.go
        ├── cmd/refactor_key.go
        └── cmd/merge_fragments.go

internal/
  locale/      YAML I/O — read entries, write keys, dedup, value linting
  source/      Source scanning — extract hardcoded strings, audit t() calls, rewrite ERB
  translate/   Google Cloud Translation — API client, cache, brand protection, writer
```

---

## `internal/locale`

**`parser.go` — `Scan(root, lang) []Entry`**

Reads all `config/locales/{lang}/**/*.{lang}.yml` files and the root
`config/locales/{lang}.yml`. Returns a flat `[]Entry` where every leaf node
in every YAML file becomes one entry with a full dot-notation key including
the lang prefix (e.g. `en.shared.buttons.close`).

**`writer.go`**

All write operations go through two primitives:

| Function | Overwrites existing? | Used by |
|---|---|---|
| `setYAMLPath(doc, path, value)` | Never | `generate`, `consolidate`, `add-key`, `refactor-key` |
| `upsertYAMLPath(doc, path, value, shouldOverwrite)` | Only if callback returns true | `translate` (passes `IsPlaceholder`) |

Higher-level write functions:
- `UpsertTopicFile(root, lang, topic, []MergeCandidate)` — writes candidates to the correct topic file (e.g. `dashboard.en.yml`), creating file and intermediate keys as needed. Returns `(filePath, changed, error)`.
- `UpsertShared(root, lang, []MergeCandidate)` — writes to `shared.{lang}.yml`.
- `DeleteKeys(filePath, []string)` — removes keys by dot-path. Used by `prune` and `refactor-key`.
- `WriteSkeleton(root, lang, baseLang)` — creates a placeholder skeleton file for a new language by copying EN structure with empty values.

**`dedup.go` — `FindDuplicates`, `FilterExistingValueCandidates`, `SuggestSharedKey`**

Identifies duplicate string values across topic files and suggests a canonical
`shared.*` key path. `SuggestSharedKey` maps semantic prefixes to shared sub-namespaces
(e.g. a value like "Save" → `shared.buttons.save`, a status word → `shared.status.*`).

**`value_lint.go`**

Heuristic checks that flag values that look like they should NOT be a locale string:
email addresses, raw HTML, kebab-case identifiers, metric codes. Used by `audit --empty-values`.

---

## `internal/source`

**`extractor.go` — `Extract(root) []HardcodedString`**

Walks `app/` for `.rb`, `.erb`, `.haml`, `.slim` files and identifies hardcoded
user-visible strings using a multi-category classifier:

| Category | Example |
|---|---|
| `tag_content` | `<h1>Welcome</h1>` |
| `placeholder` | `placeholder="Enter email"` |
| `attribute` | `title="Settings"` |
| `ruby` | `flash[:notice] = "Saved!"` |

Skips strings that look like code (camelCase, snake_case, URLs, format strings).

**`audit.go` — `Audit(root, lang, definedKeys) AuditResult`**

Cross-references every `t('key')` call in source against `definedKeys`. Returns:
- `MissingKeys` — called in source but not in YAML
- `OrphanedKeys` — in YAML but never called in source
- `EmptyKeys` — in YAML with empty value

Resolves relative keys (`t '.title'`) via `ResolveRelativeKey(relFile, leaf)` which
reconstructs the full key from the file's view path (e.g.
`app/views/dashboard/index.html.erb` + `.title` → `dashboard.index.title`).

**`fixer.go` — `InjectTCall`, `RewriteFile`**

Takes a `HardcodedString` and its suggested locale key and rewrites the source line
to use `t('key')`. Handles ERB (`<%= t(...) %>`), Ruby method calls, HTML attributes
(`data-label="..."` → `data-label="<%= t(...) %>"`).

**`fragmenter.go` — `ParseFragments(file) []FragmentLine`**

Handles the harder case: user-visible text split across multiple ERB nodes, e.g.:

```erb
<p>
  You have <%= count %> items remaining in your
  <%= link_to "cart", cart_path %>.
</p>
```

Detects sentence fragments that span ERB interpolations and constructs a single
`t('key', var: expr)` replacement. Marks lines as "complex" (skip auto-patch) when:
- ERB expression contains `||` or `&&` (boolean operators)
- Expression contains HTML tags
- Multiple comma-separated arguments

`fragmentVarName(expr)` extracts a variable name from the ERB expression
(`user.name` → `name`, `log.response_time_ms` → `response_time_ms`). Sanitizes
result to valid Ruby identifier characters only — strips operators and symbols.

---

## `internal/translate`

**`client.go` — `Client`**

Wraps the Google Cloud Translation v2 REST API. Batch-translates strings with
quota awareness. Respects `GOOGLE_APPLICATION_CREDENTIALS` env var pointing to
a service account JSON file.

**`cache.go`**

Persistent JSON cache at `reports/translate-cache.json`. Cache key is
`{sourceLang}:{targetLang}:{sourceText}`. Avoids re-translating identical strings
between runs. `--no-cache` bypasses it.

**`protector.go` — `Protect`, `Restore`**

Replaces brand names and protected words with opaque tokens (`__PROT0__`,
`__PROT1__`, ...) before sending to the API, then restores them in the result.
Input: `--protect-words` flag or `--protect-file` (newline-separated word list).

**`walker.go` — `IsPlaceholder`, `Walk`**

`IsPlaceholder(v)` — returns true for empty strings and values matching
`^(?:TODO|FIXME):\s*[a-z]`. These are the only values `translate` will overwrite.
Real translations (any non-empty, non-placeholder string) are never touched.

`Walk` traverses the EN YAML tree, finds keys with missing or placeholder values
in the target locale, and returns them as a batch for translation.

**`writer.go` — `WriteTranslations`**

Calls `locale.UpsertEntries(filePath, entries, IsPlaceholder)` to write translated
values. The `IsPlaceholder` callback is what allows overwriting placeholder values
while protecting real translations.

---

## Data flow: `generate --fix`

```
source.Extract(root)
  → []HardcodedString (all hardcoded strings with file/line)

locale.FindDuplicates(strings, existingEntries)
  → []MergeCandidate (strings with suggested shared.* key)

locale.FilterExistingValueCandidates(candidates, entries)
  → filtered (drop candidates already in shared)

locale.UpsertShared(root, lang, candidates)
  → writes new keys to shared.{lang}.yml

source.RewriteFile(file, candidates)
  → replaces hardcoded strings with t('shared.key') in source
```

## Data flow: `translate`

```
locale.Scan(root, "en")
  → []Entry (all EN keys)

translate.Walk(enEntries, targetLangEntries)
  → []string (keys missing or placeholder in target lang)

translate.Protect(strings, protectedWords)
  → strings with brand names replaced by tokens

client.Translate(strings, targetLang)
  → []string (translated)

translate.Restore(translated, tokens)
  → []string (brand names restored)

translate.WriteTranslations(root, targetLang, keyValueMap)
  → writes to target lang YAML files, skipping real translations
```

---

## Key design decisions

**No global state.** Every command is stateless — reads from disk, transforms in memory, writes to disk. No daemon, no watch mode, no shared cache between commands (except the optional translate cache file).

**Cobra for CLI.** Each command is a `*cobra.Command` registered in its file's `init()`. Shared flags (`--root`, `--lang`) are re-declared per command to avoid cross-command flag pollution.

**`gopkg.in/yaml.v3` AST, not marshal/unmarshal.** The writer operates on raw `*yaml.Node` trees to preserve comments, key ordering, and scalar style. Marshal/unmarshal round-trips would destroy YAML formatting.

**Strict-off dynamic key inference.** The scanner doesn't try to evaluate `"#{scope}.key"` interpolations. Commands that could wrongly prune dynamic keys rely on `--pattern` exclusion or the i18n-tasks `ignore_unused` config in the Rails app.
