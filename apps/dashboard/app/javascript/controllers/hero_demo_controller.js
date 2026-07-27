import { Controller } from "@hotwired/stimulus";

// Animated API playground for the hero section.
//
// Color scheme (inline styles — not Tailwind, works regardless of JIT scanning):
//   Keys        → #7dd3fc  (sky-300)
//   Strings     → #fcd34d  (amber-300)
//   Bool/num    → #c084fc  (purple-400)
//   Punctuation → #6b7280  (gray-500)
export default class extends Controller {
  static targets = ["urlPath", "demoLabel", "responseArea", "dot"];

  // Inline-style span helpers
  _k(v) {
    return `<span style="color:#7dd3fc">${v}</span>`;
  } // key
  _s(v) {
    return `<span style="color:#fcd34d">${v}</span>`;
  } // string
  _b(v) {
    return `<span style="color:#c084fc">${v}</span>`;
  } // bool/number
  _p(v) {
    return `<span style="color:#6b7280">${v}</span>`;
  } // punctuation

  get demos() {
    const k = this._k.bind(this);
    const s = this._s.bind(this);
    const b = this._b.bind(this);
    const p = this._p.bind(this);

    return [
      {
        label: "Email Validation",
        method: "POST",
        methodStyle:
          "background-color:#1e3a5f;color:#93c5fd;min-width:42px;text-align:center;",
        path: "email/validate",
        lines: [
          p("{"),
          `  ${k('"email"')}${p(":")} ${s('"jane@discardmail.con"')}${p(",")}`,
          `  ${k('"valid"')}${p(":")} ${b("false")}${p(",")}`,
          `  ${k('"syntax_valid"')}${p(":")} ${b("true")}${p(",")}`,
          `  ${k('"mx_valid"')}${p(":")} ${b("false")}${p(",")}`,
          `  ${k('"disposable"')}${p(":")} ${b("true")}${p(",")}`,
          `  ${k('"normalized"')}${p(":")} ${s('"jane@discardmail.con"')}${
            p(",")
          }`,
          `  ${k('"domain"')}${p(":")} ${s('"discardmail.con"')}${p(",")}`,
          `  ${k('"suggestion"')}${p(":")} ${s('"discardmail.com"')}`,
          p("}"),
        ],
      },
      {
        label: "Phone Validation",
        method: "GET",
        methodStyle:
          "background-color:#14532d;color:#86efac;min-width:42px;text-align:center;",
        path: "tech/validate/phone",
        lines: [
          p("{"),
          `  ${k('"number"')}${p(":")} ${s('"+14155552671"')}${p(",")}`,
          `  ${k('"valid"')}${p(":")} ${b("true")}${p(",")}`,
          `  ${k('"country"')}${p(":")} ${s('"US"')}${p(",")}`,
          `  ${k('"type"')}${p(":")} ${s('"mobile"')}${p(",")}`,
          `  ${k('"formatted"')}${p(":")} ${s('"+1 415-555-2671"')}${p(",")}`,
          `  ${k('"carrier"')}${p(": {")}`,
          `    ${k('"name"')}${p(":")} ${s('"T-Mobile"')}${p(",")}`,
          `    ${k('"source"')}${p(":")} ${s('"metadata"')}`,
          `  ${p("},")}`,
          `  ${k('"risk"')}${p(": {")}`,
          `    ${k('"is_voip"')}${p(":")} ${b("false")}${p(",")}`,
          `    ${k('"is_virtual"')}${p(":")} ${b("false")}`,
          `  ${p("}")}`,
          p("}"),
        ],
      },
      {
        label: "BIN Lookup",
        method: "GET",
        methodStyle:
          "background-color:#14532d;color:#86efac;min-width:42px;text-align:center;",
        path: "finance/bin/453980",
        lines: [
          p("{"),
          `  ${k('"bin"')}${p(":")} ${s('"453980"')}${p(",")}`,
          `  ${k('"scheme"')}${p(":")} ${s('"visa"')}${p(",")}`,
          `  ${k('"card_type"')}${p(":")} ${s('"credit"')}${p(",")}`,
          `  ${k('"card_level"')}${p(":")} ${s('"platinum"')}${p(",")}`,
          `  ${k('"issuer_name"')}${p(":")} ${s('"Chase Bank"')}${p(",")}`,
          `  ${k('"issuer_url"')}${p(":")} ${s('"https://www.chase.com"')}${
            p(",")
          }`,
          `  ${k('"issuer_phone"')}${p(":")} ${s('"+1-800-432-3117"')}${
            p(",")
          }`,
          `  ${k('"country_code"')}${p(":")} ${s('"US"')}${p(",")}`,
          `  ${k('"country_name"')}${p(":")} ${s('"United States"')}${p(",")}`,
          `  ${k('"prepaid"')}${p(":")} ${b("false")}${p(",")}`,
          `  ${k('"luhn_prefix_valid"')}${p(":")} ${b("true")}${p(",")}`,
          `  ${k('"confidence"')}${p(":")} ${b("0.97")}`,
          p("}"),
        ],
      },
    ];
  }

  connect() {
    this._index = 0;
    this._running = true;
    this._showImmediate(this.demos[0]);
    this._loop();
  }

  disconnect() {
    this._running = false;
    clearTimeout(this._timer);
  }

  _showImmediate(demo) {
    this.urlPathTarget.textContent = demo.path;
    this.demoLabelTarget.textContent = demo.label;
    this.responseAreaTarget.innerHTML = "";
    for (const line of demo.lines) {
      this.responseAreaTarget.appendChild(this._makeLine(line));
    }
    this._updateDots(0);
  }

  async _loop() {
    while (this._running) {
      await this._sleep(5000);
      if (!this._running) return;

      this._index = (this._index + 1) % this.demos.length;
      const demo = this.demos[this._index];

      // Fade out response
      this.responseAreaTarget.style.transition = "opacity 0.3s ease";
      this.responseAreaTarget.style.opacity = "0";
      await this._sleep(320);
      this.responseAreaTarget.innerHTML = "";
      this.responseAreaTarget.style.transition = "";
      this.responseAreaTarget.style.opacity = "1";

      // Update label + dots
      this.demoLabelTarget.textContent = demo.label;
      this._updateDots(this._index);

      // Partial URL rewrite
      await this._rewritePath(this.urlPathTarget.textContent, demo.path);
      await this._sleep(280);

      // Stagger lines in
      for (const line of demo.lines) {
        if (!this._running) return;
        const el = this._makeLine(line);
        el.style.opacity = "0";
        el.style.transform = "translateY(3px)";
        el.style.transition = "opacity 0.12s ease, transform 0.12s ease";
        this.responseAreaTarget.appendChild(el);
        void el.offsetWidth;
        el.style.opacity = "1";
        el.style.transform = "translateY(0)";
        await this._sleep(65);
      }
    }
  }

  async _rewritePath(oldPath, newPath) {
    let common = 0;
    while (
      common < oldPath.length && common < newPath.length &&
      oldPath[common] === newPath[common]
    ) {
      common++;
    }
    let current = oldPath;
    while (current.length > common) {
      current = current.slice(0, -1);
      this.urlPathTarget.textContent = current;
      await this._sleep(16 + Math.random() * 16);
    }
    const suffix = newPath.slice(common);
    for (const char of suffix) {
      if (!this._running) return;
      current += char;
      this.urlPathTarget.textContent = current;
      await this._sleep(28 + Math.random() * 32);
    }
  }

  _updateDots(activeIndex) {
    this.dotTargets.forEach((dot, i) => {
      dot.style.cssText = i === activeIndex
        ? "width:8px;height:8px;background-color:rgb(96,165,250);border-radius:9999px;transition:all 0.3s;"
        : "width:6px;height:6px;background-color:rgba(255,255,255,0.18);border-radius:9999px;transition:all 0.3s;";
    });
  }

  _makeLine(html) {
    const el = document.createElement("div");
    el.innerHTML = html;
    // white-space:pre preserves the leading spaces used for JSON indentation
    el.style.cssText =
      "font-family:ui-monospace,monospace;font-size:13px;line-height:1.65;white-space:pre;color:#9ca3af;";
    return el;
  }

  _sleep(ms) {
    return new Promise((resolve) => {
      this._timer = setTimeout(resolve, ms);
    });
  }
}
