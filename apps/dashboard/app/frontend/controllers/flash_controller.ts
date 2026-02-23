import { Controller } from "@hotwired/stimulus"

// Auto-dismissing flash messages controller
// Usage: data-controller="flash" data-flash-delay-value="5000"
export default class extends Controller {
  static values = {
    delay: { type: Number, default: 5000 }
  }

  declare readonly delayValue: number

  private timeout?: ReturnType<typeof setTimeout>

  connect() {
    // Auto-dismiss after delay
    this.timeout = setTimeout(() => {
      this.dismiss()
    }, this.delayValue)
  }

  disconnect() {
    if (this.timeout) {
      clearTimeout(this.timeout)
    }
  }

  dismiss() {
    // Fade out animation
    this.element.classList.add("transition-all", "duration-300", "opacity-0", "translate-y-2")

    // Remove after animation
    setTimeout(() => {
      this.element.remove()
    }, 300)
  }

  // Allow manual dismiss
  close() {
    if (this.timeout) {
      clearTimeout(this.timeout)
    }
    this.dismiss()
  }
}
