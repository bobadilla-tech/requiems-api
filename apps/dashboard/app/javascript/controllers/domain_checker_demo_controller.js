import { Controller } from "@hotwired/stimulus";

// Simulated demo controller for the Domain Checker landing page hero.
// No real API call is made — response data is mocked to illustrate the shape.
// Error messages are passed via data-values from ERB (i18n-ready).
export default class extends Controller {
  static targets = ["input", "button", "errorMessage", "result"];
  static values = { errorEmpty: String };

  check() {
    this._clearError();
    this.resultTarget.classList.add("hidden");

    const domain = this.inputTarget.value.trim();

    if (!domain) {
      this._showError(this.errorEmptyValue);
      return;
    }

    const mock = {
      domain: domain,
      status: "active",
      dns_records: true,
      ssl_valid: true,
      mx_valid: true,
      reputation: "clean",
    };

    this.resultTarget.querySelector(
      "[data-domain-checker-demo-domain]",
    ).textContent = mock.domain;

    this.resultTarget.querySelector(
      "[data-domain-checker-demo-status]",
    ).textContent = mock.status;

    this.resultTarget.querySelector(
      "[data-domain-checker-demo-dns]",
    ).textContent = mock.dns_records ? "✓ Valid" : "✗ Invalid";

    this.resultTarget.querySelector(
      "[data-domain-checker-demo-ssl]",
    ).textContent = mock.ssl_valid ? "✓ Valid" : "✗ Invalid";

    this.resultTarget.querySelector(
      "[data-domain-checker-demo-mx]",
    ).textContent = mock.mx_valid ? "✓ Valid" : "✗ Invalid";

    this.resultTarget.querySelector(
      "[data-domain-checker-demo-reputation]",
    ).textContent = mock.reputation;

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
