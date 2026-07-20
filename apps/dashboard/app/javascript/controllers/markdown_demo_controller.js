import { Controller } from "@hotwired/stimulus";
// Handles client-side validation and loading state for the Markdown to HTML demo.
// API call and result rendering are handled server-side via Turbo Frames.
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
  _showError(msg) {
    this.errorMessageTarget.textContent = msg;
    this.errorMessageTarget.classList.remove("hidden");
  }
  _clearError() {
    this.errorMessageTarget.classList.add("hidden");
    this.errorMessageTarget.textContent = "";
  }
}
