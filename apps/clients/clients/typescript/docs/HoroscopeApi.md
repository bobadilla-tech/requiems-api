# HoroscopeApi

All URIs are relative to *https://requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1EntertainmentHoroscopeBatchPost**](#v1entertainmenthoroscopebatchpost) | **POST** /v1/entertainment/horoscope/batch | Batch Daily Horoscopes|
|[**v1EntertainmentHoroscopeSignGet**](#v1entertainmenthoroscopesignget) | **GET** /v1/entertainment/horoscope/{sign} | Get Daily Horoscope|

# **v1EntertainmentHoroscopeBatchPost**
> V1EntertainmentHoroscopeBatchPost200Response v1EntertainmentHoroscopeBatchPost(v1EntertainmentHoroscopeBatchPostRequest)

Returns up to 12 daily horoscopes in a single request.

### Example

```typescript
import {
    HoroscopeApi,
    Configuration,
    V1EntertainmentHoroscopeBatchPostRequest
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new HoroscopeApi(configuration);

let v1EntertainmentHoroscopeBatchPostRequest: V1EntertainmentHoroscopeBatchPostRequest; //

const { status, data } = await apiInstance.v1EntertainmentHoroscopeBatchPost(
    v1EntertainmentHoroscopeBatchPostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **v1EntertainmentHoroscopeBatchPostRequest** | **V1EntertainmentHoroscopeBatchPostRequest**|  | |


### Return type

**V1EntertainmentHoroscopeBatchPost200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**400** | Invalid JSON, malformed request body, or unexpected field types. |  -  |
|**422** | The signs array is missing, empty, or contains more than 12 items, or includes unsupported sign values. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

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

