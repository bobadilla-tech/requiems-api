# \WorldTimeAPI

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1PlacesTimeTimezoneGet**](WorldTimeAPI.md#V1PlacesTimeTimezoneGet) | **Get** /v1/places/time/{timezone} | Get Current Time by Timezone



## V1PlacesTimeTimezoneGet

> V1PlacesTimeTimezoneGet200Response V1PlacesTimeTimezoneGet(ctx, timezone).Execute()

Get Current Time by Timezone



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
	timezone := "America/New_York" // string | IANA timezone identifier (e.g. 'America/New_York', 'Europe/London', 'Asia/Kolkata')

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorldTimeAPI.V1PlacesTimeTimezoneGet(context.Background(), timezone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorldTimeAPI.V1PlacesTimeTimezoneGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1PlacesTimeTimezoneGet`: V1PlacesTimeTimezoneGet200Response
	fmt.Fprintf(os.Stdout, "Response from `WorldTimeAPI.V1PlacesTimeTimezoneGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**timezone** | **string** | IANA timezone identifier (e.g. &#39;America/New_York&#39;, &#39;Europe/London&#39;, &#39;Asia/Kolkata&#39;) | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1PlacesTimeTimezoneGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1PlacesTimeTimezoneGet200Response**](V1PlacesTimeTimezoneGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

