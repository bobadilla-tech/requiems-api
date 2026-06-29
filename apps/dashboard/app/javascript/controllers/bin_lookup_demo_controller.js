import { Controller } from "@hotwired/stimulus";

// Handles client-side validation and loading state for the BIN Lookup demo.
// API call and result rendering are handled server-side via Turbo Frames.
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
