import { Controller } from "@hotwired/stimulus";

export default class extends Controller {
  static targets = ["input", "result", "errorMessage", "spinner", "button"];

  validate(event) {
    event.preventDefault();
    this._clearError();

    const phone = this.inputTarget.value.trim();

    if (!phone) {
      this._showError("Enter a phone number to validate.");
      return;
    }

    this._setLoading(true);

    setTimeout(() => {
      const simulated = {
        phone: phone,
        valid: true,
        format_valid: true,
        carrier: "T-Mobile",
        line_type: "mobile",
        country: "US",
        reachable: true,
      };

      this._renderResult(simulated);
      this._setLoading(false);
    }, 400);
  }

  _renderResult(data) {
    const panel = this.resultTarget;
    panel.innerHTML = `
      <div class="grid grid-cols-2 gap-x-6 gap-y-3 text-sm">
        <div class="text-gray-400">Valid format</div>
        <div class="text-right font-mono ${
          data.format_valid ? "text-green-400" : "text-red-400"
        }">${data.format_valid ? "✓ Yes" : "✗ No"}</div>

        <div class="text-gray-400">Carrier</div>
        <div class="text-right font-mono text-white">${data.carrier}</div>

        <div class="text-gray-400">Line type</div>
        <div class="text-right font-mono text-white">${data.line_type}</div>

        <div class="text-gray-400">Country</div>
        <div class="text-right font-mono text-white">${data.country}</div>

        <div class="text-gray-400">Reachable</div>
        <div class="text-right font-mono ${
          data.reachable ? "text-green-400" : "text-red-400"
        }">${data.reachable ? "✓ Yes" : "✗ No"}</div>
      </div>
    `;

    panel.classList.remove("hidden");
  }

  _setLoading(loading) {
    if (loading) {
      this.buttonTarget.disabled = true;
      this.spinnerTarget.classList.remove("hidden");
    } else {
      this.buttonTarget.disabled = false;
      this.spinnerTarget.classList.add("hidden");
    }
  }

  _showError(msg) {
    this.errorMessageTarget.textContent = msg;
    this.errorMessageTarget.classList.remove("hidden");
  }

  _clearError() {
    this.errorMessageTarget.classList.add("hidden");
    this.errorMessageTarget.textContent = "";
  }
}
