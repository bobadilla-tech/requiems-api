import { Controller } from "@hotwired/stimulus"

interface RequestData {
  [key: string]: string | number | boolean | unknown[] | Record<string, unknown>
}

// Interactive API playground for testing endpoints — routes through /api/proxy
export default class extends Controller {
  static targets = [
    "param",
    "submitButton",
    "submitText",
    "loading",
    "responseContainer",
    "responseStatus",
    "statusBadge",
    "responseTime",
    "responseBody",
    "errorContainer",
    "errorMessage"
  ]

  static values = {
    method: String,
    url: String,
    index: Number,
    modalId: String
  }

  declare readonly paramTargets: HTMLInputElement[]
  declare readonly submitButtonTarget: HTMLButtonElement
  declare readonly submitTextTarget: HTMLElement
  declare readonly loadingTarget: HTMLElement
  declare readonly responseContainerTarget: HTMLElement
  declare readonly responseStatusTarget: HTMLElement
  declare readonly statusBadgeTarget: HTMLElement
  declare readonly responseTimeTarget: HTMLElement
  declare readonly responseBodyTarget: HTMLElement
  declare readonly errorContainerTarget: HTMLElement
  declare readonly errorMessageTarget: HTMLElement

  declare readonly methodValue: string
  declare readonly urlValue: string
  declare readonly indexValue: number
  declare readonly modalIdValue: string
  declare readonly hasModalIdValue: boolean

  private startTime: number | null = null
  private hasResponse = false

  connect() {
    this.startTime = null
    this.hasResponse = false
  }

  async sendRequest(event: Event) {
    event.preventDefault()

    // Collect parameters
    const requestData = this.collectParameters()

    // Validate required parameters
    if (!this.validateParameters()) {
      this.showError("Please fill in all required fields")
      return
    }

    // Show loading state
    this.showLoading()

    try {
      this.startTime = performance.now()

      const response = await this.makeRequest(requestData)

      const endTime = performance.now()
      const clientDuration = Math.round(endTime - this.startTime)

      await this.handleResponse(response, clientDuration)
    } catch (error) {
      this.showError(error instanceof Error ? error.message : "Request failed. Please try again.")
    } finally {
      this.hideLoading()
    }
  }

  collectParameters(): RequestData {
    const data: RequestData = {}

    this.paramTargets.forEach((input) => {
      const name = input.dataset.paramName
      const type = input.dataset.paramType
      let value: string | number | boolean | unknown[] | Record<string, unknown> = input.value.trim()

      if (!value || !name) return

      // Parse arrays and objects
      if (type === "array" || type === "object") {
        try {
          value = JSON.parse(value as string)
        } catch (e) {
          console.error(`Failed to parse ${type} for ${name}:`, e)
          return
        }
      }

      // Convert numbers and booleans
      if (type === "integer" || type === "number") {
        value = Number(value)
      } else if (type === "boolean") {
        value = value === "true"
      }

      data[name] = value
    })

    return data
  }

  validateParameters(): boolean {
    const requiredParams = this.paramTargets.filter(
      (input) => input.dataset.paramRequired === "true"
    )

    for (const param of requiredParams) {
      if (!param.value.trim()) {
        param.classList.add("border-red-500")
        return false
      } else {
        param.classList.remove("border-red-500")
      }
    }

    return true
  }

  async makeRequest(data: RequestData): Promise<Response> {
    const csrfToken = (document.querySelector('meta[name="csrf-token"]') as HTMLMetaElement)?.content

    return fetch('/api/proxy', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'X-CSRF-Token': csrfToken
      },
      body: JSON.stringify({
        endpoint: this.urlValue,
        method: this.methodValue,
        params: data
      })
    })
  }

  async handleResponse(response: Response, clientDuration: number) {
    const proxyResult = await response.json()

    // Use the actual API status code from the proxy envelope
    const actualStatus: number = proxyResult.status_code || response.status

    // On 429, keep the last successful response visible and open the signup popup
    if (actualStatus === 429 && this.modalIdValue) {
      if (this.hasResponse) {
        this.responseContainerTarget.classList.remove("hidden")
      }
      this.openRateLimitModal()
      return
    }

    // Show response container
    this.hideError()
    this.responseContainerTarget.classList.remove("hidden")

    const statusOk = actualStatus >= 200 && actualStatus < 300

    // Update status badge
    const statusClass = statusOk ? "bg-green-100 text-green-800" : "bg-red-100 text-red-800"
    this.statusBadgeTarget.className = `inline-flex items-center px-2.5 py-0.5 rounded text-xs font-medium ${statusClass}`
    this.statusBadgeTarget.textContent = actualStatus.toString()

    // Update response time (prefer server-measured, fall back to client)
    const displayTime = proxyResult.response_time_ms || clientDuration
    this.responseTimeTarget.textContent = `${displayTime}ms`

    // Display the actual API response body (unwrap proxy envelope)
    const body = proxyResult.data ?? proxyResult.error ?? proxyResult
    this.responseBodyTarget.textContent = JSON.stringify(body, null, 2)

    this.hasResponse = true
  }

  openRateLimitModal() {
    const modal = document.getElementById(this.modalIdValue)
    if (!modal) return

    modal.style.display = "flex"
    document.body.classList.add("overflow-hidden")

    const closeModal = () => {
      modal.style.display = "none"
      document.body.classList.remove("overflow-hidden")
      document.removeEventListener("keydown", escHandler)
    }
    const escHandler = (e: KeyboardEvent) => { if (e.key === "Escape") closeModal() }
    document.addEventListener("keydown", escHandler)
  }

  showLoading() {
    this.submitButtonTarget.disabled = true
    this.submitTextTarget.textContent = "Sending..."
    this.loadingTarget.classList.remove("hidden")
    this.hideError()
    this.responseContainerTarget.classList.add("hidden")
  }

  hideLoading() {
    this.submitButtonTarget.disabled = false
    this.submitTextTarget.textContent = "Send Request"
    this.loadingTarget.classList.add("hidden")
  }

  showError(message: string) {
    this.errorMessageTarget.textContent = message
    this.errorContainerTarget.classList.remove("hidden")
    this.responseContainerTarget.classList.add("hidden")
  }

  hideError() {
    this.errorContainerTarget.classList.add("hidden")
  }
}
