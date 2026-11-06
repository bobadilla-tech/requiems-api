# \RandomUserAPI

All URIs are relative to *https://requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1TechnologyRandomUserBatchPost**](RandomUserAPI.md#V1TechnologyRandomUserBatchPost) | **Post** /v1/technology/random-user/batch | Batch Generate Users
[**V1TechnologyRandomUserGet**](RandomUserAPI.md#V1TechnologyRandomUserGet) | **Get** /v1/technology/random-user | Get Random User



## V1TechnologyRandomUserBatchPost

> V1TechnologyRandomUserBatchPost200Response V1TechnologyRandomUserBatchPost(ctx).V1TechnologyRandomUserBatchPostRequest(v1TechnologyRandomUserBatchPostRequest).Execute()

Batch Generate Users



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
	v1TechnologyRandomUserBatchPostRequest := *openapiclient.NewV1TechnologyRandomUserBatchPostRequest(int32(2)) // V1TechnologyRandomUserBatchPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RandomUserAPI.V1TechnologyRandomUserBatchPost(context.Background()).V1TechnologyRandomUserBatchPostRequest(v1TechnologyRandomUserBatchPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RandomUserAPI.V1TechnologyRandomUserBatchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1TechnologyRandomUserBatchPost`: V1TechnologyRandomUserBatchPost200Response
	fmt.Fprintf(os.Stdout, "Response from `RandomUserAPI.V1TechnologyRandomUserBatchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1TechnologyRandomUserBatchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1TechnologyRandomUserBatchPostRequest** | [**V1TechnologyRandomUserBatchPostRequest**](V1TechnologyRandomUserBatchPostRequest.md) |  | 

### Return type

[**V1TechnologyRandomUserBatchPost200Response**](V1TechnologyRandomUserBatchPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1TechnologyRandomUserGet

> V1TechnologyRandomUserGet200Response V1TechnologyRandomUserGet(ctx).Execute()

Get Random User



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
	resp, r, err := apiClient.RandomUserAPI.V1TechnologyRandomUserGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RandomUserAPI.V1TechnologyRandomUserGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1TechnologyRandomUserGet`: V1TechnologyRandomUserGet200Response
	fmt.Fprintf(os.Stdout, "Response from `RandomUserAPI.V1TechnologyRandomUserGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1TechnologyRandomUserGetRequest struct via the builder pattern


### Return type

[**V1TechnologyRandomUserGet200Response**](V1TechnologyRandomUserGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

