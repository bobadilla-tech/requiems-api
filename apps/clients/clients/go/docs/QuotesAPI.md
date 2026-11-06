# \QuotesAPI

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1TextQuotesRandomGet**](QuotesAPI.md#V1TextQuotesRandomGet) | **Get** /v1/text/quotes/random | Get Random Quote



## V1TextQuotesRandomGet

> V1TextQuotesRandomGet200Response V1TextQuotesRandomGet(ctx).Execute()

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
	resp, r, err := apiClient.QuotesAPI.V1TextQuotesRandomGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.V1TextQuotesRandomGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1TextQuotesRandomGet`: V1TextQuotesRandomGet200Response
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.V1TextQuotesRandomGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1TextQuotesRandomGetRequest struct via the builder pattern


### Return type

[**V1TextQuotesRandomGet200Response**](V1TextQuotesRandomGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

