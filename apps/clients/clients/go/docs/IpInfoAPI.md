# \IpInfoAPI

All URIs are relative to *https://requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1NetworkingIpGet**](IpInfoAPI.md#V1NetworkingIpGet) | **Get** /v1/networking/ip | Get IP Info (Caller IP)
[**V1NetworkingIpInfoBatchPost**](IpInfoAPI.md#V1NetworkingIpInfoBatchPost) | **Post** /v1/networking/ip/info/batch | Batch IP Geolocation
[**V1NetworkingIpIpGet**](IpInfoAPI.md#V1NetworkingIpIpGet) | **Get** /v1/networking/ip/{ip} | Get IP Info for IP



## V1NetworkingIpGet

> V1NetworkingIpGet200Response V1NetworkingIpGet(ctx).Execute()

Get IP Info (Caller IP)



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IpInfoAPI.V1NetworkingIpGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IpInfoAPI.V1NetworkingIpGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1NetworkingIpGet`: V1NetworkingIpGet200Response
	fmt.Fprintf(os.Stdout, "Response from `IpInfoAPI.V1NetworkingIpGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1NetworkingIpGetRequest struct via the builder pattern


### Return type

[**V1NetworkingIpGet200Response**](V1NetworkingIpGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1NetworkingIpInfoBatchPost

> V1NetworkingIpInfoBatchPost200Response V1NetworkingIpInfoBatchPost(ctx).V1NetworkingIpAsnBatchPostRequest(v1NetworkingIpAsnBatchPostRequest).Execute()

Batch IP Geolocation



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
	v1NetworkingIpAsnBatchPostRequest := *openapiclient.NewV1NetworkingIpAsnBatchPostRequest("Ips_example") // V1NetworkingIpAsnBatchPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IpInfoAPI.V1NetworkingIpInfoBatchPost(context.Background()).V1NetworkingIpAsnBatchPostRequest(v1NetworkingIpAsnBatchPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IpInfoAPI.V1NetworkingIpInfoBatchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1NetworkingIpInfoBatchPost`: V1NetworkingIpInfoBatchPost200Response
	fmt.Fprintf(os.Stdout, "Response from `IpInfoAPI.V1NetworkingIpInfoBatchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1NetworkingIpInfoBatchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1NetworkingIpAsnBatchPostRequest** | [**V1NetworkingIpAsnBatchPostRequest**](V1NetworkingIpAsnBatchPostRequest.md) |  | 

### Return type

[**V1NetworkingIpInfoBatchPost200Response**](V1NetworkingIpInfoBatchPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1NetworkingIpIpGet

> V1NetworkingIpIpGet200Response V1NetworkingIpIpGet(ctx, ip).Execute()

Get IP Info for IP



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
	ip := "8.8.8.8" // string | The IP address to look up (supports IPv4 and IPv6)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IpInfoAPI.V1NetworkingIpIpGet(context.Background(), ip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IpInfoAPI.V1NetworkingIpIpGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1NetworkingIpIpGet`: V1NetworkingIpIpGet200Response
	fmt.Fprintf(os.Stdout, "Response from `IpInfoAPI.V1NetworkingIpIpGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ip** | **string** | The IP address to look up (supports IPv4 and IPv6) | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1NetworkingIpIpGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1NetworkingIpIpGet200Response**](V1NetworkingIpIpGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

