# \EmojiAPI

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EntertainmentEmojiNameGet**](EmojiAPI.md#V1EntertainmentEmojiNameGet) | **Get** /v1/entertainment/emoji/{name} | Get Emoji by Name
[**V1EntertainmentEmojiRandomGet**](EmojiAPI.md#V1EntertainmentEmojiRandomGet) | **Get** /v1/entertainment/emoji/random | Get Random Emoji
[**V1EntertainmentEmojiSearchGet**](EmojiAPI.md#V1EntertainmentEmojiSearchGet) | **Get** /v1/entertainment/emoji/search | Search Emoji



## V1EntertainmentEmojiNameGet

> V1EntertainmentEmojiNameGet200Response V1EntertainmentEmojiNameGet(ctx, name).Execute()

Get Emoji by Name



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
	name := "grinning_face" // string | CLDR snake_case emoji name (e.g. grinning_face, thumbs_up)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmojiAPI.V1EntertainmentEmojiNameGet(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmojiAPI.V1EntertainmentEmojiNameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1EntertainmentEmojiNameGet`: V1EntertainmentEmojiNameGet200Response
	fmt.Fprintf(os.Stdout, "Response from `EmojiAPI.V1EntertainmentEmojiNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | CLDR snake_case emoji name (e.g. grinning_face, thumbs_up) | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EntertainmentEmojiNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1EntertainmentEmojiNameGet200Response**](V1EntertainmentEmojiNameGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EntertainmentEmojiRandomGet

> V1EntertainmentEmojiRandomGet200Response V1EntertainmentEmojiRandomGet(ctx).Execute()

Get Random Emoji



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
	resp, r, err := apiClient.EmojiAPI.V1EntertainmentEmojiRandomGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmojiAPI.V1EntertainmentEmojiRandomGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1EntertainmentEmojiRandomGet`: V1EntertainmentEmojiRandomGet200Response
	fmt.Fprintf(os.Stdout, "Response from `EmojiAPI.V1EntertainmentEmojiRandomGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1EntertainmentEmojiRandomGetRequest struct via the builder pattern


### Return type

[**V1EntertainmentEmojiRandomGet200Response**](V1EntertainmentEmojiRandomGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EntertainmentEmojiSearchGet

> V1EntertainmentEmojiSearchGet200Response V1EntertainmentEmojiSearchGet(ctx).Q(q).Execute()

Search Emoji



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
	q := "happy" // string | Search term to match against emoji names and categories (e.g. happy, heart, food)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmojiAPI.V1EntertainmentEmojiSearchGet(context.Background()).Q(q).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmojiAPI.V1EntertainmentEmojiSearchGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1EntertainmentEmojiSearchGet`: V1EntertainmentEmojiSearchGet200Response
	fmt.Fprintf(os.Stdout, "Response from `EmojiAPI.V1EntertainmentEmojiSearchGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1EntertainmentEmojiSearchGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Search term to match against emoji names and categories (e.g. happy, heart, food) | 

### Return type

[**V1EntertainmentEmojiSearchGet200Response**](V1EntertainmentEmojiSearchGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

