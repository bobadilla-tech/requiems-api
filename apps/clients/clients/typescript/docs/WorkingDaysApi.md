# WorkingDaysApi

All URIs are relative to *https://api.requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1PlacesWorkingDaysGet**](#v1placesworkingdaysget) | **GET** /v1/places/working-days | Calculate Working Days|

# **v1PlacesWorkingDaysGet**
> V1PlacesWorkingDaysGet200Response v1PlacesWorkingDaysGet()

Calculate the number of working days between two dates, optionally accounting for country-specific holidays

### Example

```typescript
import {
    WorkingDaysApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new WorkingDaysApi(configuration);

let from: string; //Start date in YYYY-MM-DD format (ISO 8601) (default to undefined)
let to: string; //End date in YYYY-MM-DD format (ISO 8601). Must be >= from date. (default to undefined)
let country: string; //ISO 3166-1 alpha-2 country code (e.g., \"US\", \"GB\", \"FR\"). When provided, country-specific holidays are excluded from working days count. (optional) (default to undefined)
let subdivision: string; //ISO 3166-2 subdivision code for state/region within the country (e.g., \"NY\" for New York, \"CA\" for California). Only used when country is provided. (optional) (default to undefined)

const { status, data } = await apiInstance.v1PlacesWorkingDaysGet(
    from,
    to,
    country,
    subdivision
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **from** | [**string**] | Start date in YYYY-MM-DD format (ISO 8601) | defaults to undefined|
| **to** | [**string**] | End date in YYYY-MM-DD format (ISO 8601). Must be &gt;&#x3D; from date. | defaults to undefined|
| **country** | [**string**] | ISO 3166-1 alpha-2 country code (e.g., \&quot;US\&quot;, \&quot;GB\&quot;, \&quot;FR\&quot;). When provided, country-specific holidays are excluded from working days count. | (optional) defaults to undefined|
| **subdivision** | [**string**] | ISO 3166-2 subdivision code for state/region within the country (e.g., \&quot;NY\&quot; for New York, \&quot;CA\&quot; for California). Only used when country is provided. | (optional) defaults to undefined|


### Return type

**V1PlacesWorkingDaysGet200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**400** | The from and to parameters are required, or to date is before from date, or invalid date format |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

