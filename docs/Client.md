# Client

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **int32** |  | [optional] 
**ServiceNumber** | Pointer to **int32** |  | [optional] 
**Developer** | Pointer to **string** |  | [optional] 
**ClientId** | Pointer to **int64** |  | [optional] 
**ClientIdAlias** | Pointer to **string** |  | [optional] 
**ClientIdAliasEnabled** | Pointer to **bool** |  | [optional] 
**ClientSecret** | Pointer to **string** |  | [optional] 
**ClientType** | Pointer to **string** |  | [optional] 
**RedirectUris** | Pointer to **[]string** |  | [optional] 
**ResponseTypes** | Pointer to **[]string** |  | [optional] 
**GrantTypes** | Pointer to **[]string** |  | [optional] 
**ApplicationType** | Pointer to **string** |  | [optional] 
**Contacts** | Pointer to **[]string** |  | [optional] 
**ClientName** | Pointer to **string** |  | [optional] 
**ClientNames** | Pointer to [**[]TaggedValue**](TaggedValue.md) |  | [optional] 
**LogoUri** | Pointer to **string** |  | [optional] 
**LogoUris** | Pointer to [**[]TaggedValue**](TaggedValue.md) |  | [optional] 
**ClientUri** | Pointer to **string** |  | [optional] 
**ClientUris** | Pointer to [**[]TaggedValue**](TaggedValue.md) |  | [optional] 
**PolicyUri** | Pointer to **string** |  | [optional] 
**PolicyUris** | Pointer to [**[]TaggedValue**](TaggedValue.md) |  | [optional] 
**TosUri** | Pointer to **string** |  | [optional] 
**TosUris** | Pointer to [**[]TaggedValue**](TaggedValue.md) |  | [optional] 
**JwksUri** | Pointer to **string** |  | [optional] 
**Jwks** | Pointer to **string** |  | [optional] 
**DerivedSectorIdentifier** | Pointer to **string** |  | [optional] 
**SectorIdentifierUri** | Pointer to **string** |  | [optional] 
**SubjectType** | Pointer to **string** |  | [optional] 
**IdTokenSignAlg** | Pointer to **string** |  | [optional] 
**IdTokenEncryptionAlg** | Pointer to **string** |  | [optional] 
**IdTokenEncryptionEnc** | Pointer to **string** |  | [optional] 
**UserInfoSignAlg** | Pointer to **string** |  | [optional] 
**UserInfoEncryptionAlg** | Pointer to **string** |  | [optional] 
**UserInfoEncryptionEnc** | Pointer to **string** |  | [optional] 
**RequestSignAlg** | Pointer to **string** |  | [optional] 
**RequestEncryptionAlg** | Pointer to **string** |  | [optional] 
**RequestEncryptionEnc** | Pointer to **string** |  | [optional] 
**TokenAuthMethod** | Pointer to **string** |  | [optional] 
**TokenAuthSignAlg** | Pointer to **string** |  | [optional] 
**DefaultMaxAge** | Pointer to **int32** |  | [optional] 
**DefaultAcrs** | Pointer to **[]string** |  | [optional] 
**AuthTimeRequired** | Pointer to **bool** |  | [optional] 
**LoginUri** | Pointer to **string** |  | [optional] 
**RequestUris** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Descriptions** | Pointer to [**[]TaggedValue**](TaggedValue.md) |  | [optional] 
**CreatedAt** | Pointer to **int64** |  | [optional] 
**ModifiedAt** | Pointer to **int64** |  | [optional] 
**Extension** | Pointer to [**ClientExtension**](ClientExtension.md) |  | [optional] 
**TlsClientAuthSubjectDn** | Pointer to **string** |  | [optional] 
**TlsClientAuthSanDns** | Pointer to **string** |  | [optional] 
**TlsClientAuthSanUri** | Pointer to **string** |  | [optional] 
**TlsClientAuthSanIp** | Pointer to **string** |  | [optional] 
**TlsClientAuthSanEmail** | Pointer to **string** |  | [optional] 
**TlsClientCertificateBoundAccessTokens** | Pointer to **bool** |  | [optional] 
**SelfSignedCertificateKeyId** | Pointer to **string** |  | [optional] 
**SoftwareId** | Pointer to **string** |  | [optional] 
**SoftwareVersion** | Pointer to **string** |  | [optional] 
**AuthorizationSignAlg** | Pointer to **string** |  | [optional] 
**AuthorizationEncryptionAlg** | Pointer to **string** |  | [optional] 
**AuthorizationEncryptionEnc** | Pointer to **string** |  | [optional] 
**BcDeliveryMode** | Pointer to **string** |  | [optional] 
**BcNotificationEndpoint** | Pointer to **string** |  | [optional] 
**BcRequestSignAlg** | Pointer to **string** |  | [optional] 
**BcUserCodeRequired** | Pointer to **bool** |  | [optional] 
**DynamicallyRegistered** | Pointer to **bool** |  | [optional] 
**RegistrationAccessTokenHash** | Pointer to **string** |  | [optional] 
**AuthorizationDetailsTypes** | Pointer to **[]string** |  | [optional] 
**ParRequired** | Pointer to **bool** |  | [optional] 
**RequestObjectRequired** | Pointer to **bool** |  | [optional] 
**Attributes** | Pointer to [**[]Pair**](Pair.md) |  | [optional] 
**CustomMetadata** | Pointer to **string** |  | [optional] 
**FrontChannelRequestObjectEncryptionRequired** | Pointer to **bool** |  | [optional] 
**RequestObjectEncryptionAlgMatchRequired** | Pointer to **bool** |  | [optional] 
**RequestObjectEncryptionEncMatchRequired** | Pointer to **bool** |  | [optional] 
**DigestAlgorithm** | Pointer to **string** |  | [optional] 
**SingleAccessTokenPerSubject** | Pointer to **bool** |  | [optional] 
**PkceRequired** | Pointer to **bool** |  | [optional] 
**PkceS256Required** | Pointer to **bool** |  | [optional] 
**RsSignedRequestKeyId** | Pointer to **string** |  | [optional] 
**RsRequestSigned** | Pointer to **bool** |  | [optional] 
**DpopRequired** | Pointer to **bool** |  | [optional] 
**EntityId** | Pointer to **string** |  | [optional] 
**TrustAnchorId** | Pointer to **string** |  | [optional] 
**TrustChain** | Pointer to **[]string** |  | [optional] 
**TrustChainExpiresAt** | Pointer to **int64** |  | [optional] 
**TrustChainUpdatedAt** | Pointer to **int64** |  | [optional] 
**OrganizationName** | Pointer to **string** |  | [optional] 
**SignedJwksUri** | Pointer to **string** |  | [optional] 
**ClientRegistrationTypes** | Pointer to **[]string** |  | [optional] 
**AutomaticallyRegistered** | Pointer to **bool** |  | [optional] 
**ExplicitlyRegistered** | Pointer to **bool** |  | [optional] 
**CredentialOfferEndpoint** | Pointer to **string** |  | [optional] 
**SectorIdentifier** | Pointer to **string** |  | [optional] 

## Methods

### NewClient

`func NewClient() *Client`

NewClient instantiates a new Client object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientWithDefaults

`func NewClientWithDefaults() *Client`

NewClientWithDefaults instantiates a new Client object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *Client) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Client) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Client) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *Client) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetServiceNumber

`func (o *Client) GetServiceNumber() int32`

GetServiceNumber returns the ServiceNumber field if non-nil, zero value otherwise.

### GetServiceNumberOk

`func (o *Client) GetServiceNumberOk() (*int32, bool)`

GetServiceNumberOk returns a tuple with the ServiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceNumber

`func (o *Client) SetServiceNumber(v int32)`

SetServiceNumber sets ServiceNumber field to given value.

### HasServiceNumber

`func (o *Client) HasServiceNumber() bool`

HasServiceNumber returns a boolean if a field has been set.

### GetDeveloper

`func (o *Client) GetDeveloper() string`

GetDeveloper returns the Developer field if non-nil, zero value otherwise.

### GetDeveloperOk

`func (o *Client) GetDeveloperOk() (*string, bool)`

GetDeveloperOk returns a tuple with the Developer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloper

`func (o *Client) SetDeveloper(v string)`

SetDeveloper sets Developer field to given value.

### HasDeveloper

`func (o *Client) HasDeveloper() bool`

HasDeveloper returns a boolean if a field has been set.

### GetClientId

`func (o *Client) GetClientId() int64`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *Client) GetClientIdOk() (*int64, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *Client) SetClientId(v int64)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *Client) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientIdAlias

`func (o *Client) GetClientIdAlias() string`

GetClientIdAlias returns the ClientIdAlias field if non-nil, zero value otherwise.

### GetClientIdAliasOk

`func (o *Client) GetClientIdAliasOk() (*string, bool)`

GetClientIdAliasOk returns a tuple with the ClientIdAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIdAlias

`func (o *Client) SetClientIdAlias(v string)`

SetClientIdAlias sets ClientIdAlias field to given value.

### HasClientIdAlias

`func (o *Client) HasClientIdAlias() bool`

HasClientIdAlias returns a boolean if a field has been set.

### GetClientIdAliasEnabled

`func (o *Client) GetClientIdAliasEnabled() bool`

GetClientIdAliasEnabled returns the ClientIdAliasEnabled field if non-nil, zero value otherwise.

### GetClientIdAliasEnabledOk

`func (o *Client) GetClientIdAliasEnabledOk() (*bool, bool)`

GetClientIdAliasEnabledOk returns a tuple with the ClientIdAliasEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIdAliasEnabled

`func (o *Client) SetClientIdAliasEnabled(v bool)`

SetClientIdAliasEnabled sets ClientIdAliasEnabled field to given value.

### HasClientIdAliasEnabled

`func (o *Client) HasClientIdAliasEnabled() bool`

HasClientIdAliasEnabled returns a boolean if a field has been set.

### GetClientSecret

`func (o *Client) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *Client) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *Client) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *Client) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetClientType

`func (o *Client) GetClientType() string`

GetClientType returns the ClientType field if non-nil, zero value otherwise.

### GetClientTypeOk

`func (o *Client) GetClientTypeOk() (*string, bool)`

GetClientTypeOk returns a tuple with the ClientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientType

`func (o *Client) SetClientType(v string)`

SetClientType sets ClientType field to given value.

### HasClientType

`func (o *Client) HasClientType() bool`

HasClientType returns a boolean if a field has been set.

### GetRedirectUris

`func (o *Client) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *Client) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *Client) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.

### HasRedirectUris

`func (o *Client) HasRedirectUris() bool`

HasRedirectUris returns a boolean if a field has been set.

### GetResponseTypes

`func (o *Client) GetResponseTypes() []string`

GetResponseTypes returns the ResponseTypes field if non-nil, zero value otherwise.

### GetResponseTypesOk

`func (o *Client) GetResponseTypesOk() (*[]string, bool)`

GetResponseTypesOk returns a tuple with the ResponseTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTypes

`func (o *Client) SetResponseTypes(v []string)`

SetResponseTypes sets ResponseTypes field to given value.

### HasResponseTypes

`func (o *Client) HasResponseTypes() bool`

HasResponseTypes returns a boolean if a field has been set.

### GetGrantTypes

`func (o *Client) GetGrantTypes() []string`

GetGrantTypes returns the GrantTypes field if non-nil, zero value otherwise.

### GetGrantTypesOk

`func (o *Client) GetGrantTypesOk() (*[]string, bool)`

GetGrantTypesOk returns a tuple with the GrantTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantTypes

`func (o *Client) SetGrantTypes(v []string)`

SetGrantTypes sets GrantTypes field to given value.

### HasGrantTypes

`func (o *Client) HasGrantTypes() bool`

HasGrantTypes returns a boolean if a field has been set.

### GetApplicationType

`func (o *Client) GetApplicationType() string`

GetApplicationType returns the ApplicationType field if non-nil, zero value otherwise.

### GetApplicationTypeOk

`func (o *Client) GetApplicationTypeOk() (*string, bool)`

GetApplicationTypeOk returns a tuple with the ApplicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationType

`func (o *Client) SetApplicationType(v string)`

SetApplicationType sets ApplicationType field to given value.

### HasApplicationType

`func (o *Client) HasApplicationType() bool`

HasApplicationType returns a boolean if a field has been set.

### GetContacts

`func (o *Client) GetContacts() []string`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *Client) GetContactsOk() (*[]string, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *Client) SetContacts(v []string)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *Client) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### GetClientName

`func (o *Client) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *Client) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *Client) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *Client) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetClientNames

`func (o *Client) GetClientNames() []TaggedValue`

GetClientNames returns the ClientNames field if non-nil, zero value otherwise.

### GetClientNamesOk

`func (o *Client) GetClientNamesOk() (*[]TaggedValue, bool)`

GetClientNamesOk returns a tuple with the ClientNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientNames

`func (o *Client) SetClientNames(v []TaggedValue)`

SetClientNames sets ClientNames field to given value.

### HasClientNames

`func (o *Client) HasClientNames() bool`

HasClientNames returns a boolean if a field has been set.

### GetLogoUri

`func (o *Client) GetLogoUri() string`

GetLogoUri returns the LogoUri field if non-nil, zero value otherwise.

### GetLogoUriOk

`func (o *Client) GetLogoUriOk() (*string, bool)`

GetLogoUriOk returns a tuple with the LogoUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUri

`func (o *Client) SetLogoUri(v string)`

SetLogoUri sets LogoUri field to given value.

### HasLogoUri

`func (o *Client) HasLogoUri() bool`

HasLogoUri returns a boolean if a field has been set.

### GetLogoUris

`func (o *Client) GetLogoUris() []TaggedValue`

GetLogoUris returns the LogoUris field if non-nil, zero value otherwise.

### GetLogoUrisOk

`func (o *Client) GetLogoUrisOk() (*[]TaggedValue, bool)`

GetLogoUrisOk returns a tuple with the LogoUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUris

`func (o *Client) SetLogoUris(v []TaggedValue)`

SetLogoUris sets LogoUris field to given value.

### HasLogoUris

`func (o *Client) HasLogoUris() bool`

HasLogoUris returns a boolean if a field has been set.

### GetClientUri

`func (o *Client) GetClientUri() string`

GetClientUri returns the ClientUri field if non-nil, zero value otherwise.

### GetClientUriOk

`func (o *Client) GetClientUriOk() (*string, bool)`

GetClientUriOk returns a tuple with the ClientUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientUri

`func (o *Client) SetClientUri(v string)`

SetClientUri sets ClientUri field to given value.

### HasClientUri

`func (o *Client) HasClientUri() bool`

HasClientUri returns a boolean if a field has been set.

### GetClientUris

`func (o *Client) GetClientUris() []TaggedValue`

GetClientUris returns the ClientUris field if non-nil, zero value otherwise.

### GetClientUrisOk

`func (o *Client) GetClientUrisOk() (*[]TaggedValue, bool)`

GetClientUrisOk returns a tuple with the ClientUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientUris

`func (o *Client) SetClientUris(v []TaggedValue)`

SetClientUris sets ClientUris field to given value.

### HasClientUris

`func (o *Client) HasClientUris() bool`

HasClientUris returns a boolean if a field has been set.

### GetPolicyUri

`func (o *Client) GetPolicyUri() string`

GetPolicyUri returns the PolicyUri field if non-nil, zero value otherwise.

### GetPolicyUriOk

`func (o *Client) GetPolicyUriOk() (*string, bool)`

GetPolicyUriOk returns a tuple with the PolicyUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyUri

`func (o *Client) SetPolicyUri(v string)`

SetPolicyUri sets PolicyUri field to given value.

### HasPolicyUri

`func (o *Client) HasPolicyUri() bool`

HasPolicyUri returns a boolean if a field has been set.

### GetPolicyUris

`func (o *Client) GetPolicyUris() []TaggedValue`

GetPolicyUris returns the PolicyUris field if non-nil, zero value otherwise.

### GetPolicyUrisOk

`func (o *Client) GetPolicyUrisOk() (*[]TaggedValue, bool)`

GetPolicyUrisOk returns a tuple with the PolicyUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyUris

`func (o *Client) SetPolicyUris(v []TaggedValue)`

SetPolicyUris sets PolicyUris field to given value.

### HasPolicyUris

`func (o *Client) HasPolicyUris() bool`

HasPolicyUris returns a boolean if a field has been set.

### GetTosUri

`func (o *Client) GetTosUri() string`

GetTosUri returns the TosUri field if non-nil, zero value otherwise.

### GetTosUriOk

`func (o *Client) GetTosUriOk() (*string, bool)`

GetTosUriOk returns a tuple with the TosUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTosUri

`func (o *Client) SetTosUri(v string)`

SetTosUri sets TosUri field to given value.

### HasTosUri

`func (o *Client) HasTosUri() bool`

HasTosUri returns a boolean if a field has been set.

### GetTosUris

`func (o *Client) GetTosUris() []TaggedValue`

GetTosUris returns the TosUris field if non-nil, zero value otherwise.

### GetTosUrisOk

`func (o *Client) GetTosUrisOk() (*[]TaggedValue, bool)`

GetTosUrisOk returns a tuple with the TosUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTosUris

`func (o *Client) SetTosUris(v []TaggedValue)`

SetTosUris sets TosUris field to given value.

### HasTosUris

`func (o *Client) HasTosUris() bool`

HasTosUris returns a boolean if a field has been set.

### GetJwksUri

`func (o *Client) GetJwksUri() string`

GetJwksUri returns the JwksUri field if non-nil, zero value otherwise.

### GetJwksUriOk

`func (o *Client) GetJwksUriOk() (*string, bool)`

GetJwksUriOk returns a tuple with the JwksUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksUri

`func (o *Client) SetJwksUri(v string)`

SetJwksUri sets JwksUri field to given value.

### HasJwksUri

`func (o *Client) HasJwksUri() bool`

HasJwksUri returns a boolean if a field has been set.

### GetJwks

`func (o *Client) GetJwks() string`

GetJwks returns the Jwks field if non-nil, zero value otherwise.

### GetJwksOk

`func (o *Client) GetJwksOk() (*string, bool)`

GetJwksOk returns a tuple with the Jwks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwks

`func (o *Client) SetJwks(v string)`

SetJwks sets Jwks field to given value.

### HasJwks

`func (o *Client) HasJwks() bool`

HasJwks returns a boolean if a field has been set.

### GetDerivedSectorIdentifier

`func (o *Client) GetDerivedSectorIdentifier() string`

GetDerivedSectorIdentifier returns the DerivedSectorIdentifier field if non-nil, zero value otherwise.

### GetDerivedSectorIdentifierOk

`func (o *Client) GetDerivedSectorIdentifierOk() (*string, bool)`

GetDerivedSectorIdentifierOk returns a tuple with the DerivedSectorIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedSectorIdentifier

`func (o *Client) SetDerivedSectorIdentifier(v string)`

SetDerivedSectorIdentifier sets DerivedSectorIdentifier field to given value.

### HasDerivedSectorIdentifier

`func (o *Client) HasDerivedSectorIdentifier() bool`

HasDerivedSectorIdentifier returns a boolean if a field has been set.

### GetSectorIdentifierUri

`func (o *Client) GetSectorIdentifierUri() string`

GetSectorIdentifierUri returns the SectorIdentifierUri field if non-nil, zero value otherwise.

### GetSectorIdentifierUriOk

`func (o *Client) GetSectorIdentifierUriOk() (*string, bool)`

GetSectorIdentifierUriOk returns a tuple with the SectorIdentifierUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectorIdentifierUri

`func (o *Client) SetSectorIdentifierUri(v string)`

SetSectorIdentifierUri sets SectorIdentifierUri field to given value.

### HasSectorIdentifierUri

`func (o *Client) HasSectorIdentifierUri() bool`

HasSectorIdentifierUri returns a boolean if a field has been set.

### GetSubjectType

`func (o *Client) GetSubjectType() string`

GetSubjectType returns the SubjectType field if non-nil, zero value otherwise.

### GetSubjectTypeOk

`func (o *Client) GetSubjectTypeOk() (*string, bool)`

GetSubjectTypeOk returns a tuple with the SubjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectType

`func (o *Client) SetSubjectType(v string)`

SetSubjectType sets SubjectType field to given value.

### HasSubjectType

`func (o *Client) HasSubjectType() bool`

HasSubjectType returns a boolean if a field has been set.

### GetIdTokenSignAlg

`func (o *Client) GetIdTokenSignAlg() string`

GetIdTokenSignAlg returns the IdTokenSignAlg field if non-nil, zero value otherwise.

### GetIdTokenSignAlgOk

`func (o *Client) GetIdTokenSignAlgOk() (*string, bool)`

GetIdTokenSignAlgOk returns a tuple with the IdTokenSignAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTokenSignAlg

`func (o *Client) SetIdTokenSignAlg(v string)`

SetIdTokenSignAlg sets IdTokenSignAlg field to given value.

### HasIdTokenSignAlg

`func (o *Client) HasIdTokenSignAlg() bool`

HasIdTokenSignAlg returns a boolean if a field has been set.

### GetIdTokenEncryptionAlg

`func (o *Client) GetIdTokenEncryptionAlg() string`

GetIdTokenEncryptionAlg returns the IdTokenEncryptionAlg field if non-nil, zero value otherwise.

### GetIdTokenEncryptionAlgOk

`func (o *Client) GetIdTokenEncryptionAlgOk() (*string, bool)`

GetIdTokenEncryptionAlgOk returns a tuple with the IdTokenEncryptionAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTokenEncryptionAlg

`func (o *Client) SetIdTokenEncryptionAlg(v string)`

SetIdTokenEncryptionAlg sets IdTokenEncryptionAlg field to given value.

### HasIdTokenEncryptionAlg

`func (o *Client) HasIdTokenEncryptionAlg() bool`

HasIdTokenEncryptionAlg returns a boolean if a field has been set.

### GetIdTokenEncryptionEnc

`func (o *Client) GetIdTokenEncryptionEnc() string`

GetIdTokenEncryptionEnc returns the IdTokenEncryptionEnc field if non-nil, zero value otherwise.

### GetIdTokenEncryptionEncOk

`func (o *Client) GetIdTokenEncryptionEncOk() (*string, bool)`

GetIdTokenEncryptionEncOk returns a tuple with the IdTokenEncryptionEnc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTokenEncryptionEnc

`func (o *Client) SetIdTokenEncryptionEnc(v string)`

SetIdTokenEncryptionEnc sets IdTokenEncryptionEnc field to given value.

### HasIdTokenEncryptionEnc

`func (o *Client) HasIdTokenEncryptionEnc() bool`

HasIdTokenEncryptionEnc returns a boolean if a field has been set.

### GetUserInfoSignAlg

`func (o *Client) GetUserInfoSignAlg() string`

GetUserInfoSignAlg returns the UserInfoSignAlg field if non-nil, zero value otherwise.

### GetUserInfoSignAlgOk

`func (o *Client) GetUserInfoSignAlgOk() (*string, bool)`

GetUserInfoSignAlgOk returns a tuple with the UserInfoSignAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInfoSignAlg

`func (o *Client) SetUserInfoSignAlg(v string)`

SetUserInfoSignAlg sets UserInfoSignAlg field to given value.

### HasUserInfoSignAlg

`func (o *Client) HasUserInfoSignAlg() bool`

HasUserInfoSignAlg returns a boolean if a field has been set.

### GetUserInfoEncryptionAlg

`func (o *Client) GetUserInfoEncryptionAlg() string`

GetUserInfoEncryptionAlg returns the UserInfoEncryptionAlg field if non-nil, zero value otherwise.

### GetUserInfoEncryptionAlgOk

`func (o *Client) GetUserInfoEncryptionAlgOk() (*string, bool)`

GetUserInfoEncryptionAlgOk returns a tuple with the UserInfoEncryptionAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInfoEncryptionAlg

`func (o *Client) SetUserInfoEncryptionAlg(v string)`

SetUserInfoEncryptionAlg sets UserInfoEncryptionAlg field to given value.

### HasUserInfoEncryptionAlg

`func (o *Client) HasUserInfoEncryptionAlg() bool`

HasUserInfoEncryptionAlg returns a boolean if a field has been set.

### GetUserInfoEncryptionEnc

`func (o *Client) GetUserInfoEncryptionEnc() string`

GetUserInfoEncryptionEnc returns the UserInfoEncryptionEnc field if non-nil, zero value otherwise.

### GetUserInfoEncryptionEncOk

`func (o *Client) GetUserInfoEncryptionEncOk() (*string, bool)`

GetUserInfoEncryptionEncOk returns a tuple with the UserInfoEncryptionEnc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInfoEncryptionEnc

`func (o *Client) SetUserInfoEncryptionEnc(v string)`

SetUserInfoEncryptionEnc sets UserInfoEncryptionEnc field to given value.

### HasUserInfoEncryptionEnc

`func (o *Client) HasUserInfoEncryptionEnc() bool`

HasUserInfoEncryptionEnc returns a boolean if a field has been set.

### GetRequestSignAlg

`func (o *Client) GetRequestSignAlg() string`

GetRequestSignAlg returns the RequestSignAlg field if non-nil, zero value otherwise.

### GetRequestSignAlgOk

`func (o *Client) GetRequestSignAlgOk() (*string, bool)`

GetRequestSignAlgOk returns a tuple with the RequestSignAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestSignAlg

`func (o *Client) SetRequestSignAlg(v string)`

SetRequestSignAlg sets RequestSignAlg field to given value.

### HasRequestSignAlg

`func (o *Client) HasRequestSignAlg() bool`

HasRequestSignAlg returns a boolean if a field has been set.

### GetRequestEncryptionAlg

`func (o *Client) GetRequestEncryptionAlg() string`

GetRequestEncryptionAlg returns the RequestEncryptionAlg field if non-nil, zero value otherwise.

### GetRequestEncryptionAlgOk

`func (o *Client) GetRequestEncryptionAlgOk() (*string, bool)`

GetRequestEncryptionAlgOk returns a tuple with the RequestEncryptionAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestEncryptionAlg

`func (o *Client) SetRequestEncryptionAlg(v string)`

SetRequestEncryptionAlg sets RequestEncryptionAlg field to given value.

### HasRequestEncryptionAlg

`func (o *Client) HasRequestEncryptionAlg() bool`

HasRequestEncryptionAlg returns a boolean if a field has been set.

### GetRequestEncryptionEnc

`func (o *Client) GetRequestEncryptionEnc() string`

GetRequestEncryptionEnc returns the RequestEncryptionEnc field if non-nil, zero value otherwise.

### GetRequestEncryptionEncOk

`func (o *Client) GetRequestEncryptionEncOk() (*string, bool)`

GetRequestEncryptionEncOk returns a tuple with the RequestEncryptionEnc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestEncryptionEnc

`func (o *Client) SetRequestEncryptionEnc(v string)`

SetRequestEncryptionEnc sets RequestEncryptionEnc field to given value.

### HasRequestEncryptionEnc

`func (o *Client) HasRequestEncryptionEnc() bool`

HasRequestEncryptionEnc returns a boolean if a field has been set.

### GetTokenAuthMethod

`func (o *Client) GetTokenAuthMethod() string`

GetTokenAuthMethod returns the TokenAuthMethod field if non-nil, zero value otherwise.

### GetTokenAuthMethodOk

`func (o *Client) GetTokenAuthMethodOk() (*string, bool)`

GetTokenAuthMethodOk returns a tuple with the TokenAuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAuthMethod

`func (o *Client) SetTokenAuthMethod(v string)`

SetTokenAuthMethod sets TokenAuthMethod field to given value.

### HasTokenAuthMethod

`func (o *Client) HasTokenAuthMethod() bool`

HasTokenAuthMethod returns a boolean if a field has been set.

### GetTokenAuthSignAlg

`func (o *Client) GetTokenAuthSignAlg() string`

GetTokenAuthSignAlg returns the TokenAuthSignAlg field if non-nil, zero value otherwise.

### GetTokenAuthSignAlgOk

`func (o *Client) GetTokenAuthSignAlgOk() (*string, bool)`

GetTokenAuthSignAlgOk returns a tuple with the TokenAuthSignAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAuthSignAlg

`func (o *Client) SetTokenAuthSignAlg(v string)`

SetTokenAuthSignAlg sets TokenAuthSignAlg field to given value.

### HasTokenAuthSignAlg

`func (o *Client) HasTokenAuthSignAlg() bool`

HasTokenAuthSignAlg returns a boolean if a field has been set.

### GetDefaultMaxAge

`func (o *Client) GetDefaultMaxAge() int32`

GetDefaultMaxAge returns the DefaultMaxAge field if non-nil, zero value otherwise.

### GetDefaultMaxAgeOk

`func (o *Client) GetDefaultMaxAgeOk() (*int32, bool)`

GetDefaultMaxAgeOk returns a tuple with the DefaultMaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMaxAge

`func (o *Client) SetDefaultMaxAge(v int32)`

SetDefaultMaxAge sets DefaultMaxAge field to given value.

### HasDefaultMaxAge

`func (o *Client) HasDefaultMaxAge() bool`

HasDefaultMaxAge returns a boolean if a field has been set.

### GetDefaultAcrs

`func (o *Client) GetDefaultAcrs() []string`

GetDefaultAcrs returns the DefaultAcrs field if non-nil, zero value otherwise.

### GetDefaultAcrsOk

`func (o *Client) GetDefaultAcrsOk() (*[]string, bool)`

GetDefaultAcrsOk returns a tuple with the DefaultAcrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAcrs

`func (o *Client) SetDefaultAcrs(v []string)`

SetDefaultAcrs sets DefaultAcrs field to given value.

### HasDefaultAcrs

`func (o *Client) HasDefaultAcrs() bool`

HasDefaultAcrs returns a boolean if a field has been set.

### GetAuthTimeRequired

`func (o *Client) GetAuthTimeRequired() bool`

GetAuthTimeRequired returns the AuthTimeRequired field if non-nil, zero value otherwise.

### GetAuthTimeRequiredOk

`func (o *Client) GetAuthTimeRequiredOk() (*bool, bool)`

GetAuthTimeRequiredOk returns a tuple with the AuthTimeRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthTimeRequired

`func (o *Client) SetAuthTimeRequired(v bool)`

SetAuthTimeRequired sets AuthTimeRequired field to given value.

### HasAuthTimeRequired

`func (o *Client) HasAuthTimeRequired() bool`

HasAuthTimeRequired returns a boolean if a field has been set.

### GetLoginUri

`func (o *Client) GetLoginUri() string`

GetLoginUri returns the LoginUri field if non-nil, zero value otherwise.

### GetLoginUriOk

`func (o *Client) GetLoginUriOk() (*string, bool)`

GetLoginUriOk returns a tuple with the LoginUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginUri

`func (o *Client) SetLoginUri(v string)`

SetLoginUri sets LoginUri field to given value.

### HasLoginUri

`func (o *Client) HasLoginUri() bool`

HasLoginUri returns a boolean if a field has been set.

### GetRequestUris

`func (o *Client) GetRequestUris() []string`

GetRequestUris returns the RequestUris field if non-nil, zero value otherwise.

### GetRequestUrisOk

`func (o *Client) GetRequestUrisOk() (*[]string, bool)`

GetRequestUrisOk returns a tuple with the RequestUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestUris

`func (o *Client) SetRequestUris(v []string)`

SetRequestUris sets RequestUris field to given value.

### HasRequestUris

`func (o *Client) HasRequestUris() bool`

HasRequestUris returns a boolean if a field has been set.

### GetDescription

`func (o *Client) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Client) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Client) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Client) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDescriptions

`func (o *Client) GetDescriptions() []TaggedValue`

GetDescriptions returns the Descriptions field if non-nil, zero value otherwise.

### GetDescriptionsOk

`func (o *Client) GetDescriptionsOk() (*[]TaggedValue, bool)`

GetDescriptionsOk returns a tuple with the Descriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptions

`func (o *Client) SetDescriptions(v []TaggedValue)`

SetDescriptions sets Descriptions field to given value.

### HasDescriptions

`func (o *Client) HasDescriptions() bool`

HasDescriptions returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Client) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Client) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Client) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Client) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *Client) GetModifiedAt() int64`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Client) GetModifiedAtOk() (*int64, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Client) SetModifiedAt(v int64)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *Client) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetExtension

`func (o *Client) GetExtension() ClientExtension`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *Client) GetExtensionOk() (*ClientExtension, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *Client) SetExtension(v ClientExtension)`

SetExtension sets Extension field to given value.

### HasExtension

`func (o *Client) HasExtension() bool`

HasExtension returns a boolean if a field has been set.

### GetTlsClientAuthSubjectDn

`func (o *Client) GetTlsClientAuthSubjectDn() string`

GetTlsClientAuthSubjectDn returns the TlsClientAuthSubjectDn field if non-nil, zero value otherwise.

### GetTlsClientAuthSubjectDnOk

`func (o *Client) GetTlsClientAuthSubjectDnOk() (*string, bool)`

GetTlsClientAuthSubjectDnOk returns a tuple with the TlsClientAuthSubjectDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsClientAuthSubjectDn

`func (o *Client) SetTlsClientAuthSubjectDn(v string)`

SetTlsClientAuthSubjectDn sets TlsClientAuthSubjectDn field to given value.

### HasTlsClientAuthSubjectDn

`func (o *Client) HasTlsClientAuthSubjectDn() bool`

HasTlsClientAuthSubjectDn returns a boolean if a field has been set.

### GetTlsClientAuthSanDns

`func (o *Client) GetTlsClientAuthSanDns() string`

GetTlsClientAuthSanDns returns the TlsClientAuthSanDns field if non-nil, zero value otherwise.

### GetTlsClientAuthSanDnsOk

`func (o *Client) GetTlsClientAuthSanDnsOk() (*string, bool)`

GetTlsClientAuthSanDnsOk returns a tuple with the TlsClientAuthSanDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsClientAuthSanDns

`func (o *Client) SetTlsClientAuthSanDns(v string)`

SetTlsClientAuthSanDns sets TlsClientAuthSanDns field to given value.

### HasTlsClientAuthSanDns

`func (o *Client) HasTlsClientAuthSanDns() bool`

HasTlsClientAuthSanDns returns a boolean if a field has been set.

### GetTlsClientAuthSanUri

`func (o *Client) GetTlsClientAuthSanUri() string`

GetTlsClientAuthSanUri returns the TlsClientAuthSanUri field if non-nil, zero value otherwise.

### GetTlsClientAuthSanUriOk

`func (o *Client) GetTlsClientAuthSanUriOk() (*string, bool)`

GetTlsClientAuthSanUriOk returns a tuple with the TlsClientAuthSanUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsClientAuthSanUri

`func (o *Client) SetTlsClientAuthSanUri(v string)`

SetTlsClientAuthSanUri sets TlsClientAuthSanUri field to given value.

### HasTlsClientAuthSanUri

`func (o *Client) HasTlsClientAuthSanUri() bool`

HasTlsClientAuthSanUri returns a boolean if a field has been set.

### GetTlsClientAuthSanIp

`func (o *Client) GetTlsClientAuthSanIp() string`

GetTlsClientAuthSanIp returns the TlsClientAuthSanIp field if non-nil, zero value otherwise.

### GetTlsClientAuthSanIpOk

`func (o *Client) GetTlsClientAuthSanIpOk() (*string, bool)`

GetTlsClientAuthSanIpOk returns a tuple with the TlsClientAuthSanIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsClientAuthSanIp

`func (o *Client) SetTlsClientAuthSanIp(v string)`

SetTlsClientAuthSanIp sets TlsClientAuthSanIp field to given value.

### HasTlsClientAuthSanIp

`func (o *Client) HasTlsClientAuthSanIp() bool`

HasTlsClientAuthSanIp returns a boolean if a field has been set.

### GetTlsClientAuthSanEmail

`func (o *Client) GetTlsClientAuthSanEmail() string`

GetTlsClientAuthSanEmail returns the TlsClientAuthSanEmail field if non-nil, zero value otherwise.

### GetTlsClientAuthSanEmailOk

`func (o *Client) GetTlsClientAuthSanEmailOk() (*string, bool)`

GetTlsClientAuthSanEmailOk returns a tuple with the TlsClientAuthSanEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsClientAuthSanEmail

`func (o *Client) SetTlsClientAuthSanEmail(v string)`

SetTlsClientAuthSanEmail sets TlsClientAuthSanEmail field to given value.

### HasTlsClientAuthSanEmail

`func (o *Client) HasTlsClientAuthSanEmail() bool`

HasTlsClientAuthSanEmail returns a boolean if a field has been set.

### GetTlsClientCertificateBoundAccessTokens

`func (o *Client) GetTlsClientCertificateBoundAccessTokens() bool`

GetTlsClientCertificateBoundAccessTokens returns the TlsClientCertificateBoundAccessTokens field if non-nil, zero value otherwise.

### GetTlsClientCertificateBoundAccessTokensOk

`func (o *Client) GetTlsClientCertificateBoundAccessTokensOk() (*bool, bool)`

GetTlsClientCertificateBoundAccessTokensOk returns a tuple with the TlsClientCertificateBoundAccessTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsClientCertificateBoundAccessTokens

`func (o *Client) SetTlsClientCertificateBoundAccessTokens(v bool)`

SetTlsClientCertificateBoundAccessTokens sets TlsClientCertificateBoundAccessTokens field to given value.

### HasTlsClientCertificateBoundAccessTokens

`func (o *Client) HasTlsClientCertificateBoundAccessTokens() bool`

HasTlsClientCertificateBoundAccessTokens returns a boolean if a field has been set.

### GetSelfSignedCertificateKeyId

`func (o *Client) GetSelfSignedCertificateKeyId() string`

GetSelfSignedCertificateKeyId returns the SelfSignedCertificateKeyId field if non-nil, zero value otherwise.

### GetSelfSignedCertificateKeyIdOk

`func (o *Client) GetSelfSignedCertificateKeyIdOk() (*string, bool)`

GetSelfSignedCertificateKeyIdOk returns a tuple with the SelfSignedCertificateKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfSignedCertificateKeyId

`func (o *Client) SetSelfSignedCertificateKeyId(v string)`

SetSelfSignedCertificateKeyId sets SelfSignedCertificateKeyId field to given value.

### HasSelfSignedCertificateKeyId

`func (o *Client) HasSelfSignedCertificateKeyId() bool`

HasSelfSignedCertificateKeyId returns a boolean if a field has been set.

### GetSoftwareId

`func (o *Client) GetSoftwareId() string`

GetSoftwareId returns the SoftwareId field if non-nil, zero value otherwise.

### GetSoftwareIdOk

`func (o *Client) GetSoftwareIdOk() (*string, bool)`

GetSoftwareIdOk returns a tuple with the SoftwareId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareId

`func (o *Client) SetSoftwareId(v string)`

SetSoftwareId sets SoftwareId field to given value.

### HasSoftwareId

`func (o *Client) HasSoftwareId() bool`

HasSoftwareId returns a boolean if a field has been set.

### GetSoftwareVersion

`func (o *Client) GetSoftwareVersion() string`

GetSoftwareVersion returns the SoftwareVersion field if non-nil, zero value otherwise.

### GetSoftwareVersionOk

`func (o *Client) GetSoftwareVersionOk() (*string, bool)`

GetSoftwareVersionOk returns a tuple with the SoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareVersion

`func (o *Client) SetSoftwareVersion(v string)`

SetSoftwareVersion sets SoftwareVersion field to given value.

### HasSoftwareVersion

`func (o *Client) HasSoftwareVersion() bool`

HasSoftwareVersion returns a boolean if a field has been set.

### GetAuthorizationSignAlg

`func (o *Client) GetAuthorizationSignAlg() string`

GetAuthorizationSignAlg returns the AuthorizationSignAlg field if non-nil, zero value otherwise.

### GetAuthorizationSignAlgOk

`func (o *Client) GetAuthorizationSignAlgOk() (*string, bool)`

GetAuthorizationSignAlgOk returns a tuple with the AuthorizationSignAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationSignAlg

`func (o *Client) SetAuthorizationSignAlg(v string)`

SetAuthorizationSignAlg sets AuthorizationSignAlg field to given value.

### HasAuthorizationSignAlg

`func (o *Client) HasAuthorizationSignAlg() bool`

HasAuthorizationSignAlg returns a boolean if a field has been set.

### GetAuthorizationEncryptionAlg

`func (o *Client) GetAuthorizationEncryptionAlg() string`

GetAuthorizationEncryptionAlg returns the AuthorizationEncryptionAlg field if non-nil, zero value otherwise.

### GetAuthorizationEncryptionAlgOk

`func (o *Client) GetAuthorizationEncryptionAlgOk() (*string, bool)`

GetAuthorizationEncryptionAlgOk returns a tuple with the AuthorizationEncryptionAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationEncryptionAlg

`func (o *Client) SetAuthorizationEncryptionAlg(v string)`

SetAuthorizationEncryptionAlg sets AuthorizationEncryptionAlg field to given value.

### HasAuthorizationEncryptionAlg

`func (o *Client) HasAuthorizationEncryptionAlg() bool`

HasAuthorizationEncryptionAlg returns a boolean if a field has been set.

### GetAuthorizationEncryptionEnc

`func (o *Client) GetAuthorizationEncryptionEnc() string`

GetAuthorizationEncryptionEnc returns the AuthorizationEncryptionEnc field if non-nil, zero value otherwise.

### GetAuthorizationEncryptionEncOk

`func (o *Client) GetAuthorizationEncryptionEncOk() (*string, bool)`

GetAuthorizationEncryptionEncOk returns a tuple with the AuthorizationEncryptionEnc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationEncryptionEnc

`func (o *Client) SetAuthorizationEncryptionEnc(v string)`

SetAuthorizationEncryptionEnc sets AuthorizationEncryptionEnc field to given value.

### HasAuthorizationEncryptionEnc

`func (o *Client) HasAuthorizationEncryptionEnc() bool`

HasAuthorizationEncryptionEnc returns a boolean if a field has been set.

### GetBcDeliveryMode

`func (o *Client) GetBcDeliveryMode() string`

GetBcDeliveryMode returns the BcDeliveryMode field if non-nil, zero value otherwise.

### GetBcDeliveryModeOk

`func (o *Client) GetBcDeliveryModeOk() (*string, bool)`

GetBcDeliveryModeOk returns a tuple with the BcDeliveryMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBcDeliveryMode

`func (o *Client) SetBcDeliveryMode(v string)`

SetBcDeliveryMode sets BcDeliveryMode field to given value.

### HasBcDeliveryMode

`func (o *Client) HasBcDeliveryMode() bool`

HasBcDeliveryMode returns a boolean if a field has been set.

### GetBcNotificationEndpoint

`func (o *Client) GetBcNotificationEndpoint() string`

GetBcNotificationEndpoint returns the BcNotificationEndpoint field if non-nil, zero value otherwise.

### GetBcNotificationEndpointOk

`func (o *Client) GetBcNotificationEndpointOk() (*string, bool)`

GetBcNotificationEndpointOk returns a tuple with the BcNotificationEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBcNotificationEndpoint

`func (o *Client) SetBcNotificationEndpoint(v string)`

SetBcNotificationEndpoint sets BcNotificationEndpoint field to given value.

### HasBcNotificationEndpoint

`func (o *Client) HasBcNotificationEndpoint() bool`

HasBcNotificationEndpoint returns a boolean if a field has been set.

### GetBcRequestSignAlg

`func (o *Client) GetBcRequestSignAlg() string`

GetBcRequestSignAlg returns the BcRequestSignAlg field if non-nil, zero value otherwise.

### GetBcRequestSignAlgOk

`func (o *Client) GetBcRequestSignAlgOk() (*string, bool)`

GetBcRequestSignAlgOk returns a tuple with the BcRequestSignAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBcRequestSignAlg

`func (o *Client) SetBcRequestSignAlg(v string)`

SetBcRequestSignAlg sets BcRequestSignAlg field to given value.

### HasBcRequestSignAlg

`func (o *Client) HasBcRequestSignAlg() bool`

HasBcRequestSignAlg returns a boolean if a field has been set.

### GetBcUserCodeRequired

`func (o *Client) GetBcUserCodeRequired() bool`

GetBcUserCodeRequired returns the BcUserCodeRequired field if non-nil, zero value otherwise.

### GetBcUserCodeRequiredOk

`func (o *Client) GetBcUserCodeRequiredOk() (*bool, bool)`

GetBcUserCodeRequiredOk returns a tuple with the BcUserCodeRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBcUserCodeRequired

`func (o *Client) SetBcUserCodeRequired(v bool)`

SetBcUserCodeRequired sets BcUserCodeRequired field to given value.

### HasBcUserCodeRequired

`func (o *Client) HasBcUserCodeRequired() bool`

HasBcUserCodeRequired returns a boolean if a field has been set.

### GetDynamicallyRegistered

`func (o *Client) GetDynamicallyRegistered() bool`

GetDynamicallyRegistered returns the DynamicallyRegistered field if non-nil, zero value otherwise.

### GetDynamicallyRegisteredOk

`func (o *Client) GetDynamicallyRegisteredOk() (*bool, bool)`

GetDynamicallyRegisteredOk returns a tuple with the DynamicallyRegistered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicallyRegistered

`func (o *Client) SetDynamicallyRegistered(v bool)`

SetDynamicallyRegistered sets DynamicallyRegistered field to given value.

### HasDynamicallyRegistered

`func (o *Client) HasDynamicallyRegistered() bool`

HasDynamicallyRegistered returns a boolean if a field has been set.

### GetRegistrationAccessTokenHash

`func (o *Client) GetRegistrationAccessTokenHash() string`

GetRegistrationAccessTokenHash returns the RegistrationAccessTokenHash field if non-nil, zero value otherwise.

### GetRegistrationAccessTokenHashOk

`func (o *Client) GetRegistrationAccessTokenHashOk() (*string, bool)`

GetRegistrationAccessTokenHashOk returns a tuple with the RegistrationAccessTokenHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationAccessTokenHash

`func (o *Client) SetRegistrationAccessTokenHash(v string)`

SetRegistrationAccessTokenHash sets RegistrationAccessTokenHash field to given value.

### HasRegistrationAccessTokenHash

`func (o *Client) HasRegistrationAccessTokenHash() bool`

HasRegistrationAccessTokenHash returns a boolean if a field has been set.

### GetAuthorizationDetailsTypes

`func (o *Client) GetAuthorizationDetailsTypes() []string`

GetAuthorizationDetailsTypes returns the AuthorizationDetailsTypes field if non-nil, zero value otherwise.

### GetAuthorizationDetailsTypesOk

`func (o *Client) GetAuthorizationDetailsTypesOk() (*[]string, bool)`

GetAuthorizationDetailsTypesOk returns a tuple with the AuthorizationDetailsTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationDetailsTypes

`func (o *Client) SetAuthorizationDetailsTypes(v []string)`

SetAuthorizationDetailsTypes sets AuthorizationDetailsTypes field to given value.

### HasAuthorizationDetailsTypes

`func (o *Client) HasAuthorizationDetailsTypes() bool`

HasAuthorizationDetailsTypes returns a boolean if a field has been set.

### GetParRequired

`func (o *Client) GetParRequired() bool`

GetParRequired returns the ParRequired field if non-nil, zero value otherwise.

### GetParRequiredOk

`func (o *Client) GetParRequiredOk() (*bool, bool)`

GetParRequiredOk returns a tuple with the ParRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParRequired

`func (o *Client) SetParRequired(v bool)`

SetParRequired sets ParRequired field to given value.

### HasParRequired

`func (o *Client) HasParRequired() bool`

HasParRequired returns a boolean if a field has been set.

### GetRequestObjectRequired

`func (o *Client) GetRequestObjectRequired() bool`

GetRequestObjectRequired returns the RequestObjectRequired field if non-nil, zero value otherwise.

### GetRequestObjectRequiredOk

`func (o *Client) GetRequestObjectRequiredOk() (*bool, bool)`

GetRequestObjectRequiredOk returns a tuple with the RequestObjectRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestObjectRequired

`func (o *Client) SetRequestObjectRequired(v bool)`

SetRequestObjectRequired sets RequestObjectRequired field to given value.

### HasRequestObjectRequired

`func (o *Client) HasRequestObjectRequired() bool`

HasRequestObjectRequired returns a boolean if a field has been set.

### GetAttributes

`func (o *Client) GetAttributes() []Pair`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Client) GetAttributesOk() (*[]Pair, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Client) SetAttributes(v []Pair)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Client) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetCustomMetadata

`func (o *Client) GetCustomMetadata() string`

GetCustomMetadata returns the CustomMetadata field if non-nil, zero value otherwise.

### GetCustomMetadataOk

`func (o *Client) GetCustomMetadataOk() (*string, bool)`

GetCustomMetadataOk returns a tuple with the CustomMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMetadata

`func (o *Client) SetCustomMetadata(v string)`

SetCustomMetadata sets CustomMetadata field to given value.

### HasCustomMetadata

`func (o *Client) HasCustomMetadata() bool`

HasCustomMetadata returns a boolean if a field has been set.

### GetFrontChannelRequestObjectEncryptionRequired

`func (o *Client) GetFrontChannelRequestObjectEncryptionRequired() bool`

GetFrontChannelRequestObjectEncryptionRequired returns the FrontChannelRequestObjectEncryptionRequired field if non-nil, zero value otherwise.

### GetFrontChannelRequestObjectEncryptionRequiredOk

`func (o *Client) GetFrontChannelRequestObjectEncryptionRequiredOk() (*bool, bool)`

GetFrontChannelRequestObjectEncryptionRequiredOk returns a tuple with the FrontChannelRequestObjectEncryptionRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontChannelRequestObjectEncryptionRequired

`func (o *Client) SetFrontChannelRequestObjectEncryptionRequired(v bool)`

SetFrontChannelRequestObjectEncryptionRequired sets FrontChannelRequestObjectEncryptionRequired field to given value.

### HasFrontChannelRequestObjectEncryptionRequired

`func (o *Client) HasFrontChannelRequestObjectEncryptionRequired() bool`

HasFrontChannelRequestObjectEncryptionRequired returns a boolean if a field has been set.

### GetRequestObjectEncryptionAlgMatchRequired

`func (o *Client) GetRequestObjectEncryptionAlgMatchRequired() bool`

GetRequestObjectEncryptionAlgMatchRequired returns the RequestObjectEncryptionAlgMatchRequired field if non-nil, zero value otherwise.

### GetRequestObjectEncryptionAlgMatchRequiredOk

`func (o *Client) GetRequestObjectEncryptionAlgMatchRequiredOk() (*bool, bool)`

GetRequestObjectEncryptionAlgMatchRequiredOk returns a tuple with the RequestObjectEncryptionAlgMatchRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestObjectEncryptionAlgMatchRequired

`func (o *Client) SetRequestObjectEncryptionAlgMatchRequired(v bool)`

SetRequestObjectEncryptionAlgMatchRequired sets RequestObjectEncryptionAlgMatchRequired field to given value.

### HasRequestObjectEncryptionAlgMatchRequired

`func (o *Client) HasRequestObjectEncryptionAlgMatchRequired() bool`

HasRequestObjectEncryptionAlgMatchRequired returns a boolean if a field has been set.

### GetRequestObjectEncryptionEncMatchRequired

`func (o *Client) GetRequestObjectEncryptionEncMatchRequired() bool`

GetRequestObjectEncryptionEncMatchRequired returns the RequestObjectEncryptionEncMatchRequired field if non-nil, zero value otherwise.

### GetRequestObjectEncryptionEncMatchRequiredOk

`func (o *Client) GetRequestObjectEncryptionEncMatchRequiredOk() (*bool, bool)`

GetRequestObjectEncryptionEncMatchRequiredOk returns a tuple with the RequestObjectEncryptionEncMatchRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestObjectEncryptionEncMatchRequired

`func (o *Client) SetRequestObjectEncryptionEncMatchRequired(v bool)`

SetRequestObjectEncryptionEncMatchRequired sets RequestObjectEncryptionEncMatchRequired field to given value.

### HasRequestObjectEncryptionEncMatchRequired

`func (o *Client) HasRequestObjectEncryptionEncMatchRequired() bool`

HasRequestObjectEncryptionEncMatchRequired returns a boolean if a field has been set.

### GetDigestAlgorithm

`func (o *Client) GetDigestAlgorithm() string`

GetDigestAlgorithm returns the DigestAlgorithm field if non-nil, zero value otherwise.

### GetDigestAlgorithmOk

`func (o *Client) GetDigestAlgorithmOk() (*string, bool)`

GetDigestAlgorithmOk returns a tuple with the DigestAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigestAlgorithm

`func (o *Client) SetDigestAlgorithm(v string)`

SetDigestAlgorithm sets DigestAlgorithm field to given value.

### HasDigestAlgorithm

`func (o *Client) HasDigestAlgorithm() bool`

HasDigestAlgorithm returns a boolean if a field has been set.

### GetSingleAccessTokenPerSubject

`func (o *Client) GetSingleAccessTokenPerSubject() bool`

GetSingleAccessTokenPerSubject returns the SingleAccessTokenPerSubject field if non-nil, zero value otherwise.

### GetSingleAccessTokenPerSubjectOk

`func (o *Client) GetSingleAccessTokenPerSubjectOk() (*bool, bool)`

GetSingleAccessTokenPerSubjectOk returns a tuple with the SingleAccessTokenPerSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleAccessTokenPerSubject

`func (o *Client) SetSingleAccessTokenPerSubject(v bool)`

SetSingleAccessTokenPerSubject sets SingleAccessTokenPerSubject field to given value.

### HasSingleAccessTokenPerSubject

`func (o *Client) HasSingleAccessTokenPerSubject() bool`

HasSingleAccessTokenPerSubject returns a boolean if a field has been set.

### GetPkceRequired

`func (o *Client) GetPkceRequired() bool`

GetPkceRequired returns the PkceRequired field if non-nil, zero value otherwise.

### GetPkceRequiredOk

`func (o *Client) GetPkceRequiredOk() (*bool, bool)`

GetPkceRequiredOk returns a tuple with the PkceRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkceRequired

`func (o *Client) SetPkceRequired(v bool)`

SetPkceRequired sets PkceRequired field to given value.

### HasPkceRequired

`func (o *Client) HasPkceRequired() bool`

HasPkceRequired returns a boolean if a field has been set.

### GetPkceS256Required

`func (o *Client) GetPkceS256Required() bool`

GetPkceS256Required returns the PkceS256Required field if non-nil, zero value otherwise.

### GetPkceS256RequiredOk

`func (o *Client) GetPkceS256RequiredOk() (*bool, bool)`

GetPkceS256RequiredOk returns a tuple with the PkceS256Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkceS256Required

`func (o *Client) SetPkceS256Required(v bool)`

SetPkceS256Required sets PkceS256Required field to given value.

### HasPkceS256Required

`func (o *Client) HasPkceS256Required() bool`

HasPkceS256Required returns a boolean if a field has been set.

### GetRsSignedRequestKeyId

`func (o *Client) GetRsSignedRequestKeyId() string`

GetRsSignedRequestKeyId returns the RsSignedRequestKeyId field if non-nil, zero value otherwise.

### GetRsSignedRequestKeyIdOk

`func (o *Client) GetRsSignedRequestKeyIdOk() (*string, bool)`

GetRsSignedRequestKeyIdOk returns a tuple with the RsSignedRequestKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRsSignedRequestKeyId

`func (o *Client) SetRsSignedRequestKeyId(v string)`

SetRsSignedRequestKeyId sets RsSignedRequestKeyId field to given value.

### HasRsSignedRequestKeyId

`func (o *Client) HasRsSignedRequestKeyId() bool`

HasRsSignedRequestKeyId returns a boolean if a field has been set.

### GetRsRequestSigned

`func (o *Client) GetRsRequestSigned() bool`

GetRsRequestSigned returns the RsRequestSigned field if non-nil, zero value otherwise.

### GetRsRequestSignedOk

`func (o *Client) GetRsRequestSignedOk() (*bool, bool)`

GetRsRequestSignedOk returns a tuple with the RsRequestSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRsRequestSigned

`func (o *Client) SetRsRequestSigned(v bool)`

SetRsRequestSigned sets RsRequestSigned field to given value.

### HasRsRequestSigned

`func (o *Client) HasRsRequestSigned() bool`

HasRsRequestSigned returns a boolean if a field has been set.

### GetDpopRequired

`func (o *Client) GetDpopRequired() bool`

GetDpopRequired returns the DpopRequired field if non-nil, zero value otherwise.

### GetDpopRequiredOk

`func (o *Client) GetDpopRequiredOk() (*bool, bool)`

GetDpopRequiredOk returns a tuple with the DpopRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpopRequired

`func (o *Client) SetDpopRequired(v bool)`

SetDpopRequired sets DpopRequired field to given value.

### HasDpopRequired

`func (o *Client) HasDpopRequired() bool`

HasDpopRequired returns a boolean if a field has been set.

### GetEntityId

`func (o *Client) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *Client) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *Client) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *Client) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetTrustAnchorId

`func (o *Client) GetTrustAnchorId() string`

GetTrustAnchorId returns the TrustAnchorId field if non-nil, zero value otherwise.

### GetTrustAnchorIdOk

`func (o *Client) GetTrustAnchorIdOk() (*string, bool)`

GetTrustAnchorIdOk returns a tuple with the TrustAnchorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustAnchorId

`func (o *Client) SetTrustAnchorId(v string)`

SetTrustAnchorId sets TrustAnchorId field to given value.

### HasTrustAnchorId

`func (o *Client) HasTrustAnchorId() bool`

HasTrustAnchorId returns a boolean if a field has been set.

### GetTrustChain

`func (o *Client) GetTrustChain() []string`

GetTrustChain returns the TrustChain field if non-nil, zero value otherwise.

### GetTrustChainOk

`func (o *Client) GetTrustChainOk() (*[]string, bool)`

GetTrustChainOk returns a tuple with the TrustChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustChain

`func (o *Client) SetTrustChain(v []string)`

SetTrustChain sets TrustChain field to given value.

### HasTrustChain

`func (o *Client) HasTrustChain() bool`

HasTrustChain returns a boolean if a field has been set.

### GetTrustChainExpiresAt

`func (o *Client) GetTrustChainExpiresAt() int64`

GetTrustChainExpiresAt returns the TrustChainExpiresAt field if non-nil, zero value otherwise.

### GetTrustChainExpiresAtOk

`func (o *Client) GetTrustChainExpiresAtOk() (*int64, bool)`

GetTrustChainExpiresAtOk returns a tuple with the TrustChainExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustChainExpiresAt

`func (o *Client) SetTrustChainExpiresAt(v int64)`

SetTrustChainExpiresAt sets TrustChainExpiresAt field to given value.

### HasTrustChainExpiresAt

`func (o *Client) HasTrustChainExpiresAt() bool`

HasTrustChainExpiresAt returns a boolean if a field has been set.

### GetTrustChainUpdatedAt

`func (o *Client) GetTrustChainUpdatedAt() int64`

GetTrustChainUpdatedAt returns the TrustChainUpdatedAt field if non-nil, zero value otherwise.

### GetTrustChainUpdatedAtOk

`func (o *Client) GetTrustChainUpdatedAtOk() (*int64, bool)`

GetTrustChainUpdatedAtOk returns a tuple with the TrustChainUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustChainUpdatedAt

`func (o *Client) SetTrustChainUpdatedAt(v int64)`

SetTrustChainUpdatedAt sets TrustChainUpdatedAt field to given value.

### HasTrustChainUpdatedAt

`func (o *Client) HasTrustChainUpdatedAt() bool`

HasTrustChainUpdatedAt returns a boolean if a field has been set.

### GetOrganizationName

`func (o *Client) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *Client) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *Client) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *Client) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### GetSignedJwksUri

`func (o *Client) GetSignedJwksUri() string`

GetSignedJwksUri returns the SignedJwksUri field if non-nil, zero value otherwise.

### GetSignedJwksUriOk

`func (o *Client) GetSignedJwksUriOk() (*string, bool)`

GetSignedJwksUriOk returns a tuple with the SignedJwksUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedJwksUri

`func (o *Client) SetSignedJwksUri(v string)`

SetSignedJwksUri sets SignedJwksUri field to given value.

### HasSignedJwksUri

`func (o *Client) HasSignedJwksUri() bool`

HasSignedJwksUri returns a boolean if a field has been set.

### GetClientRegistrationTypes

`func (o *Client) GetClientRegistrationTypes() []string`

GetClientRegistrationTypes returns the ClientRegistrationTypes field if non-nil, zero value otherwise.

### GetClientRegistrationTypesOk

`func (o *Client) GetClientRegistrationTypesOk() (*[]string, bool)`

GetClientRegistrationTypesOk returns a tuple with the ClientRegistrationTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientRegistrationTypes

`func (o *Client) SetClientRegistrationTypes(v []string)`

SetClientRegistrationTypes sets ClientRegistrationTypes field to given value.

### HasClientRegistrationTypes

`func (o *Client) HasClientRegistrationTypes() bool`

HasClientRegistrationTypes returns a boolean if a field has been set.

### GetAutomaticallyRegistered

`func (o *Client) GetAutomaticallyRegistered() bool`

GetAutomaticallyRegistered returns the AutomaticallyRegistered field if non-nil, zero value otherwise.

### GetAutomaticallyRegisteredOk

`func (o *Client) GetAutomaticallyRegisteredOk() (*bool, bool)`

GetAutomaticallyRegisteredOk returns a tuple with the AutomaticallyRegistered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticallyRegistered

`func (o *Client) SetAutomaticallyRegistered(v bool)`

SetAutomaticallyRegistered sets AutomaticallyRegistered field to given value.

### HasAutomaticallyRegistered

`func (o *Client) HasAutomaticallyRegistered() bool`

HasAutomaticallyRegistered returns a boolean if a field has been set.

### GetExplicitlyRegistered

`func (o *Client) GetExplicitlyRegistered() bool`

GetExplicitlyRegistered returns the ExplicitlyRegistered field if non-nil, zero value otherwise.

### GetExplicitlyRegisteredOk

`func (o *Client) GetExplicitlyRegisteredOk() (*bool, bool)`

GetExplicitlyRegisteredOk returns a tuple with the ExplicitlyRegistered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitlyRegistered

`func (o *Client) SetExplicitlyRegistered(v bool)`

SetExplicitlyRegistered sets ExplicitlyRegistered field to given value.

### HasExplicitlyRegistered

`func (o *Client) HasExplicitlyRegistered() bool`

HasExplicitlyRegistered returns a boolean if a field has been set.

### GetCredentialOfferEndpoint

`func (o *Client) GetCredentialOfferEndpoint() string`

GetCredentialOfferEndpoint returns the CredentialOfferEndpoint field if non-nil, zero value otherwise.

### GetCredentialOfferEndpointOk

`func (o *Client) GetCredentialOfferEndpointOk() (*string, bool)`

GetCredentialOfferEndpointOk returns a tuple with the CredentialOfferEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialOfferEndpoint

`func (o *Client) SetCredentialOfferEndpoint(v string)`

SetCredentialOfferEndpoint sets CredentialOfferEndpoint field to given value.

### HasCredentialOfferEndpoint

`func (o *Client) HasCredentialOfferEndpoint() bool`

HasCredentialOfferEndpoint returns a boolean if a field has been set.

### GetSectorIdentifier

`func (o *Client) GetSectorIdentifier() string`

GetSectorIdentifier returns the SectorIdentifier field if non-nil, zero value otherwise.

### GetSectorIdentifierOk

`func (o *Client) GetSectorIdentifierOk() (*string, bool)`

GetSectorIdentifierOk returns a tuple with the SectorIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectorIdentifier

`func (o *Client) SetSectorIdentifier(v string)`

SetSectorIdentifier sets SectorIdentifier field to given value.

### HasSectorIdentifier

`func (o *Client) HasSectorIdentifier() bool`

HasSectorIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


