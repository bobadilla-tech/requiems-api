# \WhoisAPI

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1NetworkingWhoisDomainGet**](WhoisAPI.md#V1NetworkingWhoisDomainGet) | **Get** /v1/networking/whois/{domain} | WHOIS Lookup



## V1NetworkingWhoisDomainGet

> V1NetworkingWhoisDomainGet200Response V1NetworkingWhoisDomainGet(ctx, domain).Execute()

WHOIS Lookup



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
	domain := "example.com" // string | The domain name to look up (e.g. example.com)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhoisAPI.V1NetworkingWhoisDomainGet(context.Background(), domain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhoisAPI.V1NetworkingWhoisDomainGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1NetworkingWhoisDomainGet`: V1NetworkingWhoisDomainGet200Response
	fmt.Fprintf(os.Stdout, "Response from `WhoisAPI.V1NetworkingWhoisDomainGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | The domain name to look up (e.g. example.com) | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1NetworkingWhoisDomainGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1NetworkingWhoisDomainGet200Response**](V1NetworkingWhoisDomainGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

