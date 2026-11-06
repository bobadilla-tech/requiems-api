# CitiesApi

All URIs are relative to *https://api.requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1PlacesCitiesCityGet**](#v1placescitiescityget) | **GET** /v1/places/cities/{city} | Get City Info|

# **v1PlacesCitiesCityGet**
> V1PlacesCitiesCityGet200Response v1PlacesCitiesCityGet()

Returns metadata for a city by name. Lookup is case-insensitive.

### Example

```typescript
import {
    CitiesApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new CitiesApi(configuration);

let city: string; //City name to look up (e.g. london, tokyo, new york city) (default to undefined)

const { status, data } = await apiInstance.v1PlacesCitiesCityGet(
    city
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **city** | [**string**] | City name to look up (e.g. london, tokyo, new york city) | defaults to undefined|


### Return type

**V1PlacesCitiesCityGet200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**404** | No city with that name was found in the dataset. |  -  |
|**500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

