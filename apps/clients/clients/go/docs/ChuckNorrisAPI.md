# \ChuckNorrisAPI

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EntertainmentChuckNorrisGet**](ChuckNorrisAPI.md#V1EntertainmentChuckNorrisGet) | **Get** /v1/entertainment/chuck-norris | Get Random Chuck Norris Fact



## V1EntertainmentChuckNorrisGet

> V1EntertainmentChuckNorrisGet200Response V1EntertainmentChuckNorrisGet(ctx).Execute()

Get Random Chuck Norris Fact



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
	resp, r, err := apiClient.ChuckNorrisAPI.V1EntertainmentChuckNorrisGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChuckNorrisAPI.V1EntertainmentChuckNorrisGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1EntertainmentChuckNorrisGet`: V1EntertainmentChuckNorrisGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ChuckNorrisAPI.V1EntertainmentChuckNorrisGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1EntertainmentChuckNorrisGetRequest struct via the builder pattern


### Return type

[**V1EntertainmentChuckNorrisGet200Response**](V1EntertainmentChuckNorrisGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

