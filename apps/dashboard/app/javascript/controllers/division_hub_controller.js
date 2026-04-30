import { Controller } from "@hotwired/stimulus";

// Carousel + filter tabs for division marketing pages
export default class extends Controller {
  static targets = ["slide", "dot", "filterBtn", "gridCard"];
  static values = { index: { type: Number, default: 0 } };

  connect() {
    this.indexValue = 0;
    this.activeFilter = "all";
    this.renderSlide();
    this.applyFilter();
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

  setFilter(event) {
    event.preventDefault();
    const key = event.currentTarget.dataset.filterKey;
    if (!key) return;
    this.activeFilter = key;
    this.applyFilter();
    this.filterBtnTargets.forEach((btn) => {
      const active = btn.dataset.filterKey === key;
      btn.classList.toggle("bg-blue-600", active);
      btn.classList.toggle("text-white", active);
      btn.classList.toggle("dark:bg-blue-500", active);
      btn.classList.toggle("border-blue-600", active);
      btn.classList.toggle("border-gray-300", !active);
      btn.classList.toggle("text-gray-700", !active);
      btn.classList.toggle("dark:border-gray-600", !active);
      btn.classList.toggle("dark:text-gray-200", !active);
    });
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

  applyFilter() {
    this.gridCardTargets.forEach((card) => {
      const fk = card.dataset.filterKey || "";
      const show = this.activeFilter === "all" || fk === this.activeFilter;
      card.classList.toggle("hidden", !show);
    });
  }
}
