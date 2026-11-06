# V1TechnologyUseragentGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Browser** | Pointer to **string** | Detected browser name (e.g. Chrome, Firefox, Safari, Edge, Opera, Internet Explorer, Other) | [optional] 
**BrowserVersion** | Pointer to **string** | Detected browser version (major.minor) | [optional] 
**Os** | Pointer to **string** | Detected operating system (e.g. Windows, macOS, Linux, Android, iOS, ChromeOS, Other) | [optional] 
**OsVersion** | Pointer to **string** | Detected OS version (format varies by platform) | [optional] 
**Device** | Pointer to **string** | Device type — one of desktop, mobile, tablet, bot, or unknown | [optional] 
**IsBot** | Pointer to **bool** | True when the user agent matches a known bot or crawler pattern | [optional] 

## Methods

### NewV1TechnologyUseragentGet200ResponseData

`func NewV1TechnologyUseragentGet200ResponseData() *V1TechnologyUseragentGet200ResponseData`

NewV1TechnologyUseragentGet200ResponseData instantiates a new V1TechnologyUseragentGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TechnologyUseragentGet200ResponseDataWithDefaults

`func NewV1TechnologyUseragentGet200ResponseDataWithDefaults() *V1TechnologyUseragentGet200ResponseData`

NewV1TechnologyUseragentGet200ResponseDataWithDefaults instantiates a new V1TechnologyUseragentGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrowser

`func (o *V1TechnologyUseragentGet200ResponseData) GetBrowser() string`

GetBrowser returns the Browser field if non-nil, zero value otherwise.

### GetBrowserOk

`func (o *V1TechnologyUseragentGet200ResponseData) GetBrowserOk() (*string, bool)`

GetBrowserOk returns a tuple with the Browser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowser

`func (o *V1TechnologyUseragentGet200ResponseData) SetBrowser(v string)`

SetBrowser sets Browser field to given value.

### HasBrowser

`func (o *V1TechnologyUseragentGet200ResponseData) HasBrowser() bool`

HasBrowser returns a boolean if a field has been set.

### GetBrowserVersion

`func (o *V1TechnologyUseragentGet200ResponseData) GetBrowserVersion() string`

GetBrowserVersion returns the BrowserVersion field if non-nil, zero value otherwise.

### GetBrowserVersionOk

`func (o *V1TechnologyUseragentGet200ResponseData) GetBrowserVersionOk() (*string, bool)`

GetBrowserVersionOk returns a tuple with the BrowserVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserVersion

`func (o *V1TechnologyUseragentGet200ResponseData) SetBrowserVersion(v string)`

SetBrowserVersion sets BrowserVersion field to given value.

### HasBrowserVersion

`func (o *V1TechnologyUseragentGet200ResponseData) HasBrowserVersion() bool`

HasBrowserVersion returns a boolean if a field has been set.

### GetOs

`func (o *V1TechnologyUseragentGet200ResponseData) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *V1TechnologyUseragentGet200ResponseData) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *V1TechnologyUseragentGet200ResponseData) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *V1TechnologyUseragentGet200ResponseData) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetOsVersion

`func (o *V1TechnologyUseragentGet200ResponseData) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *V1TechnologyUseragentGet200ResponseData) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *V1TechnologyUseragentGet200ResponseData) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *V1TechnologyUseragentGet200ResponseData) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetDevice

`func (o *V1TechnologyUseragentGet200ResponseData) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *V1TechnologyUseragentGet200ResponseData) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *V1TechnologyUseragentGet200ResponseData) SetDevice(v string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *V1TechnologyUseragentGet200ResponseData) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetIsBot

`func (o *V1TechnologyUseragentGet200ResponseData) GetIsBot() bool`

GetIsBot returns the IsBot field if non-nil, zero value otherwise.

### GetIsBotOk

`func (o *V1TechnologyUseragentGet200ResponseData) GetIsBotOk() (*bool, bool)`

GetIsBotOk returns a tuple with the IsBot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBot

`func (o *V1TechnologyUseragentGet200ResponseData) SetIsBot(v bool)`

SetIsBot sets IsBot field to given value.

### HasIsBot

`func (o *V1TechnologyUseragentGet200ResponseData) HasIsBot() bool`

HasIsBot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


