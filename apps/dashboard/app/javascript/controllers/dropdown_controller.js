import { Controller } from "@hotwired/stimulus";

// Dropdown menu controller
// Usage:
//   <div data-controller="dropdown">
//     <button data-action="click->dropdown#toggle">Menu</button>
//     <div data-dropdown-target="menu" class="hidden">
//       <!-- Dropdown content -->
//     </div>
//   </div>
export default class extends Controller {
  static targets = ["menu", "trigger"];

  toggle(event) {
    event.stopPropagation();
    this.menuTarget.classList.toggle("hidden");
  }

  hide(event) {
    if (this.menuTarget.classList.contains("hidden")) return;
    const t = event.target;
    if (this.menuTarget.contains(t)) return;
    if (this.hasTriggerTarget && this.triggerTarget.contains(t)) return;
    if (!this.hasTriggerTarget && this.element.contains(t)) return;
    this.menuTarget.classList.add("hidden");
  }

  connect() {
    this._boundHide = this.hide.bind(this);
    document.addEventListener("click", this._boundHide);
  }

  disconnect() {
    document.removeEventListener("click", this._boundHide);
  }
}
