import { Controller } from "@hotwired/stimulus";

// Thin controller for the Email Normalizer demo.
// Client-side responsibility: input validation and loading state only.
// The form POSTs to ToolDemosController via Turbo Frame — no fetch here.
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

  // ── private ────────────────────────────────────────────────────────────────

  _showError(msg) {
    this.errorMessageTarget.textContent = msg;
    this.errorMessageTarget.classList.remove("hidden");
  }

  _clearError() {
    this.errorMessageTarget.classList.add("hidden");
    this.errorMessageTarget.textContent = "";
  }
}
