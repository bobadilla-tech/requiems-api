# V1EntertainmentHoroscopeBatchPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**results** | **Array&lt;any&gt;** | List of daily horoscopes returned in the same order as the input signs. | [optional] [default to undefined]
**results_sign** | **string** | Zodiac sign the horoscope corresponds to. | [optional] [default to undefined]
**results_date** | **string** | Date the horoscope is for, in YYYY-MM-DD format. | [optional] [default to undefined]
**results_horoscope** | **string** | Daily horoscope reading for the given sign. | [optional] [default to undefined]
**results_lucky_number** | **number** | Lucky number for the given sign on that date. | [optional] [default to undefined]
**results_mood** | **string** | Predicted mood for the given sign on that date. | [optional] [default to undefined]
**total** | **number** | Total number of horoscopes returned. | [optional] [default to undefined]

## Example

```typescript
import { V1EntertainmentHoroscopeBatchPost200ResponseData } from 'api-client';

const instance: V1EntertainmentHoroscopeBatchPost200ResponseData = {
    results,
    results_sign,
    results_date,
    results_horoscope,
    results_lucky_number,
    results_mood,
    total,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
