# V1TechnologyBaseBatchPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | **string** | List of conversion requests (1–50 items). | 
**ItemsFrom** | **int32** | Source base (2, 8, 10, or 16). | 
**ItemsTo** | **int32** | Target base (2, 8, 10, or 16). | 
**ItemsValue** | **string** | The number as a string. Accepts optional prefixes: 0x (hex), 0b (binary), 0o (octal). | 

## Methods

### NewV1TechnologyBaseBatchPostRequest

`func NewV1TechnologyBaseBatchPostRequest(items string, itemsFrom int32, itemsTo int32, itemsValue string, ) *V1TechnologyBaseBatchPostRequest`

NewV1TechnologyBaseBatchPostRequest instantiates a new V1TechnologyBaseBatchPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TechnologyBaseBatchPostRequestWithDefaults

`func NewV1TechnologyBaseBatchPostRequestWithDefaults() *V1TechnologyBaseBatchPostRequest`

NewV1TechnologyBaseBatchPostRequestWithDefaults instantiates a new V1TechnologyBaseBatchPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *V1TechnologyBaseBatchPostRequest) GetItems() string`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *V1TechnologyBaseBatchPostRequest) GetItemsOk() (*string, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *V1TechnologyBaseBatchPostRequest) SetItems(v string)`

SetItems sets Items field to given value.


### GetItemsFrom

`func (o *V1TechnologyBaseBatchPostRequest) GetItemsFrom() int32`

GetItemsFrom returns the ItemsFrom field if non-nil, zero value otherwise.

### GetItemsFromOk

`func (o *V1TechnologyBaseBatchPostRequest) GetItemsFromOk() (*int32, bool)`

GetItemsFromOk returns a tuple with the ItemsFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsFrom

`func (o *V1TechnologyBaseBatchPostRequest) SetItemsFrom(v int32)`

SetItemsFrom sets ItemsFrom field to given value.


### GetItemsTo

`func (o *V1TechnologyBaseBatchPostRequest) GetItemsTo() int32`

GetItemsTo returns the ItemsTo field if non-nil, zero value otherwise.

### GetItemsToOk

`func (o *V1TechnologyBaseBatchPostRequest) GetItemsToOk() (*int32, bool)`

GetItemsToOk returns a tuple with the ItemsTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsTo

`func (o *V1TechnologyBaseBatchPostRequest) SetItemsTo(v int32)`

SetItemsTo sets ItemsTo field to given value.


### GetItemsValue

`func (o *V1TechnologyBaseBatchPostRequest) GetItemsValue() string`

GetItemsValue returns the ItemsValue field if non-nil, zero value otherwise.

### GetItemsValueOk

`func (o *V1TechnologyBaseBatchPostRequest) GetItemsValueOk() (*string, bool)`

GetItemsValueOk returns a tuple with the ItemsValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsValue

`func (o *V1TechnologyBaseBatchPostRequest) SetItemsValue(v string)`

SetItemsValue sets ItemsValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


