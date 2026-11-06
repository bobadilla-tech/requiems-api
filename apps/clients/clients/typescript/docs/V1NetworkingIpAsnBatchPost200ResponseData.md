# V1NetworkingIpAsnBatchPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**results** | **Array&lt;any&gt;** | Per-item results in input order. | [optional] [default to undefined]
**results_ip** | **string** | The IP address that was looked up. | [optional] [default to undefined]
**results_result** | **object** | ASN data (omitted on error). Same fields as the single-item endpoint. | [optional] [default to undefined]
**results_error** | **string** | Error message if the lookup failed (omitted on success). | [optional] [default to undefined]
**total** | **number** | Total number of items processed. | [optional] [default to undefined]

## Example

```typescript
import { V1NetworkingIpAsnBatchPost200ResponseData } from 'api-client';

const instance: V1NetworkingIpAsnBatchPost200ResponseData = {
    results,
    results_ip,
    results_result,
    results_error,
    total,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
