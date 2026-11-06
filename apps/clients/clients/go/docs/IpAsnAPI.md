# \IpAsnAPI

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1NetworkingIpAsnGet**](IpAsnAPI.md#V1NetworkingIpAsnGet) | **Get** /v1/networking/ip/asn | Lookup ASN (Caller IP)
[**V1NetworkingIpAsnIpGet**](IpAsnAPI.md#V1NetworkingIpAsnIpGet) | **Get** /v1/networking/ip/asn/{ip} | Lookup ASN for IP



## V1NetworkingIpAsnGet

> V1NetworkingIpAsnGet200Response V1NetworkingIpAsnGet(ctx).Execute()

Lookup ASN (Caller IP)



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
	resp, r, err := apiClient.IpAsnAPI.V1NetworkingIpAsnGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IpAsnAPI.V1NetworkingIpAsnGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1NetworkingIpAsnGet`: V1NetworkingIpAsnGet200Response
	fmt.Fprintf(os.Stdout, "Response from `IpAsnAPI.V1NetworkingIpAsnGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1NetworkingIpAsnGetRequest struct via the builder pattern


### Return type

[**V1NetworkingIpAsnGet200Response**](V1NetworkingIpAsnGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1NetworkingIpAsnIpGet

> V1NetworkingIpAsnIpGet200Response V1NetworkingIpAsnIpGet(ctx, ip).Execute()

Lookup ASN for IP



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
	resp, r, err := apiClient.IpAsnAPI.V1NetworkingIpAsnIpGet(context.Background(), ip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IpAsnAPI.V1NetworkingIpAsnIpGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1NetworkingIpAsnIpGet`: V1NetworkingIpAsnIpGet200Response
	fmt.Fprintf(os.Stdout, "Response from `IpAsnAPI.V1NetworkingIpAsnIpGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ip** | **string** | The IP address to look up (supports IPv4 and IPv6) | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1NetworkingIpAsnIpGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1NetworkingIpAsnIpGet200Response**](V1NetworkingIpAsnIpGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

