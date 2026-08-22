# frozen_string_literal: true

# Configures Rails' own X-Forwarded-For-aware proxy list for framework-level
# consumers of request.remote_ip (e.g. Rack::Attack's default #req.ip key).
#
# This does NOT by itself give ApiProxyController/ToolDemosController the
# "only trust a forwarded header from a verified hop" guarantee they need —
# see app/lib/trusted_proxy.rb for why, and for the resolver those two
# controllers actually use.
Rails.application.config.to_prepare do
  Rails.application.config.action_dispatch.trusted_proxies = TrustedProxy::TRUSTED_RANGES
end
