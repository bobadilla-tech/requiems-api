import { Controller } from "@hotwired/stimulus"

// Connects to data-controller="confirm-delete"
export default class extends Controller {
  static targets = ["modal", "emailInput"]

  declare readonly modalTarget: HTMLElement
  declare readonly emailInputTarget: HTMLInputElement
  declare readonly hasModalTarget: boolean
  declare readonly hasEmailInputTarget: boolean

  private boundHandleEscape!: (event: KeyboardEvent) => void

  open(event: Event) {
    event.preventDefault()

    if (this.hasModalTarget) {
      this.modalTarget.classList.remove("hidden")
      document.body.style.overflow = "hidden"

      // Focus on email input
      if (this.hasEmailInputTarget) {
        setTimeout(() => {
          this.emailInputTarget.focus()
        }, 100)
      }
    }
  }

  close(event?: Event) {
    if (event) {
      event.preventDefault()
    }

    if (this.hasModalTarget) {
      this.modalTarget.classList.add("hidden")
      document.body.style.overflow = ""

      // Clear email input
      if (this.hasEmailInputTarget) {
        this.emailInputTarget.value = ""
      }
    }
  }

  // Close modal when clicking outside
  clickOutside(event: Event) {
    if (event.target === this.modalTarget) {
      this.close()
    }
  }

  // Close modal on Escape key
  handleEscape(event: KeyboardEvent) {
    if (event.key === "Escape") {
      this.close()
    }
  }

  connect() {
    // Add event listener for Escape key
    this.boundHandleEscape = this.handleEscape.bind(this)
    document.addEventListener("keydown", this.boundHandleEscape)

    // Add event listener for clicking outside
    if (this.hasModalTarget) {
      this.modalTarget.addEventListener("click", this.clickOutside.bind(this))
    }
  }

  disconnect() {
    // Remove event listeners
    document.removeEventListener("keydown", this.boundHandleEscape)
  }
}
