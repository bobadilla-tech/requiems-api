# OpenapiClient::V1NetworkingWhoisDomainGet200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **domain** | **String** | The domain name that was looked up | [optional] |
| **registrar** | **String** | The name of the registrar holding the domain registration | [optional] |
| **name_servers** | **String** | List of authoritative name servers for the domain | [optional] |
| **status** | **String** | EPP status codes for the domain (e.g. clientTransferProhibited) | [optional] |
| **created_date** | **String** | Date the domain was first registered (ISO 8601) | [optional] |
| **updated_date** | **String** | Date the domain record was last updated (ISO 8601) | [optional] |
| **expiry_date** | **String** | Date the domain registration expires (ISO 8601) | [optional] |
| **dnssec** | **Boolean** | True when DNSSEC is enabled for the domain | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1NetworkingWhoisDomainGet200ResponseData.new(
  domain: null,
  registrar: null,
  name_servers: null,
  status: null,
  created_date: null,
  updated_date: null,
  expiry_date: null,
  dnssec: null
)
```

