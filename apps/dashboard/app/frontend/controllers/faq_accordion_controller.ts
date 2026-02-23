import { Controller } from "@hotwired/stimulus"

// FAQ accordion controller
export default class extends Controller {
  static targets = ["button", "content"]

  toggle(event: Event) {
    const button = event.currentTarget as HTMLElement
    const content = button.nextElementSibling as HTMLElement | null
    const icon = button.querySelector("svg") as SVGElement | null

    if (!content || !icon) return

    // Toggle content visibility
    content.classList.toggle("hidden")

    // Rotate icon
    if (content.classList.contains("hidden")) {
      icon.style.transform = "rotate(0deg)"
    } else {
      icon.style.transform = "rotate(180deg)"
    }
  }
}
