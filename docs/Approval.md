# Approval

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Approved** | **bool** |  | 
**Ticket** | **string** |  | 
**Scopes** | Pointer to **[]string** |  | [optional] 

## Methods

### NewApproval

`func NewApproval(approved bool, ticket string, ) *Approval`

NewApproval instantiates a new Approval object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalWithDefaults

`func NewApprovalWithDefaults() *Approval`

NewApprovalWithDefaults instantiates a new Approval object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApproved

`func (o *Approval) GetApproved() bool`

GetApproved returns the Approved field if non-nil, zero value otherwise.

### GetApprovedOk

`func (o *Approval) GetApprovedOk() (*bool, bool)`

GetApprovedOk returns a tuple with the Approved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproved

`func (o *Approval) SetApproved(v bool)`

SetApproved sets Approved field to given value.


### GetTicket

`func (o *Approval) GetTicket() string`

GetTicket returns the Ticket field if non-nil, zero value otherwise.

### GetTicketOk

`func (o *Approval) GetTicketOk() (*string, bool)`

GetTicketOk returns a tuple with the Ticket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicket

`func (o *Approval) SetTicket(v string)`

SetTicket sets Ticket field to given value.


### GetScopes

`func (o *Approval) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *Approval) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *Approval) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *Approval) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


