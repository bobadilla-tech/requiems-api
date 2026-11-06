# \DisposableEmailAPI

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1NetworkingDisposableBatchPost**](DisposableEmailAPI.md#V1NetworkingDisposableBatchPost) | **Post** /v1/networking/disposable/batch | Check Batch Emails
[**V1NetworkingDisposableCheckPost**](DisposableEmailAPI.md#V1NetworkingDisposableCheckPost) | **Post** /v1/networking/disposable/check | Check Single Email
[**V1NetworkingDisposableDomainDomainGet**](DisposableEmailAPI.md#V1NetworkingDisposableDomainDomainGet) | **Get** /v1/networking/disposable/domain/{domain} | Check Domain
[**V1NetworkingDisposableDomainsGet**](DisposableEmailAPI.md#V1NetworkingDisposableDomainsGet) | **Get** /v1/networking/disposable/domains | List Domains (Paginated)
[**V1NetworkingDisposableStatsGet**](DisposableEmailAPI.md#V1NetworkingDisposableStatsGet) | **Get** /v1/networking/disposable/stats | Get Statistics



## V1NetworkingDisposableBatchPost

> V1NetworkingDisposableBatchPost200Response V1NetworkingDisposableBatchPost(ctx).V1NetworkingDisposableBatchPostRequest(v1NetworkingDisposableBatchPostRequest).Execute()

Check Batch Emails



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
	v1NetworkingDisposableBatchPostRequest := *openapiclient.NewV1NetworkingDisposableBatchPostRequest([]interface{}{nil}) // V1NetworkingDisposableBatchPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisposableEmailAPI.V1NetworkingDisposableBatchPost(context.Background()).V1NetworkingDisposableBatchPostRequest(v1NetworkingDisposableBatchPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisposableEmailAPI.V1NetworkingDisposableBatchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1NetworkingDisposableBatchPost`: V1NetworkingDisposableBatchPost200Response
	fmt.Fprintf(os.Stdout, "Response from `DisposableEmailAPI.V1NetworkingDisposableBatchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1NetworkingDisposableBatchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1NetworkingDisposableBatchPostRequest** | [**V1NetworkingDisposableBatchPostRequest**](V1NetworkingDisposableBatchPostRequest.md) |  | 

### Return type

[**V1NetworkingDisposableBatchPost200Response**](V1NetworkingDisposableBatchPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1NetworkingDisposableCheckPost

> V1NetworkingDisposableCheckPost200Response V1NetworkingDisposableCheckPost(ctx).V1NetworkingDisposableCheckPostRequest(v1NetworkingDisposableCheckPostRequest).Execute()

Check Single Email



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
	v1NetworkingDisposableCheckPostRequest := *openapiclient.NewV1NetworkingDisposableCheckPostRequest("test@example.com") // V1NetworkingDisposableCheckPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisposableEmailAPI.V1NetworkingDisposableCheckPost(context.Background()).V1NetworkingDisposableCheckPostRequest(v1NetworkingDisposableCheckPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisposableEmailAPI.V1NetworkingDisposableCheckPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1NetworkingDisposableCheckPost`: V1NetworkingDisposableCheckPost200Response
	fmt.Fprintf(os.Stdout, "Response from `DisposableEmailAPI.V1NetworkingDisposableCheckPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1NetworkingDisposableCheckPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1NetworkingDisposableCheckPostRequest** | [**V1NetworkingDisposableCheckPostRequest**](V1NetworkingDisposableCheckPostRequest.md) |  | 

### Return type

[**V1NetworkingDisposableCheckPost200Response**](V1NetworkingDisposableCheckPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1NetworkingDisposableDomainDomainGet

> V1NetworkingDisposableDomainDomainGet200Response V1NetworkingDisposableDomainDomainGet(ctx, domain).Execute()

Check Domain



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
	domain := "tempmail.com" // string | The domain to check

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisposableEmailAPI.V1NetworkingDisposableDomainDomainGet(context.Background(), domain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisposableEmailAPI.V1NetworkingDisposableDomainDomainGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1NetworkingDisposableDomainDomainGet`: V1NetworkingDisposableDomainDomainGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DisposableEmailAPI.V1NetworkingDisposableDomainDomainGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | The domain to check | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1NetworkingDisposableDomainDomainGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1NetworkingDisposableDomainDomainGet200Response**](V1NetworkingDisposableDomainDomainGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1NetworkingDisposableDomainsGet

> V1NetworkingDisposableDomainsGet200Response V1NetworkingDisposableDomainsGet(ctx).Page(page).PerPage(perPage).Execute()

List Domains (Paginated)



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
	page := int32(1) // int32 | Page number (default: 1) (optional)
	perPage := int32(100) // int32 | Items per page (default: 100, max: 1000) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisposableEmailAPI.V1NetworkingDisposableDomainsGet(context.Background()).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisposableEmailAPI.V1NetworkingDisposableDomainsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1NetworkingDisposableDomainsGet`: V1NetworkingDisposableDomainsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DisposableEmailAPI.V1NetworkingDisposableDomainsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1NetworkingDisposableDomainsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number (default: 1) | 
 **perPage** | **int32** | Items per page (default: 100, max: 1000) | 

### Return type

[**V1NetworkingDisposableDomainsGet200Response**](V1NetworkingDisposableDomainsGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1NetworkingDisposableStatsGet

> V1NetworkingDisposableStatsGet200Response V1NetworkingDisposableStatsGet(ctx).Execute()

Get Statistics



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
	resp, r, err := apiClient.DisposableEmailAPI.V1NetworkingDisposableStatsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisposableEmailAPI.V1NetworkingDisposableStatsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1NetworkingDisposableStatsGet`: V1NetworkingDisposableStatsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DisposableEmailAPI.V1NetworkingDisposableStatsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1NetworkingDisposableStatsGetRequest struct via the builder pattern


### Return type

[**V1NetworkingDisposableStatsGet200Response**](V1NetworkingDisposableStatsGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

