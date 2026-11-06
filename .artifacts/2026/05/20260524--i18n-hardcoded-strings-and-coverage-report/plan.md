# i18n: Hardcoded Strings Extraction + Coverage Report Tool

The Rails dashboard supports English, Spanish, and French, but had two
unresolved problems: user-facing strings hardcoded in controllers and views
bypassed the translation system entirely, and there was no tooling to detect or
report which translation keys were missing per language.

---

## Initial Problem

### 1 — Hardcoded strings

Flash messages, mailer content, and several view strings were written as raw
English string literals rather than `t()` calls. This meant:

- Non-English users received English flash messages regardless of their locale
  setting.
- Mailer text templates were all English while their HTML counterparts already
  used `t()`.
- No compile-time or runtime indication that these strings were untranslated.

### 2 — No coverage visibility

Spanish was ~54% translated (1,513 / ~2,800 EN keys). French was ~50% (~1,406 /
~2,800). Three entire feature files were missing for Spanish (`systems`,
`divisions`) and three for French (`systems`, `case_studies`, `divisions`).
There was no way to run a report to see what was missing.

---

## Initial State (before this work)

| Area                                                                                    | Status                                                               |
| --------------------------------------------------------------------------------------- | -------------------------------------------------------------------- |
| EN locale                                                                               | Complete                                                             |
| ES locale                                                                               | ~54% — missing `systems`, `divisions` files entirely                 |
| FR locale                                                                               | ~50% — missing `systems`, `case_studies`, `divisions` files entirely |
| Admin controllers                                                                       | All flash strings hardcoded                                          |
| Dashboard controllers (`settings`, `api_keys`, `billing`)                               | All flash strings hardcoded                                          |
| Public controllers (`sales_inquiries`, `suggestions`, `categories`, `apis`, `examples`) | All flash strings hardcoded                                          |
| Mailers — HTML                                                                          | Already using `t()`                                                  |
| Mailers — text                                                                          | Fully hardcoded                                                      |
| i18n tooling                                                                            | None                                                                 |

---

## Work Done

### Hardcoded string extraction

Every hardcoded user-facing string in controllers and views replaced with `t()`
calls. Keys added to the appropriate EN locale file with correct namespacing. ES
and FR stubs (`"TODO: translate"`) added in parallel so the coverage report
surfaces them immediately.

**Controllers fixed:**

| Controller                                | Keys extracted |
| ----------------------------------------- | -------------- |
| `dashboard/settings_controller.rb`        | 7              |
| `dashboard/api_keys_controller.rb`        | 6              |
| `dashboard/billing_controller.rb`         | 6              |
| `sales_inquiries_controller.rb`           | 3              |
| `suggestions_controller.rb`               | 3              |
| `categories_controller.rb`                | 1              |
| `apis_controller.rb`                      | 2              |
| `examples_controller.rb`                  | 1              |
| `admin/abuse_reports_controller.rb`       | 5              |
| `admin/api_keys_controller.rb`            | 4              |
| `admin/private_deployments_controller.rb` | 7              |
| `admin/promotions_controller.rb`          | 8              |
| `admin/users_controller.rb`               | 13             |
| `admin/analytics_controller.rb`           | 1              |
| `admin/dashboard_controller.rb`           | 1              |

**Views and mailers fixed:**

- `examples/show.html.erb` — tutorial coming-soon block, action links
- `sales_inquiries/new.html.erb` — hero description paragraph
- `account_deletion_mailer/confirmation.text.erb` — mirrored to use same keys as
  the HTML version
- `private_deployment_mailer/request_received.{html,text}.erb` — both templates
  now fully translated; new keys added under
  `dashboard.mailers.private_deployments.request_received`

**New locale files created:**

- `config/locales/en/admin.en.yml` — all admin flash strings
- `config/locales/es/admin.es.yml` — stubs
- `config/locales/fr/admin.fr.yml` — stubs

### Missing-file scaffold

Five locale files that were entirely absent were generated using
`bin/scaffold_locales`, a one-time Ruby script that reads an EN source file and
writes a parallel file with all leaf values replaced by `"TODO: translate"`:

- `config/locales/es/systems.es.yml`
- `config/locales/es/divisions.es.yml`
- `config/locales/fr/systems.fr.yml`
- `config/locales/fr/case_studies.fr.yml`
- `config/locales/fr/divisions.fr.yml`

### i18n-tasks gem

Added `gem "i18n-tasks", "~> 1.0"` to the `development` group. Configured in
`config/i18n-tasks.yml`.

### Rake tasks

`lib/tasks/i18n.rake` adds two tasks:

- `rake i18n:report` / `rake i18n:report[es]` — wraps `i18n-tasks missing` with
  a header per locale.
- `rake i18n:todos` — lists every `"TODO: translate"` placeholder across all
  locale files with file path and line number.

---

## Acceptance Criteria

- [ ] `grep -rn 'notice:\s*"[A-Z]\|alert:\s*"[A-Z]\|flash\.now.*=\s*"[A-Z]'
      app/controllers/`
      returns no results (excluding internal reason strings passed to
      `revoke!`).
- [ ] `bundle exec i18n-tasks missing -l en` returns 0 missing keys.
- [ ] `bundle exec i18n-tasks missing -l es` and `fr` return only pre-existing
      content gaps — no keys from controllers or mailers.
- [ ] `rake i18n:todos` lists all scaffolded `TODO: translate` entries so
      translators know exactly what remains.
- [ ] Mailer text templates use `t()` for all prose; no English string literals.

---

## How to Generate a Translation Report

These commands are how a translator (or PM) can see exactly what is missing
before shipping a new language version.

### Missing keys per locale

```bash
# All non-English locales
cd apps/dashboard
bundle exec i18n-tasks missing

# Single locale
bundle exec i18n-tasks missing -l es
bundle exec i18n-tasks missing -l fr
```

The output table shows: locale, missing key path, and the English value (or the
source file/line if the key is only used in a view but not yet defined
anywhere). This is the authoritative gap list for a translator.

### Wrapper rake tasks

```bash
# ES + FR combined (with headers)
rake i18n:report

# Single locale
rake i18n:report[es]
rake i18n:report[fr]
```

### TODO placeholders (scaffolded files)

```bash
rake i18n:todos
```

This shows every `"TODO: translate"` entry with its file path and line number. A
translator can open each file, search for `TODO`, and fill in the translated
value.

### Adding a new locale in the future

1. Add the locale symbol to `config.i18n.available_locales` in
   `config/application.rb`.
2. Run `bin/scaffold_locales` for each existing EN source file:
   ```bash
   ruby bin/scaffold_locales config/locales/en/home.en.yml pt
   ruby bin/scaffold_locales config/locales/en/dashboard.en.yml pt
   # ... repeat for each module file
   ```
3. Run `bundle exec i18n-tasks missing -l pt` to confirm all keys are accounted
   for (even if as stubs).
4. Hand the `rake i18n:todos` output to a translator.

---

## Notes

- The `bin/scaffold_locales` script is a developer utility only. It generates
  structural stubs — it does not translate. Running it again on an existing file
  would overwrite real translations, so it is only safe to run for files that do
  not yet exist.
- Admin strings (`admin.*`) have ES and FR stubs but are lower priority — the
  admin panel is typically used in English only. They are included for
  completeness and will appear in `rake i18n:todos`.
- The 50 pre-existing EN→ES/FR gaps visible in `i18n-tasks missing` output (home
  page sections: `engine_spotlight`, `how_it_works`, `systems`, `use_cases`,
  `why_different`, plus some shared nav keys) are content translation work, not
  a code issue. They were missing before this change.
