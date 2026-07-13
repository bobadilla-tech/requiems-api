import { Controller } from "@hotwired/stimulus";

// Handles client-side validation and loading state for the Timezone demo.
// API call and result rendering are handled server-side via Turbo Frames.
export default class extends Controller {
  static targets = ["input", "lat", "lon", "button", "errorMessage", "spinner"];
  static values = { errorEmpty: String };

  onSubmitStart(event) {
    this._clearError();

    const hasCity = this.inputTarget.value.trim().length > 0;
    const hasCoords =
      this.latTarget.value.trim().length > 0 &&
      this.lonTarget.value.trim().length > 0;

    if (!hasCity && !hasCoords) {
      event.detail.formSubmission.stop();
      this._showError(this.errorEmptyValue);
      return;
    }

    this.buttonTarget.disabled = true;
    this.spinnerTarget.classList.remove("hidden");
  }

  // City and coordinates are mutually exclusive: typing in one clears and
  // disables the other, so it's always obvious which value will be used.
  onCityInput() {
    const hasCity = this.inputTarget.value.trim().length > 0;
    this.latTarget.disabled = hasCity;
    this.lonTarget.disabled = hasCity;
    if (hasCity) {
      this.latTarget.value = "";
      this.lonTarget.value = "";
    }
  }

  onCoordsInput() {
    const hasCoords =
      this.latTarget.value.trim().length > 0 ||
      this.lonTarget.value.trim().length > 0;
    this.inputTarget.disabled = hasCoords;
    if (hasCoords) {
      this.inputTarget.value = "";
    }
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
