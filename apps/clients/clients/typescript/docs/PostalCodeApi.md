# PostalCodeApi

All URIs are relative to *https://api.requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1PlacesPostalCodeGet**](#v1placespostalcodeget) | **GET** /v1/places/postal/{code} | Lookup Postal Code|

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

