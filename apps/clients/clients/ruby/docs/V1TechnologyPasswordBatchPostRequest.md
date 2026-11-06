# OpenapiClient::V1TechnologyPasswordBatchPostRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **items** | **String** | List of password generation requests (1–50 items). |  |
| **items_length** | **Integer** | Password length (8–128). Defaults to 16 when omitted. | [optional] |
| **items_uppercase** | **Boolean** | Include uppercase letters (A–Z). Defaults to false. | [optional] |
| **items_numbers** | **Boolean** | Include numbers (0–9). Defaults to false. | [optional] |
| **items_symbols** | **Boolean** | Include special characters. Defaults to false. | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1TechnologyPasswordBatchPostRequest.new(
  items: null,
  items_length: null,
  items_uppercase: null,
  items_numbers: null,
  items_symbols: null
)
```

