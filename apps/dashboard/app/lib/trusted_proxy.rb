# frozen_string_literal: true

# Mirrors apps/api/platform/httpx/trustedproxy.go's ClientIP: only trusts
# X-Forwarded-For when the immediate connection (request.remote_addr, the raw
# socket peer — not Rails' own request.remote_ip, which does NOT gate on this)
# actually comes from a trusted hop (the local Caddy or Cloudflare's edge).
#
# Rails' own ActionDispatch::RemoteIp (config.action_dispatch.trusted_proxies)
# does not provide this guarantee by itself: it filters known-proxy IPs out of
# whatever X-Forwarded-For claims, but never checks that the actual TCP peer
# is one of those proxies — a request that reaches Rails directly can still
# set X-Forwarded-For to any value RemoteIp doesn't happen to recognize as a
# trusted proxy, and RemoteIp will use it. That gap is exactly what
# CF-Connecting-IP was being read unconditionally for before this fix.
class TrustedProxy
  # Static snapshot, not fetched dynamically — same "placeholder, revisit"
  # caveat as the Go-side list this mirrors. Kept in sync manually.
  CLOUDFLARE_IP_RANGES = %w[
    173.245.48.0/20 103.21.244.0/22 103.22.200.0/22 103.31.4.0/22
    141.101.64.0/18 108.162.192.0/18 190.93.240.0/20 188.114.96.0/20
    197.234.240.0/22 198.41.128.0/17 162.158.0.0/15 104.16.0.0/13
    104.24.0.0/14 172.64.0.0/13 131.0.72.0/22
    2400:cb00::/32 2606:4700::/32 2803:f800::/32 2405:b500::/32
    2405:8100::/32 2a06:98c0::/29 2c0f:f248::/32
  ].freeze

  # These are the Docker/Kamal peer networks that can reach Rails through the
  # Caddy/Kamal proxy path. Do not use Rails' broad private-network defaults:
  # an unconfigured private peer must not gain the ability to override the
  # caller IP with X-Forwarded-For.
  CADDY_PEER_RANGES = %w[
    127.0.0.1/32
    172.18.0.0/16
    172.20.0.0/16
  ].freeze

  TRUSTED_RANGES = (
    CADDY_PEER_RANGES.map { |cidr| IPAddr.new(cidr) } +
    CLOUDFLARE_IP_RANGES.map { |cidr| IPAddr.new(cidr) }
  ).freeze

  # Returns the caller's real IP: X-Forwarded-For's first hop when
  # request.remote_addr is a trusted proxy, otherwise remote_addr itself —
  # ignoring any forwarded header a direct, untrusted connection sent.
  def self.client_ip(request)
    remote_addr = request.remote_addr

    begin
      trusted = TRUSTED_RANGES.any? { |range| range.include?(IPAddr.new(remote_addr)) }
    rescue IPAddr::Error
      trusted = false
    end

    return remote_addr unless trusted

    xff = request.headers["X-Forwarded-For"]
    return xff.split(",").first.strip if xff.present?

    remote_addr
  end
end
