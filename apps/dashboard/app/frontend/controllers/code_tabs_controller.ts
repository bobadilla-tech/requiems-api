import { Controller } from "@hotwired/stimulus"

// Code example tabs switcher
export default class extends Controller {
  static targets = ["tab", "content"]

  declare readonly tabTargets: HTMLElement[]
  declare readonly contentTargets: HTMLElement[]

  switch(event: Event) {
    const index = parseInt((event.currentTarget as HTMLElement).dataset.index ?? "0")

    // Update tab styles
    this.tabTargets.forEach((tab, i) => {
      if (i === index) {
        tab.className = "px-4 py-2 text-sm font-medium text-blue-600 border-b-2 border-blue-600 transition-colors"
      } else {
        tab.className = "px-4 py-2 text-sm font-medium text-gray-600 hover:text-gray-900 transition-colors"
      }
    })

    // Show/hide content
    this.contentTargets.forEach((content, i) => {
      if (i === index) {
        content.classList.remove("hidden")
      } else {
        content.classList.add("hidden")
      }
    })
  }
}
