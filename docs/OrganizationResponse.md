# OrganizationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Plan** | Pointer to **string** |  | [optional] 
**Members** | Pointer to [**[]OrganizationAccessResponse**](OrganizationAccessResponse.md) |  | [optional] 

## Methods

### NewOrganizationResponse

`func NewOrganizationResponse() *OrganizationResponse`

NewOrganizationResponse instantiates a new OrganizationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationResponseWithDefaults

`func NewOrganizationResponseWithDefaults() *OrganizationResponse`

NewOrganizationResponseWithDefaults instantiates a new OrganizationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationResponse) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *OrganizationResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *OrganizationResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrganizationResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlan

`func (o *OrganizationResponse) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *OrganizationResponse) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *OrganizationResponse) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *OrganizationResponse) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetMembers

`func (o *OrganizationResponse) GetMembers() []OrganizationAccessResponse`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *OrganizationResponse) GetMembersOk() (*[]OrganizationAccessResponse, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *OrganizationResponse) SetMembers(v []OrganizationAccessResponse)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *OrganizationResponse) HasMembers() bool`

HasMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


