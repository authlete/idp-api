# ApiServerAccessResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **int64** |  | [optional] 
**OrganizationId** | Pointer to **int64** |  | [optional] 
**ApiServerId** | Pointer to **int64** |  | [optional] 
**Privileges** | Pointer to **[]string** |  | [optional] 
**AllServicePrivileges** | Pointer to **[]string** |  | [optional] 
**AllClientPrivileges** | Pointer to **[]string** |  | [optional] 
**ApiServer** | Pointer to [**AuthleteApiServerResponse**](AuthleteApiServerResponse.md) |  | [optional] 
**ServiceAccess** | Pointer to [**[]ServiceAccessResponse**](ServiceAccessResponse.md) |  | [optional] 

## Methods

### NewApiServerAccessResponse

`func NewApiServerAccessResponse() *ApiServerAccessResponse`

NewApiServerAccessResponse instantiates a new ApiServerAccessResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiServerAccessResponseWithDefaults

`func NewApiServerAccessResponseWithDefaults() *ApiServerAccessResponse`

NewApiServerAccessResponseWithDefaults instantiates a new ApiServerAccessResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *ApiServerAccessResponse) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ApiServerAccessResponse) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ApiServerAccessResponse) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ApiServerAccessResponse) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *ApiServerAccessResponse) GetOrganizationId() int64`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ApiServerAccessResponse) GetOrganizationIdOk() (*int64, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ApiServerAccessResponse) SetOrganizationId(v int64)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *ApiServerAccessResponse) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetApiServerId

`func (o *ApiServerAccessResponse) GetApiServerId() int64`

GetApiServerId returns the ApiServerId field if non-nil, zero value otherwise.

### GetApiServerIdOk

`func (o *ApiServerAccessResponse) GetApiServerIdOk() (*int64, bool)`

GetApiServerIdOk returns a tuple with the ApiServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiServerId

`func (o *ApiServerAccessResponse) SetApiServerId(v int64)`

SetApiServerId sets ApiServerId field to given value.

### HasApiServerId

`func (o *ApiServerAccessResponse) HasApiServerId() bool`

HasApiServerId returns a boolean if a field has been set.

### GetPrivileges

`func (o *ApiServerAccessResponse) GetPrivileges() []string`

GetPrivileges returns the Privileges field if non-nil, zero value otherwise.

### GetPrivilegesOk

`func (o *ApiServerAccessResponse) GetPrivilegesOk() (*[]string, bool)`

GetPrivilegesOk returns a tuple with the Privileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileges

`func (o *ApiServerAccessResponse) SetPrivileges(v []string)`

SetPrivileges sets Privileges field to given value.

### HasPrivileges

`func (o *ApiServerAccessResponse) HasPrivileges() bool`

HasPrivileges returns a boolean if a field has been set.

### GetAllServicePrivileges

`func (o *ApiServerAccessResponse) GetAllServicePrivileges() []string`

GetAllServicePrivileges returns the AllServicePrivileges field if non-nil, zero value otherwise.

### GetAllServicePrivilegesOk

`func (o *ApiServerAccessResponse) GetAllServicePrivilegesOk() (*[]string, bool)`

GetAllServicePrivilegesOk returns a tuple with the AllServicePrivileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllServicePrivileges

`func (o *ApiServerAccessResponse) SetAllServicePrivileges(v []string)`

SetAllServicePrivileges sets AllServicePrivileges field to given value.

### HasAllServicePrivileges

`func (o *ApiServerAccessResponse) HasAllServicePrivileges() bool`

HasAllServicePrivileges returns a boolean if a field has been set.

### GetAllClientPrivileges

`func (o *ApiServerAccessResponse) GetAllClientPrivileges() []string`

GetAllClientPrivileges returns the AllClientPrivileges field if non-nil, zero value otherwise.

### GetAllClientPrivilegesOk

`func (o *ApiServerAccessResponse) GetAllClientPrivilegesOk() (*[]string, bool)`

GetAllClientPrivilegesOk returns a tuple with the AllClientPrivileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllClientPrivileges

`func (o *ApiServerAccessResponse) SetAllClientPrivileges(v []string)`

SetAllClientPrivileges sets AllClientPrivileges field to given value.

### HasAllClientPrivileges

`func (o *ApiServerAccessResponse) HasAllClientPrivileges() bool`

HasAllClientPrivileges returns a boolean if a field has been set.

### GetApiServer

`func (o *ApiServerAccessResponse) GetApiServer() AuthleteApiServerResponse`

GetApiServer returns the ApiServer field if non-nil, zero value otherwise.

### GetApiServerOk

`func (o *ApiServerAccessResponse) GetApiServerOk() (*AuthleteApiServerResponse, bool)`

GetApiServerOk returns a tuple with the ApiServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiServer

`func (o *ApiServerAccessResponse) SetApiServer(v AuthleteApiServerResponse)`

SetApiServer sets ApiServer field to given value.

### HasApiServer

`func (o *ApiServerAccessResponse) HasApiServer() bool`

HasApiServer returns a boolean if a field has been set.

### GetServiceAccess

`func (o *ApiServerAccessResponse) GetServiceAccess() []ServiceAccessResponse`

GetServiceAccess returns the ServiceAccess field if non-nil, zero value otherwise.

### GetServiceAccessOk

`func (o *ApiServerAccessResponse) GetServiceAccessOk() (*[]ServiceAccessResponse, bool)`

GetServiceAccessOk returns a tuple with the ServiceAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccess

`func (o *ApiServerAccessResponse) SetServiceAccess(v []ServiceAccessResponse)`

SetServiceAccess sets ServiceAccess field to given value.

### HasServiceAccess

`func (o *ApiServerAccessResponse) HasServiceAccess() bool`

HasServiceAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


