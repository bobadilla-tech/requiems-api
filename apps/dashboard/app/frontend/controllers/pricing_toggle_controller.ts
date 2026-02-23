import { Controller } from "@hotwired/stimulus"

// Toggle between monthly and yearly pricing
export default class extends Controller {
  static targets = ["monthlyButton", "yearlyButton", "monthlyPrice", "yearlyPrice", "billingCycleInput"]

  declare readonly monthlyButtonTarget: HTMLElement
  declare readonly yearlyButtonTarget: HTMLElement
  declare readonly monthlyPriceTargets: HTMLElement[]
  declare readonly yearlyPriceTargets: HTMLElement[]
  declare readonly billingCycleInputTargets: HTMLInputElement[]

  connect() {
    this.showMonthly()
  }

  showMonthly() {
    // Update button styles
    this.monthlyButtonTarget.className = "px-6 py-2 rounded-md text-sm font-medium transition-colors bg-blue-600 text-white"
    this.yearlyButtonTarget.className = "px-6 py-2 rounded-md text-sm font-medium transition-colors text-gray-700 hover:text-gray-900"

    // Show monthly prices, hide yearly prices
    this.monthlyPriceTargets.forEach(el => el.classList.remove("hidden"))
    this.yearlyPriceTargets.forEach(el => el.classList.add("hidden"))

    // Update billing cycle inputs
    this.billingCycleInputTargets.forEach(input => input.value = "monthly")
  }

  showYearly() {
    // Update button styles
    this.yearlyButtonTarget.className = "px-6 py-2 rounded-md text-sm font-medium transition-colors bg-blue-600 text-white"
    this.monthlyButtonTarget.className = "px-6 py-2 rounded-md text-sm font-medium transition-colors text-gray-700 hover:text-gray-900"

    // Show yearly prices, hide monthly prices
    this.yearlyPriceTargets.forEach(el => el.classList.remove("hidden"))
    this.monthlyPriceTargets.forEach(el => el.classList.add("hidden"))

    // Update billing cycle inputs
    this.billingCycleInputTargets.forEach(input => input.value = "yearly")
  }
}
