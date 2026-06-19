import { Controller } from "@hotwired/stimulus";

// Live demo controller for the Random Quotes tool page.
// Calls GET /v1/entertainment/quotes/random via the dashboard proxy
// and renders the result. No user input is required.
// Error messages are passed via data-values from ERB (i18n-ready).
export default class extends Controller {
  static targets = ["button", "result", "errorMessage", "spinner"];
  static values = {
    errorRateLimit: String,
    errorGeneric: String,
    errorUnexpected: String,
    errorNetwork: String,
    errorSecurityToken: String,
  };

  async fetch() {
    // Always reset UI state before any early return or async work.
    this.resultTarget.classList.add("hidden");
    this._clearError();

    this.buttonTarget.disabled = true;
    this.spinnerTarget.classList.remove("hidden");

    try {
      const csrfToken = document.querySelector('meta[name="csrf-token"]')
        ?.content;

      if (!csrfToken) {
        this._showError(this.errorSecurityTokenValue);
        return;
      }

      const response = await window.fetch("/api/proxy", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "X-CSRF-Token": csrfToken,
        },
        body: JSON.stringify({
          endpoint: "/v1/entertainment/quotes/random",
          method: "GET",
          params: { _t: Date.now() },
        }),
      });

      if (response.status === 429) {
        this._showError(this.errorRateLimitValue);
        return;
      }

      if (!response.ok) {
        this._showError(this.errorGenericValue);
        return;
      }

      const json = await response.json();
      const quote = json?.data?.data;

      if (!quote?.text || !quote?.author) {
        this._showError(this.errorUnexpectedValue);
        return;
      }

      this._renderQuote(quote);
    } catch (_err) {
      this._showError(this.errorNetworkValue);
    } finally {
      this.buttonTarget.disabled = false;
      this.spinnerTarget.classList.add("hidden");
    }
  }

  _renderQuote(quote) {
    this.resultTarget.querySelector(
      "[data-quotes-demo-text]",
    ).textContent = `“${quote.text}”`;
    this.resultTarget.querySelector(
      "[data-quotes-demo-author]",
    ).textContent = `— ${quote.author}`;

    this.resultTarget.classList.remove("hidden");
  }

  _showError(message) {
    this.errorMessageTarget.textContent = message;
    this.errorMessageTarget.classList.remove("hidden");
  }

  _clearError() {
    this.errorMessageTarget.textContent = "";
    this.errorMessageTarget.classList.add("hidden");
  }
}
