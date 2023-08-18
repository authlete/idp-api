# ApiServerAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Privileges** | Pointer to **[]string** |  | [optional] 
**AllServicePrivileges** | Pointer to **[]string** |  | [optional] 
**AllClientPrivileges** | Pointer to **[]string** |  | [optional] 
**ServiceAccess** | Pointer to [**[]ServiceAccess**](ServiceAccess.md) |  | [optional] 
**OrganizationAccessId** | Pointer to **int64** |  | [optional] 
**ApiServerId** | Pointer to **int64** |  | [optional] 

## Methods

### NewApiServerAccess

`func NewApiServerAccess() *ApiServerAccess`

NewApiServerAccess instantiates a new ApiServerAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiServerAccessWithDefaults

`func NewApiServerAccessWithDefaults() *ApiServerAccess`

NewApiServerAccessWithDefaults instantiates a new ApiServerAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiServerAccess) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiServerAccess) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiServerAccess) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ApiServerAccess) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPrivileges

`func (o *ApiServerAccess) GetPrivileges() []string`

GetPrivileges returns the Privileges field if non-nil, zero value otherwise.

### GetPrivilegesOk

`func (o *ApiServerAccess) GetPrivilegesOk() (*[]string, bool)`

GetPrivilegesOk returns a tuple with the Privileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileges

`func (o *ApiServerAccess) SetPrivileges(v []string)`

SetPrivileges sets Privileges field to given value.

### HasPrivileges

`func (o *ApiServerAccess) HasPrivileges() bool`

HasPrivileges returns a boolean if a field has been set.

### GetAllServicePrivileges

`func (o *ApiServerAccess) GetAllServicePrivileges() []string`

GetAllServicePrivileges returns the AllServicePrivileges field if non-nil, zero value otherwise.

### GetAllServicePrivilegesOk

`func (o *ApiServerAccess) GetAllServicePrivilegesOk() (*[]string, bool)`

GetAllServicePrivilegesOk returns a tuple with the AllServicePrivileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllServicePrivileges

`func (o *ApiServerAccess) SetAllServicePrivileges(v []string)`

SetAllServicePrivileges sets AllServicePrivileges field to given value.

### HasAllServicePrivileges

`func (o *ApiServerAccess) HasAllServicePrivileges() bool`

HasAllServicePrivileges returns a boolean if a field has been set.

### GetAllClientPrivileges

`func (o *ApiServerAccess) GetAllClientPrivileges() []string`

GetAllClientPrivileges returns the AllClientPrivileges field if non-nil, zero value otherwise.

### GetAllClientPrivilegesOk

`func (o *ApiServerAccess) GetAllClientPrivilegesOk() (*[]string, bool)`

GetAllClientPrivilegesOk returns a tuple with the AllClientPrivileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllClientPrivileges

`func (o *ApiServerAccess) SetAllClientPrivileges(v []string)`

SetAllClientPrivileges sets AllClientPrivileges field to given value.

### HasAllClientPrivileges

`func (o *ApiServerAccess) HasAllClientPrivileges() bool`

HasAllClientPrivileges returns a boolean if a field has been set.

### GetServiceAccess

`func (o *ApiServerAccess) GetServiceAccess() []ServiceAccess`

GetServiceAccess returns the ServiceAccess field if non-nil, zero value otherwise.

### GetServiceAccessOk

`func (o *ApiServerAccess) GetServiceAccessOk() (*[]ServiceAccess, bool)`

GetServiceAccessOk returns a tuple with the ServiceAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccess

`func (o *ApiServerAccess) SetServiceAccess(v []ServiceAccess)`

SetServiceAccess sets ServiceAccess field to given value.

### HasServiceAccess

`func (o *ApiServerAccess) HasServiceAccess() bool`

HasServiceAccess returns a boolean if a field has been set.

### GetOrganizationAccessId

`func (o *ApiServerAccess) GetOrganizationAccessId() int64`

GetOrganizationAccessId returns the OrganizationAccessId field if non-nil, zero value otherwise.

### GetOrganizationAccessIdOk

`func (o *ApiServerAccess) GetOrganizationAccessIdOk() (*int64, bool)`

GetOrganizationAccessIdOk returns a tuple with the OrganizationAccessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationAccessId

`func (o *ApiServerAccess) SetOrganizationAccessId(v int64)`

SetOrganizationAccessId sets OrganizationAccessId field to given value.

### HasOrganizationAccessId

`func (o *ApiServerAccess) HasOrganizationAccessId() bool`

HasOrganizationAccessId returns a boolean if a field has been set.

### GetApiServerId

`func (o *ApiServerAccess) GetApiServerId() int64`

GetApiServerId returns the ApiServerId field if non-nil, zero value otherwise.

### GetApiServerIdOk

`func (o *ApiServerAccess) GetApiServerIdOk() (*int64, bool)`

GetApiServerIdOk returns a tuple with the ApiServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiServerId

`func (o *ApiServerAccess) SetApiServerId(v int64)`

SetApiServerId sets ApiServerId field to given value.

### HasApiServerId

`func (o *ApiServerAccess) HasApiServerId() bool`

HasApiServerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


