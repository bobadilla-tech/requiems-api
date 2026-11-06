# HolidaysApi

All URIs are relative to *https://api.requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1PlacesHolidaysBatchPost**](#v1placesholidaysbatchpost) | **POST** /v1/places/holidays/batch | Batch Get Holidays|
|[**v1PlacesHolidaysGet**](#v1placesholidaysget) | **GET** /v1/places/holidays | Get Holidays|

# **v1PlacesHolidaysBatchPost**
> V1PlacesHolidaysBatchPost200Response v1PlacesHolidaysBatchPost(v1PlacesHolidaysBatchPostRequest)

Returns holidays for up to 50 (country, year) pairs in a single request. Each pair is processed independently — if one combination has no data, it returns found:false without failing the entire batch.

### Example

```typescript
import {
    HolidaysApi,
    Configuration,
    V1PlacesHolidaysBatchPostRequest
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new HolidaysApi(configuration);

let v1PlacesHolidaysBatchPostRequest: V1PlacesHolidaysBatchPostRequest; //

const { status, data } = await apiInstance.v1PlacesHolidaysBatchPost(
    v1PlacesHolidaysBatchPostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **v1PlacesHolidaysBatchPostRequest** | **V1PlacesHolidaysBatchPostRequest**|  | |


### Return type

**V1PlacesHolidaysBatchPost200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**400** | Malformed request body |  -  |
|**422** | queries is missing, empty, exceeds 50 items, or contains invalid country codes or years |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1PlacesHolidaysGet**
> V1PlacesHolidaysGet200Response v1PlacesHolidaysGet()

Returns a list of public holidays for the specified country and year

### Example

```typescript
import {
    HolidaysApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new HolidaysApi(configuration);

let country: string; //ISO 3166-1 alpha-2 country code (e.g., \"US\", \"GB\", \"DE\") (default to undefined)
let year: number; //Year for which to retrieve holidays (e.g., 2025) (default to undefined)

const { status, data } = await apiInstance.v1PlacesHolidaysGet(
    country,
    year
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **country** | [**string**] | ISO 3166-1 alpha-2 country code (e.g., \&quot;US\&quot;, \&quot;GB\&quot;, \&quot;DE\&quot;) | defaults to undefined|
| **year** | [**number**] | Year for which to retrieve holidays (e.g., 2025) | defaults to undefined|


### Return type

**V1PlacesHolidaysGet200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**400** | Missing or invalid country code or year parameter |  -  |
|**404** | No holidays found for the specified country and year |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

