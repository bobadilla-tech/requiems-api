import { Controller } from "@hotwired/stimulus"

// Dropdown menu controller
// Usage:
//   <div data-controller="dropdown">
//     <button data-action="click->dropdown#toggle">Menu</button>
//     <div data-dropdown-target="menu" class="hidden">
//       <!-- Dropdown content -->
//     </div>
//   </div>
export default class extends Controller {
  static targets = ["menu"]

  declare readonly menuTarget: HTMLElement

  private boundHide!: (event: Event) => void

  toggle(event: Event) {
    event.stopPropagation()
    this.menuTarget.classList.toggle("hidden")
  }

  hide(event: Event) {
    if (!this.element.contains(event.target as Node)) {
      this.menuTarget.classList.add("hidden")
    }
  }

  connect() {
    // Close dropdown when clicking outside
    this.boundHide = this.hide.bind(this)
    document.addEventListener("click", this.boundHide)
  }

  disconnect() {
    document.removeEventListener("click", this.boundHide)
  }
}
