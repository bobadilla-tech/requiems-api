import { Controller } from "@hotwired/stimulus";

// Handles the live email validation demo in the hero section.
// Calls /api/proxy → POST /v1/validation/email and renders the result inline.
export default class extends Controller {
  static targets = [
    "input",
    "button",
    "result",
    "errorMessage",
    "spinner",
  ];

  connect() {
    this._pending = false;
  }

  async verify(event) {
    event.preventDefault();
    if (this._pending) return;

    this._clearError();
    this.resultTarget.classList.add("hidden");
    const email = this.inputTarget.value.trim();

    if (!email) {
      this._showError("Enter an email address to verify.");
      return;
    }

    this._setLoading(true);

    try {
      const response = await fetch("/api/proxy", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "X-CSRF-Token": this._csrfToken(),
        },
        body: JSON.stringify({
          endpoint: "/v1/validation/email",
          method: "POST",
          params: { email },
        }),
      });

      if (response.status === 429) {
        this._showError(
          "Too many requests. Wait a moment and try again.",
        );
        return;
      }

      const json = await response.json();

      if (!response.ok || json.error) {
        this._showError("Something went wrong. Try again.");
        return;
      }

      this._renderResult(json.data?.data ?? json.data);
    } catch (_err) {
      this._showError("Could not reach the API. Check your connection.");
    } finally {
      this._setLoading(false);
    }
  }

  // ── private ────────────────────────────────────────────────────────────────

  _renderResult(data) {
    if (!data) {
      this._showError("No data returned.");
      return;
    }

    const rows = [
      { label: "Syntax", value: data.syntax_valid, type: "bool" },
      { label: "MX Record", value: data.mx_valid, type: "bool" },
      { label: "Disposable domain", value: data.disposable, type: "bool_inv" },
      { label: "Valid", value: data.valid, type: "bool" },
      { label: "Normalized", value: data.normalized, type: "text" },
      { label: "Domain", value: data.domain, type: "text" },
    ];

    if (data.suggestion) {
      rows.push({
        label: "Suggestion",
        value: data.suggestion,
        type: "suggestion",
      });
    }

    const items = rows
      .map((r) => `
        <div class="flex items-center justify-between py-2 border-b border-gray-100 dark:border-gray-700 last:border-0">
          <span class="text-sm text-gray-500 dark:text-gray-400">${r.label}</span>
          <span class="text-sm font-medium ${
        this._valueClass(r.type, r.value)
      }">${this._valueLabel(r.type, r.value)}</span>
        </div>`)
      .join("");

    this.resultTarget.innerHTML = `
      <div class="rounded-xl border border-gray-200 dark:border-gray-700 bg-white dark:bg-gray-800 divide-y divide-gray-100 dark:divide-gray-700 overflow-hidden">
        <div class="px-4 py-3 bg-gray-50 dark:bg-gray-900 flex items-center gap-2">
          <span class="text-xs font-mono text-gray-500 dark:text-gray-400 truncate">${
      this._escapeHtml(data.email ?? this.inputTarget.value)
    }</span>
          <span class="ml-auto shrink-0 inline-flex items-center px-2 py-0.5 rounded-full text-xs font-medium ${
      data.valid
        ? "bg-emerald-100 text-emerald-700 dark:bg-emerald-900 dark:text-emerald-300"
        : "bg-red-100 text-red-700 dark:bg-red-900 dark:text-red-300"
    }">
            ${data.valid ? "Valid" : "Invalid"}
          </span>
        </div>
        <div class="px-4 py-1">${items}</div>
      </div>`;

    this.resultTarget.classList.remove("hidden");
  }

  _valueClass(type, value) {
    if (type === "bool") {
      return value
        ? "text-emerald-600 dark:text-emerald-400"
        : "text-red-500 dark:text-red-400";
    }
    if (type === "bool_inv") {
      return value
        ? "text-red-500 dark:text-red-400"
        : "text-emerald-600 dark:text-emerald-400";
    }
    if (type === "suggestion") return "text-amber-600 dark:text-amber-400";
    return "text-gray-800 dark:text-gray-200";
  }

  _valueLabel(type, value) {
    if (type === "bool") return value ? "✓ Yes" : "✗ No";
    if (type === "bool_inv") return value ? "✗ Disposable" : "✓ Not disposable";
    if (type === "suggestion") {
      return `Did you mean ${this._escapeHtml(value)}?`;
    }
    return value ? this._escapeHtml(String(value)) : "—";
  }

  _setLoading(state) {
    this._pending = state;
    this.buttonTarget.disabled = state;
    this.spinnerTarget.classList.toggle("hidden", !state);
  }

  _showError(msg) {
    this.errorMessageTarget.textContent = msg;
    this.errorMessageTarget.classList.remove("hidden");
  }

  _clearError() {
    this.errorMessageTarget.classList.add("hidden");
    this.errorMessageTarget.textContent = "";
  }

  _csrfToken() {
    return document.querySelector('meta[name="csrf-token"]')?.content ?? "";
  }

  _escapeHtml(str) {
    return String(str)
      .replace(/&/g, "&amp;")
      .replace(/</g, "&lt;")
      .replace(/>/g, "&gt;")
      .replace(/"/g, "&quot;");
  }
}
