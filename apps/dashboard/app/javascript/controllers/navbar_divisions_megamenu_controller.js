import { Controller } from "@hotwired/stimulus";

// Hover/focus flyout: keeps panel open while moving from hub link into the panel (relatedTarget + delay).
export default class extends Controller {
  static targets = ["panel", "trigger"];

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
    if (this.hasTriggerTarget) {
      this.triggerTarget.setAttribute("aria-expanded", "true");
    }
  }

  _hidePanel() {
    const p = this.panelTarget;
    p.classList.add("invisible", "opacity-0", "pointer-events-none");
    p.classList.remove("opacity-100");
    if (this.hasTriggerTarget) {
      this.triggerTarget.setAttribute("aria-expanded", "false");
    }
  }

  _scheduleHide() {
    this._clearHideTimer();
    this._hideTimer = window.setTimeout(() => {
      this._hidePanel();
      this._hideTimer = null;
    }, 280);
  }

  _clearHideTimer() {
    if (this._hideTimer != null) {
      window.clearTimeout(this._hideTimer);
      this._hideTimer = null;
    }
  }
}
