# \DadJokesAPI

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EntertainmentJokesDadGet**](DadJokesAPI.md#V1EntertainmentJokesDadGet) | **Get** /v1/entertainment/jokes/dad | Get Random Dad Joke



## V1EntertainmentJokesDadGet

> V1EntertainmentJokesDadGet200Response V1EntertainmentJokesDadGet(ctx).Execute()

Get Random Dad Joke



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
	resp, r, err := apiClient.DadJokesAPI.V1EntertainmentJokesDadGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DadJokesAPI.V1EntertainmentJokesDadGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1EntertainmentJokesDadGet`: V1EntertainmentJokesDadGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DadJokesAPI.V1EntertainmentJokesDadGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1EntertainmentJokesDadGetRequest struct via the builder pattern


### Return type

[**V1EntertainmentJokesDadGet200Response**](V1EntertainmentJokesDadGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

