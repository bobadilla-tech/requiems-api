import { Controller } from "@hotwired/stimulus"

export default class extends Controller {
  connect() {
    this.links = Array.from(this.element.querySelectorAll("a[href^='#']"))
    this.targets = this.links.map(a => document.querySelector(a.getAttribute("href"))).filter(Boolean)
    this._onScroll = this._update.bind(this)
    window.addEventListener("scroll", this._onScroll, { passive: true })
    this._update()
  }

  disconnect() {
    window.removeEventListener("scroll", this._onScroll)
  }

  _update() {
    const scrollY = window.scrollY + 120
    let active = this.targets[0]

    for (const target of this.targets) {
      if (target.offsetTop <= scrollY) active = target
    }

    this.links.forEach(link => {
      const isCurrent = link.getAttribute("href") === `#${active?.id}`
      link.classList.toggle("bg-gray-100", isCurrent)
      link.classList.toggle("dark:bg-gray-700", isCurrent)
      link.classList.toggle("text-gray-900", isCurrent)
      link.classList.toggle("dark:text-gray-100", isCurrent)
      link.classList.toggle("font-semibold", isCurrent)
    })
  }
}
