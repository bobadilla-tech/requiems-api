import { Controller } from "@hotwired/stimulus"

// Connects to data-controller="copy-key"
export default class extends Controller {
  static values = { key: String }
  static targets = ["defaultText", "copiedText"]

  declare readonly keyValue: string
  declare readonly hasKeyValue: boolean
  declare readonly defaultTextTarget: HTMLElement
  declare readonly copiedTextTarget: HTMLElement
  declare readonly hasDefaultTextTarget: boolean
  declare readonly hasCopiedTextTarget: boolean

  copy(event: Event) {
    event.preventDefault()

    const textToCopy = this.hasKeyValue ? this.keyValue : (this.element as HTMLElement).dataset.copyKeyValue ?? ""

    // Use modern clipboard API
    if (navigator.clipboard && navigator.clipboard.writeText) {
      navigator.clipboard.writeText(textToCopy).then(() => {
        this.showCopiedFeedback()
      }).catch(err => {
        console.error('Failed to copy text: ', err)
        this.fallbackCopy(textToCopy)
      })
    } else {
      this.fallbackCopy(textToCopy)
    }
  }

  fallbackCopy(text: string) {
    // Fallback for older browsers
    const textArea = document.createElement("textarea")
    textArea.value = text
    textArea.style.position = "fixed"
    textArea.style.left = "-999999px"
    textArea.style.top = "-999999px"
    document.body.appendChild(textArea)
    textArea.focus()
    textArea.select()

    try {
      document.execCommand('copy')
      this.showCopiedFeedback()
    } catch (err) {
      console.error('Fallback: Could not copy text: ', err)
    }

    document.body.removeChild(textArea)
  }

  showCopiedFeedback() {
    // Toggle text if targets exist
    if (this.hasDefaultTextTarget && this.hasCopiedTextTarget) {
      this.defaultTextTarget.classList.add("hidden")
      this.copiedTextTarget.classList.remove("hidden")

      setTimeout(() => {
        this.defaultTextTarget.classList.remove("hidden")
        this.copiedTextTarget.classList.add("hidden")
      }, 2000)
    }

    // Show visual feedback
    const button = this.element as HTMLElement
    const originalClasses = button.className
    button.classList.add("bg-green-600")
    button.classList.remove("bg-blue-600")

    setTimeout(() => {
      button.className = originalClasses
    }, 2000)
  }
}
