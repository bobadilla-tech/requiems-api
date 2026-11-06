# FitnessExercisesApi

All URIs are relative to *https://api.requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1HealthBodyPartsGet**](#v1healthbodypartsget) | **GET** /v1/health/body-parts | List Body Parts|
|[**v1HealthEquipmentGet**](#v1healthequipmentget) | **GET** /v1/health/equipment | List Equipment|
|[**v1HealthExercisesGet**](#v1healthexercisesget) | **GET** /v1/health/exercises | List Exercises|
|[**v1HealthExercisesIdGet**](#v1healthexercisesidget) | **GET** /v1/health/exercises/{id} | Get Exercise by ID|
|[**v1HealthExercisesRandomGet**](#v1healthexercisesrandomget) | **GET** /v1/health/exercises/random | Random Exercise|
|[**v1HealthMusclesGet**](#v1healthmusclesget) | **GET** /v1/health/muscles | List Muscles|

# **v1HealthBodyPartsGet**
> V1HealthBodyPartsGet200Response v1HealthBodyPartsGet()

Returns a sorted list of all distinct body part values present in the dataset. Use these as valid values for the body_part filter.

### Example

```typescript
import {
    FitnessExercisesApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new FitnessExercisesApi(configuration);

const { status, data } = await apiInstance.v1HealthBodyPartsGet();
```

### Parameters
This endpoint does not have any parameters.


### Return type

**V1HealthBodyPartsGet200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1HealthEquipmentGet**
> V1HealthEquipmentGet200Response v1HealthEquipmentGet()

Returns a sorted list of all distinct equipment types. Use these as valid values for the equipment filter.

### Example

```typescript
import {
    FitnessExercisesApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new FitnessExercisesApi(configuration);

const { status, data } = await apiInstance.v1HealthEquipmentGet();
```

### Parameters
This endpoint does not have any parameters.


### Return type

**V1HealthEquipmentGet200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1HealthExercisesGet**
> V1HealthExercisesGet200Response v1HealthExercisesGet()

Returns a paginated list of exercises. All filter parameters are optional and combinable.

### Example

```typescript
import {
    FitnessExercisesApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new FitnessExercisesApi(configuration);

let bodyPart: string; //Filter by body part (e.g. chest, back, upper legs). Use /v1/health/body-parts for valid values. (optional) (default to undefined)
let equipment: string; //Filter by equipment type (e.g. barbell, dumbbell, body weight). Use /v1/health/equipment for valid values. (optional) (default to undefined)
let muscle: string; //Filter by target or secondary muscle (e.g. biceps, glutes). Use /v1/health/muscles for valid values. (optional) (default to undefined)
let search: string; //Full-text search on exercise name. (optional) (default to undefined)
let page: number; //Page number (default: 1) (optional) (default to undefined)
let perPage: number; //Results per page, 1–100 (default: 20) (optional) (default to undefined)

const { status, data } = await apiInstance.v1HealthExercisesGet(
    bodyPart,
    equipment,
    muscle,
    search,
    page,
    perPage
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **bodyPart** | [**string**] | Filter by body part (e.g. chest, back, upper legs). Use /v1/health/body-parts for valid values. | (optional) defaults to undefined|
| **equipment** | [**string**] | Filter by equipment type (e.g. barbell, dumbbell, body weight). Use /v1/health/equipment for valid values. | (optional) defaults to undefined|
| **muscle** | [**string**] | Filter by target or secondary muscle (e.g. biceps, glutes). Use /v1/health/muscles for valid values. | (optional) defaults to undefined|
| **search** | [**string**] | Full-text search on exercise name. | (optional) defaults to undefined|
| **page** | [**number**] | Page number (default: 1) | (optional) defaults to undefined|
| **perPage** | [**number**] | Results per page, 1–100 (default: 20) | (optional) defaults to undefined|


### Return type

**V1HealthExercisesGet200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**400** | A query parameter has an invalid value (e.g. per_page out of range). |  -  |
|**500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1HealthExercisesIdGet**
> V1HealthExercisesIdGet200Response v1HealthExercisesIdGet()

Returns a single exercise by its numeric ID.

### Example

```typescript
import {
    FitnessExercisesApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new FitnessExercisesApi(configuration);

let id: number; //Numeric exercise ID (default to undefined)

const { status, data } = await apiInstance.v1HealthExercisesIdGet(
    id
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **id** | [**number**] | Numeric exercise ID | defaults to undefined|


### Return type

**V1HealthExercisesIdGet200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**400** | The id parameter is not a positive integer. |  -  |
|**404** | No exercise exists with the given ID. |  -  |
|**500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1HealthExercisesRandomGet**
> V1HealthExercisesRandomGet200Response v1HealthExercisesRandomGet()

Returns a single randomly selected exercise. Accepts the same filter parameters as the list endpoint, so you can get a random chest exercise, a random bodyweight exercise, etc.

### Example

```typescript
import {
    FitnessExercisesApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new FitnessExercisesApi(configuration);

let bodyPart: string; //Restrict random selection to this body part. (optional) (default to undefined)
let equipment: string; //Restrict random selection to this equipment type. (optional) (default to undefined)
let muscle: string; //Restrict random selection to exercises targeting this muscle. (optional) (default to undefined)
let search: string; //Restrict random selection to exercises matching this search term. (optional) (default to undefined)

const { status, data } = await apiInstance.v1HealthExercisesRandomGet(
    bodyPart,
    equipment,
    muscle,
    search
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **bodyPart** | [**string**] | Restrict random selection to this body part. | (optional) defaults to undefined|
| **equipment** | [**string**] | Restrict random selection to this equipment type. | (optional) defaults to undefined|
| **muscle** | [**string**] | Restrict random selection to exercises targeting this muscle. | (optional) defaults to undefined|
| **search** | [**string**] | Restrict random selection to exercises matching this search term. | (optional) defaults to undefined|


### Return type

**V1HealthExercisesRandomGet200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**404** | No exercises match the given filters. |  -  |
|**500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1HealthMusclesGet**
> V1HealthMusclesGet200Response v1HealthMusclesGet()

Returns a sorted list of all distinct muscle names (combining target and secondary muscles). Use these as valid values for the muscle filter.

### Example

```typescript
import {
    FitnessExercisesApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new FitnessExercisesApi(configuration);

const { status, data } = await apiInstance.v1HealthMusclesGet();
```

### Parameters
This endpoint does not have any parameters.


### Return type

**V1HealthMusclesGet200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

