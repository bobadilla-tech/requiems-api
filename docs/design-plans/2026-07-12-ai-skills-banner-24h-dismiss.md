# AI Skills Banner: 24h Dismiss Persistence

The AI Skills promotional banner (`partials/shared/ai_skills_banner`, rendered unconditionally on every page in `application.html.erb`) had no persistence: `ai_skills_banner_controller.js`'s only method, `dismiss()`, just called `this.element.remove()`. The banner reappeared on every page load or navigation, even seconds after a user closed it — the file's own comment admitted this ("Reappears on every page reload"). The fix: closing the banner should suppress it for 24 hours, fully client-side, no server or cookie involvement.

---

## What Changed

### 1. Persisting the dismissal

`ai_skills_banner_controller.js`'s `dismiss()` now writes a "muted until" timestamp to `localStorage` in addition to removing the element from the DOM:

```js
const STORAGE_KEY = "ai-skills-banner-dismissed-until";
const MUTE_DURATION_MS = 24 * 60 * 60 * 1000;

dismiss() {
  localStorage.setItem(STORAGE_KEY, Date.now() + MUTE_DURATION_MS);
  this.element.remove();
}
```

The value is a plain numeric string (`Date.now() + 24h`, in ms), not JSON — matching the existing convention used by `theme_controller.js` for the dark-mode preference (`localStorage.setItem("theme", ...)`, no namespacing, no serialization). Each dismiss resets the 24h window from that moment; after it elapses, the banner shows again on the next load.

### 2. Avoiding a flash on reload

A Stimulus controller only connects after the DOM has parsed, so checking the mute state inside `connect()` would let the banner flash visible for a moment on every page load before being hidden — even on days it should stay muted. The app already solves this exact problem for dark mode: `application.html.erb`'s `<head>` has a synchronous inline `<script>` that reads `localStorage` and adds a class to `<html>` *before* the body paints.

This change adds a second inline script right after it, following the same shape:

```erb
<script>
  (function() {
    var mutedUntil = localStorage.getItem('ai-skills-banner-dismissed-until');
    if (mutedUntil && Date.now() < Number(mutedUntil)) {
      document.documentElement.classList.add('hide-ai-skills-banner');
    }
  })();
</script>
<style>.hide-ai-skills-banner #ai-skills-banner { display: none; }</style>
```

Same file, same IIFE-in-`<head>` idiom, same "add a class to `<html>` before paint" mechanism as the theme script — no new pattern introduced.

### Known tradeoff

This is intentionally 100% client-side (no cookie, no server session state), matching the request. Like the existing dark-mode script it mirrors, it depends on JavaScript running before first paint — with JS disabled or blocked, the banner will show regardless of prior dismissal. This is an accepted tradeoff already present in this codebase for the same class of problem, not a new risk introduced here.
