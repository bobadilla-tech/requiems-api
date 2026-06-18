# Adding a Tool Page

> file layout, demo form pattern, i18n, input types, and shared components

Tool pages are marketing pages with a live interactive demo. Each page shows
what an API does, why someone would use it, and lets visitors try it without an
account. This guide covers how to build one correctly.

---

## 1. Goals

1. **No innerHTML in JavaScript** — Result rendering lives in Rails ERB, not JS
   template strings.
2. **All strings via `t()`** — Every user-facing string must come from a locale
   file. Hard-coding English is not acceptable.
3. **Correct input types** — Use `<select>` for enumerated options, `<textarea>`
   for long text, `type="number"` for numeric input.
4. **Shared components** — Use `render "partials/shared/button"`, not custom
   inline `link_to` classes.
5. **No escaping your own backend** — Do not call `escapeHtml()` on data
   returned by our own API. The backend already sanitizes output.

---

## 2. File layout

```
apps/dashboard/app/
├── controllers/
│   └── tool_demos_controller.rb          # one action per tool
├── javascript/controllers/
│   └── {tool_name}_demo_controller.js    # client-side validation + loading state only
├── views/
│   ├── tool_demos/
│   │   └── {tool_name}.html.erb          # Turbo Frame wrapper rendered by the demo action
│   └── partials/tools/
│       ├── shared/
│       │   └── _cta.html.erb             # shared CTA section (accepts: heading, body, docs_url)
│       └── {tool_name}/
│           ├── _hero.html.erb            # live demo form + Turbo Frame slot
│           ├── _what_it_does.html.erb
│           ├── _use_cases.html.erb
│           ├── _api_combinations.html.erb
│           ├── _faq.html.erb
│           └── _cta.html.erb             # wraps partials/tools/shared/_cta
├── tools/
│   └── {tool_name}/
│       └── show.html.erb
config/locales/en/
└── tools.en.yml                          # all tool strings live here
```

---

## 3. Demo form pattern (Turbo Frame)

The demo form submits to a Rails action. The action calls the API and renders a
result partial inside a `<turbo-frame>`. **Never render results via `innerHTML`
in JavaScript.**

### 3a. Hero partial — the form

```erb
<%# partials/tools/{tool_name}/_hero.html.erb %>

<div class="mt-12 max-w-xl mx-auto"
     data-controller="{tool-name}-demo"
     data-{tool-name}-demo-error-empty-value="<%= t('tools.{tool_name}.demo.error_empty') %>">

  <form data-action="submit->{tool-name}-demo#beforeSubmit turbo:submit-start->{tool-name}-demo#onSubmitStart turbo:submit-end->{tool-name}-demo#onSubmitEnd"
        action="<%= tool_demo_{tool_name}_path %>"
        method="post"
        data-turbo-frame="{tool_name}-demo-result"
        novalidate>
    <%= hidden_field_tag :authenticity_token, form_authenticity_token %>

    <%# form fields with name= attributes — these become params in the controller %>

    <button type="submit" data-{tool-name}-demo-target="button" ...>
      <span data-{tool-name}-demo-target="spinner" class="hidden">...</span>
      Submit
    </button>

    <p data-{tool-name}-demo-target="errorMessage"
       class="hidden mt-2 text-sm text-red-500"
       role="alert"></p>
  </form>

  <turbo-frame id="{tool_name}-demo-result"></turbo-frame>
</div>
```

Key points:

- **`name=` attributes on inputs** — Turbo submits the form as a normal POST;
  Rails reads `params[:field_name]`.
- **`data-turbo-frame`** must match the `id` of the `<turbo-frame>` below the
  form.
- **`hidden_field_tag :authenticity_token`** — required; the demo routes are
  outside the locale scope.
- Pass only the error strings that JS needs (empty-field check) as Stimulus
  value attributes. All API error messages come from the server via the result
  partial.

### 3b. Routes

```ruby
# config/routes.rb — outside the locale scope block
post "tools/demos/{tool-name}", to: "tool_demos#{action_name}", as: :tool_demo_{tool_name}
```

### 3c. Controller action

```ruby
# app/controllers/tool_demos_controller.rb
class ToolDemosController < ApplicationController
  layout false

  def {tool_name}
    field = params[:field].to_s.strip

    if field.blank?
      return render_demo_error("{tool_name}", t("tools.{tool_name}.demo.error_empty"))
    end

    result = api_call(endpoint: "/v1/...", method: "POST", params: { field: field })

    if result.status_code == 429
      return render_demo_error("{tool_name}", t("tools.{tool_name}.demo.error_rate_limit"))
    end

    unless result.status_code == 200
      return render_demo_error("{tool_name}", t("tools.{tool_name}.demo.error_generic"))
    end

    data = result.data&.dig("data", "data") || result.data&.dig("data")
    return render_demo_error("{tool_name}", t("tools.{tool_name}.demo.error_no_data")) if data.nil?

    render "tool_demos/{tool_name}", locals: { data: data }
  end

  private

  def api_call(endpoint:, method:, params:)
    ApiProxyService.call(
      endpoint: endpoint,
      method: method,
      params: params,
      forwarded_for: request.headers["CF-Connecting-IP"] || request.remote_ip
    )
  rescue StandardError => e
    Rails.logger.error("ToolDemosController error: #{e.message}")
    ApiProxyService::Result.new(status_code: 500, data: nil, error: e.message)
  end

  def render_demo_error(tool, message)
    render "tool_demos/demo_error", locals: { tool: tool, message: message }
  end
end
```

Use `ApiProxyService.call` — do not duplicate HTTP call logic.

### 3d. Result view

```erb
<%# app/views/tool_demos/{tool_name}.html.erb %>
<turbo-frame id="{tool_name}-demo-result">
  <div class="mt-4 text-left rounded-xl border border-gray-200 dark:border-gray-700 bg-white dark:bg-gray-800 overflow-hidden">
    <div class="px-4 py-3 bg-gray-50 dark:bg-gray-900">
      <span class="text-sm font-semibold text-gray-900 dark:text-white">
        <%= t("tools.{tool_name}.demo.result_heading") %>
      </span>
    </div>
    <div class="px-4 py-1">
      <%# render result data here using t() for labels %>
    </div>
  </div>
</turbo-frame>
```

The `<turbo-frame id>` must match the slot in the hero partial exactly.

### 3e. Stimulus controller — thin

```javascript
// app/javascript/controllers/{tool_name}_demo_controller.js
import { Controller } from "@hotwired/stimulus";

export default class extends Controller {
  static targets = ["input", "button", "errorMessage", "spinner"];
  static values = { errorEmpty: String };

  beforeSubmit(event) {
    this._clearError();
    if (!this.inputTarget.value.trim()) {
      event.preventDefault();
      this._showError(this.errorEmptyValue);
    }
  }

  onSubmitStart() {
    this.buttonTarget.disabled = true;
    this.spinnerTarget.classList.remove("hidden");
  }

  onSubmitEnd() {
    this.buttonTarget.disabled = false;
    this.spinnerTarget.classList.add("hidden");
  }

  _showError(msg) {
    this.errorMessageTarget.textContent = msg;
    this.errorMessageTarget.classList.remove("hidden");
  }

  _clearError() {
    this.errorMessageTarget.classList.add("hidden");
    this.errorMessageTarget.textContent = "";
  }
}
```

The controller does **only** client-side validation and loading state. No
`fetch`, no `innerHTML`, no result rendering.

---

## 4. i18n requirements

All user-facing strings in tool pages must use `t()`. Add keys to
`apps/dashboard/config/locales/en/tools.en.yml` and stub `TODO: translate` in
the matching `es/tools.es.yml` and `fr/tools.fr.yml` files.

### Key structure

```yaml
en:
  tools:
    { tool_name }:
      demo:
        error_empty: "..." # client-side empty-field check (passed via data attribute)
        error_rate_limit: "..." # 429 from API
        error_generic: "..." # any other non-200
        error_network: "..." # network error (optional — only for JS fetch pattern)
        error_no_data: "..." # API returned 200 but empty data
        result_heading: "..." # card header in result view
        result_failed_heading: "..." # card header in error result
        label_*: "..." # row labels in the result table
```

### Passing strings to JS (for empty-field check only)

Quotes pattern — data attribute on the controller element:

```erb
data-{tool-name}-demo-error-empty-value="<%= t('tools.{tool_name}.demo.error_empty') %>"
```

JS reads it via:

```javascript
static values = { errorEmpty: String };
// use: this.errorEmptyValue
```

All other error messages come from the server-rendered result partial and use
`t()` directly — no JS needed.

---

## 5. Input types

| Data type                                     | Use                                         |
| --------------------------------------------- | ------------------------------------------- |
| Enumerated set (units, categories, countries) | `<select>` with `<optgroup>` for categories |
| Free-form text                                | `<textarea>` with `rows` set                |
| Email                                         | `<input type="email">`                      |
| Numbers                                       | `<input type="number">`                     |
| Generic text with no enumeration              | `<input type="text">` — last resort         |

Never use `type="text"` when the valid values are a known set. The user cannot
discover options from a text box.

---

## 6. Shared components

### Button

```erb
<%= render "partials/shared/button",
      text:    "Get API Key",
      variant: "brand",
      href:    new_user_registration_path,
      size:    "lg" %>
```

Available variants: `brand`, `primary`, `secondary`, `success`, `danger`,
`warning`, `outline`, `outline-primary`, `ghost`, `link`.

Do not write inline `link_to` with custom Tailwind classes for CTA buttons.

### CTA section

```erb
<%# partials/tools/{tool_name}/_cta.html.erb %>
<%= render "partials/tools/shared/cta",
      heading:  "Your call-to-action headline.",
      body:     "Supporting sentence with value proposition.",
      docs_url: api_path("your-api-slug") %>
```

---

## 7. HTML escaping rules

**Do not escape data from our own backend.** The internal API already sanitizes
output. Calling `escapeHtml()` on API response fields causes double-encoding
(users see `&amp;` literally) and signals distrust of our own infrastructure.

**Do escape** genuinely untrusted user input if it is ever inserted into the DOM
directly (outside of a Turbo Frame response). The correct library is
**DOMPurify**, not hand-rolled regex replacements.

In practice: with the Turbo Frame pattern, result data goes through Rails ERB
(`<%= value %>`) which HTML-escapes automatically. No JS escaping is needed.

---

## 8. Anti-patterns

| Anti-pattern                                       | Correct approach                               |
| -------------------------------------------------- | ---------------------------------------------- |
| `innerHTML = \`<div>...\``                         | Render result via Rails + Turbo Frame          |
| `_escapeHtml()` on API response fields             | Remove it — Rails ERB escapes at render time   |
| Hardcoded English error strings in JS              | Pass via `data-*-value` from ERB using `t()`   |
| `type="text"` for unit/category selectors          | Use `<select>` with `<optgroup>`               |
| Raw `link_to` with inline Tailwind classes for CTA | Use `render "partials/shared/button"`          |
| Copy-pasting the CTA section HTML per tool         | Use `render "partials/tools/shared/cta"`       |
| Defining `def method_name` inside ERB              | Use a local lambda: `helper = ->(arg) { ... }` |
| `fetch()` to `/api/proxy` + JS result rendering    | Submit form to `ToolDemosController` action    |

---

## 9. Pre-merge checklist

- [ ] `grep -r "innerHTML" app/javascript/controllers/{tool_name}*` returns
      nothing
- [ ] `grep -r "_escapeHtml\|escapeHtml" app/javascript/controllers/{tool_name}*`
      returns nothing
- [ ] All user-facing strings go through `t()` — no hardcoded English in JS or
      ERB
- [ ] New locale keys added to `en/tools.en.yml`, `es/tools.es.yml` (TODO stub),
      `fr/tools.fr.yml` (TODO stub)
- [ ] Input types are semantically correct (no `type="text"` for enumerated
      values)
- [ ] CTA buttons use `render "partials/shared/button"` — not raw `link_to` with
      inline classes
- [ ] CTA section uses `render "partials/tools/shared/cta"` (or a deliberate
      custom design with justification)
- [ ] Demo visible in both light mode and dark mode with sufficient contrast
- [ ] Routes added to `config/routes.rb` outside the locale scope block
- [ ] `ApiProxyService.call` used — no duplicate HTTP logic
- [ ] `layout false` on `ToolDemosController` actions
