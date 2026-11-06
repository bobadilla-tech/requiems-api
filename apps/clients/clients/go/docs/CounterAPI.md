# \CounterAPI

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1TechnologyCounterNamespaceGet**](CounterAPI.md#V1TechnologyCounterNamespaceGet) | **Get** /v1/technology/counter/{namespace} | Get Counter Value
[**V1TechnologyCounterNamespacePost**](CounterAPI.md#V1TechnologyCounterNamespacePost) | **Post** /v1/technology/counter/{namespace} | Increment Counter



## V1TechnologyCounterNamespaceGet

> V1TechnologyCounterNamespaceGet200Response V1TechnologyCounterNamespaceGet(ctx, namespace).Execute()

Get Counter Value



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
	namespace := "page-views" // string | Counter namespace (1-64 chars: alphanumeric, hyphen, underscore)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CounterAPI.V1TechnologyCounterNamespaceGet(context.Background(), namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CounterAPI.V1TechnologyCounterNamespaceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1TechnologyCounterNamespaceGet`: V1TechnologyCounterNamespaceGet200Response
	fmt.Fprintf(os.Stdout, "Response from `CounterAPI.V1TechnologyCounterNamespaceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | Counter namespace (1-64 chars: alphanumeric, hyphen, underscore) | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1TechnologyCounterNamespaceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1TechnologyCounterNamespaceGet200Response**](V1TechnologyCounterNamespaceGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1TechnologyCounterNamespacePost

> V1TechnologyCounterNamespacePost200Response V1TechnologyCounterNamespacePost(ctx, namespace).Execute()

Increment Counter



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
	namespace := "page-views" // string | Counter namespace (1-64 chars: alphanumeric, hyphen, underscore)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CounterAPI.V1TechnologyCounterNamespacePost(context.Background(), namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CounterAPI.V1TechnologyCounterNamespacePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1TechnologyCounterNamespacePost`: V1TechnologyCounterNamespacePost200Response
	fmt.Fprintf(os.Stdout, "Response from `CounterAPI.V1TechnologyCounterNamespacePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | Counter namespace (1-64 chars: alphanumeric, hyphen, underscore) | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1TechnologyCounterNamespacePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1TechnologyCounterNamespacePost200Response**](V1TechnologyCounterNamespacePost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

