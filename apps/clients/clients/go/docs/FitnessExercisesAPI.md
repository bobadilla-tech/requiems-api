# \FitnessExercisesAPI

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1HealthBodyPartsGet**](FitnessExercisesAPI.md#V1HealthBodyPartsGet) | **Get** /v1/health/body-parts | List Body Parts
[**V1HealthEquipmentGet**](FitnessExercisesAPI.md#V1HealthEquipmentGet) | **Get** /v1/health/equipment | List Equipment
[**V1HealthExercisesGet**](FitnessExercisesAPI.md#V1HealthExercisesGet) | **Get** /v1/health/exercises | List Exercises
[**V1HealthExercisesIdGet**](FitnessExercisesAPI.md#V1HealthExercisesIdGet) | **Get** /v1/health/exercises/{id} | Get Exercise by ID
[**V1HealthExercisesRandomGet**](FitnessExercisesAPI.md#V1HealthExercisesRandomGet) | **Get** /v1/health/exercises/random | Random Exercise
[**V1HealthMusclesGet**](FitnessExercisesAPI.md#V1HealthMusclesGet) | **Get** /v1/health/muscles | List Muscles



## V1HealthBodyPartsGet

> V1HealthBodyPartsGet200Response V1HealthBodyPartsGet(ctx).Execute()

List Body Parts



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
	resp, r, err := apiClient.FitnessExercisesAPI.V1HealthBodyPartsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FitnessExercisesAPI.V1HealthBodyPartsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1HealthBodyPartsGet`: V1HealthBodyPartsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `FitnessExercisesAPI.V1HealthBodyPartsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1HealthBodyPartsGetRequest struct via the builder pattern


### Return type

[**V1HealthBodyPartsGet200Response**](V1HealthBodyPartsGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1HealthEquipmentGet

> V1HealthEquipmentGet200Response V1HealthEquipmentGet(ctx).Execute()

List Equipment



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
	resp, r, err := apiClient.FitnessExercisesAPI.V1HealthEquipmentGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FitnessExercisesAPI.V1HealthEquipmentGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1HealthEquipmentGet`: V1HealthEquipmentGet200Response
	fmt.Fprintf(os.Stdout, "Response from `FitnessExercisesAPI.V1HealthEquipmentGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1HealthEquipmentGetRequest struct via the builder pattern


### Return type

[**V1HealthEquipmentGet200Response**](V1HealthEquipmentGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1HealthExercisesGet

> V1HealthExercisesGet200Response V1HealthExercisesGet(ctx).BodyPart(bodyPart).Equipment(equipment).Muscle(muscle).Search(search).Page(page).PerPage(perPage).Execute()

List Exercises



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
	bodyPart := "chest" // string | Filter by body part (e.g. chest, back, upper legs). Use /v1/health/body-parts for valid values. (optional)
	equipment := "barbell" // string | Filter by equipment type (e.g. barbell, dumbbell, body weight). Use /v1/health/equipment for valid values. (optional)
	muscle := "biceps" // string | Filter by target or secondary muscle (e.g. biceps, glutes). Use /v1/health/muscles for valid values. (optional)
	search := "bench press" // string | Full-text search on exercise name. (optional)
	page := int32(1) // int32 | Page number (default: 1) (optional)
	perPage := int32(20) // int32 | Results per page, 1–100 (default: 20) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FitnessExercisesAPI.V1HealthExercisesGet(context.Background()).BodyPart(bodyPart).Equipment(equipment).Muscle(muscle).Search(search).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FitnessExercisesAPI.V1HealthExercisesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1HealthExercisesGet`: V1HealthExercisesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `FitnessExercisesAPI.V1HealthExercisesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1HealthExercisesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bodyPart** | **string** | Filter by body part (e.g. chest, back, upper legs). Use /v1/health/body-parts for valid values. | 
 **equipment** | **string** | Filter by equipment type (e.g. barbell, dumbbell, body weight). Use /v1/health/equipment for valid values. | 
 **muscle** | **string** | Filter by target or secondary muscle (e.g. biceps, glutes). Use /v1/health/muscles for valid values. | 
 **search** | **string** | Full-text search on exercise name. | 
 **page** | **int32** | Page number (default: 1) | 
 **perPage** | **int32** | Results per page, 1–100 (default: 20) | 

### Return type

[**V1HealthExercisesGet200Response**](V1HealthExercisesGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1HealthExercisesIdGet

> V1HealthExercisesIdGet200Response V1HealthExercisesIdGet(ctx, id).Execute()

Get Exercise by ID



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
	id := int32(1) // int32 | Numeric exercise ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FitnessExercisesAPI.V1HealthExercisesIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FitnessExercisesAPI.V1HealthExercisesIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1HealthExercisesIdGet`: V1HealthExercisesIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `FitnessExercisesAPI.V1HealthExercisesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Numeric exercise ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1HealthExercisesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1HealthExercisesIdGet200Response**](V1HealthExercisesIdGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1HealthExercisesRandomGet

> V1HealthExercisesRandomGet200Response V1HealthExercisesRandomGet(ctx).BodyPart(bodyPart).Equipment(equipment).Muscle(muscle).Search(search).Execute()

Random Exercise



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
	bodyPart := "back" // string | Restrict random selection to this body part. (optional)
	equipment := "body weight" // string | Restrict random selection to this equipment type. (optional)
	muscle := "glutes" // string | Restrict random selection to exercises targeting this muscle. (optional)
	search := "curl" // string | Restrict random selection to exercises matching this search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FitnessExercisesAPI.V1HealthExercisesRandomGet(context.Background()).BodyPart(bodyPart).Equipment(equipment).Muscle(muscle).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FitnessExercisesAPI.V1HealthExercisesRandomGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1HealthExercisesRandomGet`: V1HealthExercisesRandomGet200Response
	fmt.Fprintf(os.Stdout, "Response from `FitnessExercisesAPI.V1HealthExercisesRandomGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1HealthExercisesRandomGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bodyPart** | **string** | Restrict random selection to this body part. | 
 **equipment** | **string** | Restrict random selection to this equipment type. | 
 **muscle** | **string** | Restrict random selection to exercises targeting this muscle. | 
 **search** | **string** | Restrict random selection to exercises matching this search term. | 

### Return type

[**V1HealthExercisesRandomGet200Response**](V1HealthExercisesRandomGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1HealthMusclesGet

> V1HealthMusclesGet200Response V1HealthMusclesGet(ctx).Execute()

List Muscles



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
	resp, r, err := apiClient.FitnessExercisesAPI.V1HealthMusclesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FitnessExercisesAPI.V1HealthMusclesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1HealthMusclesGet`: V1HealthMusclesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `FitnessExercisesAPI.V1HealthMusclesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1HealthMusclesGetRequest struct via the builder pattern


### Return type

[**V1HealthMusclesGet200Response**](V1HealthMusclesGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

