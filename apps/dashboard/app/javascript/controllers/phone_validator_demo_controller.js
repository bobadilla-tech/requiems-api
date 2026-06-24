import { Controller } from "@hotwired/stimulus";

export default class extends Controller {
  static targets = ["input", "button", "errorMessage", "spinner"];
  static values = { errorEmpty: String };

  beforeSubmit(event) {
    this._clearError();
    if (!this.inputTarget.value.trim()) {
      event.preventDefault();
      this._showError(this.errorEmptyValue || "Enter a phone number to validate.");
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
