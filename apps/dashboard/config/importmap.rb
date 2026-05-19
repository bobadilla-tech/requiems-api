# frozen_string_literal: true

pin "application"
pin "@hotwired/turbo-rails", to: "turbo.min.js"
pin "@hotwired/stimulus", to: "stimulus.min.js"
pin "@hotwired/stimulus-loading", to: "stimulus-loading.js"
pin_all_from "app/javascript/controllers", under: "controllers"

pin "chartkick", to: "https://cdn.jsdelivr.net/npm/chartkick@5.0.1/dist/chartkick.esm.js"
pin "Chart.bundle", to: "https://cdn.jsdelivr.net/npm/chart.js@4.4.1/+esm"

pin "highlight.js", to: "https://cdn.jsdelivr.net/npm/highlight.js@11.9.0/+esm"

pin "flatpickr", to: "https://cdn.jsdelivr.net/npm/flatpickr@4.6.13/+esm"
