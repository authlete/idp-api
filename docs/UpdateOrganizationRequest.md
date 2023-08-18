# UpdateOrganizationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plan** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 

## Methods

### NewUpdateOrganizationRequest

`func NewUpdateOrganizationRequest(name string, ) *UpdateOrganizationRequest`

NewUpdateOrganizationRequest instantiates a new UpdateOrganizationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOrganizationRequestWithDefaults

`func NewUpdateOrganizationRequestWithDefaults() *UpdateOrganizationRequest`

NewUpdateOrganizationRequestWithDefaults instantiates a new UpdateOrganizationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlan

`func (o *UpdateOrganizationRequest) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *UpdateOrganizationRequest) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *UpdateOrganizationRequest) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *UpdateOrganizationRequest) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetName

`func (o *UpdateOrganizationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateOrganizationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateOrganizationRequest) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


