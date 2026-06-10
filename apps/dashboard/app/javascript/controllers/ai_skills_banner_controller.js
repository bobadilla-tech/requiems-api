import { Controller } from "@hotwired/stimulus"

// Manages the AI Skills promotional banner.
// Hides the banner when dismissed and remembers the choice in localStorage.
export default class extends Controller {
  connect() {
    if (localStorage.getItem("ai_skills_banner_dismissed") === "true") {
      this.element.remove()
    }
  }

  dismiss() {
    localStorage.setItem("ai_skills_banner_dismissed", "true")
    this.element.remove()
  }
}
