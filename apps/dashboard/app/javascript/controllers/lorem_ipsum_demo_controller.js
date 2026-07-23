import { Controller } from "@hotwired/stimulus";

// Handles client-side validation and loading state for the Lorem Ipsum
// Generator demo. API call and result rendering are handled server-side
// via Turbo Frames.
export default class extends Controller {
  static targets = ["inputParagraphs", "inputSentences", "button", "errorMessage", "spinner"];
  static values = { errorInvalid: String };

  onSubmitStart(event) {
    this._clearError();

    const paragraphs = this.inputParagraphsTarget.value.trim();
    const sentences = this.inputSentencesTarget.value.trim();

    const paragraphsOk = paragraphs === "" || this._inRange(paragraphs);
    const sentencesOk = sentences === "" || this._inRange(sentences);

    if (!paragraphsOk || !sentencesOk) {
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

  _inRange(value) {
    const n = parseInt(value, 10);
    return !isNaN(n) && n >= 1 && n <= 20;
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
