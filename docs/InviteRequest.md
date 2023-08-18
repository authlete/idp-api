# InviteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationId** | **int64** |  | 
**Email** | **string** |  | 
**OrganizationPrivileges** | Pointer to **[]string** |  | [optional] 
**ApiServerPrivileges** | Pointer to **[]string** |  | [optional] 
**ServicePrivileges** | Pointer to **[]string** |  | [optional] 
**ClientPrivileges** | Pointer to **[]string** |  | [optional] 

## Methods

### NewInviteRequest

`func NewInviteRequest(organizationId int64, email string, ) *InviteRequest`

NewInviteRequest instantiates a new InviteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInviteRequestWithDefaults

`func NewInviteRequestWithDefaults() *InviteRequest`

NewInviteRequestWithDefaults instantiates a new InviteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationId

`func (o *InviteRequest) GetOrganizationId() int64`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *InviteRequest) GetOrganizationIdOk() (*int64, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *InviteRequest) SetOrganizationId(v int64)`

SetOrganizationId sets OrganizationId field to given value.


### GetEmail

`func (o *InviteRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InviteRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InviteRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetOrganizationPrivileges

`func (o *InviteRequest) GetOrganizationPrivileges() []string`

GetOrganizationPrivileges returns the OrganizationPrivileges field if non-nil, zero value otherwise.

### GetOrganizationPrivilegesOk

`func (o *InviteRequest) GetOrganizationPrivilegesOk() (*[]string, bool)`

GetOrganizationPrivilegesOk returns a tuple with the OrganizationPrivileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationPrivileges

`func (o *InviteRequest) SetOrganizationPrivileges(v []string)`

SetOrganizationPrivileges sets OrganizationPrivileges field to given value.

### HasOrganizationPrivileges

`func (o *InviteRequest) HasOrganizationPrivileges() bool`

HasOrganizationPrivileges returns a boolean if a field has been set.

### GetApiServerPrivileges

`func (o *InviteRequest) GetApiServerPrivileges() []string`

GetApiServerPrivileges returns the ApiServerPrivileges field if non-nil, zero value otherwise.

### GetApiServerPrivilegesOk

`func (o *InviteRequest) GetApiServerPrivilegesOk() (*[]string, bool)`

GetApiServerPrivilegesOk returns a tuple with the ApiServerPrivileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiServerPrivileges

`func (o *InviteRequest) SetApiServerPrivileges(v []string)`

SetApiServerPrivileges sets ApiServerPrivileges field to given value.

### HasApiServerPrivileges

`func (o *InviteRequest) HasApiServerPrivileges() bool`

HasApiServerPrivileges returns a boolean if a field has been set.

### GetServicePrivileges

`func (o *InviteRequest) GetServicePrivileges() []string`

GetServicePrivileges returns the ServicePrivileges field if non-nil, zero value otherwise.

### GetServicePrivilegesOk

`func (o *InviteRequest) GetServicePrivilegesOk() (*[]string, bool)`

GetServicePrivilegesOk returns a tuple with the ServicePrivileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePrivileges

`func (o *InviteRequest) SetServicePrivileges(v []string)`

SetServicePrivileges sets ServicePrivileges field to given value.

### HasServicePrivileges

`func (o *InviteRequest) HasServicePrivileges() bool`

HasServicePrivileges returns a boolean if a field has been set.

### GetClientPrivileges

`func (o *InviteRequest) GetClientPrivileges() []string`

GetClientPrivileges returns the ClientPrivileges field if non-nil, zero value otherwise.

### GetClientPrivilegesOk

`func (o *InviteRequest) GetClientPrivilegesOk() (*[]string, bool)`

GetClientPrivilegesOk returns a tuple with the ClientPrivileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientPrivileges

`func (o *InviteRequest) SetClientPrivileges(v []string)`

SetClientPrivileges sets ClientPrivileges field to given value.

### HasClientPrivileges

`func (o *InviteRequest) HasClientPrivileges() bool`

HasClientPrivileges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


