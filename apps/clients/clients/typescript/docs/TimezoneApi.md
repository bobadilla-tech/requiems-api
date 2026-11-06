# TimezoneApi

All URIs are relative to *https://requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1PlacesTimezoneBatchPost**](#v1placestimezonebatchpost) | **POST** /v1/places/timezone/batch | Batch Timezone Lookup|
|[**v1PlacesTimezoneGet**](#v1placestimezoneget) | **GET** /v1/places/timezone | Get Timezone|

# **v1PlacesTimezoneBatchPost**
> V1PlacesTimezoneBatchPost200Response v1PlacesTimezoneBatchPost(v1PlacesTimezoneBatchPostRequest)

Look up timezone information for up to 50 locations in a single request. Each item can specify a timezone name, city, or lat/lon coordinates. Priority is timezone name > city > coordinates. Results are returned in input order.

### Example

```typescript
import {
    TimezoneApi,
    Configuration,
    V1PlacesTimezoneBatchPostRequest
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new TimezoneApi(configuration);

let v1PlacesTimezoneBatchPostRequest: V1PlacesTimezoneBatchPostRequest; //

const { status, data } = await apiInstance.v1PlacesTimezoneBatchPost(
    v1PlacesTimezoneBatchPostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **v1PlacesTimezoneBatchPostRequest** | **V1PlacesTimezoneBatchPostRequest**|  | |


### Return type

**V1PlacesTimezoneBatchPost200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**400** | The request body is missing or malformed. |  -  |
|**422** | The items array is missing, empty, or exceeds 50 items. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

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

