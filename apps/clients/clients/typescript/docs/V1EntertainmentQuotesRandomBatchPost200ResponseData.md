# V1EntertainmentQuotesRandomBatchPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**results** | **Array&lt;any&gt;** | Array of random quotes, one per requested count. | [optional] [default to undefined]
**results_id** | **number** | Unique identifier for the quote. | [optional] [default to undefined]
**results_text** | **string** | The quote text. | [optional] [default to undefined]
**results_author** | **string** | Name of the person who said or wrote the quote. May be empty. | [optional] [default to undefined]
**total** | **number** | Total number of quotes returned. Always equals the requested count. | [optional] [default to undefined]

## Example

```typescript
import { V1EntertainmentQuotesRandomBatchPost200ResponseData } from 'api-client';

const instance: V1EntertainmentQuotesRandomBatchPost200ResponseData = {
    results,
    results_id,
    results_text,
    results_author,
    total,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
