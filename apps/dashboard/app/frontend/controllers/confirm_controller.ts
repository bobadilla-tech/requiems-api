import { Controller } from "@hotwired/stimulus"

// Confirmation dialog controller
// Usage:
//   <button
//     data-controller="confirm"
//     data-confirm-message-value="Are you sure?"
//     data-action="click->confirm#confirm"
//   >
//     Delete
//   </button>
export default class extends Controller {
  static values = {
    message: { type: String, default: "Are you sure?" }
  }

  declare readonly messageValue: string

  confirm(event: Event) {
    if (!window.confirm(this.messageValue)) {
      event.preventDefault()
      event.stopImmediatePropagation()
    }
  }
}
