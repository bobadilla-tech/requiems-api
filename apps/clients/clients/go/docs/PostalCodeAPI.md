# \PostalCodeAPI

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1PlacesPostalCodeGet**](PostalCodeAPI.md#V1PlacesPostalCodeGet) | **Get** /v1/places/postal/{code} | Lookup Postal Code



## V1PlacesPostalCodeGet

> V1PlacesPostalCodeGet200Response V1PlacesPostalCodeGet(ctx, code).Country(country).Execute()

Lookup Postal Code



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
	code := "10001" // string | The postal code to look up (e.g. 10001 for New York, SW1A 1AA for London)
	country := "US" // string | ISO 3166-1 alpha-2 country code (default: US) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostalCodeAPI.V1PlacesPostalCodeGet(context.Background(), code).Country(country).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostalCodeAPI.V1PlacesPostalCodeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1PlacesPostalCodeGet`: V1PlacesPostalCodeGet200Response
	fmt.Fprintf(os.Stdout, "Response from `PostalCodeAPI.V1PlacesPostalCodeGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | The postal code to look up (e.g. 10001 for New York, SW1A 1AA for London) | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1PlacesPostalCodeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **country** | **string** | ISO 3166-1 alpha-2 country code (default: US) | 

### Return type

[**V1PlacesPostalCodeGet200Response**](V1PlacesPostalCodeGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

