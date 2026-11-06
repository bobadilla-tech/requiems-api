# \AdviceAPI

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1TextAdviceGet**](AdviceAPI.md#V1TextAdviceGet) | **Get** /v1/text/advice | Get Random Advice



## V1TextAdviceGet

> V1TextAdviceGet200Response V1TextAdviceGet(ctx).Execute()

Get Random Advice



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
	resp, r, err := apiClient.AdviceAPI.V1TextAdviceGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdviceAPI.V1TextAdviceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1TextAdviceGet`: V1TextAdviceGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AdviceAPI.V1TextAdviceGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1TextAdviceGetRequest struct via the builder pattern


### Return type

[**V1TextAdviceGet200Response**](V1TextAdviceGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

