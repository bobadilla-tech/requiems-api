# \QrCodeAPI

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1TechnologyQrBase64Get**](QrCodeAPI.md#V1TechnologyQrBase64Get) | **Get** /v1/technology/qr/base64 | Generate QR Code (Base64 JSON)
[**V1TechnologyQrGet**](QrCodeAPI.md#V1TechnologyQrGet) | **Get** /v1/technology/qr | Generate QR Code (PNG)



## V1TechnologyQrBase64Get

> V1TechnologyQrBase64Get200Response V1TechnologyQrBase64Get(ctx).Data(data).Size(size).Recovery(recovery).Execute()

Generate QR Code (Base64 JSON)



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
	data := "https://example.com" // string | The text or URL to encode in the QR code
	size := int32(200) // int32 | Image size in pixels (default: 256, min: 50, max: 1000) (optional)
	recovery := "highest" // string | Error-correction level: low (7%), medium (15%), high (25%), highest (30%). Default: medium (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QrCodeAPI.V1TechnologyQrBase64Get(context.Background()).Data(data).Size(size).Recovery(recovery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QrCodeAPI.V1TechnologyQrBase64Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1TechnologyQrBase64Get`: V1TechnologyQrBase64Get200Response
	fmt.Fprintf(os.Stdout, "Response from `QrCodeAPI.V1TechnologyQrBase64Get`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1TechnologyQrBase64GetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | **string** | The text or URL to encode in the QR code | 
 **size** | **int32** | Image size in pixels (default: 256, min: 50, max: 1000) | 
 **recovery** | **string** | Error-correction level: low (7%), medium (15%), high (25%), highest (30%). Default: medium | 

### Return type

[**V1TechnologyQrBase64Get200Response**](V1TechnologyQrBase64Get200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1TechnologyQrGet

> V1TechnologyQrGet(ctx).Data(data).Size(size).Recovery(recovery).Execute()

Generate QR Code (PNG)



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
	data := "https://example.com" // string | The text or URL to encode in the QR code
	size := int32(200) // int32 | Image size in pixels (default: 256, min: 50, max: 1000) (optional)
	recovery := "high" // string | Error-correction level: low (7%), medium (15%), high (25%), highest (30%). Higher levels are more robust to physical damage but produce larger images. Default: medium (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.QrCodeAPI.V1TechnologyQrGet(context.Background()).Data(data).Size(size).Recovery(recovery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QrCodeAPI.V1TechnologyQrGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1TechnologyQrGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | **string** | The text or URL to encode in the QR code | 
 **size** | **int32** | Image size in pixels (default: 256, min: 50, max: 1000) | 
 **recovery** | **string** | Error-correction level: low (7%), medium (15%), high (25%), highest (30%). Higher levels are more robust to physical damage but produce larger images. Default: medium | 

### Return type

 (empty response body)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

