import { Controller } from "@hotwired/stimulus"
import hljs from "highlight.js"

// Connects to data-controller="highlight"
export default class extends Controller {
  connect() {
    // Highlight all code blocks within this element
    this.element.querySelectorAll('pre code').forEach((block) => {
      hljs.highlightElement(block as HTMLElement)
    })
  }
}
