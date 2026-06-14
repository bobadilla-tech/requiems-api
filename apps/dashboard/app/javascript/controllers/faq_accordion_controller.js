import { Controller } from "@hotwired/stimulus";

// FAQ accordion controller
export default class extends Controller {
  static targets = ["button", "content"];

  toggle(event) {
    const button = event.currentTarget;
    const content = button.nextElementSibling;
    const icon = button.querySelector("svg");

    // Toggle content visibility
    content.classList.toggle("hidden");

    // Keep aria-expanded in sync with visibility state
    const isOpen = !content.classList.contains("hidden");
    button.setAttribute("aria-expanded", String(isOpen));

    // Rotate icon
    if (content.classList.contains("hidden")) {
      icon.style.transform = "rotate(0deg)";
    } else {
      icon.style.transform = "rotate(180deg)";
    }
  }
}
