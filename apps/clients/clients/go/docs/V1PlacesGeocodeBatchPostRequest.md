# V1PlacesGeocodeBatchPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | **string** | List of addresses to geocode (1–20 items, each non-empty). Max 20 due to Nominatim usage policy. | 

## Methods

### NewV1PlacesGeocodeBatchPostRequest

`func NewV1PlacesGeocodeBatchPostRequest(addresses string, ) *V1PlacesGeocodeBatchPostRequest`

NewV1PlacesGeocodeBatchPostRequest instantiates a new V1PlacesGeocodeBatchPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1PlacesGeocodeBatchPostRequestWithDefaults

`func NewV1PlacesGeocodeBatchPostRequestWithDefaults() *V1PlacesGeocodeBatchPostRequest`

NewV1PlacesGeocodeBatchPostRequestWithDefaults instantiates a new V1PlacesGeocodeBatchPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *V1PlacesGeocodeBatchPostRequest) GetAddresses() string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *V1PlacesGeocodeBatchPostRequest) GetAddressesOk() (*string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *V1PlacesGeocodeBatchPostRequest) SetAddresses(v string)`

SetAddresses sets Addresses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


