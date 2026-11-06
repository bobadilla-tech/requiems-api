# \TriviaAPI

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EntertainmentTriviaGet**](TriviaAPI.md#V1EntertainmentTriviaGet) | **Get** /v1/entertainment/trivia | Get Trivia Question



## V1EntertainmentTriviaGet

> V1EntertainmentTriviaGet200Response V1EntertainmentTriviaGet(ctx).Category(category).Difficulty(difficulty).Execute()

Get Trivia Question



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
	category := "category_example" // string | Filter by category. One of: science, history, geography, sports, music, movies, literature, math, technology, nature. (optional)
	difficulty := "difficulty_example" // string | Filter by difficulty. One of: easy, medium, hard. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TriviaAPI.V1EntertainmentTriviaGet(context.Background()).Category(category).Difficulty(difficulty).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TriviaAPI.V1EntertainmentTriviaGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1EntertainmentTriviaGet`: V1EntertainmentTriviaGet200Response
	fmt.Fprintf(os.Stdout, "Response from `TriviaAPI.V1EntertainmentTriviaGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1EntertainmentTriviaGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **category** | **string** | Filter by category. One of: science, history, geography, sports, music, movies, literature, math, technology, nature. | 
 **difficulty** | **string** | Filter by difficulty. One of: easy, medium, hard. | 

### Return type

[**V1EntertainmentTriviaGet200Response**](V1EntertainmentTriviaGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

