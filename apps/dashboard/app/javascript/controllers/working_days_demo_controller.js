import { Controller } from "@hotwired/stimulus";

// Handles client-side validation and loading state for the Working Days
// Calculator demo. API call and result rendering are handled server-side
// via Turbo Frames.
export default class extends Controller {
  static targets = ["inputFrom", "inputTo", "inputCountry", "button", "errorMessage", "spinner"];
  static values = { errorEmpty: String, errorInvalid: String };

  onSubmitStart(event) {
    this._clearError();

    const from = this.inputFromTarget.value;
    const to = this.inputToTarget.value;

    if (!from || !to) {
      event.detail.formSubmission.stop();
      this._showError(this.errorEmptyValue);
      return;
    }

    if (new Date(to) < new Date(from)) {
      event.detail.formSubmission.stop();
      this._showError(this.errorInvalidValue);
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
