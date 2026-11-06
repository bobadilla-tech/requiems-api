# \CitiesAPI

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1PlacesCitiesCityGet**](CitiesAPI.md#V1PlacesCitiesCityGet) | **Get** /v1/places/cities/{city} | Get City Info



## V1PlacesCitiesCityGet

> V1PlacesCitiesCityGet200Response V1PlacesCitiesCityGet(ctx, city).Execute()

Get City Info



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
	city := "london" // string | City name to look up (e.g. london, tokyo, new york city)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CitiesAPI.V1PlacesCitiesCityGet(context.Background(), city).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CitiesAPI.V1PlacesCitiesCityGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1PlacesCitiesCityGet`: V1PlacesCitiesCityGet200Response
	fmt.Fprintf(os.Stdout, "Response from `CitiesAPI.V1PlacesCitiesCityGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**city** | **string** | City name to look up (e.g. london, tokyo, new york city) | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1PlacesCitiesCityGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1PlacesCitiesCityGet200Response**](V1PlacesCitiesCityGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

