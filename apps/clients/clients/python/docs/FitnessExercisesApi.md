# openapi_client.FitnessExercisesApi

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_health_body_parts_get**](FitnessExercisesApi.md#v1_health_body_parts_get) | **GET** /v1/health/body-parts | List Body Parts
[**v1_health_equipment_get**](FitnessExercisesApi.md#v1_health_equipment_get) | **GET** /v1/health/equipment | List Equipment
[**v1_health_exercises_get**](FitnessExercisesApi.md#v1_health_exercises_get) | **GET** /v1/health/exercises | List Exercises
[**v1_health_exercises_id_get**](FitnessExercisesApi.md#v1_health_exercises_id_get) | **GET** /v1/health/exercises/{id} | Get Exercise by ID
[**v1_health_exercises_random_get**](FitnessExercisesApi.md#v1_health_exercises_random_get) | **GET** /v1/health/exercises/random | Random Exercise
[**v1_health_muscles_get**](FitnessExercisesApi.md#v1_health_muscles_get) | **GET** /v1/health/muscles | List Muscles


# **v1_health_body_parts_get**
> V1HealthBodyPartsGet200Response v1_health_body_parts_get()

List Body Parts

Returns a sorted list of all distinct body part values present in the dataset. Use these as valid values for the body_part filter.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_health_body_parts_get200_response import V1HealthBodyPartsGet200Response
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.requiems.xyz
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://api.requiems.xyz"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: requiems-api-key
configuration.api_key['requiems-api-key'] = os.environ["API_KEY"]

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['requiems-api-key'] = 'Bearer'

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.FitnessExercisesApi(api_client)

    try:
        # List Body Parts
        api_response = api_instance.v1_health_body_parts_get()
        print("The response of FitnessExercisesApi->v1_health_body_parts_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FitnessExercisesApi->v1_health_body_parts_get: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

[**V1HealthBodyPartsGet200Response**](V1HealthBodyPartsGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1_health_equipment_get**
> V1HealthEquipmentGet200Response v1_health_equipment_get()

List Equipment

Returns a sorted list of all distinct equipment types. Use these as valid values for the equipment filter.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_health_equipment_get200_response import V1HealthEquipmentGet200Response
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.requiems.xyz
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://api.requiems.xyz"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: requiems-api-key
configuration.api_key['requiems-api-key'] = os.environ["API_KEY"]

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['requiems-api-key'] = 'Bearer'

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.FitnessExercisesApi(api_client)

    try:
        # List Equipment
        api_response = api_instance.v1_health_equipment_get()
        print("The response of FitnessExercisesApi->v1_health_equipment_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FitnessExercisesApi->v1_health_equipment_get: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

[**V1HealthEquipmentGet200Response**](V1HealthEquipmentGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1_health_exercises_get**
> V1HealthExercisesGet200Response v1_health_exercises_get(body_part=body_part, equipment=equipment, muscle=muscle, search=search, page=page, per_page=per_page)

List Exercises

Returns a paginated list of exercises. All filter parameters are optional and combinable.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_health_exercises_get200_response import V1HealthExercisesGet200Response
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.requiems.xyz
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://api.requiems.xyz"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: requiems-api-key
configuration.api_key['requiems-api-key'] = os.environ["API_KEY"]

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['requiems-api-key'] = 'Bearer'

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.FitnessExercisesApi(api_client)
    body_part = 'chest' # str | Filter by body part (e.g. chest, back, upper legs). Use /v1/health/body-parts for valid values. (optional)
    equipment = 'barbell' # str | Filter by equipment type (e.g. barbell, dumbbell, body weight). Use /v1/health/equipment for valid values. (optional)
    muscle = 'biceps' # str | Filter by target or secondary muscle (e.g. biceps, glutes). Use /v1/health/muscles for valid values. (optional)
    search = 'bench press' # str | Full-text search on exercise name. (optional)
    page = 1 # int | Page number (default: 1) (optional)
    per_page = 20 # int | Results per page, 1–100 (default: 20) (optional)

    try:
        # List Exercises
        api_response = api_instance.v1_health_exercises_get(body_part=body_part, equipment=equipment, muscle=muscle, search=search, page=page, per_page=per_page)
        print("The response of FitnessExercisesApi->v1_health_exercises_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FitnessExercisesApi->v1_health_exercises_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body_part** | **str**| Filter by body part (e.g. chest, back, upper legs). Use /v1/health/body-parts for valid values. | [optional] 
 **equipment** | **str**| Filter by equipment type (e.g. barbell, dumbbell, body weight). Use /v1/health/equipment for valid values. | [optional] 
 **muscle** | **str**| Filter by target or secondary muscle (e.g. biceps, glutes). Use /v1/health/muscles for valid values. | [optional] 
 **search** | **str**| Full-text search on exercise name. | [optional] 
 **page** | **int**| Page number (default: 1) | [optional] 
 **per_page** | **int**| Results per page, 1–100 (default: 20) | [optional] 

### Return type

[**V1HealthExercisesGet200Response**](V1HealthExercisesGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**400** | A query parameter has an invalid value (e.g. per_page out of range). |  -  |
**500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1_health_exercises_id_get**
> V1HealthExercisesIdGet200Response v1_health_exercises_id_get(id)

Get Exercise by ID

Returns a single exercise by its numeric ID.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_health_exercises_id_get200_response import V1HealthExercisesIdGet200Response
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.requiems.xyz
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://api.requiems.xyz"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: requiems-api-key
configuration.api_key['requiems-api-key'] = os.environ["API_KEY"]

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['requiems-api-key'] = 'Bearer'

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.FitnessExercisesApi(api_client)
    id = 1 # int | Numeric exercise ID

    try:
        # Get Exercise by ID
        api_response = api_instance.v1_health_exercises_id_get(id)
        print("The response of FitnessExercisesApi->v1_health_exercises_id_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FitnessExercisesApi->v1_health_exercises_id_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Numeric exercise ID | 

### Return type

[**V1HealthExercisesIdGet200Response**](V1HealthExercisesIdGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**400** | The id parameter is not a positive integer. |  -  |
**404** | No exercise exists with the given ID. |  -  |
**500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1_health_exercises_random_get**
> V1HealthExercisesRandomGet200Response v1_health_exercises_random_get(body_part=body_part, equipment=equipment, muscle=muscle, search=search)

Random Exercise

Returns a single randomly selected exercise. Accepts the same filter parameters as the list endpoint, so you can get a random chest exercise, a random bodyweight exercise, etc.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_health_exercises_random_get200_response import V1HealthExercisesRandomGet200Response
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.requiems.xyz
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://api.requiems.xyz"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: requiems-api-key
configuration.api_key['requiems-api-key'] = os.environ["API_KEY"]

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['requiems-api-key'] = 'Bearer'

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.FitnessExercisesApi(api_client)
    body_part = 'back' # str | Restrict random selection to this body part. (optional)
    equipment = 'body weight' # str | Restrict random selection to this equipment type. (optional)
    muscle = 'glutes' # str | Restrict random selection to exercises targeting this muscle. (optional)
    search = 'curl' # str | Restrict random selection to exercises matching this search term. (optional)

    try:
        # Random Exercise
        api_response = api_instance.v1_health_exercises_random_get(body_part=body_part, equipment=equipment, muscle=muscle, search=search)
        print("The response of FitnessExercisesApi->v1_health_exercises_random_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FitnessExercisesApi->v1_health_exercises_random_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body_part** | **str**| Restrict random selection to this body part. | [optional] 
 **equipment** | **str**| Restrict random selection to this equipment type. | [optional] 
 **muscle** | **str**| Restrict random selection to exercises targeting this muscle. | [optional] 
 **search** | **str**| Restrict random selection to exercises matching this search term. | [optional] 

### Return type

[**V1HealthExercisesRandomGet200Response**](V1HealthExercisesRandomGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**404** | No exercises match the given filters. |  -  |
**500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1_health_muscles_get**
> V1HealthMusclesGet200Response v1_health_muscles_get()

List Muscles

Returns a sorted list of all distinct muscle names (combining target and secondary muscles). Use these as valid values for the muscle filter.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_health_muscles_get200_response import V1HealthMusclesGet200Response
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.requiems.xyz
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://api.requiems.xyz"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: requiems-api-key
configuration.api_key['requiems-api-key'] = os.environ["API_KEY"]

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['requiems-api-key'] = 'Bearer'

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.FitnessExercisesApi(api_client)

    try:
        # List Muscles
        api_response = api_instance.v1_health_muscles_get()
        print("The response of FitnessExercisesApi->v1_health_muscles_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FitnessExercisesApi->v1_health_muscles_get: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

[**V1HealthMusclesGet200Response**](V1HealthMusclesGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

