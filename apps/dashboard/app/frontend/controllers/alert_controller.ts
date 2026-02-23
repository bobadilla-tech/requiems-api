import { Controller } from "@hotwired/stimulus"

// Dismissible alert controller
// Usage: data-controller="alert"
export default class extends Controller {
  dismiss() {
    this.element.remove()
  }

  // Auto-dismiss after delay (in milliseconds)
  connect() {
    const delay = (this.element as HTMLElement).dataset.alertDelay
    if (delay) {
      setTimeout(() => {
        this.fadeOut()
      }, parseInt(delay))
    }
  }

  fadeOut() {
    this.element.classList.add("transition-opacity", "duration-300", "opacity-0")
    setTimeout(() => {
      this.element.remove()
    }, 300)
  }
}
