# Tools Refactor — Audit, Fix, and Guide

The dashboard has 4 tool pages (Unit Conversion, Sentiment Analysis, Email
Validator, Quotes). A code audit revealed 6 recurring pitfalls introduced during
AI-assisted development. This plan documents the pitfalls, fixes each tool, then
codifies the correct approach in a new `docs/core/adding-tools.md` guide so
future engineers ship clean tools from the start.

---

## Part 1 — Current State Audit

### Tools inventory

| Tool               | Controller (JS)                      | Hero partial                                       | Show view                                |
| ------------------ | ------------------------------------ | -------------------------------------------------- | ---------------------------------------- |
| Unit Conversion    | `unit_conversion_demo_controller.js` | `partials/tools/unit_conversion/_hero.html.erb`    | `tools/unit_conversion/show.html.erb`    |
| Sentiment Analysis | `sentiment_demo_controller.js`       | `partials/tools/sentiment_analysis/_hero.html.erb` | `tools/sentiment_analysis/show.html.erb` |
| Email Validator    | `email_validator_demo_controller.js` | `partials/tools/email_validator/_hero.html.erb`    | `tools/email_validator/show.html.erb`    |
| Quotes             | `quotes_demo_controller.js`          | `partials/tools/quotes/_hero.html.erb`             | `tools/quotes/show.html.erb`             |

All files live under `apps/dashboard/`.

---

### Pitfall 1 — Hardcoded strings instead of `t()`

**Severity: High** — breaks i18n. App supports EN/ES/FR.

3 of 4 controllers have all error messages hardcoded in English:

```javascript
// BAD — unit_conversion_demo_controller.js, sentiment_demo_controller.js, email_validator_demo_controller.js
this._showError("Too many requests. Wait a moment and try again.");
this._showError("Could not reach the API. Check your connection.");
this._showError("No data returned.");
```

Only Quotes uses `t()` correctly — errors passed via data attributes on the hero
partial:

```erb
data-quotes-demo-error-rate-limit-value="<%= t('tools.quotes.demo.error_rate_limit') %>"
```

`tools.en.yml` only has keys for `tools.quotes.demo.*` (5 keys, 399 bytes
total).

**Fix:** Add locale keys for all 3 remaining tools; update hero partials to pass
errors as Stimulus values; update JS controllers to read
`this.errorRateLimitValue` etc. (Quotes pattern).

---

### Pitfall 2 — Unnecessary `_escapeHtml()` on own-backend data

**Severity: Medium** — dead code, misleading to readers, signals distrust of our
own API.

`unit_conversion_demo_controller.js` and `email_validator_demo_controller.js`
both define and call `_escapeHtml()` on fields returned by our own backend
(`data.from`, `data.to`, `data.result`, `data.formula`, email string,
suggestion). Our backend already sanitizes output. Double-escaping also corrupts
output (e.g. `&amp;` visible to users).

The method is copied verbatim in both files:

```javascript
_escapeHtml(str) {
  return String(str)
    .replace(/&/g, "&amp;").replace(/</g, "&lt;")
    .replace(/>/g, "&gt;").replace(/"/g, "&quot;");
}
```

**Fix:** Remove `_escapeHtml()` method and all call sites that wrap API response
fields. User-typed input never goes into `innerHTML` directly (it goes into the
POST body), so no escape is needed there either.

---

### Pitfall 3 — Direct `innerHTML` assignment with template literals

**Severity: High** — XSS vector, couples markup to JS, makes i18n impossible.

All 3 non-Quotes controllers assign multi-line HTML template strings to
`this.resultTarget.innerHTML`. Labels like "Conversion Result", "Breakdown",
"From", "To" are baked into JS strings — untranslatable, hard to style, not
reviewable by designers.

```javascript
// BAD — unit_conversion_demo_controller.js line 109
this.resultTarget.innerHTML = `
  <div class="rounded-xl border ...">
    <span class="text-sm font-semibold">Conversion Result</span>
    ${items}
  </div>`;
```

**Fix — migrate to Turbo Frames:**

1. Wrap the result `<div>` in each hero partial with
   `<turbo-frame id="tool-result">`.
2. Give the form a `data-turbo-frame="tool-result"` attribute and point `action`
   to a new Rails route.
3. Add lightweight demo actions to `tools_controller.rb` (or a
   `Tools::DemoController`):
   - `POST /tools/unit-conversion/demo` → calls service → renders
     `_result.html.erb` or `_error.html.erb` partial inside the Turbo Frame.
4. Result partials live in `partials/tools/{tool_name}/` alongside existing
   partials.
5. Labels and strings in partials use `t()` normally.
6. Remove `_renderResult` / `_renderError` JS methods entirely. Stimulus
   controller shrinks to form submission + loading state toggle.

Loading state: keep `data-unit-conversion-demo-target="spinner"` and toggle it
on `turbo:submit-start` / `turbo:submit-end` events in the controller.

This approach eliminates innerHTML, makes result markup reviewable, enables
i18n, and cuts each JS controller by ~60 lines.

---

### Pitfall 4 — Wrong input types (Unit Conversion)

**Severity: Medium** — poor UX; user has no idea what unit strings the API
accepts.

`unit_conversion/_hero.html.erb` lines 33-49 use `type="text"` for "Source unit"
and "Target unit". The API expects specific strings (e.g. `miles`, `km`,
`celsius`).

**Fix:** Replace both text inputs with `<select>` dropdowns. Group options by
category. Define the unit list in the hero partial (or a helper) so the Stimulus
controller just reads `.value` as before — no JS change needed.

```erb
<select data-unit-conversion-demo-target="from" ...>
  <optgroup label="Length">
    <option value="miles">Miles</option>
    <option value="km">Kilometers</option>
    ...
  </optgroup>
  <optgroup label="Temperature">
    <option value="celsius">Celsius</option>
    ...
  </optgroup>
</select>
```

---

### Pitfall 5 — Repetitive HTML not using shared components

**Severity: Medium** — maintenance burden; style drift between tools.

Every tool hero partial hand-rolls its own CTA buttons with inline Tailwind
classes instead of using `render "shared/button"`:

```erb
<!-- BAD — unit_conversion/_hero.html.erb lines 13-23 -->
<%= link_to new_user_registration_path,
      class: "inline-flex items-center px-6 py-3 rounded-lg bg-brand-primary text-white ..." do %>
  Get API Key
<% end %>
```

The shared `_button.html.erb` partial exists with 9 variants but tools ignore
it.

Additionally the `_cta.html.erb` section at the bottom of each tool is identical
across all 4 tools — same copy, same buttons.

**Fix:**

1. In all hero partials, replace raw `link_to`/`<button>` with
   `render "shared/button"`. Note: the shared button uses `bg-blue-600` for
   `primary` — add a `brand` variant that uses `bg-brand-primary` to match
   current tool styling, or align on blue.
2. Extract a shared `partials/tools/shared/_cta.html.erb` that all 4 tool show
   views render instead of their own copy.
3. The demo submit button inside the form should also use the button partial (or
   remain a raw `<button>` since it needs Stimulus `data-target` — acceptable
   exception if partial supports data pass-through, which it does via
   `btn_data`).

---

### Pitfall 6 — Button visibility in light mode

**Severity: High** — trust/reputation issue; buttons invisible to light-mode
users.

The shared `_button.html.erb` `outline` variant is:

```
bg-white dark:bg-gray-800 border-2 border-gray-300 dark:border-gray-600 text-gray-700 dark:text-gray-300
```

This is fine for light mode. However tool hero sections use raw `link_to` with
custom classes that may omit light-mode text/border colors. Audit each tool's
CTA link classes and ensure every button is visible in both themes.

**Fix:** Enforced automatically when tools switch to `render "shared/button"`
with the correct variant.

---

## Part 2 — Fix Plan (Execution Order)

### Step 1 — i18n keys (all tools)

Add to `apps/dashboard/config/locales/en/tools.en.yml`:

```yaml
en:
  tools:
    quotes:
      demo: # existing
        ...
    unit_conversion:
      demo:
        error_fill_all_fields: "Fill in all fields."
        error_rate_limit: "Too many requests. Wait a moment and try again."
        error_network: "Could not reach the API. Check your connection."
        error_no_data: "No data returned."
    sentiment_analysis:
      demo:
        error_empty: "Enter some text to analyze."
        error_rate_limit: "Too many requests. Wait a moment and try again."
        error_generic: "Something went wrong. Try again."
        error_network: "Could not reach the API. Check your connection."
        error_no_data: "No data returned."
    email_validator:
      demo:
        error_empty: "Enter an email address to verify."
        error_rate_limit: "Too many requests. Wait a moment and try again."
        error_generic: "Something went wrong. Try again."
        error_network: "Could not reach the API. Check your connection."
```

Add stub keys to `es/` and `fr/` locale files with `TODO: translate` values.

Update each tool's `_hero.html.erb` to pass error strings as Stimulus value
attributes (Quotes pattern). Update each JS controller to read
`this.error*Value` instead of hardcoded strings.

### Step 2 — Turbo Frame migration (all tools)

For each tool:

1. **Route** — add `post "/tools/{tool}/demo", to: "tools/demos#{tool_action}"`
   (or use one `demos_controller` with `tool` param).
2. **Controller action** — thin: read params, call existing service/proxy logic
   (reuse what the `/api/proxy` route does), render partial.
3. **Result partials** — `_demo_result.html.erb` and `_demo_error.html.erb` per
   tool in `partials/tools/{tool_name}/`.
4. **Hero partial** — wrap result div in
   `<turbo-frame id="{tool}-demo-result">`, add `action` + `data-turbo-frame` to
   form.
5. **JS controller** — strip `_renderResult`, `_renderError`, `_escapeHtml`.
   Keep loading state toggle wired to `turbo:submit-start`/`turbo:submit-end`.

### Step 3 — Unit Conversion inputs

Replace text inputs with grouped `<select>` elements. Unit list defined in the
partial (not in JS). Keep `data-unit-conversion-demo-target="from"` and `"to"`
so the controller reads `.value` unchanged.

### Step 4 — Button standardization

Replace all raw `link_to`/`<button>` CTA elements in tool hero partials with
`render "shared/button"`. Add `brand` variant to `_button.html.erb` for emerald
brand-primary color. Extract shared tool CTA block.

### Step 5 — Verify light/dark button visibility

For every tool, view in both light and dark mode. Confirm all interactive
elements are visible with sufficient contrast.

---

## Part 3 — New Guide: `docs/core/adding-tools.md`

A step-by-step guide (format matching `batch-apis.md` and
`adding-go-endpoints.md`) covering:

1. **What a tool page is** — marketing + interactive demo, not a feature page
2. **Directory layout** — file tree with annotations
3. **Demo form pattern** — Turbo Frame standard (with code example), why not
   `fetch` + `innerHTML`
4. **i18n requirements** — all strings via `t()`, locale key naming convention,
   stub ES/FR on add
5. **Input types** — use semantic HTML (`<select>`, `<textarea>`,
   `type="number"`) not generic `type="text"`
6. **Shared components** — always use `render "shared/button"`,
   `render "shared/card"`, etc.
7. **When NOT to escape HTML** — our backend is trusted; only escape genuinely
   untrusted user input rendered in the DOM, and use DOMPurify (not hand-rolled
   regex) if so
8. **Anti-patterns table** — don't → do
9. **Pre-merge checklist**

---

## Verification

After implementation:

1. `bin/rails server` — open each tool page in light + dark mode.
2. Unit Conversion: dropdowns show correct grouped units; submit with valid pair
   returns result in-page via Turbo.
3. Sentiment / Email Validator: demo forms work; result renders via Turbo Frame
   with no `innerHTML` in JS.
4. Switch browser to Spanish/French locale — error messages should show
   translated strings (or TODO stubs, not English hardcode).
5. `grep -r "innerHTML" apps/dashboard/app/javascript/controllers/` — should
   return 0 results for tool controllers.
6. `grep -r "_escapeHtml" apps/dashboard/app/javascript/controllers/` — should
   return 0 results.
7. `bin/rails test` / existing test suite green.
