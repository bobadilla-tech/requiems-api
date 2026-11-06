# \TimezoneAPI

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1PlacesTimezoneGet**](TimezoneAPI.md#V1PlacesTimezoneGet) | **Get** /v1/places/timezone | Get Timezone



## V1PlacesTimezoneGet

> V1PlacesTimezoneGet200Response V1PlacesTimezoneGet(ctx).Lat(lat).Lon(lon).City(city).Execute()

Get Timezone



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	lat := "lat_example" // string | Latitude of the location (-90 to 90). Required when using coordinate-based lookup. (optional)
	lon := "lon_example" // string | Longitude of the location (-180 to 180). Required when using coordinate-based lookup. (optional)
	city := "city_example" // string | City name for city-based lookup (e.g. 'Tokyo', 'London'). Required when not using coordinates. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TimezoneAPI.V1PlacesTimezoneGet(context.Background()).Lat(lat).Lon(lon).City(city).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimezoneAPI.V1PlacesTimezoneGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1PlacesTimezoneGet`: V1PlacesTimezoneGet200Response
	fmt.Fprintf(os.Stdout, "Response from `TimezoneAPI.V1PlacesTimezoneGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1PlacesTimezoneGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lat** | **string** | Latitude of the location (-90 to 90). Required when using coordinate-based lookup. | 
 **lon** | **string** | Longitude of the location (-180 to 180). Required when using coordinate-based lookup. | 
 **city** | **string** | City name for city-based lookup (e.g. &#39;Tokyo&#39;, &#39;London&#39;). Required when not using coordinates. | 

### Return type

[**V1PlacesTimezoneGet200Response**](V1PlacesTimezoneGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

