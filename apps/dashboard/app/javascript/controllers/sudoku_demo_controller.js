import { Controller } from "@hotwired/stimulus";

// Handles loading state for the Sudoku demo.
// Difficulty is a required <select> with a default option, so empty-field
// validation is unnecessary. API call and result rendering are server-side.
export default class extends Controller {
  static targets = ["button", "errorMessage", "spinner"];

  onSubmitStart() {
    this._clearError();
    this.buttonTarget.disabled = true;
    this.spinnerTarget.classList.remove("hidden");
  }

  onSubmitEnd() {
    this.buttonTarget.disabled = false;
    this.spinnerTarget.classList.add("hidden");
  }

  _clearError() {
    if (!this.hasErrorMessageTarget) return;
    this.errorMessageTarget.classList.add("hidden");
    this.errorMessageTarget.textContent = "";
  }
}
