# openapi_client.IdentityRiskApi

All URIs are relative to *https://requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_systems_risk_score_post**](IdentityRiskApi.md#v1_systems_risk_score_post) | **POST** /v1/systems/risk/score | Score Risk
[**v1_systems_signup_protect_post**](IdentityRiskApi.md#v1_systems_signup_protect_post) | **POST** /v1/systems/signup/protect | Protect Signup
[**v1_systems_user_verify_post**](IdentityRiskApi.md#v1_systems_user_verify_post) | **POST** /v1/systems/user/verify | Verify User


# **v1_systems_risk_score_post**
> V1SystemsRiskScorePost200Response v1_systems_risk_score_post(v1_systems_risk_score_post_request)

Score Risk

Score a user for risk without the full signal breakdown. Lower latency than /signup/protect, suited for background re-scoring and rate limiting.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_systems_risk_score_post200_response import V1SystemsRiskScorePost200Response
from openapi_client.models.v1_systems_risk_score_post_request import V1SystemsRiskScorePostRequest
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://requiems.xyz
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://requiems.xyz"
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
    api_instance = openapi_client.IdentityRiskApi(api_client)
    v1_systems_risk_score_post_request = openapi_client.V1SystemsRiskScorePostRequest() # V1SystemsRiskScorePostRequest | 

    try:
        # Score Risk
        api_response = api_instance.v1_systems_risk_score_post(v1_systems_risk_score_post_request)
        print("The response of IdentityRiskApi->v1_systems_risk_score_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling IdentityRiskApi->v1_systems_risk_score_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1_systems_risk_score_post_request** | [**V1SystemsRiskScorePostRequest**](V1SystemsRiskScorePostRequest.md)|  | 

### Return type

[**V1SystemsRiskScorePost200Response**](V1SystemsRiskScorePost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**401** | Missing or invalid API key |  -  |
**422** | All of email, phone, and ip_address were omitted |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1_systems_signup_protect_post**
> V1SystemsSignupProtectPost200Response v1_systems_signup_protect_post(v1_systems_signup_protect_post_request)

Protect Signup

Evaluate a new user at signup. Returns a full risk decision with per-signal breakdown across email, phone, and IP.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_systems_signup_protect_post200_response import V1SystemsSignupProtectPost200Response
from openapi_client.models.v1_systems_signup_protect_post_request import V1SystemsSignupProtectPostRequest
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://requiems.xyz
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://requiems.xyz"
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
    api_instance = openapi_client.IdentityRiskApi(api_client)
    v1_systems_signup_protect_post_request = openapi_client.V1SystemsSignupProtectPostRequest() # V1SystemsSignupProtectPostRequest | 

    try:
        # Protect Signup
        api_response = api_instance.v1_systems_signup_protect_post(v1_systems_signup_protect_post_request)
        print("The response of IdentityRiskApi->v1_systems_signup_protect_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling IdentityRiskApi->v1_systems_signup_protect_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1_systems_signup_protect_post_request** | [**V1SystemsSignupProtectPostRequest**](V1SystemsSignupProtectPostRequest.md)|  | 

### Return type

[**V1SystemsSignupProtectPost200Response**](V1SystemsSignupProtectPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**401** | Missing or invalid API key in the requiems-api-key header |  -  |
**422** | All of email, phone, and ip_address were omitted. At least one is required. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1_systems_user_verify_post**
> V1SystemsUserVerifyPost200Response v1_systems_user_verify_post(v1_systems_user_verify_post_request)

Verify User

Deep-verify an email address using domain-level signals including WHOIS age, MX records, and domain availability. Optional IP check. Use for high-value or suspicious accounts.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_systems_user_verify_post200_response import V1SystemsUserVerifyPost200Response
from openapi_client.models.v1_systems_user_verify_post_request import V1SystemsUserVerifyPostRequest
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://requiems.xyz
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://requiems.xyz"
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
    api_instance = openapi_client.IdentityRiskApi(api_client)
    v1_systems_user_verify_post_request = openapi_client.V1SystemsUserVerifyPostRequest() # V1SystemsUserVerifyPostRequest | 

    try:
        # Verify User
        api_response = api_instance.v1_systems_user_verify_post(v1_systems_user_verify_post_request)
        print("The response of IdentityRiskApi->v1_systems_user_verify_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling IdentityRiskApi->v1_systems_user_verify_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1_systems_user_verify_post_request** | [**V1SystemsUserVerifyPostRequest**](V1SystemsUserVerifyPostRequest.md)|  | 

### Return type

[**V1SystemsUserVerifyPost200Response**](V1SystemsUserVerifyPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**401** | Missing or invalid API key |  -  |
**422** | email field is missing or empty |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

