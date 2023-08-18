# AuthzDetailsElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Locations** | Pointer to **[]string** |  | [optional] 
**Actions** | Pointer to **[]string** |  | [optional] 
**DataTypes** | Pointer to **[]string** |  | [optional] 
**Identifier** | Pointer to **string** |  | [optional] 
**Privileges** | Pointer to **[]string** |  | [optional] 
**OtherFields** | Pointer to **string** |  | [optional] 
**OtherFieldsFromMap** | Pointer to [**AuthzDetailsElement**](AuthzDetailsElement.md) |  | [optional] 
**OtherFieldsAsMap** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewAuthzDetailsElement

`func NewAuthzDetailsElement() *AuthzDetailsElement`

NewAuthzDetailsElement instantiates a new AuthzDetailsElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthzDetailsElementWithDefaults

`func NewAuthzDetailsElementWithDefaults() *AuthzDetailsElement`

NewAuthzDetailsElementWithDefaults instantiates a new AuthzDetailsElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AuthzDetailsElement) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthzDetailsElement) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthzDetailsElement) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AuthzDetailsElement) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLocations

`func (o *AuthzDetailsElement) GetLocations() []string`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *AuthzDetailsElement) GetLocationsOk() (*[]string, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *AuthzDetailsElement) SetLocations(v []string)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *AuthzDetailsElement) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetActions

`func (o *AuthzDetailsElement) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *AuthzDetailsElement) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *AuthzDetailsElement) SetActions(v []string)`

SetActions sets Actions field to given value.

### HasActions

`func (o *AuthzDetailsElement) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetDataTypes

`func (o *AuthzDetailsElement) GetDataTypes() []string`

GetDataTypes returns the DataTypes field if non-nil, zero value otherwise.

### GetDataTypesOk

`func (o *AuthzDetailsElement) GetDataTypesOk() (*[]string, bool)`

GetDataTypesOk returns a tuple with the DataTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTypes

`func (o *AuthzDetailsElement) SetDataTypes(v []string)`

SetDataTypes sets DataTypes field to given value.

### HasDataTypes

`func (o *AuthzDetailsElement) HasDataTypes() bool`

HasDataTypes returns a boolean if a field has been set.

### GetIdentifier

`func (o *AuthzDetailsElement) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *AuthzDetailsElement) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *AuthzDetailsElement) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *AuthzDetailsElement) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetPrivileges

`func (o *AuthzDetailsElement) GetPrivileges() []string`

GetPrivileges returns the Privileges field if non-nil, zero value otherwise.

### GetPrivilegesOk

`func (o *AuthzDetailsElement) GetPrivilegesOk() (*[]string, bool)`

GetPrivilegesOk returns a tuple with the Privileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileges

`func (o *AuthzDetailsElement) SetPrivileges(v []string)`

SetPrivileges sets Privileges field to given value.

### HasPrivileges

`func (o *AuthzDetailsElement) HasPrivileges() bool`

HasPrivileges returns a boolean if a field has been set.

### GetOtherFields

`func (o *AuthzDetailsElement) GetOtherFields() string`

GetOtherFields returns the OtherFields field if non-nil, zero value otherwise.

### GetOtherFieldsOk

`func (o *AuthzDetailsElement) GetOtherFieldsOk() (*string, bool)`

GetOtherFieldsOk returns a tuple with the OtherFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherFields

`func (o *AuthzDetailsElement) SetOtherFields(v string)`

SetOtherFields sets OtherFields field to given value.

### HasOtherFields

`func (o *AuthzDetailsElement) HasOtherFields() bool`

HasOtherFields returns a boolean if a field has been set.

### GetOtherFieldsFromMap

`func (o *AuthzDetailsElement) GetOtherFieldsFromMap() AuthzDetailsElement`

GetOtherFieldsFromMap returns the OtherFieldsFromMap field if non-nil, zero value otherwise.

### GetOtherFieldsFromMapOk

`func (o *AuthzDetailsElement) GetOtherFieldsFromMapOk() (*AuthzDetailsElement, bool)`

GetOtherFieldsFromMapOk returns a tuple with the OtherFieldsFromMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherFieldsFromMap

`func (o *AuthzDetailsElement) SetOtherFieldsFromMap(v AuthzDetailsElement)`

SetOtherFieldsFromMap sets OtherFieldsFromMap field to given value.

### HasOtherFieldsFromMap

`func (o *AuthzDetailsElement) HasOtherFieldsFromMap() bool`

HasOtherFieldsFromMap returns a boolean if a field has been set.

### GetOtherFieldsAsMap

`func (o *AuthzDetailsElement) GetOtherFieldsAsMap() map[string]map[string]interface{}`

GetOtherFieldsAsMap returns the OtherFieldsAsMap field if non-nil, zero value otherwise.

### GetOtherFieldsAsMapOk

`func (o *AuthzDetailsElement) GetOtherFieldsAsMapOk() (*map[string]map[string]interface{}, bool)`

GetOtherFieldsAsMapOk returns a tuple with the OtherFieldsAsMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherFieldsAsMap

`func (o *AuthzDetailsElement) SetOtherFieldsAsMap(v map[string]map[string]interface{})`

SetOtherFieldsAsMap sets OtherFieldsAsMap field to given value.

### HasOtherFieldsAsMap

`func (o *AuthzDetailsElement) HasOtherFieldsAsMap() bool`

HasOtherFieldsAsMap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


