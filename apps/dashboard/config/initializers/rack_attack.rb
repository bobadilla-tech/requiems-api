# frozen_string_literal: true

class Rack::Attack
  Rack::Attack.cache.store = Rails.cache

  # Throttle API playground requests for anonymous users: 1 req/min
  throttle("api_proxy/ip", limit: 1, period: 1.minute) do |req|
    if req.path == "/api/proxy" && req.post?
      # Only throttle anonymous users — authenticated users have a separate throttle
      req.ip unless req.env["warden"]&.user
    end
  end

  # Throttle API playground requests for free authenticated users: 4 req/min
  # Paying users (any plan other than free) are not rate limited in the playground
  throttle("api_proxy/user", limit: 4, period: 1.minute) do |req|
    if req.path == "/api/proxy" && req.post?
      user = req.env["warden"]&.user
      "user:#{user.id}" if user && user.current_plan == "free"
    end
  end

  # Throttle tool demo requests for anonymous users: 5 req/min
  throttle("tool_demos/ip", limit: 5, period: 1.minute) do |req|
    if req.path.start_with?("/tools/demos/") && req.post?
      # Only throttle anonymous users — authenticated users have a separate throttle
      req.ip unless req.env["warden"]&.user
    end
  end

  # Throttle tool demo requests for free authenticated users: 15 req/min
  # Paying users (any plan other than free) are not rate limited on tool demos
  throttle("tool_demos/user", limit: 15, period: 1.minute) do |req|
    if req.path.start_with?("/tools/demos/") && req.post?
      user = req.env["warden"]&.user
      "user:#{user.id}" if user && user.current_plan == "free"
    end
  end

  # Throttle login attempts by IP address
  throttle("logins/ip", limit: 5, period: 20.seconds) do |req|
    if req.path == "/users/sign_in" && req.post?
      req.ip
    end
  end

  # Optional: Throttle login attempts by email
  throttle("logins/email", limit: 5, period: 20.seconds) do |req|
    if req.path == "/users/sign_in" && req.post?
      # Normalize email
      req.params["user"]&.dig("email")&.to_s&.downcase&.gsub(/\s+/, "")
    end
  end

  ### Custom Throttle Response ###

  # Customize the response when throttled
  self.throttled_responder = lambda do |request|
    match_data = request.env["rack.attack.match_data"]
    match_name = request.env["rack.attack.matched"]
    now = match_data[:epoch_time]
    retry_after = match_data[:period] - (now % match_data[:period])

    headers = {
      "RateLimit-Limit" => match_data[:limit].to_s,
      "RateLimit-Remaining" => "0",
      "RateLimit-Reset" => (now + retry_after).to_s,
      "Content-Type" => "application/json"
    }

    message = if match_name == "api_proxy/ip" || match_name == "tool_demos/ip"
      "You're testing too fast. Create a free account to get a higher limit."
    else
      "Too many requests. Please slow down and try again shortly."
    end

    body = {
      error: "Rate limit exceeded",
      message: message,
      retry_after: retry_after
    }.to_json

    [ 429, headers, [ body ] ]
  end

  ActiveSupport::Notifications.subscribe("throttle.rack_attack") do |_name, _start, _finish, _request_id, payload|
    req = payload[:request]
    Rails.logger.warn("[Rack::Attack][Throttled] IP: #{req.ip} | Path: #{req.path}")
  end
end
