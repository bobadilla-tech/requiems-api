# \Base64API

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1TechnologyBase64DecodePost**](Base64API.md#V1TechnologyBase64DecodePost) | **Post** /v1/technology/base64/decode | Decode
[**V1TechnologyBase64EncodePost**](Base64API.md#V1TechnologyBase64EncodePost) | **Post** /v1/technology/base64/encode | Encode



## V1TechnologyBase64DecodePost

> V1TechnologyBase64DecodePost200Response V1TechnologyBase64DecodePost(ctx).V1TechnologyBase64DecodePostRequest(v1TechnologyBase64DecodePostRequest).Execute()

Decode



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
	v1TechnologyBase64DecodePostRequest := *openapiclient.NewV1TechnologyBase64DecodePostRequest("SGVsbG8sIHdvcmxkIQ==") // V1TechnologyBase64DecodePostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Base64API.V1TechnologyBase64DecodePost(context.Background()).V1TechnologyBase64DecodePostRequest(v1TechnologyBase64DecodePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Base64API.V1TechnologyBase64DecodePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1TechnologyBase64DecodePost`: V1TechnologyBase64DecodePost200Response
	fmt.Fprintf(os.Stdout, "Response from `Base64API.V1TechnologyBase64DecodePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1TechnologyBase64DecodePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1TechnologyBase64DecodePostRequest** | [**V1TechnologyBase64DecodePostRequest**](V1TechnologyBase64DecodePostRequest.md) |  | 

### Return type

[**V1TechnologyBase64DecodePost200Response**](V1TechnologyBase64DecodePost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1TechnologyBase64EncodePost

> V1TechnologyBase64EncodePost200Response V1TechnologyBase64EncodePost(ctx).V1TechnologyBase64EncodePostRequest(v1TechnologyBase64EncodePostRequest).Execute()

Encode



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
	v1TechnologyBase64EncodePostRequest := *openapiclient.NewV1TechnologyBase64EncodePostRequest("Hello, world!") // V1TechnologyBase64EncodePostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Base64API.V1TechnologyBase64EncodePost(context.Background()).V1TechnologyBase64EncodePostRequest(v1TechnologyBase64EncodePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Base64API.V1TechnologyBase64EncodePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1TechnologyBase64EncodePost`: V1TechnologyBase64EncodePost200Response
	fmt.Fprintf(os.Stdout, "Response from `Base64API.V1TechnologyBase64EncodePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1TechnologyBase64EncodePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1TechnologyBase64EncodePostRequest** | [**V1TechnologyBase64EncodePostRequest**](V1TechnologyBase64EncodePostRequest.md) |  | 

### Return type

[**V1TechnologyBase64EncodePost200Response**](V1TechnologyBase64EncodePost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

