# V1SystemsTimezoneFromIpIpGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ip** | **string** | The IP address as provided | [optional] [default to undefined]
**city** | **string** | City resolved from the IP | [optional] [default to undefined]
**country_code** | **string** | ISO 3166-1 alpha-2 country code | [optional] [default to undefined]
**timezone** | **string** | IANA timezone identifier. Null when timezone resolution fails. | [optional] [default to undefined]
**utc_offset** | **string** | UTC offset string. Null when timezone resolution fails. | [optional] [default to undefined]
**dst_active** | **string** | True when daylight saving time is currently active. Null when timezone resolution fails. | [optional] [default to undefined]
**current_time** | **string** | Current local time in RFC3339 format | [optional] [default to undefined]

## Example

```typescript
import { V1SystemsTimezoneFromIpIpGet200ResponseData } from 'api-client';

const instance: V1SystemsTimezoneFromIpIpGet200ResponseData = {
    ip,
    city,
    country_code,
    timezone,
    utc_offset,
    dst_active,
    current_time,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
