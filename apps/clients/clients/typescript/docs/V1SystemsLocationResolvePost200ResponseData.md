# V1SystemsLocationResolvePost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**address** | **string** | Normalized full address returned by the geocoder | [optional] [default to undefined]
**city** | **string** | City name | [optional] [default to undefined]
**country** | **string** | Full country name | [optional] [default to undefined]
**country_code** | **string** | ISO 3166-1 alpha-2 country code | [optional] [default to undefined]
**coordinates** | **object** | Resolved lat/lng coordinates | [optional] [default to undefined]
**timezone** | **string** | IANA timezone identifier (e.g. Europe/Berlin). Null when timezone resolution fails. | [optional] [default to undefined]
**utc_offset** | **string** | UTC offset string (e.g. +02:00). Null when timezone resolution fails. | [optional] [default to undefined]
**current_time** | **string** | Current local time in RFC3339 format. Null when timezone resolution fails. | [optional] [default to undefined]
**is_holiday_today** | **boolean** | True when today is a public holiday in the resolved country | [optional] [default to undefined]
**working_days_this_month** | **number** | Number of working days (excluding weekends and public holidays) in the current calendar month | [optional] [default to undefined]
**next_holiday** | **string** | Next public holiday within 90 days. Shape: {date: YYYY-MM-DD, name: string, type: national} | [optional] [default to undefined]
**flags** | **Array&lt;any&gt;** | Non-fatal warnings. Possible values: timezone_unavailable, calendar_unavailable | [optional] [default to undefined]

## Example

```typescript
import { V1SystemsLocationResolvePost200ResponseData } from 'api-client';

const instance: V1SystemsLocationResolvePost200ResponseData = {
    address,
    city,
    country,
    country_code,
    coordinates,
    timezone,
    utc_offset,
    current_time,
    is_holiday_today,
    working_days_this_month,
    next_holiday,
    flags,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
