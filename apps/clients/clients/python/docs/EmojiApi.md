# openapi_client.EmojiApi

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_entertainment_emoji_name_get**](EmojiApi.md#v1_entertainment_emoji_name_get) | **GET** /v1/entertainment/emoji/{name} | Get Emoji by Name
[**v1_entertainment_emoji_random_get**](EmojiApi.md#v1_entertainment_emoji_random_get) | **GET** /v1/entertainment/emoji/random | Get Random Emoji
[**v1_entertainment_emoji_search_get**](EmojiApi.md#v1_entertainment_emoji_search_get) | **GET** /v1/entertainment/emoji/search | Search Emoji


# **v1_entertainment_emoji_name_get**
> V1EntertainmentEmojiNameGet200Response v1_entertainment_emoji_name_get(name)

Get Emoji by Name

Returns a specific emoji by its CLDR snake_case name. The name is case-insensitive.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_entertainment_emoji_name_get200_response import V1EntertainmentEmojiNameGet200Response
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
    api_instance = openapi_client.EmojiApi(api_client)
    name = 'grinning_face' # str | CLDR snake_case emoji name (e.g. grinning_face, thumbs_up)

    try:
        # Get Emoji by Name
        api_response = api_instance.v1_entertainment_emoji_name_get(name)
        print("The response of EmojiApi->v1_entertainment_emoji_name_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling EmojiApi->v1_entertainment_emoji_name_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **str**| CLDR snake_case emoji name (e.g. grinning_face, thumbs_up) | 

### Return type

[**V1EntertainmentEmojiNameGet200Response**](V1EntertainmentEmojiNameGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**404** | No emoji found with the given name. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1_entertainment_emoji_random_get**
> V1EntertainmentEmojiRandomGet200Response v1_entertainment_emoji_random_get()

Get Random Emoji

Returns a randomly selected emoji with its full metadata.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_entertainment_emoji_random_get200_response import V1EntertainmentEmojiRandomGet200Response
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
    api_instance = openapi_client.EmojiApi(api_client)

    try:
        # Get Random Emoji
        api_response = api_instance.v1_entertainment_emoji_random_get()
        print("The response of EmojiApi->v1_entertainment_emoji_random_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling EmojiApi->v1_entertainment_emoji_random_get: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

[**V1EntertainmentEmojiRandomGet200Response**](V1EntertainmentEmojiRandomGet200Response.md)

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

# **v1_entertainment_emoji_search_get**
> V1EntertainmentEmojiSearchGet200Response v1_entertainment_emoji_search_get(q)

Search Emoji

Search for emojis whose name or category contains the given query string (case-insensitive). Returns a list of all matches.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_entertainment_emoji_search_get200_response import V1EntertainmentEmojiSearchGet200Response
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
    api_instance = openapi_client.EmojiApi(api_client)
    q = 'happy' # str | Search term to match against emoji names and categories (e.g. happy, heart, food)

    try:
        # Search Emoji
        api_response = api_instance.v1_entertainment_emoji_search_get(q)
        print("The response of EmojiApi->v1_entertainment_emoji_search_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling EmojiApi->v1_entertainment_emoji_search_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **str**| Search term to match against emoji names and categories (e.g. happy, heart, food) | 

### Return type

[**V1EntertainmentEmojiSearchGet200Response**](V1EntertainmentEmojiSearchGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**400** | The q query parameter is missing or empty. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

