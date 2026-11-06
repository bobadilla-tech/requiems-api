# HoroscopeApi

All URIs are relative to *https://api.requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1EntertainmentHoroscopeSignGet**](#v1entertainmenthoroscopesignget) | **GET** /v1/entertainment/horoscope/{sign} | Get Daily Horoscope|

# **v1EntertainmentHoroscopeSignGet**
> V1EntertainmentHoroscopeSignGet200Response v1EntertainmentHoroscopeSignGet()

Returns a daily horoscope reading for the specified zodiac sign.

### Example

```typescript
import {
    HoroscopeApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new HoroscopeApi(configuration);

let sign: string; //Zodiac sign (case-insensitive). Supported values: aries, taurus, gemini, cancer, leo, virgo, libra, scorpio, sagittarius, capricorn, aquarius, pisces (default to undefined)

const { status, data } = await apiInstance.v1EntertainmentHoroscopeSignGet(
    sign
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **sign** | [**string**] | Zodiac sign (case-insensitive). Supported values: aries, taurus, gemini, cancer, leo, virgo, libra, scorpio, sagittarius, capricorn, aquarius, pisces | defaults to undefined|


### Return type

**V1EntertainmentHoroscopeSignGet200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

