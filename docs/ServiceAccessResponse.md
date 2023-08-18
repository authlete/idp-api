# ServiceAccessResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **int64** |  | [optional] 
**OrganizationId** | Pointer to **int64** |  | [optional] 
**ApiServerId** | Pointer to **int64** |  | [optional] 
**ServiceId** | Pointer to **int64** |  | [optional] 
**AllClientPrivileges** | Pointer to **[]string** |  | [optional] 
**Privileges** | Pointer to **[]string** |  | [optional] 
**ClientAccess** | Pointer to [**[]ClientAccessResponse**](ClientAccessResponse.md) |  | [optional] 

## Methods

### NewServiceAccessResponse

`func NewServiceAccessResponse() *ServiceAccessResponse`

NewServiceAccessResponse instantiates a new ServiceAccessResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAccessResponseWithDefaults

`func NewServiceAccessResponseWithDefaults() *ServiceAccessResponse`

NewServiceAccessResponseWithDefaults instantiates a new ServiceAccessResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *ServiceAccessResponse) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ServiceAccessResponse) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ServiceAccessResponse) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ServiceAccessResponse) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *ServiceAccessResponse) GetOrganizationId() int64`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ServiceAccessResponse) GetOrganizationIdOk() (*int64, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ServiceAccessResponse) SetOrganizationId(v int64)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *ServiceAccessResponse) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetApiServerId

`func (o *ServiceAccessResponse) GetApiServerId() int64`

GetApiServerId returns the ApiServerId field if non-nil, zero value otherwise.

### GetApiServerIdOk

`func (o *ServiceAccessResponse) GetApiServerIdOk() (*int64, bool)`

GetApiServerIdOk returns a tuple with the ApiServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiServerId

`func (o *ServiceAccessResponse) SetApiServerId(v int64)`

SetApiServerId sets ApiServerId field to given value.

### HasApiServerId

`func (o *ServiceAccessResponse) HasApiServerId() bool`

HasApiServerId returns a boolean if a field has been set.

### GetServiceId

`func (o *ServiceAccessResponse) GetServiceId() int64`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ServiceAccessResponse) GetServiceIdOk() (*int64, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ServiceAccessResponse) SetServiceId(v int64)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *ServiceAccessResponse) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetAllClientPrivileges

`func (o *ServiceAccessResponse) GetAllClientPrivileges() []string`

GetAllClientPrivileges returns the AllClientPrivileges field if non-nil, zero value otherwise.

### GetAllClientPrivilegesOk

`func (o *ServiceAccessResponse) GetAllClientPrivilegesOk() (*[]string, bool)`

GetAllClientPrivilegesOk returns a tuple with the AllClientPrivileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllClientPrivileges

`func (o *ServiceAccessResponse) SetAllClientPrivileges(v []string)`

SetAllClientPrivileges sets AllClientPrivileges field to given value.

### HasAllClientPrivileges

`func (o *ServiceAccessResponse) HasAllClientPrivileges() bool`

HasAllClientPrivileges returns a boolean if a field has been set.

### GetPrivileges

`func (o *ServiceAccessResponse) GetPrivileges() []string`

GetPrivileges returns the Privileges field if non-nil, zero value otherwise.

### GetPrivilegesOk

`func (o *ServiceAccessResponse) GetPrivilegesOk() (*[]string, bool)`

GetPrivilegesOk returns a tuple with the Privileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileges

`func (o *ServiceAccessResponse) SetPrivileges(v []string)`

SetPrivileges sets Privileges field to given value.

### HasPrivileges

`func (o *ServiceAccessResponse) HasPrivileges() bool`

HasPrivileges returns a boolean if a field has been set.

### GetClientAccess

`func (o *ServiceAccessResponse) GetClientAccess() []ClientAccessResponse`

GetClientAccess returns the ClientAccess field if non-nil, zero value otherwise.

### GetClientAccessOk

`func (o *ServiceAccessResponse) GetClientAccessOk() (*[]ClientAccessResponse, bool)`

GetClientAccessOk returns a tuple with the ClientAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAccess

`func (o *ServiceAccessResponse) SetClientAccess(v []ClientAccessResponse)`

SetClientAccess sets ClientAccess field to given value.

### HasClientAccess

`func (o *ServiceAccessResponse) HasClientAccess() bool`

HasClientAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


