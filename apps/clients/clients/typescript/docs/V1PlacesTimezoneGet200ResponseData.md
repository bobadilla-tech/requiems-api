# V1PlacesTimezoneGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**timezone** | **string** | IANA timezone identifier (e.g. \&quot;Europe/London\&quot;, \&quot;Asia/Tokyo\&quot;) | [optional] [default to undefined]
**offset** | **string** | UTC offset in +HH:MM or -HH:MM format (e.g. \&#39;+05:30\&#39;, \&#39;-05:00\&#39;) | [optional] [default to undefined]
**current_time** | **string** | Current time in UTC, formatted as RFC 3339 (e.g. \&quot;2024-12-15T14:30:00Z\&quot;) | [optional] [default to undefined]
**is_dst** | **boolean** | Whether the location is currently observing daylight saving time | [optional] [default to undefined]

## Example

```typescript
import { V1PlacesTimezoneGet200ResponseData } from 'api-client';

const instance: V1PlacesTimezoneGet200ResponseData = {
    timezone,
    offset,
    current_time,
    is_dst,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
