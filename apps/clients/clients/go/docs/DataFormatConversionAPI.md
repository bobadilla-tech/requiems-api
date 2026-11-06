# \DataFormatConversionAPI

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1TechnologyFormatPost**](DataFormatConversionAPI.md#V1TechnologyFormatPost) | **Post** /v1/technology/format | Convert Format



## V1TechnologyFormatPost

> V1TechnologyFormatPost200Response V1TechnologyFormatPost(ctx).V1TechnologyFormatPostRequest(v1TechnologyFormatPostRequest).Execute()

Convert Format



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
	v1TechnologyFormatPostRequest := *openapiclient.NewV1TechnologyFormatPostRequest("json", "yaml", "{\"name\":\"Alice\",\"age\":30}") // V1TechnologyFormatPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataFormatConversionAPI.V1TechnologyFormatPost(context.Background()).V1TechnologyFormatPostRequest(v1TechnologyFormatPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataFormatConversionAPI.V1TechnologyFormatPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1TechnologyFormatPost`: V1TechnologyFormatPost200Response
	fmt.Fprintf(os.Stdout, "Response from `DataFormatConversionAPI.V1TechnologyFormatPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1TechnologyFormatPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1TechnologyFormatPostRequest** | [**V1TechnologyFormatPostRequest**](V1TechnologyFormatPostRequest.md) |  | 

### Return type

[**V1TechnologyFormatPost200Response**](V1TechnologyFormatPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

