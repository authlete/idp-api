# UpdateApiServerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiServerUrl** | **string** |  | 
**Description** | **string** |  | 
**OwnedBy** | Pointer to **int64** |  | [optional] 

## Methods

### NewUpdateApiServerRequest

`func NewUpdateApiServerRequest(apiServerUrl string, description string, ) *UpdateApiServerRequest`

NewUpdateApiServerRequest instantiates a new UpdateApiServerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateApiServerRequestWithDefaults

`func NewUpdateApiServerRequestWithDefaults() *UpdateApiServerRequest`

NewUpdateApiServerRequestWithDefaults instantiates a new UpdateApiServerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiServerUrl

`func (o *UpdateApiServerRequest) GetApiServerUrl() string`

GetApiServerUrl returns the ApiServerUrl field if non-nil, zero value otherwise.

### GetApiServerUrlOk

`func (o *UpdateApiServerRequest) GetApiServerUrlOk() (*string, bool)`

GetApiServerUrlOk returns a tuple with the ApiServerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiServerUrl

`func (o *UpdateApiServerRequest) SetApiServerUrl(v string)`

SetApiServerUrl sets ApiServerUrl field to given value.


### GetDescription

`func (o *UpdateApiServerRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateApiServerRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateApiServerRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetOwnedBy

`func (o *UpdateApiServerRequest) GetOwnedBy() int64`

GetOwnedBy returns the OwnedBy field if non-nil, zero value otherwise.

### GetOwnedByOk

`func (o *UpdateApiServerRequest) GetOwnedByOk() (*int64, bool)`

GetOwnedByOk returns a tuple with the OwnedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedBy

`func (o *UpdateApiServerRequest) SetOwnedBy(v int64)`

SetOwnedBy sets OwnedBy field to given value.

### HasOwnedBy

`func (o *UpdateApiServerRequest) HasOwnedBy() bool`

HasOwnedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


