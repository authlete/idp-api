# Organization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Xid** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Plan** | Pointer to **string** |  | [optional] 
**ServiceListing** | Pointer to [**[]ServiceInstance**](ServiceInstance.md) |  | [optional] 
**Members** | Pointer to [**[]OrganizationAccess**](OrganizationAccess.md) |  | [optional] 

## Methods

### NewOrganization

`func NewOrganization() *Organization`

NewOrganization instantiates a new Organization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationWithDefaults

`func NewOrganizationWithDefaults() *Organization`

NewOrganizationWithDefaults instantiates a new Organization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Organization) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Organization) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Organization) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Organization) HasId() bool`

HasId returns a boolean if a field has been set.

### GetXid

`func (o *Organization) GetXid() int64`

GetXid returns the Xid field if non-nil, zero value otherwise.

### GetXidOk

`func (o *Organization) GetXidOk() (*int64, bool)`

GetXidOk returns a tuple with the Xid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXid

`func (o *Organization) SetXid(v int64)`

SetXid sets Xid field to given value.

### HasXid

`func (o *Organization) HasXid() bool`

HasXid returns a boolean if a field has been set.

### GetName

`func (o *Organization) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Organization) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Organization) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Organization) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlan

`func (o *Organization) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *Organization) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *Organization) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *Organization) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetServiceListing

`func (o *Organization) GetServiceListing() []ServiceInstance`

GetServiceListing returns the ServiceListing field if non-nil, zero value otherwise.

### GetServiceListingOk

`func (o *Organization) GetServiceListingOk() (*[]ServiceInstance, bool)`

GetServiceListingOk returns a tuple with the ServiceListing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceListing

`func (o *Organization) SetServiceListing(v []ServiceInstance)`

SetServiceListing sets ServiceListing field to given value.

### HasServiceListing

`func (o *Organization) HasServiceListing() bool`

HasServiceListing returns a boolean if a field has been set.

### GetMembers

`func (o *Organization) GetMembers() []OrganizationAccess`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *Organization) GetMembersOk() (*[]OrganizationAccess, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *Organization) SetMembers(v []OrganizationAccess)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *Organization) HasMembers() bool`

HasMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


