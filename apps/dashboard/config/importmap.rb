# frozen_string_literal: true

pin "application"
pin "@hotwired/turbo-rails", to: "turbo.min.js"
pin "@hotwired/stimulus", to: "stimulus.min.js"
pin "@hotwired/stimulus-loading", to: "stimulus-loading.js"
pin_all_from "app/javascript/controllers", under: "controllers"

pin "chartkick", to: "https://esm.sh/chartkick@5.0.1?target=es2020"
pin "Chart.bundle", to: "https://esm.sh/chart.js@4.4.1?bundle&target=es2020"

pin "highlight.js", to: "https://esm.sh/highlight.js@11.9.0?target=es2020"

pin "flatpickr", to: "https://esm.sh/flatpickr@4.6.13?target=es2020"
