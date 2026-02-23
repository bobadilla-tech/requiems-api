import { Controller } from "@hotwired/stimulus"

// Modal dialog controller
// Usage: data-controller="modal"
export default class extends Controller {
  private boundCloseWithKeyboard!: (event: KeyboardEvent) => void

  open() {
    this.element.classList.remove("hidden")
    document.body.classList.add("overflow-hidden")
  }

  close() {
    this.element.classList.add("hidden")
    document.body.classList.remove("overflow-hidden")
  }

  // Close on Escape key
  closeWithKeyboard(event: KeyboardEvent) {
    if (event.key === "Escape") {
      this.close()
    }
  }

  connect() {
    // Listen for Escape key
    this.boundCloseWithKeyboard = this.closeWithKeyboard.bind(this)
    document.addEventListener("keydown", this.boundCloseWithKeyboard)
  }

  disconnect() {
    document.removeEventListener("keydown", this.boundCloseWithKeyboard)
    document.body.classList.remove("overflow-hidden")
  }
}
