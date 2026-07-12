import { Controller } from "@hotwired/stimulus";

const STORAGE_KEY = "ai-skills-banner-dismissed-until";
const MUTE_DURATION_MS = 24 * 60 * 60 * 1000;

// Manages the AI Skills promotional banner.
// Hides the banner when dismissed, and mutes it for 24h via localStorage
// (see the inline script in application.html.erb's <head> that reads this
// back before paint, mirroring the dark-mode FOUC-avoidance pattern).
export default class extends Controller {
  dismiss() {
    localStorage.setItem(STORAGE_KEY, Date.now() + MUTE_DURATION_MS);
    this.element.remove();
  }
}
