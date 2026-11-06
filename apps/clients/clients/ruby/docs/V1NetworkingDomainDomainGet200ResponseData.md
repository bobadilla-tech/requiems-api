# OpenapiClient::V1NetworkingDomainDomainGet200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **domain** | **String** | The domain that was looked up | [optional] |
| **available** | **Boolean** | True when the domain appears to be unregistered (NS lookup returns NXDOMAIN). False when name servers are delegated. | [optional] |
| **dns_a** | **String** | IPv4 addresses (A records) | [optional] |
| **dns_aaaa** | **String** | IPv6 addresses (AAAA records) | [optional] |
| **dns_mx** | **String** | Mail exchange records, each with host and priority fields | [optional] |
| **dns_ns** | **String** | Authoritative name server hostnames | [optional] |
| **dns_txt** | **String** | TXT record values (SPF, DKIM, verification tokens, etc.) | [optional] |
| **dns_cname** | **String** | CNAME alias target, if the domain is an alias. Empty string when no alias exists. | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1NetworkingDomainDomainGet200ResponseData.new(
  domain: null,
  available: null,
  dns_a: null,
  dns_aaaa: null,
  dns_mx: null,
  dns_ns: null,
  dns_txt: null,
  dns_cname: null
)
```

