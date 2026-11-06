# V1EntertainmentTriviaBatchPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**results** | **Array&lt;any&gt;** | List of trivia questions returned in the same order as the input filters. | [optional] [default to undefined]
**results_category** | **string** | The topic filter used to retrieve the trivia question. | [optional] [default to undefined]
**results_difficulty** | **string** | The complexity level filter used to retrieve the trivia question. | [optional] [default to undefined]
**results_error** | **string** | Describes why the question could not be retrieved for this filter. Omitted if the query was successful. | [optional] [default to undefined]
**results_data** | **object** | The trivia question returned for this filter. Empty if the query failed. | [optional] [default to undefined]
**results_data_question** | **string** | The trivia question text. | [optional] [default to undefined]
**results_data_options** | **Array&lt;any&gt;** | List of possible answers for the question. | [optional] [default to undefined]
**results_data_answer** | **string** | The correct answer to the trivia question. | [optional] [default to undefined]
**results_data_category** | **string** | The category the trivia question belongs to. | [optional] [default to undefined]
**results_data_difficulty** | **string** | The difficulty level of the trivia question. | [optional] [default to undefined]
**total** | **number** | Total number of trivia results returned. | [optional] [default to undefined]

## Example

```typescript
import { V1EntertainmentTriviaBatchPost200ResponseData } from 'api-client';

const instance: V1EntertainmentTriviaBatchPost200ResponseData = {
    results,
    results_category,
    results_difficulty,
    results_error,
    results_data,
    results_data_question,
    results_data_options,
    results_data_answer,
    results_data_category,
    results_data_difficulty,
    total,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
