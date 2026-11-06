# V1PlacesTimezoneBatchPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**results** | **Array&lt;any&gt;** | Per-item results in input order. | [optional] [default to undefined]
**results_info_timezone** | **string** | IANA timezone identifier. | [optional] [default to undefined]
**results_info_offset** | **string** | UTC offset in +HH:MM or -HH:MM format. | [optional] [default to undefined]
**results_info_current_time** | **string** | Current time in UTC (RFC 3339). | [optional] [default to undefined]
**results_info_is_dst** | **boolean** | Whether daylight saving time is currently active. | [optional] [default to undefined]
**results_error** | **string** | Error message if this item failed (omitted on success). | [optional] [default to undefined]
**total** | **number** | Total number of items processed. | [optional] [default to undefined]

## Example

```typescript
import { V1PlacesTimezoneBatchPost200ResponseData } from 'api-client';

const instance: V1PlacesTimezoneBatchPost200ResponseData = {
    results,
    results_info_timezone,
    results_info_offset,
    results_info_current_time,
    results_info_is_dst,
    results_error,
    total,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
