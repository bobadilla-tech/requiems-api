import { Controller } from "@hotwired/stimulus";

// Carousel + use-case grid for division marketing pages
export default class extends Controller {
  static targets = ["slide", "dot", "gridCard", "slideViewport"];
  static values = { index: { type: Number, default: 0 } };

  connect() {
    this.indexValue = 0;
    this.renderSlide();
    this.gridCardTargets.forEach((card) => card.classList.remove("hidden"));
    this._onResize = () => this.scheduleSyncSlideHeights();
    window.addEventListener("resize", this._onResize);
    requestAnimationFrame(() => {
      requestAnimationFrame(() => this.syncSlideHeights());
    });
    if (typeof document !== "undefined" && document.fonts?.ready) {
      document.fonts.ready.then(() => this.syncSlideHeights());
    }
  }

  disconnect() {
    window.removeEventListener("resize", this._onResize);
    clearTimeout(this._resizeTimer);
  }

  scheduleSyncSlideHeights() {
    clearTimeout(this._resizeTimer);
    this._resizeTimer = setTimeout(() => this.syncSlideHeights(), 120);
  }

  /** Keep carousel panel height stable when switching slides (max of all slides). */
  syncSlideHeights() {
    if (!this.hasSlideViewportTarget) return;
    const slides = this.slideTargets;
    if (slides.length <= 1) {
      this.slideViewportTarget.style.minHeight = "";
      return;
    }

    let max = 0;
    slides.forEach((el) => {
      slides.forEach((s) => s.classList.add("hidden"));
      el.classList.remove("hidden");
      max = Math.max(max, el.offsetHeight);
    });

    this.renderSlide();
    this.slideViewportTarget.style.minHeight = `${Math.ceil(max)}px`;
  }

  next(event) {
    event.preventDefault();
    const n = this.slideTargets.length;
    if (n === 0) return;
    this.indexValue = (this.indexValue + 1) % n;
    this.renderSlide();
  }

  prev(event) {
    event.preventDefault();
    const n = this.slideTargets.length;
    if (n === 0) return;
    this.indexValue = (this.indexValue - 1 + n) % n;
    this.renderSlide();
  }

  goTo(event) {
    event.preventDefault();
    const i = Number.parseInt(event.currentTarget.dataset.slideIndex, 10);
    if (Number.isNaN(i)) return;
    this.indexValue = i;
    this.renderSlide();
  }

  renderSlide() {
    this.slideTargets.forEach((el, i) => {
      const on = i === this.indexValue;
      el.classList.toggle("hidden", !on);
      el.setAttribute("aria-hidden", on ? "false" : "true");
    });
    this.dotTargets.forEach((dot, i) => {
      const on = i === this.indexValue;
      dot.classList.toggle("bg-blue-600", on);
      dot.classList.toggle("dark:bg-blue-500", on);
      dot.classList.toggle("bg-gray-300", !on);
      dot.classList.toggle("dark:bg-gray-600", !on);
      dot.setAttribute("aria-selected", on ? "true" : "false");
    });
  }

}
