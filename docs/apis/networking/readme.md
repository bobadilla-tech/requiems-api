# Networking & Internet APIs

## Overview

IP intelligence, domain data, and network lookup tools for security and
infrastructure work.

## Endpoints

### [IP Geolocation](./ip-geolocation.md) - ✅ MVP

Get geolocation data for any IP address: country, city, ISP, and VPN detection.

- **Status:** mvp
- **Endpoints:** `GET /v1/networking/ip` (caller IP),
  `GET /v1/networking/ip/{ip}`
- **Credit Cost:** 1

### [ASN Lookup](./asn-lookup.md) - ✅ MVP

Look up ASN, organization, ISP, and network route information for any IP
address.

- **Status:** mvp
- **Endpoint:** `GET /v1/networking/ip/asn/{ip}`
- **Credit Cost:** 1

### [VPN & Proxy Detection](./vpn-detection.md) - ✅ MVP

Detect if an IP belongs to a VPN, proxy, Tor exit node, or hosting provider.

- **Status:** mvp
- **Endpoint:** `GET /v1/networking/ip/vpn/{ip}`
- **Credit Cost:** 1

### [WHOIS Lookup](./whois.md) - ✅ MVP

Get domain registration details: registrar, name servers, creation/expiry dates.

- **Status:** mvp
- **Endpoint:** `GET /v1/networking/whois/{domain}`
- **Credit Cost:** 1

### [Domain Info](./domain-info.md) - ✅ MVP

Look up DNS records (A, AAAA, MX, NS, TXT, CNAME) and check domain availability.

- **Status:** mvp
- **Endpoint:** `GET /v1/networking/domain/{domain}`
- **Credit Cost:** 1

### [MX Lookup](./mx-lookup.md) - ✅ MVP

Look up MX records for any domain, sorted by priority.

- **Status:** mvp
- **Endpoint:** `GET /v1/networking/mx/{domain}`
- **Credit Cost:** 1

### [Disposable Domain Checker](./disposable-email.md) - ✅ MVP

Check whether an email domain belongs to a known disposable/temporary provider.

- **Status:** mvp
- **Endpoints:**
  - `POST /v1/networking/disposable/check`
  - `POST /v1/networking/disposable/batch`
  - `GET /v1/networking/disposable/domain/{domain}`
  - `GET /v1/networking/disposable/domains`
  - `GET /v1/networking/disposable/stats`
- **Credit Cost:** 1

## Category Statistics

- Total Endpoints: 7
- Live: 7
