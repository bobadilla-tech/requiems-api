# OpenapiClient::V1FinanceIbanBatchPostRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **numbers** | **Array&lt;Object&gt;** | Array of iban numbers to validate (min: 1, max: 50). |  |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1FinanceIbanBatchPostRequest.new(
  numbers: [&quot;GB29NWBK60161331926819&quot;, &quot;DE89370400440532013000&quot;]
)
```

