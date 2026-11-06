# OpenapiClient::V1TextThesaurusWordGet200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **word** | **String** | The normalized (lowercased) word that was looked up | [optional] |
| **synonyms** | **String** | List of words with similar meaning | [optional] |
| **antonyms** | **String** | List of words with opposite meaning | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1TextThesaurusWordGet200ResponseData.new(
  word: null,
  synonyms: null,
  antonyms: null
)
```

