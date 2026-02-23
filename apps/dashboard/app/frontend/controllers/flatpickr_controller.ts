import { Controller } from "@hotwired/stimulus"
import flatpickr from "flatpickr"

type FlatpickrOptions = NonNullable<Parameters<typeof flatpickr>[1]>
type FlatpickrInstance = ReturnType<typeof flatpickr>

// Connects to data-controller="flatpickr"
export default class extends Controller {
  static values = {
    mode: { type: String, default: "single" }, // single, range, multiple
    enableTime: { type: Boolean, default: false },
    dateFormat: { type: String, default: "Y-m-d" },
    minDate: String,
    maxDate: String,
    defaultDate: String
  }

  declare readonly modeValue: string
  declare readonly enableTimeValue: boolean
  declare readonly dateFormatValue: string
  declare readonly minDateValue: string
  declare readonly maxDateValue: string
  declare readonly defaultDateValue: string
  declare readonly hasMinDateValue: boolean
  declare readonly hasMaxDateValue: boolean
  declare readonly hasDefaultDateValue: boolean

  private picker?: FlatpickrInstance

  connect() {
    const options: FlatpickrOptions = {
      mode: this.modeValue as FlatpickrOptions["mode"],
      enableTime: this.enableTimeValue,
      dateFormat: this.dateFormatValue
    }

    if (this.hasMinDateValue) {
      options.minDate = this.minDateValue
    }

    if (this.hasMaxDateValue) {
      options.maxDate = this.maxDateValue
    }

    if (this.hasDefaultDateValue) {
      options.defaultDate = this.defaultDateValue
    }

    // Initialize flatpickr on this element
    this.picker = flatpickr(this.element as HTMLElement, options)
  }

  disconnect() {
    if (this.picker) {
      const instance = Array.isArray(this.picker) ? this.picker[0] : this.picker
      instance?.destroy()
    }
  }
}
