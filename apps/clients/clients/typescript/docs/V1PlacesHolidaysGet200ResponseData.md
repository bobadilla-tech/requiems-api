# V1PlacesHolidaysGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**country** | **string** | ISO 3166-1 alpha-2 country code | [optional] [default to undefined]
**year** | **number** | Year for which holidays are returned | [optional] [default to undefined]
**holidays** | **Array&lt;any&gt;** | Array of holiday objects | [optional] [default to undefined]
**holidays_date** | **string** | Holiday date in YYYY-MM-DD format | [optional] [default to undefined]
**holidays_name** | **string** | Name of the holiday | [optional] [default to undefined]
**total** | **number** | Total number of holidays for the country/year | [optional] [default to undefined]

## Example

```typescript
import { V1PlacesHolidaysGet200ResponseData } from 'api-client';

const instance: V1PlacesHolidaysGet200ResponseData = {
    country,
    year,
    holidays,
    holidays_date,
    holidays_name,
    total,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
