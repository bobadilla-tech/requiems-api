# \DataIntegrityAPI

All URIs are relative to *https://requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1SystemsContentModeratePost**](DataIntegrityAPI.md#V1SystemsContentModeratePost) | **Post** /v1/systems/content/moderate | Moderate Content
[**V1SystemsTextNormalizePost**](DataIntegrityAPI.md#V1SystemsTextNormalizePost) | **Post** /v1/systems/text/normalize | Normalize Text



## V1SystemsContentModeratePost

> V1SystemsContentModeratePost200Response V1SystemsContentModeratePost(ctx).V1SystemsContentModeratePostRequest(v1SystemsContentModeratePostRequest).Execute()

Moderate Content



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
	v1SystemsContentModeratePostRequest := *openapiclient.NewV1SystemsContentModeratePostRequest("This is some user-generated content.") // V1SystemsContentModeratePostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataIntegrityAPI.V1SystemsContentModeratePost(context.Background()).V1SystemsContentModeratePostRequest(v1SystemsContentModeratePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataIntegrityAPI.V1SystemsContentModeratePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SystemsContentModeratePost`: V1SystemsContentModeratePost200Response
	fmt.Fprintf(os.Stdout, "Response from `DataIntegrityAPI.V1SystemsContentModeratePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1SystemsContentModeratePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1SystemsContentModeratePostRequest** | [**V1SystemsContentModeratePostRequest**](V1SystemsContentModeratePostRequest.md) |  | 

### Return type

[**V1SystemsContentModeratePost200Response**](V1SystemsContentModeratePost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SystemsTextNormalizePost

> V1SystemsTextNormalizePost200Response V1SystemsTextNormalizePost(ctx).V1SystemsTextNormalizePostRequest(v1SystemsTextNormalizePostRequest).Execute()

Normalize Text



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
	v1SystemsTextNormalizePostRequest := *openapiclient.NewV1SystemsTextNormalizePostRequest("  Hello,   World!  ", []interface{}{nil}) // V1SystemsTextNormalizePostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataIntegrityAPI.V1SystemsTextNormalizePost(context.Background()).V1SystemsTextNormalizePostRequest(v1SystemsTextNormalizePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataIntegrityAPI.V1SystemsTextNormalizePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SystemsTextNormalizePost`: V1SystemsTextNormalizePost200Response
	fmt.Fprintf(os.Stdout, "Response from `DataIntegrityAPI.V1SystemsTextNormalizePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1SystemsTextNormalizePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1SystemsTextNormalizePostRequest** | [**V1SystemsTextNormalizePostRequest**](V1SystemsTextNormalizePostRequest.md) |  | 

### Return type

[**V1SystemsTextNormalizePost200Response**](V1SystemsTextNormalizePost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

