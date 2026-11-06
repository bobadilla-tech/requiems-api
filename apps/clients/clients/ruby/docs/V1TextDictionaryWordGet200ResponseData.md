# OpenapiClient::V1TextDictionaryWordGet200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **word** | **String** | The normalized (lowercased) word that was looked up | [optional] |
| **phonetic** | **String** | IPA phonetic transcription of the word (may be omitted if unavailable) | [optional] |
| **definitions** | **String** | One or more definitions for the word, each with partOfSpeech, definition, and an optional example | [optional] |
| **definitions_part_of_speech** | **String** | Grammatical category (e.g. noun, verb, adjective) | [optional] |
| **definitions_definition** | **String** | Plain-text definition of the word | [optional] |
| **definitions_example** | **String** | Example sentence using the word (may be omitted) | [optional] |
| **synonyms** | **String** | List of words with similar meaning | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1TextDictionaryWordGet200ResponseData.new(
  word: null,
  phonetic: null,
  definitions: null,
  definitions_part_of_speech: null,
  definitions_definition: null,
  definitions_example: null,
  synonyms: null
)
```

