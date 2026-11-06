# \VpnDetectionAPI

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1NetworkingIpVpnIpGet**](VpnDetectionAPI.md#V1NetworkingIpVpnIpGet) | **Get** /v1/networking/ip/vpn/{ip} | Check IP Address



## V1NetworkingIpVpnIpGet

> V1NetworkingIpVpnIpGet200Response V1NetworkingIpVpnIpGet(ctx, ip).Execute()

Check IP Address



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
	ip := "8.8.8.8" // string | The IP address to check (supports IPv4 and IPv6)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpnDetectionAPI.V1NetworkingIpVpnIpGet(context.Background(), ip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpnDetectionAPI.V1NetworkingIpVpnIpGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1NetworkingIpVpnIpGet`: V1NetworkingIpVpnIpGet200Response
	fmt.Fprintf(os.Stdout, "Response from `VpnDetectionAPI.V1NetworkingIpVpnIpGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ip** | **string** | The IP address to check (supports IPv4 and IPv6) | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1NetworkingIpVpnIpGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1NetworkingIpVpnIpGet200Response**](V1NetworkingIpVpnIpGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

