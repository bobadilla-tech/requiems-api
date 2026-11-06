# V1TechnologyQrBase64BatchPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | **string** | List of QR generation requests (1–50 items). | 
**ItemsData** | **string** | The text or URL to encode. | 
**ItemsSize** | Pointer to **int32** | Image size in pixels (50–1000). Defaults to 256. | [optional] 
**ItemsRecovery** | Pointer to **string** | Error-correction level: low, medium, high, highest. Defaults to medium. | [optional] 

## Methods

### NewV1TechnologyQrBase64BatchPostRequest

`func NewV1TechnologyQrBase64BatchPostRequest(items string, itemsData string, ) *V1TechnologyQrBase64BatchPostRequest`

NewV1TechnologyQrBase64BatchPostRequest instantiates a new V1TechnologyQrBase64BatchPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TechnologyQrBase64BatchPostRequestWithDefaults

`func NewV1TechnologyQrBase64BatchPostRequestWithDefaults() *V1TechnologyQrBase64BatchPostRequest`

NewV1TechnologyQrBase64BatchPostRequestWithDefaults instantiates a new V1TechnologyQrBase64BatchPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *V1TechnologyQrBase64BatchPostRequest) GetItems() string`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *V1TechnologyQrBase64BatchPostRequest) GetItemsOk() (*string, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *V1TechnologyQrBase64BatchPostRequest) SetItems(v string)`

SetItems sets Items field to given value.


### GetItemsData

`func (o *V1TechnologyQrBase64BatchPostRequest) GetItemsData() string`

GetItemsData returns the ItemsData field if non-nil, zero value otherwise.

### GetItemsDataOk

`func (o *V1TechnologyQrBase64BatchPostRequest) GetItemsDataOk() (*string, bool)`

GetItemsDataOk returns a tuple with the ItemsData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsData

`func (o *V1TechnologyQrBase64BatchPostRequest) SetItemsData(v string)`

SetItemsData sets ItemsData field to given value.


### GetItemsSize

`func (o *V1TechnologyQrBase64BatchPostRequest) GetItemsSize() int32`

GetItemsSize returns the ItemsSize field if non-nil, zero value otherwise.

### GetItemsSizeOk

`func (o *V1TechnologyQrBase64BatchPostRequest) GetItemsSizeOk() (*int32, bool)`

GetItemsSizeOk returns a tuple with the ItemsSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsSize

`func (o *V1TechnologyQrBase64BatchPostRequest) SetItemsSize(v int32)`

SetItemsSize sets ItemsSize field to given value.

### HasItemsSize

`func (o *V1TechnologyQrBase64BatchPostRequest) HasItemsSize() bool`

HasItemsSize returns a boolean if a field has been set.

### GetItemsRecovery

`func (o *V1TechnologyQrBase64BatchPostRequest) GetItemsRecovery() string`

GetItemsRecovery returns the ItemsRecovery field if non-nil, zero value otherwise.

### GetItemsRecoveryOk

`func (o *V1TechnologyQrBase64BatchPostRequest) GetItemsRecoveryOk() (*string, bool)`

GetItemsRecoveryOk returns a tuple with the ItemsRecovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsRecovery

`func (o *V1TechnologyQrBase64BatchPostRequest) SetItemsRecovery(v string)`

SetItemsRecovery sets ItemsRecovery field to given value.

### HasItemsRecovery

`func (o *V1TechnologyQrBase64BatchPostRequest) HasItemsRecovery() bool`

HasItemsRecovery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


