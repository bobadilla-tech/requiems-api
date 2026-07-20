import { Controller } from "@hotwired/stimulus";
export default class extends Controller {
  static targets = [
    "inputPrincipal",
    "inputRate",
    "inputYears",
    "button",
    "errorMessage",
    "spinner",
  ];
  static values = { errorEmpty: String, errorInvalid: String };

  onSubmitStart(event) {
    this._clearError();

    const principal = parseFloat(this.inputPrincipalTarget.value);
    const rate = parseFloat(this.inputRateTarget.value);
    const years = parseInt(this.inputYearsTarget.value, 10);

    if (
      !this.inputPrincipalTarget.value.trim() ||
      !this.inputRateTarget.value.trim() ||
      !this.inputYearsTarget.value.trim()
    ) {
      event.detail.formSubmission.stop();
      this._showError(this.errorEmptyValue);
      return;
    }

    if (
      isNaN(principal) ||
      principal <= 0 ||
      isNaN(rate) ||
      rate <= 0 ||
      isNaN(years) ||
      years < 1 ||
      years > 50
    ) {
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
