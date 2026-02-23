import { Controller } from "@hotwired/stimulus"

// Tabs controller
// Usage:
//   <div data-controller="tabs" data-tabs-index-value="0">
//     <nav>
//       <button data-action="click->tabs#select" data-tabs-index-param="0">Tab 1</button>
//       <button data-action="click->tabs#select" data-tabs-index-param="1">Tab 2</button>
//     </nav>
//     <div data-tabs-target="panel" class="">Content 1</div>
//     <div data-tabs-target="panel" class="hidden">Content 2</div>
//   </div>
export default class extends Controller {
  static targets = ["panel", "tab"]
  static values = { index: { type: Number, default: 0 } }

  declare readonly panelTargets: HTMLElement[]
  declare readonly tabTargets: HTMLElement[]
  declare readonly hasTabTarget: boolean
  declare indexValue: number

  connect() {
    this.showTab(this.indexValue)
  }

  select(event: Event) {
    event.preventDefault()
    const index = parseInt((event.currentTarget as HTMLElement).dataset.tabsIndexParam ?? "0")
    this.indexValue = index
    this.showTab(index)
  }

  showTab(index: number) {
    // Hide all panels
    this.panelTargets.forEach((panel, i) => {
      if (i === index) {
        panel.classList.remove("hidden")
      } else {
        panel.classList.add("hidden")
      }
    })

    // Update tab styles
    if (this.hasTabTarget) {
      this.tabTargets.forEach((tab, i) => {
        if (i === index) {
          tab.classList.add("border-blue-500", "text-blue-600")
          tab.classList.remove("border-transparent", "text-gray-500")
        } else {
          tab.classList.remove("border-blue-500", "text-blue-600")
          tab.classList.add("border-transparent", "text-gray-500")
        }
      })
    }
  }
}
