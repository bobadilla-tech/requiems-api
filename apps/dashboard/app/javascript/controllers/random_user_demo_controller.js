import { Controller } from "@hotwired/stimulus";

// Live demo controller for the Random User tool page.
// Calls GET /v1/technology/random-user via a Turbo Frame form submission
// and renders the result. No user input is required — just click and fetch.
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
    this.errorMessageTarget.classList.add("hidden");
    this.errorMessageTarget.textContent = "";
  }
}
