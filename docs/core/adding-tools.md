# Adding a Tool Page

> Full end-to-end guide: planning, sections, demo form, i18n, copy, and
> validation

Tool pages are marketing landing pages with a live interactive demo. Each page
shows what an API does, why someone would use it, and lets visitors try it
without an account.

---

## 1. Workflow Overview

Building a tool page follows these stages:

| Stage | What you do |
|---|---|
| **1. Scaffold** | Create show view, register route + controller, create all 7 section partials with placeholder content |
| **2. Demo form** | Implement the live demo in the hero using Turbo Frame + Stimulus |
| **3. Section content** | Fill in each section partial with real copy, use cases, and code blocks |
| **4. i18n** | Extract all strings with `rori18n`, add EN keys and ES/FR stubs |
| **5. Copy review** | Audit voice, CTAs, superlatives, technical terms across all sections |
| **6. Validation** | Test responsive layout (375px/768px/1280px), keyboard navigation, contrast |

---

## 2. Goals

1. **No innerHTML in JavaScript** — Result rendering lives in Rails ERB, not JS
   template strings.
2. **All strings via `t()`** — Every user-facing string must come from a locale
   file. Hard-coding English is not acceptable.
3. **Correct input types** — Use `<select>` for enumerated options, `<textarea>`
   for long text, `type="tel"` for phone, `type="email"` for email,
   `type="number"` for numeric input.
4. **Shared components** — Use `render "partials/shared/button"`, not custom
   inline `link_to` classes.
5. **No escaping your own backend** — Do not call `escapeHtml()` on data
   returned by our own API. The backend already sanitizes output.

---

## 3. File layout

```
apps/dashboard/app/
├── controllers/
│   ├── tool_demos_controller.rb          # one action per tool (already exists,
│   │                                     #   append a new action)
│   └── tools_controller.rb               # adds tool to SUPPORTED_TOOLS constant
├── javascript/controllers/
│   └── {tool_name}_demo_controller.js    # client-side validation + loading state only
├── views/
│   ├── tool_demos/
│   │   ├── demo_error.html.erb           # shared error partial (already exists)
│   │   └── {tool_name}.html.erb          # Turbo Frame result view
│   └── tools/
│       ├── index.html.erb                # already exists
│       └── {tool_name}/
│           └── show.html.erb             # composes all sections
├── views/partials/
│   ├── shared/
│   │   ├── _button.html.erb              # variant: brand, outline-white, etc.
│   │   ├── _card.html.erb
│   │   ├── _badge.html.erb
│   │   ├── _submit_button.html.erb       # demo form submit button with spinner
│   │   └── _logo_marquee.html.erb        # "Trusted by teams at" strip
│   └── tools/
│       ├── shared/
│       │   └── _cta.html.erb             # shared CTA section (accepts: heading, body, docs_url)
│       └── {tool_name}/
│           ├── _hero.html.erb            # headline + live demo form + Turbo Frame slot
│           ├── _what_it_does.html.erb    # accuracy + simplicity cards with code blocks
│           ├── _use_cases.html.erb       # 6 scenario cards
│           ├── _api_combinations.html.erb# 3 cross-API pairing cards
│           ├── _faq.html.erb             # 5 Q&A accordion items
│           └── _cta.html.erb             # wraps partials/tools/shared/_cta
config/locales/
├── en/
│   └── tools.en.yml                      # all tool strings (EN)
├── es/
│   └── tools.es.yml                      # Spanish translations (stubs on first addition)
└── fr/
    └── tools.fr.yml                      # French translations (stubs on first addition)
```

---

## 4. Stage 1 — Scaffold the Show View

### 4a. Register in ToolsController

Add the tool's dash-slug to `SUPPORTED_TOOLS` and add metadata in
`apps/dashboard/app/controllers/tools_controller.rb`:

```ruby
SUPPORTED_TOOLS = %w[
  email-validator sentiment-analysis ... phone-validator
  bin-lookup inflation qr-code profanity-filter trivia YOUR_TOOL
].freeze

TOOLS_METADATA = {
  # ...
  "your-tool" => {
    name: "Your Tool Name",
    description: "One-line description for the tools index grid.",
    icon_classes: "bg-indigo-50 dark:bg-indigo-900/20 text-indigo-600 dark:text-indigo-400"
  }
}.freeze
```

The route is already handled by `resources :tools, only: [ :index, :show ]`
inside the locale scope. The show action renders `tools/{id}/show` where `id`
uses underscores (e.g. `phone_validator` for `phone-validator`).

### 4b. Create the show view

```erb
<%# apps/dashboard/app/views/tools/{tool_name}/show.html.erb %>

<% content_for :title, t("tools.{tool_name}.show.{seo_title_key}") %>
<% content_for :description, "One-sentence meta description for search results." %>

<%= render "partials/tools/{tool_name}/hero" %>
<%= render "partials/shared/logo_marquee",
      logos: [
        { file: "companies/bobadilla-tech.png",   alt: "Bobadilla Tech"   },
        { file: "companies/compile-strength.png", alt: "Compile Strength" },
        { file: "companies/flyver.jpeg",           alt: "Flyver"           },
        { file: "companies/korealexa.png",         alt: "Korealexa"        },
        { file: "companies/shareweave.jpeg",       alt: "Shareweave"       },
      ],
      label: t("tools.{tool_name}.logo_marquee.trusted_by"),
      speed: "34s" %>
<%= render "partials/tools/{tool_name}/what_it_does" %>
<%= render "partials/tools/{tool_name}/use_cases" %>
<%= render "partials/tools/{tool_name}/api_combinations" %>
<%= render "partials/tools/{tool_name}/faq" %>
<%= render "partials/tools/{tool_name}/cta" %>
```

### 4c. Section partials — standard 7 sections

Each tool page must have these sections (the hero partial goes in
`partials/tools/{tool_name}/`, not `partials/tools/shared/`):

#### _hero — Live demo form

The hero is the most elaborate section. It contains the headline, subtext, CTA
buttons, and the interactive demo form. See section 5 for the full pattern.

Key elements:

- API badge label (e.g. "Phone Validation API")
- H1 headline with highlight span
- Subtitle paragraph
- Feature badges (JSON response, uptime, latency, REST API — use shared strings)
- CTA buttons using `render "partials/shared/button"` with `variant: "brand"`
- Demo form with `<turbo-frame>` slot (see section 5)

The hero uses a gradient background `linear-gradient(135deg, ...)` with an
SVG grid pattern overlay. Get the gradient values from the design spec.

#### _what_it_does — Accuracy + Simplicity cards

Two cards side-by-side on desktop, stacked on mobile:

- **Accuracy card** (`bg-white dark:bg-gray-800 rounded-xl shadow p-6`):
  Explanatory paragraph + mock JSON response in a `<pre><code>` block
- **Simplicity card** (`bg-gray-50 dark:bg-gray-900 rounded-xl shadow p-6`):
  Explanatory paragraph + mock HTTP request in a `<pre><code>` block

Use `data-controller="highlight"` on the code block wrapper for syntax
highlighting.

Layout: `grid grid-cols-1 lg:grid-cols-2 gap-8`

Section wrapper: `bg-white dark:bg-gray-950 py-20 border-t border-gray-100 dark:border-gray-800`

#### _use_cases — 6 scenario cards

6 cards in a `grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6`.

Each card:

- Context label: `text-xs font-semibold uppercase tracking-wide text-brand-primary`
- Result-oriented title (outcome first, not feature-first)
- 1–2 sentence description

Card style: `bg-white dark:bg-gray-800 rounded-xl p-6 shadow-sm border border-gray-100 dark:border-gray-700`

Reference the Phone Validator's 6 use cases for the pattern (auth, fraud,
integrations, SMS/marketing, onboarding, support).

#### _api_combinations — 3 cross-API pairs

3 cards in a `grid grid-cols-1 md:grid-cols-3 gap-6` section.

Each card shows a pair name, benefit, and outcome. Visually lighter than use
cases: `bg-gray-50 dark:bg-gray-800 rounded-xl p-5 border border-gray-200 dark:border-gray-700`.

Section has an intro sentence explaining the API composition concept. Choose 3
complementary APIs from `api_catalog.yml`.

#### _faq — 5 Q&A accordion

5 questions using the `faq-accordion` Stimulus controller
(`apps/dashboard/app/javascript/controllers/faq_accordion_controller.js`):

```erb
<div data-controller="faq-accordion">
  <button id="faq-{tool}-btn-<%= i %>"
          class="w-full flex justify-between items-center text-left gap-4 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-brand-primary rounded-lg"
          aria-expanded="false"
          aria-controls="faq-{tool}-panel-<%= i %>"
          data-action="click->faq-accordion#toggle">
    <span class="text-base font-semibold text-gray-900 dark:text-gray-100">
      <%= item[:q] %>
    </span>
    <svg class="h-5 w-5 shrink-0 text-gray-400 transition-transform duration-200" ...>
      ...
    </svg>
  </button>

  <div id="faq-{tool}-panel-<%= i %>"
       role="region"
       aria-labelledby="faq-{tool}-btn-<%= i %>"
       class="hidden prose text-sm text-gray-600 dark:text-gray-400 mt-3">
    <%= item[:a] %>
  </div>
</div>
```

Wrap each Q&A pair in `render "partials/shared/card", padding: true do ...
end`. All answers start collapsed (`hidden` class).

Include a "Still have questions? Contact support" section at the bottom.

#### _cta — Closing call-to-action

This is a thin wrapper that delegates to the shared CTA partial:

```erb
<%= render "partials/tools/shared/cta",
      heading:  t("tools.{tool_name}.cta.heading"),
      body:     t("tools.{tool_name}.cta.body"),
      docs_url: api_path("your-api-slug") %>
```

The shared CTA renders a `bg-gray-900 dark:bg-gray-950 py-20` section with:

- Headline: "Automate this with Requiems."
- Body: "One API key. ... plus 50+ other APIs..."
- Primary button: "Get your free API key" → `new_user_registration_path`
- Secondary button: "Read the docs →" → `docs_url`

---

## 5. Stage 2 — Demo Form Pattern (Turbo Frame)

The demo form submits to a Rails action. The action calls the API and renders a
result partial inside a `<turbo-frame>`. **Never render results via `innerHTML`
in JavaScript.**

### 5a. Hero partial — the form

```erb
<%# partials/tools/{tool_name}/_hero.html.erb %>

<div data-controller="{tool-name}-demo"
     data-{tool-name}-demo-error-empty-value="<%= t('tools.{tool_name}.demo.error_empty') %>">

  <form data-action="turbo:submit-start->{tool-name}-demo#onSubmitStart turbo:submit-end->{tool-name}-demo#onSubmitEnd"
        action="<%= tool_demo_{tool_name}_path %>"
        method="post"
        data-turbo-frame="{tool_name}-demo-result"
        novalidate>
    <%= hidden_field_tag :authenticity_token, form_authenticity_token %>

    <%# form fields with name= attributes — these become params in the controller %>

    <%= render "partials/shared/submit_button",
          text: t("tools.{tool_name}.hero.submit"),
          controller: "{tool-name}-demo" %>

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
- Use `render "partials/shared/submit_button"` for the submit button — it
  handles the spinner and Stimulus target wiring automatically.

### 5b. Route — outside locale scope

```ruby
# config/routes.rb — outside the locale scope block, with other tool demo routes
post "tools/demos/{tool-name}", to: "tool_demos#{action_name}", as: :tool_demo_{tool_name}
```

### 5c. Controller action — append to ToolDemosController

```ruby
# app/controllers/tool_demos_controller.rb
class ToolDemosController < ApplicationController
  layout false

  def {tool_name}
    field = params[:field].to_s.strip

    if field.blank?
      return render_demo_error("{tool_name}", t("tools.{tool_name}.demo.error_empty"))
    end

    result = api_call(endpoint: "/v1/...", method: "GET", params: { field: field })

    return render_demo_error("{tool_name}", t("tools.{tool_name}.demo.error_rate_limit")) if result.status_code == 429
    return render_demo_error("{tool_name}", t("tools.{tool_name}.demo.error_generic")) unless result.status_code == 200

    data = result.data&.dig("data", "data") || result.data&.dig("data")
    return render_demo_error("{tool_name}", t("tools.{tool_name}.demo.error_no_data")) if data.nil?

    render "tool_demos/{tool_name}", locals: { data: data, field: field }
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

### 5d. Result view

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

### 5e. Stimulus controller — thin

```javascript
// app/javascript/controllers/{tool_name}_demo_controller.js
import { Controller } from "@hotwired/stimulus";

export default class extends Controller {
  static targets = ["input", "button", "errorMessage", "spinner"];
  static values = { errorEmpty: String };

  onSubmitStart(event) {
    this._clearError();

    if (!this.inputTarget.value.trim()) {
      event.detail.formSubmission.stop();
      this._showError(this.errorEmptyValue);
      return;
    }

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

## 6. Stage 3 — Section Content Guidelines

### What It Does

- **Accuracy block**: Explain what data the API returns. Include a mock JSON
  response in a dark code block (`bg-gray-900 text-green-400`).
- **Simplicity block**: Show how easy it is to call. Include a mock HTTP request
  showing the `X-API-Key` header.

### Use Cases

Write 6 cards covering these angles (adapt to your tool):

1. **Auth/Signup** — validate at registration
2. **Fraud Prevention** — detect bad actors
3. **Integrations** — clean data in workflows
4. **Marketing/outreach** — campaign quality
5. **Onboarding** — standardize data
6. **Support/ops** — real-time lookups

Title format: result-oriented, active voice, no superlatives.

- Good: "Verify phone numbers at registration"
- Bad: "The most powerful phone validation solution"

### API Combinations

Choose 3 complementary APIs from the existing catalog. Each card:

- Pair name: "Your Tool + Complementary API"
- Benefit: What combining them achieves
- Outcome: The result for the user

Good pairs usually involve: Email Validator, Domain Checker, IP Reputation,
Email Normalizer, BIN Lookup, Sentiment Analysis.

### FAQ

Cover these 5 angles (adapt to your tool):

1. What information does the API return?
2. Does it work with international data?
3. Can it detect edge cases (VOIP, virtual, risk signals)?
4. How accurate is it?
5. Is it free to try?

Answers should be 2–4 sentences, honest, avoid marketing language.

---

## 7. Stage 4 — i18n Requirements

All user-facing strings in tool pages must use `t()`. Use **rori18n** to extract
strings and generate the locale files.

**Install once** ([rori18n repo](https://github.com/bobadilla-tech/rori18n)):

```sh
go install github.com/bobadilla-tech/rori18n@latest
# add ~/go/bin to PATH if not already there
export PATH="$PATH:$(go env GOPATH)/bin"
```

Run all commands from inside `apps/dashboard/`:

#### Step 1 — find hardcoded strings

```sh
git diff --name-only origin/main | \
  rori18n report -r . --changed-files -
```

Lists every hardcoded user-visible string in your changed files. Fix them with
`t()` calls and EN keys before moving on.

#### Step 2 — extract + generate locale skeletons

```sh
git diff --name-only origin/main | \
  rori18n generate -r . --fix --languages es,fr --changed-files -
```

One pass does three things:

1. Finds remaining hardcoded strings in your changed files
2. Writes the EN key to the correct YAML file and injects the `t()` call
3. Creates matching empty entries in `es/tools.es.yml` and `fr/tools.fr.yml`

For `t()` calls you wrote manually, also add the key with an empty value `""` to
the ES and FR files, the maintainer fills them after merge.

#### Step 3 — lint before opening the PR

```sh
rori18n lint -r .
```

Exits 1 with `file:line: missing key "..."` if any `t()` call has no matching
YAML entry. All lint errors must be fixed before opening the PR.

> **Translation** (filling ES/FR empty values with real text) is run by the
> maintainer after merge. Do not write translations manually.

### Key structure

```yaml
en:
  tools:
    { tool_name }:
      seo:
        title: "..."   # used in content_for :title
        description: "..."  # used in content_for :description
      logo_marquee:
        trusted_by: "Trusted by teams at"
      hero:
        badge: "..."
        title: "..."
        title_highlight: "..."
        subtitle: "..."
        cta_primary: "Get started free"
        cta_docs: "View docs"
        try_live: "Try it live"
        label_input: "..."
        placeholder: "..."
        submit: "..."
        no_account: "No account required"
      demo:
        error_empty: "..."
        error_rate_limit: "..."
        error_generic: "..."
        error_no_data: "..."
        result_heading: "..."
        result_failed_heading: "..."
        label_*: "..."
        badge_*: "..."
      what_it_does:
        heading: "..."
        subheading: "..."
        accuracy_heading: "..."
        accuracy_body: "..."
        simplicity_heading: "..."
        simplicity_body: "..."
      use_cases:
        heading: "..."
        subheading: "..."
        { slug }_title: "..."
        { slug }_body: "..."
        { slug }_context: "..."
      api_combinations:
        heading: "..."
        subheading: "..."
        label_benefit: "Benefit:"
        label_outcome: "Outcome:"
        { slug }_name: "..."
        { slug }_benefit: "..."
        { slug }_outcome: "..."
      faq:
        heading: "..."
        support_prompt: "..."
        support_link: "..."
        q1: "..."
        a1: "..."
        q2: "..."
        a2: "..."
        q3: "..."
        a3: "..."
        q4: "..."
        a4: "..."
        q5: "..."
        a5: "..."
      cta:
        heading: "Automate this with Requiems."
        body: "One API key. ..."
        cta_primary: "Get your free API key"
        cta_docs: "Read the docs →"
      show:
        { seo_title_key }: "... | Requiems API"
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
`t()` directly, no JS needed.

---

## 8. Stage 5 — Copy Review Checklist

Before opening the PR, audit every string in your partials:

### Headings use active voice

- Good: "Validate Phone Numbers in Real Time"
- Good: "Check Any Domain in Milliseconds"
- Avoid passive: "Phone Numbers Can Be Validated"

### No superlatives

- Avoid: "best", "most powerful", "cutting-edge", "revolutionary", "ultimate"
- Stick to factual descriptions

### CTAs use action verbs

- Good: "Get your free API key", "View docs", "Validate", "Check", "Normalize"
- Avoid: "Learn More", "Click here"

### Technical terms explained inline

- First use of terms like E.164, VOIP, MX record, ISO country code must have
  inline explanation or parenthetical
- Example: "Submit numbers in E.164 format (e.g. +44 or +52 prefix)"

### SEO meta tags

- `content_for :title`: 50–60 characters, keyword-relevant
- `content_for :description`: 150–160 characters, one sentence

### FAQ answers

- 2–4 sentences, no marketing copy
- If the answer is "no" (e.g. does not check active line), say so directly

### Tone consistency

- Same voice across all sections
- No sudden switches from "we" to "you" to "the API"

---

## 9. Stage 6 — Responsive & Accessibility Validation

### Viewport testing

Test all sections at these widths:

- **375px** (iPhone SE) — no horizontal scroll, single-column grid
- **768px** (iPad) — 2-column grids active, readable text
- **1280px** (desktop) — full layout, no orphaned elements

### Tailwind responsive classes

Verify these are correct in your partials:

- Card grids use `md:grid-cols-2` and `lg:grid-cols-3`
- Button groups use `flex-col sm:flex-row`
- What It Does cards use `lg:grid-cols-2`
- Container: `max-w-7xl mx-auto px-4 sm:px-6 lg:px-8`
- FAQ max width: `max-w-3xl`

### Keyboard navigation

Walk through with Tab/Enter/Space only:

- Hero form: Tab to input → Tab to submit → Enter to submit → result appears
- FAQ: Tab to question button → Enter/Space to toggle answer → answer visible
- All interactive elements have visible focus indicators

### Color contrast (WCAG AA)

Must pass 4.5:1 for normal text:

- White text on `bg-[#1D9E75]` (brand primary)
- White text on `bg-gray-900` (CTA section)
- Body text on light/dark backgrounds

### Lighthouse audit

- Run Lighthouse accessibility audit
- Resolve all critical and serious findings
- Check landmark structure: h1 → h2 → h3 hierarchy

---

## 10. Input types

| Data type                                     | Use                                         |
| --------------------------------------------- | ------------------------------------------- |
| Enumerated set (units, categories, countries) | `<select>` with `<optgroup>` for categories |
| Free-form text                                | `<textarea>` with `rows` set                |
| Email                                         | `<input type="email">`                      |
| Phone number                                  | `<input type="tel">`                        |
| Numbers                                       | `<input type="number">`                     |
| Generic text with no enumeration              | `<input type="text">` — last resort         |

Never use `type="text"` when the valid values are a known set. The user cannot
discover options from a text box.

---

## 11. Shared components

### Button

```erb
<%= render "partials/shared/button",
      text:    "Get API Key",
      variant: "brand",
      href:    new_user_registration_path,
      size:    "lg" %>
```

Available variants: `brand`, `primary`, `secondary`, `success`, `danger`,
`warning`, `outline`, `outline-primary`, `outline-white`, `ghost`, `link`.

Do not write inline `link_to` with custom Tailwind classes for CTA buttons.

### Submit button (demo form)

```erb
<%= render "partials/shared/submit_button",
      text: t("tools.{tool_name}.hero.submit"),
      controller: "{tool-name}-demo" %>
```

Handles spinner visibility and Stimulus target wiring automatically.

### Card

```erb
<%= render "partials/shared/card", padding: true do %>
  ...content...
<% end %>
```

### Logo marquee

```erb
<%= render "partials/shared/logo_marquee",
      logos: [
        { file: "companies/bobadilla-tech.png", alt: "Bobadilla Tech" },
        ...
      ],
      label: "Trusted by teams at",
      speed: "34s" %>
```

### CTA section

```erb
<%# partials/tools/{tool_name}/_cta.html.erb %>
<%= render "partials/tools/shared/cta",
      heading:  "Your call-to-action headline.",
      body:     "Supporting sentence with value proposition.",
      docs_url: api_path("your-api-slug") %>
```

---

## 12. HTML escaping rules

**Do not escape data from our own backend.** The internal API already sanitizes
output. Calling `escapeHtml()` on API response fields causes double-encoding
(users see `&amp;` literally) and signals distrust of our own infrastructure.

**Do escape** genuinely untrusted user input if it is ever inserted into the DOM
directly (outside of a Turbo Frame response). The correct library is
**DOMPurify**, not hand-rolled regex replacements.

In practice: with the Turbo Frame pattern, result data goes through Rails ERB
(`<%= value %>`) which HTML-escapes automatically. No JS escaping is needed.

---

## 13. Anti-patterns

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
| Omitting `content_for :title` / `:description`     | Set both in the show view                      |
| Superlatives in headings ("best", "most powerful") | Use factual, active-voice descriptions         |
| Unresponsive grid (no `md:` / `lg:` prefixes)      | Add responsive breakpoints to all card grids   |
| No keyboard focus indicators on interactive elems  | Use `focus:ring-2 focus:ring-brand-primary`    |
| Skipping ES/FR locale stubs                        | Run `rori18n generate --fix --languages es,fr` |

---

## 14. Pre-merge checklist

- [ ] Tool ID added to `SUPPORTED_TOOLS` and `TOOLS_METADATA` in
      `tools_controller.rb`
- [ ] Show view created at `tools/{tool_name}/show.html.erb` with SEO meta tags
- [ ] All 7 section partials created under `partials/tools/{tool_name}/`
- [ ] Demo route added in `config/routes.rb` outside locale scope
- [ ] Controller action added to `ToolDemosController`
- [ ] Demo result view created at `tool_demos/{tool_name}.html.erb`
- [ ] Stimulus controller created in `app/javascript/controllers/`
- [ ] `grep -r "innerHTML" app/javascript/controllers/{tool_name}*` returns
      nothing
- [ ] `grep -r "_escapeHtml\|escapeHtml" app/javascript/controllers/{tool_name}*`
      returns nothing
- [ ] All user-facing strings go through `t()` — no hardcoded English in JS or
      ERB (`rori18n report --changed-files -` returns nothing)
- [ ] `rori18n generate --fix --languages es,fr --changed-files -` run — EN keys
      written, ES/FR skeleton entries created
- [ ] `rori18n lint` exits 0
- [ ] Input types are semantically correct (no `type="text"` for enumerated
      values)
- [ ] CTA buttons use `render "partials/shared/button"` — not raw `link_to` with
      inline classes
- [ ] CTA section uses `render "partials/tools/shared/cta"` (or a deliberate
      custom design with justification)
- [ ] Demo visible in both light mode and dark mode with sufficient contrast
- [ ] `ApiProxyService.call` used — no duplicate HTTP logic
- [ ] `layout false` on `ToolDemosController` actions
- [ ] Copy review: no superlatives, active voice headings, action-verb CTAs
- [ ] Technical terms explained inline on first use
- [ ] Page renders without horizontal scroll at 375px, 768px, 1280px
- [ ] Card grids collapse to correct column counts at each breakpoint
- [ ] Hero form is fully operable by keyboard: input, button, result all
      reachable via Tab/Enter
- [ ] FAQ accordion is keyboard-operable: Tab to question, Enter/Space to toggle
- [ ] All interactive elements have visible focus indicators
- [ ] Color contrast for primary button text and CTA section text passes
      WCAG AA (4.5:1)
- [ ] No critical or serious findings in Lighthouse accessibility audit
