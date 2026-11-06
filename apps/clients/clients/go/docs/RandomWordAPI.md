# \RandomWordAPI

All URIs are relative to *https://requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1TextWordsBatchPost**](RandomWordAPI.md#V1TextWordsBatchPost) | **Post** /v1/text/words/batch | Batch Define Words
[**V1TextWordsRandomGet**](RandomWordAPI.md#V1TextWordsRandomGet) | **Get** /v1/text/words/random | Get Random Word



## V1TextWordsBatchPost

> V1TextWordsBatchPost200Response V1TextWordsBatchPost(ctx).V1TextWordsBatchPostRequest(v1TextWordsBatchPostRequest).Execute()

Batch Define Words



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
	v1TextWordsBatchPostRequest := *openapiclient.NewV1TextWordsBatchPostRequest([]interface{}{nil}) // V1TextWordsBatchPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RandomWordAPI.V1TextWordsBatchPost(context.Background()).V1TextWordsBatchPostRequest(v1TextWordsBatchPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RandomWordAPI.V1TextWordsBatchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1TextWordsBatchPost`: V1TextWordsBatchPost200Response
	fmt.Fprintf(os.Stdout, "Response from `RandomWordAPI.V1TextWordsBatchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1TextWordsBatchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1TextWordsBatchPostRequest** | [**V1TextWordsBatchPostRequest**](V1TextWordsBatchPostRequest.md) |  | 

### Return type

[**V1TextWordsBatchPost200Response**](V1TextWordsBatchPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1TextWordsRandomGet

> V1TextWordsRandomGet200Response V1TextWordsRandomGet(ctx).Execute()

Get Random Word



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
	resp, r, err := apiClient.RandomWordAPI.V1TextWordsRandomGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RandomWordAPI.V1TextWordsRandomGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1TextWordsRandomGet`: V1TextWordsRandomGet200Response
	fmt.Fprintf(os.Stdout, "Response from `RandomWordAPI.V1TextWordsRandomGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1TextWordsRandomGetRequest struct via the builder pattern


### Return type

[**V1TextWordsRandomGet200Response**](V1TextWordsRandomGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

