# GeocodeApi

All URIs are relative to *https://requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1PlacesGeocodeBatchPost**](#v1placesgeocodebatchpost) | **POST** /v1/places/geocode/batch | Batch Geocode Addresses|
|[**v1PlacesGeocodeGet**](#v1placesgeocodeget) | **GET** /v1/places/geocode | Geocode Address|
|[**v1PlacesReverseGeocodeBatchPost**](#v1placesreversegeocodebatchpost) | **POST** /v1/places/reverse-geocode/batch | Batch Reverse Geocode|
|[**v1PlacesReverseGeocodeGet**](#v1placesreversegeocodeget) | **GET** /v1/places/reverse-geocode | Reverse Geocode|

# **v1PlacesGeocodeBatchPost**
> V1PlacesGeocodeBatchPost200Response v1PlacesGeocodeBatchPost(v1PlacesGeocodeBatchPostRequest)

Convert up to 20 addresses to coordinates in a single request. Processed concurrently; results are returned in input order. Per-item errors are reported inline. Cache hits avoid outbound network calls.

### Example

```typescript
import {
    GeocodeApi,
    Configuration,
    V1PlacesGeocodeBatchPostRequest
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new GeocodeApi(configuration);

let v1PlacesGeocodeBatchPostRequest: V1PlacesGeocodeBatchPostRequest; //

const { status, data } = await apiInstance.v1PlacesGeocodeBatchPost(
    v1PlacesGeocodeBatchPostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **v1PlacesGeocodeBatchPostRequest** | **V1PlacesGeocodeBatchPostRequest**|  | |


### Return type

**V1PlacesGeocodeBatchPost200Response**

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
|**422** | The addresses array is missing, empty, or exceeds 20 items. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1PlacesGeocodeGet**
> V1PlacesGeocodeGet200Response v1PlacesGeocodeGet()

Converts a free-text address into latitude and longitude coordinates.

### Example

```typescript
import {
    GeocodeApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new GeocodeApi(configuration);

let address: string; //The address to geocode (street, city, country, or any combination) (default to undefined)

const { status, data } = await apiInstance.v1PlacesGeocodeGet(
    address
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **address** | [**string**] | The address to geocode (street, city, country, or any combination) | defaults to undefined|


### Return type

**V1PlacesGeocodeGet200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**400** | The address parameter is missing. |  -  |
|**404** | No results found for the given address. |  -  |
|**503** | The geocoding service is temporarily unavailable. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1PlacesReverseGeocodeBatchPost**
> V1PlacesReverseGeocodeBatchPost200Response v1PlacesReverseGeocodeBatchPost(v1PlacesReverseGeocodeBatchPostRequest)

Convert up to 20 coordinate pairs to addresses in a single request. Processed concurrently; results are returned in input order. Per-item errors are reported inline.

### Example

```typescript
import {
    GeocodeApi,
    Configuration,
    V1PlacesReverseGeocodeBatchPostRequest
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new GeocodeApi(configuration);

let v1PlacesReverseGeocodeBatchPostRequest: V1PlacesReverseGeocodeBatchPostRequest; //

const { status, data } = await apiInstance.v1PlacesReverseGeocodeBatchPost(
    v1PlacesReverseGeocodeBatchPostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **v1PlacesReverseGeocodeBatchPostRequest** | **V1PlacesReverseGeocodeBatchPostRequest**|  | |


### Return type

**V1PlacesReverseGeocodeBatchPost200Response**

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
|**422** | The items array is missing, empty, or exceeds 20 items. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1PlacesReverseGeocodeGet**
> V1PlacesReverseGeocodeGet200Response v1PlacesReverseGeocodeGet()

Converts geographic coordinates into a human-readable address.

### Example

```typescript
import {
    GeocodeApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new GeocodeApi(configuration);

let lat: number; //Latitude of the location (-90 to 90) (default to undefined)
let lon: number; //Longitude of the location (-180 to 180) (default to undefined)

const { status, data } = await apiInstance.v1PlacesReverseGeocodeGet(
    lat,
    lon
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **lat** | [**number**] | Latitude of the location (-90 to 90) | defaults to undefined|
| **lon** | [**number**] | Longitude of the location (-180 to 180) | defaults to undefined|


### Return type

**V1PlacesReverseGeocodeGet200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**400** | lat or lon is missing or out of range (lat: -90..90, lon: -180..180). |  -  |
|**404** | No address found for the given coordinates. |  -  |
|**503** | The geocoding service is temporarily unavailable. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

