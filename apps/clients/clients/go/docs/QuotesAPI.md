# \QuotesAPI

All URIs are relative to *https://requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EntertainmentQuotesRandomBatchPost**](QuotesAPI.md#V1EntertainmentQuotesRandomBatchPost) | **Post** /v1/entertainment/quotes/random/batch | Get Random Quotes (Batch)
[**V1EntertainmentQuotesRandomGet**](QuotesAPI.md#V1EntertainmentQuotesRandomGet) | **Get** /v1/entertainment/quotes/random | Get Random Quote



## V1EntertainmentQuotesRandomBatchPost

> V1EntertainmentQuotesRandomBatchPost200Response V1EntertainmentQuotesRandomBatchPost(ctx).V1EntertainmentQuotesRandomBatchPostRequest(v1EntertainmentQuotesRandomBatchPostRequest).Execute()

Get Random Quotes (Batch)



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
	v1EntertainmentQuotesRandomBatchPostRequest := *openapiclient.NewV1EntertainmentQuotesRandomBatchPostRequest(int32(5)) // V1EntertainmentQuotesRandomBatchPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotesAPI.V1EntertainmentQuotesRandomBatchPost(context.Background()).V1EntertainmentQuotesRandomBatchPostRequest(v1EntertainmentQuotesRandomBatchPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.V1EntertainmentQuotesRandomBatchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1EntertainmentQuotesRandomBatchPost`: V1EntertainmentQuotesRandomBatchPost200Response
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.V1EntertainmentQuotesRandomBatchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1EntertainmentQuotesRandomBatchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1EntertainmentQuotesRandomBatchPostRequest** | [**V1EntertainmentQuotesRandomBatchPostRequest**](V1EntertainmentQuotesRandomBatchPostRequest.md) |  | 

### Return type

[**V1EntertainmentQuotesRandomBatchPost200Response**](V1EntertainmentQuotesRandomBatchPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EntertainmentQuotesRandomGet

> V1EntertainmentQuotesRandomGet200Response V1EntertainmentQuotesRandomGet(ctx).Execute()

Get Random Quote



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
	resp, r, err := apiClient.QuotesAPI.V1EntertainmentQuotesRandomGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.V1EntertainmentQuotesRandomGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1EntertainmentQuotesRandomGet`: V1EntertainmentQuotesRandomGet200Response
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.V1EntertainmentQuotesRandomGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1EntertainmentQuotesRandomGetRequest struct via the builder pattern


### Return type

[**V1EntertainmentQuotesRandomGet200Response**](V1EntertainmentQuotesRandomGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

