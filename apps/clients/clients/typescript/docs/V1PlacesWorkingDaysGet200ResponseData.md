# V1PlacesWorkingDaysGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**working_days** | **number** | Number of working days between the two dates (excluding weekends and optionally holidays) | [optional] [default to undefined]
**from** | **string** | Start date (echoed from request) | [optional] [default to undefined]
**to** | **string** | End date (echoed from request) | [optional] [default to undefined]
**country** | **string** | Country code (echoed from request, empty string if not provided) | [optional] [default to undefined]
**subdivision** | **string** | Subdivision code (echoed from request, empty string if not provided) | [optional] [default to undefined]

## Example

```typescript
import { V1PlacesWorkingDaysGet200ResponseData } from 'api-client';

const instance: V1PlacesWorkingDaysGet200ResponseData = {
    working_days,
    from,
    to,
    country,
    subdivision,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
