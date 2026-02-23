import { Controller } from "@hotwired/stimulus"

// Copy to clipboard controller
// Usage:
//   <button data-controller="copy" data-copy-text-value="text to copy" data-action="click->copy#copy">
//     Copy
//   </button>
export default class extends Controller {
  static values = {
    text: String,
    targetValue: String,
    successMessage: { type: String, default: "Copied!" },
    errorMessage: { type: String, default: "Failed to copy" }
  }

  static targets = ["button", "feedback"]

  declare readonly textValue: string
  declare readonly targetValueValue: string
  declare readonly successMessageValue: string
  declare readonly errorMessageValue: string

  async copy(event: Event) {
    event.preventDefault()

    try {
      await navigator.clipboard.writeText(this.textValue)
      this.showSuccess()
    } catch (err) {
      console.error("Failed to copy:", err)
      this.showError()
    }
  }

  async copyFromTarget(event: Event) {
    event.preventDefault()

    try {
      const targetName = this.targetValueValue || (this.element as HTMLElement).dataset.copyTargetValue
      const targetElement = document.querySelector(`[data-api-playground-target="${targetName}"]`)

      if (targetElement) {
        const text = targetElement.textContent ?? ""
        await navigator.clipboard.writeText(text)
        this.showSuccess()
      }
    } catch (err) {
      console.error("Failed to copy from target:", err)
      this.showError()
    }
  }

  showSuccess() {
    const originalText = this.element.textContent

    // Show feedback
    this.element.textContent = this.successMessageValue
    this.element.classList.add("text-green-600")

    // Reset after 2 seconds
    setTimeout(() => {
      this.element.textContent = originalText
      this.element.classList.remove("text-green-600")
    }, 2000)
  }

  showError() {
    const originalText = this.element.textContent

    // Show error
    this.element.textContent = this.errorMessageValue
    this.element.classList.add("text-red-600")

    // Reset after 2 seconds
    setTimeout(() => {
      this.element.textContent = originalText
      this.element.classList.remove("text-red-600")
    }, 2000)
  }
}
