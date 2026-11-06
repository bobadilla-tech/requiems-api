# OpenapiClient::V1ValidationEmailPost200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **email** | **String** | The email address exactly as supplied in the request body; null when syntax is invalid | [optional] |
| **valid** | **Boolean** | Overall validity. True only when the address passes syntax validation and the domain has at least one MX record | [optional] |
| **syntax_valid** | **Boolean** | Whether the address is syntactically valid according to RFC 5322 | [optional] |
| **mx_valid** | **Boolean** | Whether the domain has at least one MX record, meaning it can receive email | [optional] |
| **disposable** | **Boolean** | Whether the address uses a known disposable or temporary email domain | [optional] |
| **normalized** | **String** | The canonical form of the address after normalization (lowercase, plus-tag removal, alias-domain resolution). Null when syntax is invalid | [optional] |
| **domain** | **String** | The domain part of the address (after @). Null when syntax is invalid | [optional] |
| **suggestion** | **String** | A corrected domain name when the supplied domain looks like a typo of a well-known provider (e.g. gmial.com → gmail.com). Null when no close match is found or the domain is already correct | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1ValidationEmailPost200ResponseData.new(
  email: null,
  valid: null,
  syntax_valid: null,
  mx_valid: null,
  disposable: null,
  normalized: null,
  domain: null,
  suggestion: null
)
```

