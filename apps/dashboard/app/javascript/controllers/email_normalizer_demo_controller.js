import { Controller } from "@hotwired/stimulus";

// Handles the Email Normalizer demo in the hero section.
// Simulates a normalization response client-side — no real API call required.
export default class extends Controller {
  static targets = [
    "input",
    "button",
    "result",
    "errorMessage",
    "spinner",
  ];

  connect() {
    this._pending = false;
  }

  normalize(event) {
    event.preventDefault();
    if (this._pending) return;

    this._clearError();
    this.resultTarget.classList.add("hidden");
    const raw = this.inputTarget.value.trim();

    if (!raw) {
      this._showError("Enter an email address to normalize.");
      return;
    }

    this._setLoading(true);

    // Simulate async response so the spinner is visible.
    setTimeout(() => {
      try {
        this._renderResult(this._simulateNormalization(raw));
      } catch (_err) {
        this._showError("Could not normalize the address. Check the format.");
      } finally {
        this._setLoading(false);
      }
    }, 400);
  }

  // ── private ────────────────────────────────────────────────────────────────

  _simulateNormalization(raw) {
    const trimmed = raw.trim().toLowerCase();
    const atIndex = trimmed.lastIndexOf("@");
    if (atIndex === -1) throw new Error("invalid email");

    let local = trimmed.slice(0, atIndex);
    let domain = trimmed.slice(atIndex + 1);
    const notes = [];

    if (raw !== trimmed) notes.push("lowercased");

    // Resolve Googlemail alias.
    if (domain === "googlemail.com") {
      domain = "gmail.com";
      notes.push("alias resolved");
    }

    // Strip Gmail/Fastmail plus-addressing tags.
    const plusProviders = ["gmail.com", "fastmail.com"];
    if (plusProviders.includes(domain) && local.includes("+")) {
      local = local.split("+")[0];
      notes.push("tag removed");
    }

    // Strip Gmail dots.
    if (domain === "gmail.com") {
      local = local.replace(/\./g, "");
      if (trimmed.slice(0, atIndex).includes(".")) {
        notes.push("dots removed");
      }
    }

    const provider = this._detectProvider(domain);
    const normalized = `${local}@${domain}`;

    return {
      original_email: raw,
      normalized_email: normalized,
      local_part: local,
      domain,
      provider,
      notes: notes.length ? notes.join(", ") : "no changes",
    };
  }

  _detectProvider(domain) {
    const map = {
      "gmail.com": "Gmail",
      "googlemail.com": "Gmail",
      "yahoo.com": "Yahoo",
      "yahoo.es": "Yahoo",
      "outlook.com": "Outlook",
      "hotmail.com": "Outlook",
      "live.com": "Outlook",
      "icloud.com": "iCloud",
      "me.com": "iCloud",
      "protonmail.com": "ProtonMail",
      "proton.me": "ProtonMail",
      "fastmail.com": "Fastmail",
    };
    return map[domain] ?? "Other";
  }

  _renderResult(data) {
    const rows = [
      { label: "Normalized email", value: data.normalized_email, type: "text" },
      { label: "Provider", value: data.provider, type: "text" },
      { label: "Local part", value: data.local_part, type: "text" },
      { label: "Domain", value: data.domain, type: "text" },
      { label: "Notes", value: data.notes, type: "notes" },
    ];

    const changed = data.original_email.toLowerCase().trim() !==
      data.normalized_email;

    const items = rows
      .map((r) => `
        <div class="flex items-center justify-between py-2 border-b border-gray-100 dark:border-gray-700 last:border-0">
          <span class="text-sm text-gray-500 dark:text-gray-400">${r.label}</span>
          <span class="text-sm font-medium ${
        this._valueClass(r.type, r.value)
      }">${this._escapeHtml(r.value)}</span>
        </div>`)
      .join("");

    this.resultTarget.innerHTML = `
      <div class="rounded-xl border border-gray-200 dark:border-gray-700 bg-white dark:bg-gray-800 overflow-hidden">
        <div class="px-4 py-3 bg-gray-50 dark:bg-gray-900 flex items-center gap-2">
          <span class="text-xs font-mono text-gray-500 dark:text-gray-400 truncate">${
      this._escapeHtml(data.original_email)
    }</span>
          <span class="ml-auto shrink-0 inline-flex items-center px-2 py-0.5 rounded-full text-xs font-medium ${
      changed
        ? "bg-amber-100 text-amber-700 dark:bg-amber-900 dark:text-amber-300"
        : "bg-emerald-100 text-emerald-700 dark:bg-emerald-900 dark:text-emerald-300"
    }">
            ${changed ? "Normalized" : "Already canonical"}
          </span>
        </div>
        <div class="px-4 py-1">${items}</div>
      </div>`;

    this.resultTarget.classList.remove("hidden");
  }

  _valueClass(type, value) {
    if (type === "notes") {
      return value === "no changes"
        ? "text-gray-400 dark:text-gray-500"
        : "text-amber-600 dark:text-amber-400";
    }
    return "text-gray-800 dark:text-gray-200";
  }

  _setLoading(state) {
    this._pending = state;
    this.buttonTarget.disabled = state;
    this.spinnerTarget.classList.toggle("hidden", !state);
  }

  _showError(msg) {
    this.errorMessageTarget.textContent = msg;
    this.errorMessageTarget.classList.remove("hidden");
  }

  _clearError() {
    this.errorMessageTarget.classList.add("hidden");
    this.errorMessageTarget.textContent = "";
  }

  _escapeHtml(str) {
    return String(str)
      .replace(/&/g, "&amp;")
      .replace(/</g, "&lt;")
      .replace(/>/g, "&gt;")
      .replace(/"/g, "&quot;");
  }
}
