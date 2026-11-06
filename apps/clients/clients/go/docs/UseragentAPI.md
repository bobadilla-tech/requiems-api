# \UseragentAPI

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1TechnologyUseragentGet**](UseragentAPI.md#V1TechnologyUseragentGet) | **Get** /v1/technology/useragent | Parse User Agent



## V1TechnologyUseragentGet

> V1TechnologyUseragentGet200Response V1TechnologyUseragentGet(ctx).Ua(ua).Execute()

Parse User Agent



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
	ua := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36" // string | The user agent string to parse.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UseragentAPI.V1TechnologyUseragentGet(context.Background()).Ua(ua).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UseragentAPI.V1TechnologyUseragentGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1TechnologyUseragentGet`: V1TechnologyUseragentGet200Response
	fmt.Fprintf(os.Stdout, "Response from `UseragentAPI.V1TechnologyUseragentGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1TechnologyUseragentGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ua** | **string** | The user agent string to parse. | 

### Return type

[**V1TechnologyUseragentGet200Response**](V1TechnologyUseragentGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

