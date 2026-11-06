# \ThesaurusAPI

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1TextThesaurusWordGet**](ThesaurusAPI.md#V1TextThesaurusWordGet) | **Get** /v1/text/thesaurus/{word} | Thesaurus Lookup



## V1TextThesaurusWordGet

> V1TextThesaurusWordGet200Response V1TextThesaurusWordGet(ctx, word).Execute()

Thesaurus Lookup



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
	word := "happy" // string | The word to look up in the thesaurus

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThesaurusAPI.V1TextThesaurusWordGet(context.Background(), word).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThesaurusAPI.V1TextThesaurusWordGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1TextThesaurusWordGet`: V1TextThesaurusWordGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ThesaurusAPI.V1TextThesaurusWordGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**word** | **string** | The word to look up in the thesaurus | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1TextThesaurusWordGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1TextThesaurusWordGet200Response**](V1TextThesaurusWordGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

