# \FactsAPI

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EntertainmentFactsGet**](FactsAPI.md#V1EntertainmentFactsGet) | **Get** /v1/entertainment/facts | Get Random Fact



## V1EntertainmentFactsGet

> V1EntertainmentFactsGet200Response V1EntertainmentFactsGet(ctx).Category(category).Execute()

Get Random Fact



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
	category := "science" // string | Filter by category. Valid values: science, history, technology, nature, space, food (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FactsAPI.V1EntertainmentFactsGet(context.Background()).Category(category).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FactsAPI.V1EntertainmentFactsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1EntertainmentFactsGet`: V1EntertainmentFactsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `FactsAPI.V1EntertainmentFactsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1EntertainmentFactsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **category** | **string** | Filter by category. Valid values: science, history, technology, nature, space, food | 

### Return type

[**V1EntertainmentFactsGet200Response**](V1EntertainmentFactsGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

