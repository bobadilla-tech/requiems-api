# openapi_client.DisposableEmailApi

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_networking_disposable_batch_post**](DisposableEmailApi.md#v1_networking_disposable_batch_post) | **POST** /v1/networking/disposable/batch | Check Batch Emails
[**v1_networking_disposable_check_post**](DisposableEmailApi.md#v1_networking_disposable_check_post) | **POST** /v1/networking/disposable/check | Check Single Email
[**v1_networking_disposable_domain_domain_get**](DisposableEmailApi.md#v1_networking_disposable_domain_domain_get) | **GET** /v1/networking/disposable/domain/{domain} | Check Domain
[**v1_networking_disposable_domains_get**](DisposableEmailApi.md#v1_networking_disposable_domains_get) | **GET** /v1/networking/disposable/domains | List Domains (Paginated)
[**v1_networking_disposable_stats_get**](DisposableEmailApi.md#v1_networking_disposable_stats_get) | **GET** /v1/networking/disposable/stats | Get Statistics


# **v1_networking_disposable_batch_post**
> V1NetworkingDisposableBatchPost200Response v1_networking_disposable_batch_post(v1_networking_disposable_batch_post_request)

Check Batch Emails

Validate multiple email addresses in a single request (max 100 emails)

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_networking_disposable_batch_post200_response import V1NetworkingDisposableBatchPost200Response
from openapi_client.models.v1_networking_disposable_batch_post_request import V1NetworkingDisposableBatchPostRequest
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
    api_instance = openapi_client.DisposableEmailApi(api_client)
    v1_networking_disposable_batch_post_request = openapi_client.V1NetworkingDisposableBatchPostRequest() # V1NetworkingDisposableBatchPostRequest | 

    try:
        # Check Batch Emails
        api_response = api_instance.v1_networking_disposable_batch_post(v1_networking_disposable_batch_post_request)
        print("The response of DisposableEmailApi->v1_networking_disposable_batch_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DisposableEmailApi->v1_networking_disposable_batch_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1_networking_disposable_batch_post_request** | [**V1NetworkingDisposableBatchPostRequest**](V1NetworkingDisposableBatchPostRequest.md)|  | 

### Return type

[**V1NetworkingDisposableBatchPost200Response**](V1NetworkingDisposableBatchPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**400** | The request body is missing or malformed; The emails field is missing; Too many emails in the request |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1_networking_disposable_check_post**
> V1NetworkingDisposableCheckPost200Response v1_networking_disposable_check_post(v1_networking_disposable_check_post_request)

Check Single Email

Validate whether an email address uses a disposable domain

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_networking_disposable_check_post200_response import V1NetworkingDisposableCheckPost200Response
from openapi_client.models.v1_networking_disposable_check_post_request import V1NetworkingDisposableCheckPostRequest
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
    api_instance = openapi_client.DisposableEmailApi(api_client)
    v1_networking_disposable_check_post_request = openapi_client.V1NetworkingDisposableCheckPostRequest() # V1NetworkingDisposableCheckPostRequest | 

    try:
        # Check Single Email
        api_response = api_instance.v1_networking_disposable_check_post(v1_networking_disposable_check_post_request)
        print("The response of DisposableEmailApi->v1_networking_disposable_check_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DisposableEmailApi->v1_networking_disposable_check_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1_networking_disposable_check_post_request** | [**V1NetworkingDisposableCheckPostRequest**](V1NetworkingDisposableCheckPostRequest.md)|  | 

### Return type

[**V1NetworkingDisposableCheckPost200Response**](V1NetworkingDisposableCheckPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**400** | The request body is missing or malformed; The email address format is invalid |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1_networking_disposable_domain_domain_get**
> V1NetworkingDisposableDomainDomainGet200Response v1_networking_disposable_domain_domain_get(domain)

Check Domain

Check if a specific domain is in the disposable blocklist

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_networking_disposable_domain_domain_get200_response import V1NetworkingDisposableDomainDomainGet200Response
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
    api_instance = openapi_client.DisposableEmailApi(api_client)
    domain = 'tempmail.com' # str | The domain to check

    try:
        # Check Domain
        api_response = api_instance.v1_networking_disposable_domain_domain_get(domain)
        print("The response of DisposableEmailApi->v1_networking_disposable_domain_domain_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DisposableEmailApi->v1_networking_disposable_domain_domain_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domain** | **str**| The domain to check | 

### Return type

[**V1NetworkingDisposableDomainDomainGet200Response**](V1NetworkingDisposableDomainDomainGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**400** | The domain parameter is missing |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1_networking_disposable_domains_get**
> V1NetworkingDisposableDomainsGet200Response v1_networking_disposable_domains_get(page=page, per_page=per_page)

List Domains (Paginated)

Get a paginated list of all disposable domains in the blocklist

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_networking_disposable_domains_get200_response import V1NetworkingDisposableDomainsGet200Response
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
    api_instance = openapi_client.DisposableEmailApi(api_client)
    page = 1 # int | Page number (default: 1) (optional)
    per_page = 100 # int | Items per page (default: 100, max: 1000) (optional)

    try:
        # List Domains (Paginated)
        api_response = api_instance.v1_networking_disposable_domains_get(page=page, per_page=per_page)
        print("The response of DisposableEmailApi->v1_networking_disposable_domains_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DisposableEmailApi->v1_networking_disposable_domains_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int**| Page number (default: 1) | [optional] 
 **per_page** | **int**| Items per page (default: 100, max: 1000) | [optional] 

### Return type

[**V1NetworkingDisposableDomainsGet200Response**](V1NetworkingDisposableDomainsGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1_networking_disposable_stats_get**
> V1NetworkingDisposableStatsGet200Response v1_networking_disposable_stats_get()

Get Statistics

Get statistics about the disposable email blocklist

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_networking_disposable_stats_get200_response import V1NetworkingDisposableStatsGet200Response
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
    api_instance = openapi_client.DisposableEmailApi(api_client)

    try:
        # Get Statistics
        api_response = api_instance.v1_networking_disposable_stats_get()
        print("The response of DisposableEmailApi->v1_networking_disposable_stats_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DisposableEmailApi->v1_networking_disposable_stats_get: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

[**V1NetworkingDisposableStatsGet200Response**](V1NetworkingDisposableStatsGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

