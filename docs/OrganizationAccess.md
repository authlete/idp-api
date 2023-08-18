# OrganizationAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Privileges** | Pointer to **[]string** |  | [optional] 
**AllApiServerPrivileges** | Pointer to **[]string** |  | [optional] 
**AllServicePrivileges** | Pointer to **[]string** |  | [optional] 
**AllClientPrivileges** | Pointer to **[]string** |  | [optional] 
**ApiServerAccess** | Pointer to [**[]ApiServerAccess**](ApiServerAccess.md) |  | [optional] 
**OrganizationId** | Pointer to **int64** |  | [optional] 
**UserId** | Pointer to **int64** |  | [optional] 

## Methods

### NewOrganizationAccess

`func NewOrganizationAccess() *OrganizationAccess`

NewOrganizationAccess instantiates a new OrganizationAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationAccessWithDefaults

`func NewOrganizationAccessWithDefaults() *OrganizationAccess`

NewOrganizationAccessWithDefaults instantiates a new OrganizationAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationAccess) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationAccess) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationAccess) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *OrganizationAccess) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPrivileges

`func (o *OrganizationAccess) GetPrivileges() []string`

GetPrivileges returns the Privileges field if non-nil, zero value otherwise.

### GetPrivilegesOk

`func (o *OrganizationAccess) GetPrivilegesOk() (*[]string, bool)`

GetPrivilegesOk returns a tuple with the Privileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileges

`func (o *OrganizationAccess) SetPrivileges(v []string)`

SetPrivileges sets Privileges field to given value.

### HasPrivileges

`func (o *OrganizationAccess) HasPrivileges() bool`

HasPrivileges returns a boolean if a field has been set.

### GetAllApiServerPrivileges

`func (o *OrganizationAccess) GetAllApiServerPrivileges() []string`

GetAllApiServerPrivileges returns the AllApiServerPrivileges field if non-nil, zero value otherwise.

### GetAllApiServerPrivilegesOk

`func (o *OrganizationAccess) GetAllApiServerPrivilegesOk() (*[]string, bool)`

GetAllApiServerPrivilegesOk returns a tuple with the AllApiServerPrivileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllApiServerPrivileges

`func (o *OrganizationAccess) SetAllApiServerPrivileges(v []string)`

SetAllApiServerPrivileges sets AllApiServerPrivileges field to given value.

### HasAllApiServerPrivileges

`func (o *OrganizationAccess) HasAllApiServerPrivileges() bool`

HasAllApiServerPrivileges returns a boolean if a field has been set.

### GetAllServicePrivileges

`func (o *OrganizationAccess) GetAllServicePrivileges() []string`

GetAllServicePrivileges returns the AllServicePrivileges field if non-nil, zero value otherwise.

### GetAllServicePrivilegesOk

`func (o *OrganizationAccess) GetAllServicePrivilegesOk() (*[]string, bool)`

GetAllServicePrivilegesOk returns a tuple with the AllServicePrivileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllServicePrivileges

`func (o *OrganizationAccess) SetAllServicePrivileges(v []string)`

SetAllServicePrivileges sets AllServicePrivileges field to given value.

### HasAllServicePrivileges

`func (o *OrganizationAccess) HasAllServicePrivileges() bool`

HasAllServicePrivileges returns a boolean if a field has been set.

### GetAllClientPrivileges

`func (o *OrganizationAccess) GetAllClientPrivileges() []string`

GetAllClientPrivileges returns the AllClientPrivileges field if non-nil, zero value otherwise.

### GetAllClientPrivilegesOk

`func (o *OrganizationAccess) GetAllClientPrivilegesOk() (*[]string, bool)`

GetAllClientPrivilegesOk returns a tuple with the AllClientPrivileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllClientPrivileges

`func (o *OrganizationAccess) SetAllClientPrivileges(v []string)`

SetAllClientPrivileges sets AllClientPrivileges field to given value.

### HasAllClientPrivileges

`func (o *OrganizationAccess) HasAllClientPrivileges() bool`

HasAllClientPrivileges returns a boolean if a field has been set.

### GetApiServerAccess

`func (o *OrganizationAccess) GetApiServerAccess() []ApiServerAccess`

GetApiServerAccess returns the ApiServerAccess field if non-nil, zero value otherwise.

### GetApiServerAccessOk

`func (o *OrganizationAccess) GetApiServerAccessOk() (*[]ApiServerAccess, bool)`

GetApiServerAccessOk returns a tuple with the ApiServerAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiServerAccess

`func (o *OrganizationAccess) SetApiServerAccess(v []ApiServerAccess)`

SetApiServerAccess sets ApiServerAccess field to given value.

### HasApiServerAccess

`func (o *OrganizationAccess) HasApiServerAccess() bool`

HasApiServerAccess returns a boolean if a field has been set.

### GetOrganizationId

`func (o *OrganizationAccess) GetOrganizationId() int64`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *OrganizationAccess) GetOrganizationIdOk() (*int64, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *OrganizationAccess) SetOrganizationId(v int64)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *OrganizationAccess) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetUserId

`func (o *OrganizationAccess) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *OrganizationAccess) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *OrganizationAccess) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *OrganizationAccess) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


