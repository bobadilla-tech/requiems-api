# V1NetworkingWhoisBatchPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**results** | **Array&lt;any&gt;** | One entry per domain, in the same order as the input | [optional] [default to undefined]
**results_domain** | **string** | Domain name requested in the batch | [optional] [default to undefined]
**results_found** | **boolean** | False when WHOIS data could not be found | [optional] [default to undefined]
**results_error** | **string** | Error message when found is false | [optional] [default to undefined]
**results_data** | **object** | WHOIS information for the domain when found is true | [optional] [default to undefined]
**total** | **number** | Total number of results returned | [optional] [default to undefined]

## Example

```typescript
import { V1NetworkingWhoisBatchPost200ResponseData } from 'api-client';

const instance: V1NetworkingWhoisBatchPost200ResponseData = {
    results,
    results_domain,
    results_found,
    results_error,
    results_data,
    total,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
