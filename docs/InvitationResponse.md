# InvitationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Organization** | Pointer to [**Organization**](Organization.md) |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Inviter** | Pointer to [**LogInResponse**](LogInResponse.md) |  | [optional] 
**OrganizationPrivileges** | Pointer to **[]string** |  | [optional] 
**ApiServerPrivileges** | Pointer to **[]string** |  | [optional] 
**ServicePrivileges** | Pointer to **[]string** |  | [optional] 
**ClientPrivileges** | Pointer to **[]string** |  | [optional] 

## Methods

### NewInvitationResponse

`func NewInvitationResponse() *InvitationResponse`

NewInvitationResponse instantiates a new InvitationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvitationResponseWithDefaults

`func NewInvitationResponseWithDefaults() *InvitationResponse`

NewInvitationResponseWithDefaults instantiates a new InvitationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvitationResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvitationResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvitationResponse) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InvitationResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrganization

`func (o *InvitationResponse) GetOrganization() Organization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *InvitationResponse) GetOrganizationOk() (*Organization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *InvitationResponse) SetOrganization(v Organization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *InvitationResponse) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetEmail

`func (o *InvitationResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InvitationResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InvitationResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *InvitationResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetInviter

`func (o *InvitationResponse) GetInviter() LogInResponse`

GetInviter returns the Inviter field if non-nil, zero value otherwise.

### GetInviterOk

`func (o *InvitationResponse) GetInviterOk() (*LogInResponse, bool)`

GetInviterOk returns a tuple with the Inviter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviter

`func (o *InvitationResponse) SetInviter(v LogInResponse)`

SetInviter sets Inviter field to given value.

### HasInviter

`func (o *InvitationResponse) HasInviter() bool`

HasInviter returns a boolean if a field has been set.

### GetOrganizationPrivileges

`func (o *InvitationResponse) GetOrganizationPrivileges() []string`

GetOrganizationPrivileges returns the OrganizationPrivileges field if non-nil, zero value otherwise.

### GetOrganizationPrivilegesOk

`func (o *InvitationResponse) GetOrganizationPrivilegesOk() (*[]string, bool)`

GetOrganizationPrivilegesOk returns a tuple with the OrganizationPrivileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationPrivileges

`func (o *InvitationResponse) SetOrganizationPrivileges(v []string)`

SetOrganizationPrivileges sets OrganizationPrivileges field to given value.

### HasOrganizationPrivileges

`func (o *InvitationResponse) HasOrganizationPrivileges() bool`

HasOrganizationPrivileges returns a boolean if a field has been set.

### GetApiServerPrivileges

`func (o *InvitationResponse) GetApiServerPrivileges() []string`

GetApiServerPrivileges returns the ApiServerPrivileges field if non-nil, zero value otherwise.

### GetApiServerPrivilegesOk

`func (o *InvitationResponse) GetApiServerPrivilegesOk() (*[]string, bool)`

GetApiServerPrivilegesOk returns a tuple with the ApiServerPrivileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiServerPrivileges

`func (o *InvitationResponse) SetApiServerPrivileges(v []string)`

SetApiServerPrivileges sets ApiServerPrivileges field to given value.

### HasApiServerPrivileges

`func (o *InvitationResponse) HasApiServerPrivileges() bool`

HasApiServerPrivileges returns a boolean if a field has been set.

### GetServicePrivileges

`func (o *InvitationResponse) GetServicePrivileges() []string`

GetServicePrivileges returns the ServicePrivileges field if non-nil, zero value otherwise.

### GetServicePrivilegesOk

`func (o *InvitationResponse) GetServicePrivilegesOk() (*[]string, bool)`

GetServicePrivilegesOk returns a tuple with the ServicePrivileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePrivileges

`func (o *InvitationResponse) SetServicePrivileges(v []string)`

SetServicePrivileges sets ServicePrivileges field to given value.

### HasServicePrivileges

`func (o *InvitationResponse) HasServicePrivileges() bool`

HasServicePrivileges returns a boolean if a field has been set.

### GetClientPrivileges

`func (o *InvitationResponse) GetClientPrivileges() []string`

GetClientPrivileges returns the ClientPrivileges field if non-nil, zero value otherwise.

### GetClientPrivilegesOk

`func (o *InvitationResponse) GetClientPrivilegesOk() (*[]string, bool)`

GetClientPrivilegesOk returns a tuple with the ClientPrivileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientPrivileges

`func (o *InvitationResponse) SetClientPrivileges(v []string)`

SetClientPrivileges sets ClientPrivileges field to given value.

### HasClientPrivileges

`func (o *InvitationResponse) HasClientPrivileges() bool`

HasClientPrivileges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


