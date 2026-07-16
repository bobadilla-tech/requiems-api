import { Controller } from "@hotwired/stimulus";

// Handles client-side validation and loading state for the number base
// conversion demo. API call and result rendering are handled server-side
// via Turbo Frames.
export default class extends Controller {
  static targets = ["from", "to", "value", "button", "errorMessage", "spinner"];
  static values = { errorFillAllFields: String };

  onSubmitStart(event) {
    this._clearError();

    const from = this.fromTarget.value;
    const to = this.toTarget.value;
    const value = this.valueTarget.value.trim();

    if (!from || !to || !value) {
      event.detail.formSubmission.stop();
      this._showError(this.errorFillAllFieldsValue);
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
