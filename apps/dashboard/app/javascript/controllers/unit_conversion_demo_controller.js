import { Controller } from "@hotwired/stimulus";

// Handles the live unit conversion demo in the hero section.
// Calls /api/proxy → GET /v1/technology/convert and renders the result inline.
export default class extends Controller {
  static targets = [
    "from",
    "to",
    "value",
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

    const from = this.fromTarget.value.trim();
    const to = this.toTarget.value.trim();
    const value = this.valueTarget.value.trim();

    if (!from || !to || !value) {
      this._showError("Fill in all fields.");
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
          endpoint: "/v1/technology/convert",
          method: "GET",
          params: {
            from,
            to,
            value,
          },
        }),
      });

      if (response.status === 429) {
        this._showError(
          "Too many requests. Wait a moment and try again.",
        );
        return;
      }


      const json = await response.json();

      if (!response.ok) {
        this._renderError(
          json.data?.message ||
          json.message ||
          json.error ||
          `Request failed (${response.status})`
        );
        return;
      }

      this._renderResult(json.data?.data ?? json.data);
    } catch (_err) {
      this._showError("Could not reach the API. Check your connection.");
    } finally {
      this._setLoading(false);
    }
  }

  _renderResult(data) {
    if (!data) {
      this._showError("No data returned.");
      return;
    }
    
    const rows = [
      { label: "From", value: data.from },
      { label: "To", value: data.to },
      { label: "Input", value: data.input },
      { label: "Result", value: data.result },
      { label: "Formula", value: data.formula },
    ];

    const items = rows.map((row) => `
      <div class="flex items-center justify-between py-2 border-b border-gray-100 dark:border-gray-700 last:border-0">
        <span class="text-sm text-gray-500 dark:text-gray-400">
          ${row.label}
        </span>
        <span class="text-sm font-medium text-gray-900 dark:text-gray-100">
          ${this._escapeHtml(String(row.value ?? "—"))}
        </span>
      </div>
    `).join("");

    this.resultTarget.innerHTML = `
      <div class="rounded-xl border border-gray-200 dark:border-gray-700 bg-white dark:bg-gray-800 overflow-hidden">
        <div class="px-4 py-3 bg-gray-50 dark:bg-gray-900">
          <span class="text-sm font-semibold text-gray-900 dark:text-white">
            Conversion Result
          </span>
        </div>

        <div class="px-4 py-1">
          ${items}
        </div>
      </div>
    `;

    this.resultTarget.classList.remove("hidden");
  }
  _renderError(message) {
    this.resultTarget.innerHTML = `
    <div class="rounded-xl border border-gray-200 dark:border-gray-700 bg-white dark:bg-gray-800 overflow-hidden">
        <div class="px-4 py-3 bg-gray-50 dark:bg-gray-900">
          <span class="text-sm font-semibold text-gray-900 dark:text-white">
            Conversion Failed
          </span>
        </div>

      <div class="px-4 py-4">
        <p class="text-sm text-red-600 dark:text-red-400">
          ${this._escapeHtml(message)}
        </p>
      </div>
    </div>
  `;

    this.resultTarget.classList.remove("hidden");
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