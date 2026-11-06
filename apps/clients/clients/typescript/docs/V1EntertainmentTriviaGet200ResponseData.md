# V1EntertainmentTriviaGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**question** | **string** | The trivia question text | [optional] [default to undefined]
**_options** | **string** | Four multiple-choice answer options | [optional] [default to undefined]
**answer** | **string** | The correct answer — always one of the values in options | [optional] [default to undefined]
**category** | **string** | The category the question belongs to | [optional] [default to undefined]
**difficulty** | **string** | The difficulty level of the question (easy, medium, or hard) | [optional] [default to undefined]

## Example

```typescript
import { V1EntertainmentTriviaGet200ResponseData } from 'api-client';

const instance: V1EntertainmentTriviaGet200ResponseData = {
    question,
    _options,
    answer,
    category,
    difficulty,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
