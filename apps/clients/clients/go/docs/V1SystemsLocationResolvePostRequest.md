# V1SystemsLocationResolvePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Human-readable address string to geocode. Either address or coordinates is required. | [optional] 
**Coordinates** | Pointer to **map[string]interface{}** | Coordinates object for reverse geocoding. Either address or coordinates is required. Shape: {lat: float, lng: float} | [optional] 
**CountryCode** | Pointer to **string** | ISO 3166-1 alpha-2 country code hint. Overrides the geocoder-inferred country code when provided. | [optional] 

## Methods

### NewV1SystemsLocationResolvePostRequest

`func NewV1SystemsLocationResolvePostRequest() *V1SystemsLocationResolvePostRequest`

NewV1SystemsLocationResolvePostRequest instantiates a new V1SystemsLocationResolvePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SystemsLocationResolvePostRequestWithDefaults

`func NewV1SystemsLocationResolvePostRequestWithDefaults() *V1SystemsLocationResolvePostRequest`

NewV1SystemsLocationResolvePostRequestWithDefaults instantiates a new V1SystemsLocationResolvePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *V1SystemsLocationResolvePostRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *V1SystemsLocationResolvePostRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *V1SystemsLocationResolvePostRequest) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *V1SystemsLocationResolvePostRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCoordinates

`func (o *V1SystemsLocationResolvePostRequest) GetCoordinates() map[string]interface{}`

GetCoordinates returns the Coordinates field if non-nil, zero value otherwise.

### GetCoordinatesOk

`func (o *V1SystemsLocationResolvePostRequest) GetCoordinatesOk() (*map[string]interface{}, bool)`

GetCoordinatesOk returns a tuple with the Coordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinates

`func (o *V1SystemsLocationResolvePostRequest) SetCoordinates(v map[string]interface{})`

SetCoordinates sets Coordinates field to given value.

### HasCoordinates

`func (o *V1SystemsLocationResolvePostRequest) HasCoordinates() bool`

HasCoordinates returns a boolean if a field has been set.

### GetCountryCode

`func (o *V1SystemsLocationResolvePostRequest) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *V1SystemsLocationResolvePostRequest) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *V1SystemsLocationResolvePostRequest) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *V1SystemsLocationResolvePostRequest) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


