# OpenapiClient::V1TextNormalizePost200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **original** | **String** | The email address exactly as supplied in the request body | [optional] |
| **normalized** | **String** | The canonical form of the address after all transformations | [optional] |
| **local** | **String** | The local part (before @) of the normalized address | [optional] |
| **domain** | **String** | The domain part (after @) of the normalized address | [optional] |
| **changes** | **Array&lt;Object&gt;** | Ordered list of transformations applied. Possible values: lowercased, trimmed_whitespace, removed_dots, removed_plus_tag, canonicalised_domain. Empty array when no changes were needed. | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1TextNormalizePost200ResponseData.new(
  original: null,
  normalized: null,
  local: null,
  domain: null,
  changes: null
)
```

