# \MxLookupAPI

All URIs are relative to *https://requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1NetworkingMxBatchPost**](MxLookupAPI.md#V1NetworkingMxBatchPost) | **Post** /v1/networking/mx/batch | Batch MX Lookup
[**V1NetworkingMxDomainGet**](MxLookupAPI.md#V1NetworkingMxDomainGet) | **Get** /v1/networking/mx/{domain} | MX Lookup



## V1NetworkingMxBatchPost

> V1NetworkingMxBatchPost(ctx).V1NetworkingMxBatchPostRequest(v1NetworkingMxBatchPostRequest).Execute()

Batch MX Lookup



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
	v1NetworkingMxBatchPostRequest := *openapiclient.NewV1NetworkingMxBatchPostRequest([]interface{}{nil}) // V1NetworkingMxBatchPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MxLookupAPI.V1NetworkingMxBatchPost(context.Background()).V1NetworkingMxBatchPostRequest(v1NetworkingMxBatchPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MxLookupAPI.V1NetworkingMxBatchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1NetworkingMxBatchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1NetworkingMxBatchPostRequest** | [**V1NetworkingMxBatchPostRequest**](V1NetworkingMxBatchPostRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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

