import { Controller } from "@hotwired/stimulus"

// Global keyboard shortcuts for search
// Listens for "/" or "Cmd+K"/"Ctrl+K" to focus the navbar search
export default class extends Controller {
  private boundHandleKeydown!: (event: KeyboardEvent) => void

  connect() {
    this.boundHandleKeydown = this.handleKeydown.bind(this)
    document.addEventListener("keydown", this.boundHandleKeydown)
  }

  disconnect() {
    document.removeEventListener("keydown", this.boundHandleKeydown)
  }

  handleKeydown(event: KeyboardEvent) {
    // Only trigger if not in an input/textarea
    const target = event.target as HTMLElement
    const inInput = target.tagName === 'INPUT' ||
                    target.tagName === 'TEXTAREA' ||
                    target.isContentEditable

    if (inInput) return

    // "/" key or "Cmd+K" / "Ctrl+K"
    const isSlashKey = event.key === '/'
    const isCmdK = (event.metaKey || event.ctrlKey) && event.key === 'k'

    if (isSlashKey || isCmdK) {
      event.preventDefault()

      // Find and focus navbar search
      const searchElement = document.querySelector('[data-controller~="navbar-search"]')
      if (searchElement) {
        const searchController = this.application.getControllerForElementAndIdentifier(
          searchElement as HTMLElement,
          'navbar-search'
        )

        if (searchController && typeof (searchController as { focus?: () => void }).focus === 'function') {
          (searchController as { focus: () => void }).focus()
        }
      }
    }
  }
}
