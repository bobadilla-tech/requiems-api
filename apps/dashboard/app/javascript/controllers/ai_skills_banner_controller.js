import { Controller } from "@hotwired/stimulus"

// Manages the AI Skills promotional banner.
// Hides the banner when dismissed. Reappears on every page reload.
export default class extends Controller {
  dismiss() {
    this.element.remove()
  }
}
