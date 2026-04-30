import { Controller } from "@hotwired/stimulus";

// Keeps the divisions flyout open while moving from the trigger link into the panel
// (avoids mouseleave when the pointer crosses the visual gap).
export default class extends Controller {
  static targets = ["panel"];

  connect() {
    this._hideTimer = null;
  }

  disconnect() {
    this._clearHideTimer();
  }

  onEnter() {
    this._clearHideTimer();
    this._showPanel();
  }

  onLeave(event) {
    const next = event.relatedTarget;
    if (next && this.element.contains(next)) return;
    this._scheduleHide();
  }

  _showPanel() {
    const p = this.panelTarget;
    p.classList.remove("invisible", "opacity-0", "pointer-events-none");
    p.classList.add("opacity-100");
  }

  _hidePanel() {
    const p = this.panelTarget;
    p.classList.add("invisible", "opacity-0", "pointer-events-none");
    p.classList.remove("opacity-100");
  }

  _scheduleHide() {
    this._clearHideTimer();
    this._hideTimer = window.setTimeout(() => {
      this._hidePanel();
      this._hideTimer = null;
    }, 120);
  }

  _clearHideTimer() {
    if (this._hideTimer != null) {
      window.clearTimeout(this._hideTimer);
      this._hideTimer = null;
    }
  }
}
