# \PasswordGeneratorAPI

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1TechnologyPasswordGet**](PasswordGeneratorAPI.md#V1TechnologyPasswordGet) | **Get** /v1/technology/password | Generate Password



## V1TechnologyPasswordGet

> V1TechnologyPasswordGet200Response V1TechnologyPasswordGet(ctx).Length(length).Uppercase(uppercase).Numbers(numbers).Symbols(symbols).Execute()

Generate Password



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
	length := int32(16) // int32 | Password length (8-128 characters) (optional)
	uppercase := true // bool | Include uppercase letters (A-Z) (optional)
	numbers := true // bool | Include numbers (0-9) (optional)
	symbols := true // bool | Include special characters (!@#$%^&*()-_=+[]{}|;:,.<>?) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PasswordGeneratorAPI.V1TechnologyPasswordGet(context.Background()).Length(length).Uppercase(uppercase).Numbers(numbers).Symbols(symbols).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PasswordGeneratorAPI.V1TechnologyPasswordGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1TechnologyPasswordGet`: V1TechnologyPasswordGet200Response
	fmt.Fprintf(os.Stdout, "Response from `PasswordGeneratorAPI.V1TechnologyPasswordGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1TechnologyPasswordGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **length** | **int32** | Password length (8-128 characters) | 
 **uppercase** | **bool** | Include uppercase letters (A-Z) | 
 **numbers** | **bool** | Include numbers (0-9) | 
 **symbols** | **bool** | Include special characters (!@#$%^&amp;*()-_&#x3D;+[]{}|;:,.&lt;&gt;?) | 

### Return type

[**V1TechnologyPasswordGet200Response**](V1TechnologyPasswordGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

