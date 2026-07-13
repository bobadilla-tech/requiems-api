import { Controller } from "@hotwired/stimulus";

// Live demo controller for the Random User tool page.
// Calls GET /v1/technology/random-user via a Turbo Frame form submission
// and renders the result. No user input is required — just click and fetch.
export default class extends Controller {
  static targets = ["button", "errorMessage", "spinner"];
  static values = {
    errorNetwork: String,
  };

  onSubmitStart() {
    this._clearError();
    this.buttonTarget.disabled = true;
    this.spinnerTarget.classList.remove("hidden");
  }

  onSubmitEnd(event) {
    this.buttonTarget.disabled = false;
    this.spinnerTarget.classList.add("hidden");

    if (!event.detail.success) {
      this.errorMessageTarget.textContent = this.errorNetworkValue;
      this.errorMessageTarget.classList.remove("hidden");
    }
  }

  _clearError() {
    this.errorMessageTarget.classList.add("hidden");
    this.errorMessageTarget.textContent = "";
  }
}
