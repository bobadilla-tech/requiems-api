# GlobalDataApi

All URIs are relative to *https://requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1SystemsBusinessCalendarCountryGet**](#v1systemsbusinesscalendarcountryget) | **GET** /v1/systems/business-calendar/{country} | Business Calendar|
|[**v1SystemsLocationResolvePost**](#v1systemslocationresolvepost) | **POST** /v1/systems/location/resolve | Resolve Location|
|[**v1SystemsTimezoneFromIpIpGet**](#v1systemstimezonefromipipget) | **GET** /v1/systems/timezone/from-ip/{ip} | Timezone from IP|

# **v1SystemsBusinessCalendarCountryGet**
> V1SystemsBusinessCalendarCountryGet200Response v1SystemsBusinessCalendarCountryGet()

Retrieve working days, public holidays, and next upcoming holiday for any country and optional year/month. Covers 100+ countries.

### Example

```typescript
import {
    GlobalDataApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new GlobalDataApi(configuration);

let country: string; //ISO 3166-1 alpha-2 country code (e.g. DE, US, FR) (default to undefined)
let year: number; //Year to retrieve the calendar for. Defaults to the current year. (optional) (default to undefined)
let month: number; //Month (1–12). When provided, returns month-scoped working day count, total days, weekend days, and filtered holidays. (optional) (default to undefined)

const { status, data } = await apiInstance.v1SystemsBusinessCalendarCountryGet(
    country,
    year,
    month
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **country** | [**string**] | ISO 3166-1 alpha-2 country code (e.g. DE, US, FR) | defaults to undefined|
| **year** | [**number**] | Year to retrieve the calendar for. Defaults to the current year. | (optional) defaults to undefined|
| **month** | [**number**] | Month (1–12). When provided, returns month-scoped working day count, total days, weekend days, and filtered holidays. | (optional) defaults to undefined|


### Return type

**V1SystemsBusinessCalendarCountryGet200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**401** | Missing or invalid API key |  -  |
|**422** | Country is not a 2-letter ISO code, or year/month is out of the valid range |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1SystemsLocationResolvePost**
> V1SystemsLocationResolvePost200Response v1SystemsLocationResolvePost(v1SystemsLocationResolvePostRequest)

Resolve an address or coordinates into a full location profile — city, country, timezone, UTC offset, current time, working days this month, and next holiday.

### Example

```typescript
import {
    GlobalDataApi,
    Configuration,
    V1SystemsLocationResolvePostRequest
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new GlobalDataApi(configuration);

let v1SystemsLocationResolvePostRequest: V1SystemsLocationResolvePostRequest; //

const { status, data } = await apiInstance.v1SystemsLocationResolvePost(
    v1SystemsLocationResolvePostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **v1SystemsLocationResolvePostRequest** | **V1SystemsLocationResolvePostRequest**|  | |


### Return type

**V1SystemsLocationResolvePost200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**401** | Missing or invalid API key |  -  |
|**422** | Neither address nor coordinates was provided; Geocoding service was unavailable |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1SystemsTimezoneFromIpIpGet**
> V1SystemsTimezoneFromIpIpGet200Response v1SystemsTimezoneFromIpIpGet()

Resolve the timezone, UTC offset, and DST status for any IPv4 or IPv6 address without needing an address string.

### Example

```typescript
import {
    GlobalDataApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new GlobalDataApi(configuration);

let ip: string; //IPv4 or IPv6 address to resolve (default to undefined)

const { status, data } = await apiInstance.v1SystemsTimezoneFromIpIpGet(
    ip
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **ip** | [**string**] | IPv4 or IPv6 address to resolve | defaults to undefined|


### Return type

**V1SystemsTimezoneFromIpIpGet200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**401** | Missing or invalid API key |  -  |
|**422** | The IP address could not be geolocated |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

