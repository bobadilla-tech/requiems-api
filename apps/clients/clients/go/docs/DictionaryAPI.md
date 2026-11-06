# \DictionaryAPI

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1TextDictionaryWordGet**](DictionaryAPI.md#V1TextDictionaryWordGet) | **Get** /v1/text/dictionary/{word} | Dictionary Lookup



## V1TextDictionaryWordGet

> V1TextDictionaryWordGet200Response V1TextDictionaryWordGet(ctx, word).Execute()

Dictionary Lookup



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
	word := "ephemeral" // string | The word to look up in the dictionary

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DictionaryAPI.V1TextDictionaryWordGet(context.Background(), word).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DictionaryAPI.V1TextDictionaryWordGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1TextDictionaryWordGet`: V1TextDictionaryWordGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DictionaryAPI.V1TextDictionaryWordGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**word** | **string** | The word to look up in the dictionary | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1TextDictionaryWordGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1TextDictionaryWordGet200Response**](V1TextDictionaryWordGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

