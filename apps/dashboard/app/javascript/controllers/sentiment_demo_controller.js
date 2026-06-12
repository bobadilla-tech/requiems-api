import { Controller } from "@hotwired/stimulus";

// Handles the live sentiment analysis demo in the hero section.
// Calls /api/proxy → POST /v1/text/sentiment and renders the result inline.
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

  async analyze(event) {
    event.preventDefault();
    if (this._pending) return;

    const text = this.inputTarget.value.trim();
    this.resultTarget.classList.add("hidden");

    if (!text) {
      this._showError("Enter some text to analyze.");
      return;
    }

    this._setLoading(true);
    this._clearError();

    try {
      const response = await fetch("/api/proxy", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "X-CSRF-Token": this._csrfToken(),
        },
        body: JSON.stringify({
          endpoint: "/v1/text/sentiment",
          method: "POST",
          params: { text },
        }),
      });

      if (response.status === 429) {
        this._showError("Too many requests. Wait a moment and try again.");
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

    const VALID_SENTIMENTS = ["positive", "negative", "neutral"];
    const sentiment = VALID_SENTIMENTS.includes(data.sentiment)
      ? data.sentiment
      : "neutral";
    const score = Math.round((data.score ?? 0) * 100);
    const breakdown = data.breakdown ?? {};

    const colors = this._sentimentColors(sentiment);
    const label = this._capitalize(sentiment);

    const breakdownRows = ["positive", "negative", "neutral"]
      .map((key) => {
        const pct = Math.round((breakdown[key] ?? 0) * 100);
        const barColors = this._sentimentColors(key);
        return `
        <div class="flex items-center gap-3 py-1">
          <span class="w-16 text-xs text-gray-500 dark:text-gray-400 capitalize">${key}</span>
          <div class="flex-1 bg-gray-100 dark:bg-gray-700 rounded-full h-1.5 overflow-hidden">
            <div class="h-1.5 rounded-full ${barColors.bar}" style="width:${pct}%"></div>
          </div>
          <span class="w-8 text-xs text-right font-medium text-gray-700 dark:text-gray-300">${pct}%</span>
        </div>`;
      })
      .join("");

    this.resultTarget.innerHTML = `
      <div class="rounded-xl border border-gray-200 dark:border-gray-700 bg-white dark:bg-gray-800 overflow-hidden">
        <div class="px-5 py-4 flex items-center gap-3 bg-gray-50 dark:bg-gray-900">
          <span class="inline-flex items-center px-3 py-1 rounded-full text-sm font-semibold ${colors.badge}">
            ${label}
          </span>
          <span class="text-sm text-gray-500 dark:text-gray-400">${score}% confident</span>
        </div>
        <div class="px-5 py-4">
          <p class="text-xs font-medium text-gray-400 dark:text-gray-500 uppercase tracking-wide mb-3">Breakdown</p>
          ${breakdownRows}
        </div>
      </div>`;

    this.resultTarget.classList.remove("hidden");
  }

  _sentimentColors(sentiment) {
    const map = {
      positive: {
        badge:
          "bg-emerald-100 text-emerald-700 dark:bg-emerald-900 dark:text-emerald-300",
        bar: "bg-emerald-500",
      },
      negative: {
        badge: "bg-red-100 text-red-700 dark:bg-red-900 dark:text-red-300",
        bar: "bg-red-500",
      },
      neutral: {
        badge: "bg-gray-100 text-gray-600 dark:bg-gray-700 dark:text-gray-300",
        bar: "bg-gray-400",
      },
    };
    return map[sentiment] ?? map.neutral;
  }

  _capitalize(str) {
    return String(str).charAt(0).toUpperCase() + String(str).slice(1);
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
}
