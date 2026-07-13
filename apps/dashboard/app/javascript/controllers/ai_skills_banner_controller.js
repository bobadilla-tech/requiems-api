import { Controller } from "@hotwired/stimulus";

// Keep in sync with the literal in application.html.erb's prepaint <head>
// script — can't share a single constant without either deferring that
// script (reintroducing the flash it exists to avoid) or adding a .js.erb
// build step this app doesn't otherwise use.
const STORAGE_KEY = "ai-skills-banner-dismissed-until";
const MUTE_DURATION_MS = 24 * 60 * 60 * 1000;
const HIDE_CLASS = "hide-ai-skills-banner";

// Manages the AI Skills promotional banner.
// Hides the banner when dismissed, and mutes it for 24h via localStorage
// (see the inline script in application.html.erb's <head> that reads this
// back before paint, mirroring the dark-mode FOUC-avoidance pattern).
//
// connect() re-applies the same hide class on every Turbo visit (Stimulus
// reconnects controllers whenever their element is (re)inserted by a Turbo
// body swap). Verified against Turbo's own head-diffing that in practice a
// byte-identical inline <head> script does get re-run on each visit here —
// but that's an internal implementation detail Turbo doesn't document as
// guaranteed, so this makes the sync explicit rather than relying on it.
export default class extends Controller {
  connect() {
    let mutedUntil;
    try {
      mutedUntil = localStorage.getItem(STORAGE_KEY);
    } catch {
      return;
    }
    const isMuted = Boolean(mutedUntil) && Date.now() < Number(mutedUntil);
    document.documentElement.classList.toggle(HIDE_CLASS, isMuted);
  }

  dismiss() {
    try {
      localStorage.setItem(STORAGE_KEY, Date.now() + MUTE_DURATION_MS);
    } catch {
      // Storage unavailable (private browsing, quota, disabled, etc.) —
      // still dismiss for this page view even though it won't persist.
    }
    this.element.remove();
  }
}
