# TimezoneApi

All URIs are relative to *https://api.requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1PlacesTimezoneGet**](#v1placestimezoneget) | **GET** /v1/places/timezone | Get Timezone|

# **v1PlacesTimezoneGet**
> V1PlacesTimezoneGet200Response v1PlacesTimezoneGet()

Returns timezone information for the given coordinates or city name. Provide either `city` or both `lat` and `lon`.

### Example

```typescript
import {
    TimezoneApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new TimezoneApi(configuration);

let lat: string; //Latitude of the location (-90 to 90). Required when using coordinate-based lookup. (optional) (default to undefined)
let lon: string; //Longitude of the location (-180 to 180). Required when using coordinate-based lookup. (optional) (default to undefined)
let city: string; //City name for city-based lookup (e.g. \'Tokyo\', \'London\'). Required when not using coordinates. (optional) (default to undefined)

const { status, data } = await apiInstance.v1PlacesTimezoneGet(
    lat,
    lon,
    city
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **lat** | [**string**] | Latitude of the location (-90 to 90). Required when using coordinate-based lookup. | (optional) defaults to undefined|
| **lon** | [**string**] | Longitude of the location (-180 to 180). Required when using coordinate-based lookup. | (optional) defaults to undefined|
| **city** | [**string**] | City name for city-based lookup (e.g. \&#39;Tokyo\&#39;, \&#39;London\&#39;). Required when not using coordinates. | (optional) defaults to undefined|


### Return type

**V1PlacesTimezoneGet200Response**

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

