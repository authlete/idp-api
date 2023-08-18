# OrganizationMembershipResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationId** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Plan** | Pointer to **string** |  | [optional] 
**Privileges** | Pointer to **[]string** |  | [optional] 
**ApiServers** | Pointer to [**[]ApiServerMembershipResponse**](ApiServerMembershipResponse.md) |  | [optional] 

## Methods

### NewOrganizationMembershipResponse

`func NewOrganizationMembershipResponse() *OrganizationMembershipResponse`

NewOrganizationMembershipResponse instantiates a new OrganizationMembershipResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationMembershipResponseWithDefaults

`func NewOrganizationMembershipResponseWithDefaults() *OrganizationMembershipResponse`

NewOrganizationMembershipResponseWithDefaults instantiates a new OrganizationMembershipResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationId

`func (o *OrganizationMembershipResponse) GetOrganizationId() int64`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *OrganizationMembershipResponse) GetOrganizationIdOk() (*int64, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *OrganizationMembershipResponse) SetOrganizationId(v int64)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *OrganizationMembershipResponse) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetName

`func (o *OrganizationMembershipResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationMembershipResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationMembershipResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrganizationMembershipResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlan

`func (o *OrganizationMembershipResponse) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *OrganizationMembershipResponse) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *OrganizationMembershipResponse) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *OrganizationMembershipResponse) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetPrivileges

`func (o *OrganizationMembershipResponse) GetPrivileges() []string`

GetPrivileges returns the Privileges field if non-nil, zero value otherwise.

### GetPrivilegesOk

`func (o *OrganizationMembershipResponse) GetPrivilegesOk() (*[]string, bool)`

GetPrivilegesOk returns a tuple with the Privileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileges

`func (o *OrganizationMembershipResponse) SetPrivileges(v []string)`

SetPrivileges sets Privileges field to given value.

### HasPrivileges

`func (o *OrganizationMembershipResponse) HasPrivileges() bool`

HasPrivileges returns a boolean if a field has been set.

### GetApiServers

`func (o *OrganizationMembershipResponse) GetApiServers() []ApiServerMembershipResponse`

GetApiServers returns the ApiServers field if non-nil, zero value otherwise.

### GetApiServersOk

`func (o *OrganizationMembershipResponse) GetApiServersOk() (*[]ApiServerMembershipResponse, bool)`

GetApiServersOk returns a tuple with the ApiServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiServers

`func (o *OrganizationMembershipResponse) SetApiServers(v []ApiServerMembershipResponse)`

SetApiServers sets ApiServers field to given value.

### HasApiServers

`func (o *OrganizationMembershipResponse) HasApiServers() bool`

HasApiServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


