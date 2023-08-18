# UpdateOrganizationPrivilegesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationId** | **int64** |  | 
**UserId** | **int64** |  | 
**Privileges** | Pointer to **[]string** |  | [optional] 
**AllApiServerPrivileges** | Pointer to **[]string** |  | [optional] 
**AllServicePrivileges** | Pointer to **[]string** |  | [optional] 
**AllClientPrivileges** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUpdateOrganizationPrivilegesRequest

`func NewUpdateOrganizationPrivilegesRequest(organizationId int64, userId int64, ) *UpdateOrganizationPrivilegesRequest`

NewUpdateOrganizationPrivilegesRequest instantiates a new UpdateOrganizationPrivilegesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOrganizationPrivilegesRequestWithDefaults

`func NewUpdateOrganizationPrivilegesRequestWithDefaults() *UpdateOrganizationPrivilegesRequest`

NewUpdateOrganizationPrivilegesRequestWithDefaults instantiates a new UpdateOrganizationPrivilegesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationId

`func (o *UpdateOrganizationPrivilegesRequest) GetOrganizationId() int64`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *UpdateOrganizationPrivilegesRequest) GetOrganizationIdOk() (*int64, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *UpdateOrganizationPrivilegesRequest) SetOrganizationId(v int64)`

SetOrganizationId sets OrganizationId field to given value.


### GetUserId

`func (o *UpdateOrganizationPrivilegesRequest) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UpdateOrganizationPrivilegesRequest) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UpdateOrganizationPrivilegesRequest) SetUserId(v int64)`

SetUserId sets UserId field to given value.


### GetPrivileges

`func (o *UpdateOrganizationPrivilegesRequest) GetPrivileges() []string`

GetPrivileges returns the Privileges field if non-nil, zero value otherwise.

### GetPrivilegesOk

`func (o *UpdateOrganizationPrivilegesRequest) GetPrivilegesOk() (*[]string, bool)`

GetPrivilegesOk returns a tuple with the Privileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileges

`func (o *UpdateOrganizationPrivilegesRequest) SetPrivileges(v []string)`

SetPrivileges sets Privileges field to given value.

### HasPrivileges

`func (o *UpdateOrganizationPrivilegesRequest) HasPrivileges() bool`

HasPrivileges returns a boolean if a field has been set.

### GetAllApiServerPrivileges

`func (o *UpdateOrganizationPrivilegesRequest) GetAllApiServerPrivileges() []string`

GetAllApiServerPrivileges returns the AllApiServerPrivileges field if non-nil, zero value otherwise.

### GetAllApiServerPrivilegesOk

`func (o *UpdateOrganizationPrivilegesRequest) GetAllApiServerPrivilegesOk() (*[]string, bool)`

GetAllApiServerPrivilegesOk returns a tuple with the AllApiServerPrivileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllApiServerPrivileges

`func (o *UpdateOrganizationPrivilegesRequest) SetAllApiServerPrivileges(v []string)`

SetAllApiServerPrivileges sets AllApiServerPrivileges field to given value.

### HasAllApiServerPrivileges

`func (o *UpdateOrganizationPrivilegesRequest) HasAllApiServerPrivileges() bool`

HasAllApiServerPrivileges returns a boolean if a field has been set.

### GetAllServicePrivileges

`func (o *UpdateOrganizationPrivilegesRequest) GetAllServicePrivileges() []string`

GetAllServicePrivileges returns the AllServicePrivileges field if non-nil, zero value otherwise.

### GetAllServicePrivilegesOk

`func (o *UpdateOrganizationPrivilegesRequest) GetAllServicePrivilegesOk() (*[]string, bool)`

GetAllServicePrivilegesOk returns a tuple with the AllServicePrivileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllServicePrivileges

`func (o *UpdateOrganizationPrivilegesRequest) SetAllServicePrivileges(v []string)`

SetAllServicePrivileges sets AllServicePrivileges field to given value.

### HasAllServicePrivileges

`func (o *UpdateOrganizationPrivilegesRequest) HasAllServicePrivileges() bool`

HasAllServicePrivileges returns a boolean if a field has been set.

### GetAllClientPrivileges

`func (o *UpdateOrganizationPrivilegesRequest) GetAllClientPrivileges() []string`

GetAllClientPrivileges returns the AllClientPrivileges field if non-nil, zero value otherwise.

### GetAllClientPrivilegesOk

`func (o *UpdateOrganizationPrivilegesRequest) GetAllClientPrivilegesOk() (*[]string, bool)`

GetAllClientPrivilegesOk returns a tuple with the AllClientPrivileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllClientPrivileges

`func (o *UpdateOrganizationPrivilegesRequest) SetAllClientPrivileges(v []string)`

SetAllClientPrivileges sets AllClientPrivileges field to given value.

### HasAllClientPrivileges

`func (o *UpdateOrganizationPrivilegesRequest) HasAllClientPrivileges() bool`

HasAllClientPrivileges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


