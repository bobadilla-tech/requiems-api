# Caddy

Caddy is the authenticated origin boundary for the public Go API. Cloudflare
Origin Pulls presents a client certificate signed by the configured
Cloudflare origin-pull CA; direct TLS requests without that certificate fail
with the TLS `certificate_required` alert.

Caddy forwards the API to the Kamal Go proxy on the host network. The Go API
still performs all API-key authentication and rate/quota/usage enforcement.
There is no `X-Backend-Secret` matcher.

Keep the Cloudflare proxy enabled and test both public routing and direct-origin
certificate rejection after Caddy or Kamal changes.

