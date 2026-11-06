# V1SystemsBusinessCalendarCountryGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**country_code** | **string** | ISO 3166-1 alpha-2 country code echoed from the path parameter | [optional] [default to undefined]
**year** | **number** | Year the calendar covers | [optional] [default to undefined]
**month** | **string** | Month scope. Present only when the month query parameter was provided. | [optional] [default to undefined]
**working_days** | **number** | Number of working days (excluding weekends and public holidays) in the scope | [optional] [default to undefined]
**total_days** | **string** | Total calendar days in the month. Present only when month is provided. | [optional] [default to undefined]
**weekend_days** | **string** | Number of weekend days in the month. Present only when month is provided. | [optional] [default to undefined]
**holidays** | **Array&lt;any&gt;** | List of public holidays in the scope. Each item: {date: YYYY-MM-DD, name: string, type: national} | [optional] [default to undefined]
**holiday_count** | **number** | Count of holidays in the holidays array | [optional] [default to undefined]
**next_holiday** | **string** | Next public holiday within 90 days from today. Shape: {date, name, type} | [optional] [default to undefined]

## Example

```typescript
import { V1SystemsBusinessCalendarCountryGet200ResponseData } from 'api-client';

const instance: V1SystemsBusinessCalendarCountryGet200ResponseData = {
    country_code,
    year,
    month,
    working_days,
    total_days,
    weekend_days,
    holidays,
    holiday_count,
    next_holiday,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
