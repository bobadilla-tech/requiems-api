# frozen_string_literal: true

# Application-wide constants
# These constants are loaded once when the application starts

# Observer emails for form submissions and notifications
# Can be configured via OBSERVER_EMAILS environment variable (comma-separated)
# Example: OBSERVER_EMAILS=eliaz@bobadilla.tech,support@bobadilla.tech
OBSERVER_EMAILS = ENV.fetch("OBSERVER_EMAILS", "eliaz@bobadilla.tech,ale@bobadilla.tech").split(",").map(&:strip)
