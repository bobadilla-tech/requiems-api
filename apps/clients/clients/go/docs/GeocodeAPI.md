# \GeocodeAPI

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1PlacesGeocodeGet**](GeocodeAPI.md#V1PlacesGeocodeGet) | **Get** /v1/places/geocode | Geocode Address
[**V1PlacesReverseGeocodeGet**](GeocodeAPI.md#V1PlacesReverseGeocodeGet) | **Get** /v1/places/reverse-geocode | Reverse Geocode



## V1PlacesGeocodeGet

> V1PlacesGeocodeGet200Response V1PlacesGeocodeGet(ctx).Address(address).Execute()

Geocode Address



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
	address := "1600 Pennsylvania Ave NW, Washington DC" // string | The address to geocode (street, city, country, or any combination)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GeocodeAPI.V1PlacesGeocodeGet(context.Background()).Address(address).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeocodeAPI.V1PlacesGeocodeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1PlacesGeocodeGet`: V1PlacesGeocodeGet200Response
	fmt.Fprintf(os.Stdout, "Response from `GeocodeAPI.V1PlacesGeocodeGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1PlacesGeocodeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **address** | **string** | The address to geocode (street, city, country, or any combination) | 

### Return type

[**V1PlacesGeocodeGet200Response**](V1PlacesGeocodeGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1PlacesReverseGeocodeGet

> V1PlacesReverseGeocodeGet200Response V1PlacesReverseGeocodeGet(ctx).Lat(lat).Lon(lon).Execute()

Reverse Geocode



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
	lat := float32(38.8977) // float32 | Latitude of the location (-90 to 90)
	lon := float32(-77.0365) // float32 | Longitude of the location (-180 to 180)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GeocodeAPI.V1PlacesReverseGeocodeGet(context.Background()).Lat(lat).Lon(lon).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeocodeAPI.V1PlacesReverseGeocodeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1PlacesReverseGeocodeGet`: V1PlacesReverseGeocodeGet200Response
	fmt.Fprintf(os.Stdout, "Response from `GeocodeAPI.V1PlacesReverseGeocodeGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1PlacesReverseGeocodeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lat** | **float32** | Latitude of the location (-90 to 90) | 
 **lon** | **float32** | Longitude of the location (-180 to 180) | 

### Return type

[**V1PlacesReverseGeocodeGet200Response**](V1PlacesReverseGeocodeGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

