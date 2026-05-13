import { Controller } from "@hotwired/stimulus";

// Full-screen mobile overlay + scroll body; `overflow-hidden` on `<html>` while open.
export default class extends Controller {
  static targets = ["panel", "trigger"];

  connect() {
    this._boundHide = this.hide.bind(this);
    document.addEventListener("click", this._boundHide);
  }

  toggle(event) {
    event.stopPropagation();
    this.panelTarget.classList.toggle("hidden");
    const open = !this.panelTarget.classList.contains("hidden");
    document.documentElement.classList.toggle("overflow-hidden", open);
  }

  hide(event) {
    if (this.panelTarget.classList.contains("hidden")) return;
    const t = event.target;
    if (this.panelTarget.contains(t)) return;
    if (this.triggerTarget.contains(t)) return;
    this.panelTarget.classList.add("hidden");
    document.documentElement.classList.remove("overflow-hidden");
  }

  disconnect() {
    document.removeEventListener("click", this._boundHide);
    document.documentElement.classList.remove("overflow-hidden");
  }
}
