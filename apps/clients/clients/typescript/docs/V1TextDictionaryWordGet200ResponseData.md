# V1TextDictionaryWordGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**word** | **string** | The normalized (lowercased) word that was looked up | [optional] [default to undefined]
**phonetic** | **string** | IPA phonetic transcription of the word (may be omitted if unavailable) | [optional] [default to undefined]
**definitions** | **string** | One or more definitions for the word, each with partOfSpeech, definition, and an optional example | [optional] [default to undefined]
**definitions_partOfSpeech** | **string** | Grammatical category (e.g. noun, verb, adjective) | [optional] [default to undefined]
**definitions_definition** | **string** | Plain-text definition of the word | [optional] [default to undefined]
**definitions_example** | **string** | Example sentence using the word (may be omitted) | [optional] [default to undefined]
**synonyms** | **string** | List of words with similar meaning | [optional] [default to undefined]

## Example

```typescript
import { V1TextDictionaryWordGet200ResponseData } from 'api-client';

const instance: V1TextDictionaryWordGet200ResponseData = {
    word,
    phonetic,
    definitions,
    definitions_partOfSpeech,
    definitions_definition,
    definitions_example,
    synonyms,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
