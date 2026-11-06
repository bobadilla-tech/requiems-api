# \SudokuAPI

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EntertainmentSudokuBatchPost**](SudokuAPI.md#V1EntertainmentSudokuBatchPost) | **Post** /v1/entertainment/sudoku/batch | Batch Generate Sudoku Puzzles
[**V1EntertainmentSudokuGet**](SudokuAPI.md#V1EntertainmentSudokuGet) | **Get** /v1/entertainment/sudoku | Get Sudoku Puzzle



## V1EntertainmentSudokuBatchPost

> V1EntertainmentSudokuBatchPost200Response V1EntertainmentSudokuBatchPost(ctx).V1EntertainmentSudokuBatchPostRequest(v1EntertainmentSudokuBatchPostRequest).Execute()

Batch Generate Sudoku Puzzles



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
	v1EntertainmentSudokuBatchPostRequest := *openapiclient.NewV1EntertainmentSudokuBatchPostRequest([]interface{}{nil}) // V1EntertainmentSudokuBatchPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SudokuAPI.V1EntertainmentSudokuBatchPost(context.Background()).V1EntertainmentSudokuBatchPostRequest(v1EntertainmentSudokuBatchPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SudokuAPI.V1EntertainmentSudokuBatchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1EntertainmentSudokuBatchPost`: V1EntertainmentSudokuBatchPost200Response
	fmt.Fprintf(os.Stdout, "Response from `SudokuAPI.V1EntertainmentSudokuBatchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1EntertainmentSudokuBatchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1EntertainmentSudokuBatchPostRequest** | [**V1EntertainmentSudokuBatchPostRequest**](V1EntertainmentSudokuBatchPostRequest.md) |  | 

### Return type

[**V1EntertainmentSudokuBatchPost200Response**](V1EntertainmentSudokuBatchPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EntertainmentSudokuGet

> V1EntertainmentSudokuGet200Response V1EntertainmentSudokuGet(ctx).Difficulty(difficulty).Execute()

Get Sudoku Puzzle



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
	difficulty := "difficulty_example" // string | Puzzle difficulty level. One of: easy, medium, hard. Defaults to medium. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SudokuAPI.V1EntertainmentSudokuGet(context.Background()).Difficulty(difficulty).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SudokuAPI.V1EntertainmentSudokuGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1EntertainmentSudokuGet`: V1EntertainmentSudokuGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SudokuAPI.V1EntertainmentSudokuGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1EntertainmentSudokuGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **difficulty** | **string** | Puzzle difficulty level. One of: easy, medium, hard. Defaults to medium. | 

### Return type

[**V1EntertainmentSudokuGet200Response**](V1EntertainmentSudokuGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

