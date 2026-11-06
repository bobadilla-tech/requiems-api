# \IdentityRiskAPI

All URIs are relative to *https://requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1SystemsRiskScorePost**](IdentityRiskAPI.md#V1SystemsRiskScorePost) | **Post** /v1/systems/risk/score | Score Risk
[**V1SystemsSignupProtectPost**](IdentityRiskAPI.md#V1SystemsSignupProtectPost) | **Post** /v1/systems/signup/protect | Protect Signup
[**V1SystemsUserVerifyPost**](IdentityRiskAPI.md#V1SystemsUserVerifyPost) | **Post** /v1/systems/user/verify | Verify User



## V1SystemsRiskScorePost

> V1SystemsRiskScorePost200Response V1SystemsRiskScorePost(ctx).V1SystemsRiskScorePostRequest(v1SystemsRiskScorePostRequest).Execute()

Score Risk



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
	v1SystemsRiskScorePostRequest := *openapiclient.NewV1SystemsRiskScorePostRequest() // V1SystemsRiskScorePostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityRiskAPI.V1SystemsRiskScorePost(context.Background()).V1SystemsRiskScorePostRequest(v1SystemsRiskScorePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityRiskAPI.V1SystemsRiskScorePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SystemsRiskScorePost`: V1SystemsRiskScorePost200Response
	fmt.Fprintf(os.Stdout, "Response from `IdentityRiskAPI.V1SystemsRiskScorePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1SystemsRiskScorePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1SystemsRiskScorePostRequest** | [**V1SystemsRiskScorePostRequest**](V1SystemsRiskScorePostRequest.md) |  | 

### Return type

[**V1SystemsRiskScorePost200Response**](V1SystemsRiskScorePost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SystemsSignupProtectPost

> V1SystemsSignupProtectPost200Response V1SystemsSignupProtectPost(ctx).V1SystemsSignupProtectPostRequest(v1SystemsSignupProtectPostRequest).Execute()

Protect Signup



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
	v1SystemsSignupProtectPostRequest := *openapiclient.NewV1SystemsSignupProtectPostRequest() // V1SystemsSignupProtectPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityRiskAPI.V1SystemsSignupProtectPost(context.Background()).V1SystemsSignupProtectPostRequest(v1SystemsSignupProtectPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityRiskAPI.V1SystemsSignupProtectPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SystemsSignupProtectPost`: V1SystemsSignupProtectPost200Response
	fmt.Fprintf(os.Stdout, "Response from `IdentityRiskAPI.V1SystemsSignupProtectPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1SystemsSignupProtectPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1SystemsSignupProtectPostRequest** | [**V1SystemsSignupProtectPostRequest**](V1SystemsSignupProtectPostRequest.md) |  | 

### Return type

[**V1SystemsSignupProtectPost200Response**](V1SystemsSignupProtectPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SystemsUserVerifyPost

> V1SystemsUserVerifyPost200Response V1SystemsUserVerifyPost(ctx).V1SystemsUserVerifyPostRequest(v1SystemsUserVerifyPostRequest).Execute()

Verify User



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
	v1SystemsUserVerifyPostRequest := *openapiclient.NewV1SystemsUserVerifyPostRequest("user@example.com") // V1SystemsUserVerifyPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityRiskAPI.V1SystemsUserVerifyPost(context.Background()).V1SystemsUserVerifyPostRequest(v1SystemsUserVerifyPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityRiskAPI.V1SystemsUserVerifyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SystemsUserVerifyPost`: V1SystemsUserVerifyPost200Response
	fmt.Fprintf(os.Stdout, "Response from `IdentityRiskAPI.V1SystemsUserVerifyPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1SystemsUserVerifyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1SystemsUserVerifyPostRequest** | [**V1SystemsUserVerifyPostRequest**](V1SystemsUserVerifyPostRequest.md) |  | 

### Return type

[**V1SystemsUserVerifyPost200Response**](V1SystemsUserVerifyPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

