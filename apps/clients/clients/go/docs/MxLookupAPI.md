# \MxLookupAPI

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1NetworkingMxDomainGet**](MxLookupAPI.md#V1NetworkingMxDomainGet) | **Get** /v1/networking/mx/{domain} | MX Lookup



## V1NetworkingMxDomainGet

> V1NetworkingMxDomainGet200Response V1NetworkingMxDomainGet(ctx, domain).Execute()

MX Lookup



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
	domain := "gmail.com" // string | The domain name to look up MX records for (e.g. gmail.com)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MxLookupAPI.V1NetworkingMxDomainGet(context.Background(), domain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MxLookupAPI.V1NetworkingMxDomainGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1NetworkingMxDomainGet`: V1NetworkingMxDomainGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MxLookupAPI.V1NetworkingMxDomainGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | The domain name to look up MX records for (e.g. gmail.com) | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1NetworkingMxDomainGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1NetworkingMxDomainGet200Response**](V1NetworkingMxDomainGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

