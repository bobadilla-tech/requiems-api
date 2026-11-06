# OpenapiClient::V1NetworkingMxDomainGet200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **domain** | **String** | The domain that was queried | [optional] |
| **records** | **Array&lt;Object&gt;** | List of MX records, sorted by priority ascending (lowest number &#x3D; highest priority) | [optional] |
| **records_host** | **String** | Fully-qualified hostname of the mail server (typically ends with a trailing dot) | [optional] |
| **records_priority** | **Integer** | MX priority value. Lower values have higher delivery priority per RFC 5321. | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1NetworkingMxDomainGet200ResponseData.new(
  domain: null,
  records: null,
  records_host: null,
  records_priority: null
)
```

