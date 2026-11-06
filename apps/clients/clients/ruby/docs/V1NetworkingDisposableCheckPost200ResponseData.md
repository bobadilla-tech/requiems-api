# OpenapiClient::V1NetworkingDisposableCheckPost200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **email** | **String** | The email address that was checked | [optional] |
| **is_disposable** | **Boolean** | Whether the email uses a disposable domain | [optional] |
| **domain** | **String** | The domain part of the email address | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1NetworkingDisposableCheckPost200ResponseData.new(
  email: null,
  is_disposable: null,
  domain: null
)
```

