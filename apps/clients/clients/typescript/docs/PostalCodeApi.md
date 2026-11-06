# PostalCodeApi

All URIs are relative to *https://requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1PlacesPostalBatchPost**](#v1placespostalbatchpost) | **POST** /v1/places/postal/batch | Batch Lookup Postal Codes|
|[**v1PlacesPostalCodeGet**](#v1placespostalcodeget) | **GET** /v1/places/postal/{code} | Lookup Postal Code|

# **v1PlacesPostalBatchPost**
> V1PlacesPostalBatchPost200Response v1PlacesPostalBatchPost(v1PlacesPostalBatchPostRequest)

Look up city, state, and coordinates for up to 50 postal codes in a single request. Results are returned in input order.

### Example

```typescript
import {
    PostalCodeApi,
    Configuration,
    V1PlacesPostalBatchPostRequest
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new PostalCodeApi(configuration);

let v1PlacesPostalBatchPostRequest: V1PlacesPostalBatchPostRequest; //

const { status, data } = await apiInstance.v1PlacesPostalBatchPost(
    v1PlacesPostalBatchPostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **v1PlacesPostalBatchPostRequest** | **V1PlacesPostalBatchPostRequest**|  | |


### Return type

**V1PlacesPostalBatchPost200Response**

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

# **v1PlacesPostalCodeGet**
> V1PlacesPostalCodeGet200Response v1PlacesPostalCodeGet()

Returns city, state, country, and coordinates for the given postal code.

### Example

```typescript
import {
    PostalCodeApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new PostalCodeApi(configuration);

let code: string; //The postal code to look up (e.g. 10001 for New York, SW1A 1AA for London) (default to undefined)
let country: string; //ISO 3166-1 alpha-2 country code (default: US) (optional) (default to undefined)

const { status, data } = await apiInstance.v1PlacesPostalCodeGet(
    code,
    country
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **code** | [**string**] | The postal code to look up (e.g. 10001 for New York, SW1A 1AA for London) | defaults to undefined|
| **country** | [**string**] | ISO 3166-1 alpha-2 country code (default: US) | (optional) defaults to undefined|


### Return type

**V1PlacesPostalCodeGet200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**404** | The postal code was not found for the given country. |  -  |
|**500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

