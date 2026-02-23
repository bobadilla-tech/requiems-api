import { Controller } from "@hotwired/stimulus"

// Client-side API search and filtering
export default class extends Controller {
  static targets = ["input", "card", "section", "empty"]

  declare readonly inputTarget: HTMLInputElement
  declare readonly cardTargets: HTMLElement[]
  declare readonly sectionTargets: HTMLElement[]
  declare readonly emptyTarget: HTMLElement
  declare readonly hasSectionTarget: boolean
  declare readonly hasEmptyTarget: boolean

  private totalCount = 0

  connect() {
    this.totalCount = this.cardTargets.length

    // Auto-filter if input has value (from URL param)
    if (this.inputTarget.value.trim() !== '') {
      this.filter()
    }
  }

  filter() {
    const query = this.inputTarget.value.toLowerCase().trim()
    let visibleCount = 0

    // Filter individual cards
    this.cardTargets.forEach((card) => {
      const name = card.dataset.apiName || ""
      const description = card.dataset.apiDescription || ""
      const tags = card.dataset.apiTags || ""

      const matches = name.includes(query) ||
                     description.includes(query) ||
                     tags.includes(query)

      if (query === "" || matches) {
        card.classList.remove("hidden")
        visibleCount++
      } else {
        card.classList.add("hidden")
      }
    })

    // Hide/show sections based on visible cards
    if (this.hasSectionTarget) {
      this.sectionTargets.forEach((section) => {
        const sectionCards = section.querySelectorAll('[data-api-search-target="card"]')
        const visibleInSection = Array.from(sectionCards).some(card => !card.classList.contains("hidden"))

        if (visibleInSection) {
          section.classList.remove("hidden")
        } else {
          section.classList.add("hidden")
        }
      })
    }

    // Show/hide empty state
    if (this.hasEmptyTarget) {
      if (visibleCount === 0 && query !== "") {
        this.emptyTarget.classList.remove("hidden")
      } else {
        this.emptyTarget.classList.add("hidden")
      }
    }
  }
}
