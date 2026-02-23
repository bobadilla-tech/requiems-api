import { Controller } from "@hotwired/stimulus"

// Connects to data-controller="date-range"
export default class extends Controller {
  static targets = ["form"]

  declare readonly formTarget: HTMLElement
  declare readonly hasFormTarget: boolean

  toggle(event: Event) {
    event.preventDefault()

    if (this.hasFormTarget) {
      this.formTarget.classList.toggle("hidden")
    }
  }
}
