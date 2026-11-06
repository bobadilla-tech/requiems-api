# OpenapiClient::V1NetworkingDisposableDomainsGet200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **domains** | **Array&lt;Object&gt;** | Array of domain names | [optional] |
| **total** | **Integer** | Total number of domains in blocklist | [optional] |
| **page** | **Integer** | Current page number | [optional] |
| **per_page** | **Integer** | Number of items per page | [optional] |
| **has_more** | **Boolean** | Whether there are more pages available | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1NetworkingDisposableDomainsGet200ResponseData.new(
  domains: null,
  total: null,
  page: null,
  per_page: null,
  has_more: null
)
```

