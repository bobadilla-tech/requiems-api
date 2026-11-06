# V1TextSpellcheckPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**corrected** | **string** | The full input text with all misspelled words replaced by their suggested corrections | [optional] [default to undefined]
**corrections** | **string** | List of individual corrections. Each item contains: original (the misspelled word), suggested (the correction), and position (0-based character offset in the original text) | [optional] [default to undefined]

## Example

```typescript
import { V1TextSpellcheckPost200ResponseData } from 'api-client';

const instance: V1TextSpellcheckPost200ResponseData = {
    corrected,
    corrections,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
