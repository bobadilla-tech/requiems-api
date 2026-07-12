import { Controller } from "@hotwired/stimulus";

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
