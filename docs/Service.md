# Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **int32** |  | [optional] 
**ServiceOwnerNumber** | Pointer to **int32** |  | [optional] 
**ServiceName** | Pointer to **string** |  | [optional] 
**ApiKey** | Pointer to **int64** |  | [optional] 
**ApiSecret** | Pointer to **string** |  | [optional] 
**Issuer** | Pointer to **string** |  | [optional] 
**AuthorizationEndpoint** | Pointer to **string** |  | [optional] 
**TokenEndpoint** | Pointer to **string** |  | [optional] 
**RevocationEndpoint** | Pointer to **string** |  | [optional] 
**SupportedRevocationAuthMethods** | Pointer to **[]string** |  | [optional] 
**UserInfoEndpoint** | Pointer to **string** |  | [optional] 
**JwksUri** | Pointer to **string** |  | [optional] 
**Jwks** | Pointer to **string** |  | [optional] 
**RegistrationEndpoint** | Pointer to **string** |  | [optional] 
**RegistrationManagementEndpoint** | Pointer to **string** |  | [optional] 
**SupportedScopes** | Pointer to [**[]Scope**](Scope.md) |  | [optional] 
**SupportedResponseTypes** | Pointer to **[]string** |  | [optional] 
**SupportedGrantTypes** | Pointer to **[]string** |  | [optional] 
**SupportedAcrs** | Pointer to **[]string** |  | [optional] 
**SupportedTokenAuthMethods** | Pointer to **[]string** |  | [optional] 
**SupportedDisplays** | Pointer to **[]string** |  | [optional] 
**SupportedClaimTypes** | Pointer to **[]string** |  | [optional] 
**SupportedClaims** | Pointer to **[]string** |  | [optional] 
**ServiceDocumentation** | Pointer to **string** |  | [optional] 
**SupportedClaimLocales** | Pointer to **[]string** |  | [optional] 
**SupportedUiLocales** | Pointer to **[]string** |  | [optional] 
**PolicyUri** | Pointer to **string** |  | [optional] 
**TosUri** | Pointer to **string** |  | [optional] 
**AuthenticationCallbackEndpoint** | Pointer to **string** |  | [optional] 
**AuthenticationCallbackApiKey** | Pointer to **string** |  | [optional] 
**AuthenticationCallbackApiSecret** | Pointer to **string** |  | [optional] 
**SupportedSnses** | Pointer to **[]string** |  | [optional] 
**SnsCredentials** | Pointer to [**[]SnsCredentials**](SnsCredentials.md) |  | [optional] 
**CreatedAt** | Pointer to **int64** |  | [optional] 
**ModifiedAt** | Pointer to **int64** |  | [optional] 
**DeveloperAuthenticationCallbackEndpoint** | Pointer to **string** |  | [optional] 
**DeveloperAuthenticationCallbackApiKey** | Pointer to **string** |  | [optional] 
**DeveloperAuthenticationCallbackApiSecret** | Pointer to **string** |  | [optional] 
**SupportedDeveloperSnses** | Pointer to **[]string** |  | [optional] 
**DeveloperSnsCredentials** | Pointer to [**[]SnsCredentials**](SnsCredentials.md) |  | [optional] 
**ClientsPerDeveloper** | Pointer to **int32** |  | [optional] 
**DirectAuthorizationEndpointEnabled** | Pointer to **bool** |  | [optional] 
**DirectTokenEndpointEnabled** | Pointer to **bool** |  | [optional] 
**DirectRevocationEndpointEnabled** | Pointer to **bool** |  | [optional] 
**DirectUserInfoEndpointEnabled** | Pointer to **bool** |  | [optional] 
**DirectJwksEndpointEnabled** | Pointer to **bool** |  | [optional] 
**DirectIntrospectionEndpointEnabled** | Pointer to **bool** |  | [optional] 
**SingleAccessTokenPerSubject** | Pointer to **bool** |  | [optional] 
**PkceRequired** | Pointer to **bool** |  | [optional] 
**PkceS256Required** | Pointer to **bool** |  | [optional] 
**RefreshTokenKept** | Pointer to **bool** |  | [optional] 
**RefreshTokenDurationKept** | Pointer to **bool** |  | [optional] 
**RefreshTokenDurationReset** | Pointer to **bool** |  | [optional] 
**ErrorDescriptionOmitted** | Pointer to **bool** |  | [optional] 
**ErrorUriOmitted** | Pointer to **bool** |  | [optional] 
**ClientIdAliasEnabled** | Pointer to **bool** |  | [optional] 
**SupportedServiceProfiles** | Pointer to **[]string** |  | [optional] 
**TlsClientCertificateBoundAccessTokens** | Pointer to **bool** |  | [optional] 
**IntrospectionEndpoint** | Pointer to **string** |  | [optional] 
**SupportedIntrospectionAuthMethods** | Pointer to **[]string** |  | [optional] 
**MutualTlsValidatePkiCertChain** | Pointer to **bool** |  | [optional] 
**TrustedRootCertificates** | Pointer to **[]string** |  | [optional] 
**DynamicRegistrationSupported** | Pointer to **bool** |  | [optional] 
**EndSessionEndpoint** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**AccessTokenType** | Pointer to **string** |  | [optional] 
**AccessTokenSignAlg** | Pointer to **string** |  | [optional] 
**AccessTokenDuration** | Pointer to **int64** |  | [optional] 
**RefreshTokenDuration** | Pointer to **int64** |  | [optional] 
**IdTokenDuration** | Pointer to **int64** |  | [optional] 
**AuthorizationResponseDuration** | Pointer to **int64** |  | [optional] 
**PushedAuthReqDuration** | Pointer to **int64** |  | [optional] 
**Metadata** | Pointer to [**[]Pair**](Pair.md) |  | [optional] 
**AccessTokenSignatureKeyId** | Pointer to **string** |  | [optional] 
**AuthorizationSignatureKeyId** | Pointer to **string** |  | [optional] 
**IdTokenSignatureKeyId** | Pointer to **string** |  | [optional] 
**UserInfoSignatureKeyId** | Pointer to **string** |  | [optional] 
**SupportedBackchannelTokenDeliveryModes** | Pointer to **[]string** |  | [optional] 
**BackchannelAuthenticationEndpoint** | Pointer to **string** |  | [optional] 
**BackchannelUserCodeParameterSupported** | Pointer to **bool** |  | [optional] 
**BackchannelAuthReqIdDuration** | Pointer to **int32** |  | [optional] 
**BackchannelPollingInterval** | Pointer to **int32** |  | [optional] 
**BackchannelBindingMessageRequiredInFapi** | Pointer to **bool** |  | [optional] 
**AllowableClockSkew** | Pointer to **int32** |  | [optional] 
**DeviceAuthorizationEndpoint** | Pointer to **string** |  | [optional] 
**DeviceVerificationUri** | Pointer to **string** |  | [optional] 
**DeviceVerificationUriComplete** | Pointer to **string** |  | [optional] 
**DeviceFlowCodeDuration** | Pointer to **int32** |  | [optional] 
**DeviceFlowPollingInterval** | Pointer to **int32** |  | [optional] 
**UserCodeCharset** | Pointer to **string** |  | [optional] 
**UserCodeLength** | Pointer to **int32** |  | [optional] 
**PushedAuthReqEndpoint** | Pointer to **string** |  | [optional] 
**MtlsEndpointAliases** | Pointer to [**[]NamedUri**](NamedUri.md) |  | [optional] 
**SupportedAuthorizationDetailsTypes** | Pointer to **[]string** |  | [optional] 
**SupportedTrustFrameworks** | Pointer to **[]string** |  | [optional] 
**SupportedEvidence** | Pointer to **[]string** |  | [optional] 
**SupportedIdentityDocuments** | Pointer to **[]string** |  | [optional] 
**SupportedDocuments** | Pointer to **[]string** |  | [optional] 
**SupportedVerificationMethods** | Pointer to **[]string** |  | [optional] 
**SupportedDocumentsMethods** | Pointer to **[]string** |  | [optional] 
**SupportedDocumentsValidationMethods** | Pointer to **[]string** |  | [optional] 
**SupportedDocumentsVerificationMethods** | Pointer to **[]string** |  | [optional] 
**SupportedDocumentsCheckMethods** | Pointer to **[]string** |  | [optional] 
**SupportedElectronicRecords** | Pointer to **[]string** |  | [optional] 
**SupportedVerifiedClaims** | Pointer to **[]string** |  | [optional] 
**SupportedAttachments** | Pointer to **[]string** |  | [optional] 
**SupportedDigestAlgorithms** | Pointer to **[]string** |  | [optional] 
**MissingClientIdAllowed** | Pointer to **bool** |  | [optional] 
**ParRequired** | Pointer to **bool** |  | [optional] 
**RequestObjectRequired** | Pointer to **bool** |  | [optional] 
**TraditionalRequestObjectProcessingApplied** | Pointer to **bool** |  | [optional] 
**ClaimShortcutRestrictive** | Pointer to **bool** |  | [optional] 
**ScopeRequired** | Pointer to **bool** |  | [optional] 
**NbfOptional** | Pointer to **bool** |  | [optional] 
**IssSuppressed** | Pointer to **bool** |  | [optional] 
**Attributes** | Pointer to [**[]Pair**](Pair.md) |  | [optional] 
**SupportedCustomClientMetadata** | Pointer to **[]string** |  | [optional] 
**TokenExpirationLinked** | Pointer to **bool** |  | [optional] 
**FrontChannelRequestObjectEncryptionRequired** | Pointer to **bool** |  | [optional] 
**RequestObjectEncryptionAlgMatchRequired** | Pointer to **bool** |  | [optional] 
**RequestObjectEncryptionEncMatchRequired** | Pointer to **bool** |  | [optional] 
**HsmEnabled** | Pointer to **bool** |  | [optional] 
**Hsks** | Pointer to [**[]Hsk**](Hsk.md) |  | [optional] 
**GrantManagementEndpoint** | Pointer to **string** |  | [optional] 
**GrantManagementActionRequired** | Pointer to **bool** |  | [optional] 
**UnauthorizedOnClientConfigSupported** | Pointer to **bool** |  | [optional] 
**DcrScopeUsedAsRequestable** | Pointer to **bool** |  | [optional] 
**PredefinedTransformedClaims** | Pointer to **string** |  | [optional] 
**LoopbackRedirectionUriVariable** | Pointer to **bool** |  | [optional] 
**RequestObjectAudienceChecked** | Pointer to **bool** |  | [optional] 
**AccessTokenForExternalAttachmentEmbedded** | Pointer to **bool** |  | [optional] 
**RefreshTokenIdempotent** | Pointer to **bool** |  | [optional] 
**FederationEnabled** | Pointer to **bool** |  | [optional] 
**OrganizationName** | Pointer to **string** |  | [optional] 
**AuthorityHints** | Pointer to **[]string** |  | [optional] 
**TrustAnchors** | Pointer to [**[]TrustAnchor**](TrustAnchor.md) |  | [optional] 
**FederationJwks** | Pointer to **string** |  | [optional] 
**FederationSignatureKeyId** | Pointer to **string** |  | [optional] 
**FederationConfigurationDuration** | Pointer to **int32** |  | [optional] 
**SignedJwksUri** | Pointer to **string** |  | [optional] 
**FederationRegistrationEndpoint** | Pointer to **string** |  | [optional] 
**SupportedClientRegistrationTypes** | Pointer to **[]string** |  | [optional] 
**TokenExchangeByIdentifiableClientsOnly** | Pointer to **bool** |  | [optional] 
**TokenExchangeByConfidentialClientsOnly** | Pointer to **bool** |  | [optional] 
**TokenExchangeByPermittedClientsOnly** | Pointer to **bool** |  | [optional] 
**TokenExchangeEncryptedJwtRejected** | Pointer to **bool** |  | [optional] 
**TokenExchangeUnsignedJwtRejected** | Pointer to **bool** |  | [optional] 
**JwtGrantByIdentifiableClientsOnly** | Pointer to **bool** |  | [optional] 
**JwtGrantEncryptedJwtRejected** | Pointer to **bool** |  | [optional] 
**JwtGrantUnsignedJwtRejected** | Pointer to **bool** |  | [optional] 
**DcrDuplicateSoftwareIdBlocked** | Pointer to **bool** |  | [optional] 
**ResourceSignatureKeyId** | Pointer to **string** |  | [optional] 
**RsResponseSigned** | Pointer to **bool** |  | [optional] 
**OpenidDroppedOnRefreshWithoutOfflineAccess** | Pointer to **bool** |  | [optional] 
**VerifiableCredentialsEnabled** | Pointer to **bool** |  | [optional] 
**CredentialIssuerMetadata** | Pointer to [**CredentialIssuerMetadata**](CredentialIssuerMetadata.md) |  | [optional] 
**CredentialOfferDuration** | Pointer to **int64** |  | [optional] 
**UserPinLength** | Pointer to **int32** |  | [optional] 
**IdTokenAudType** | Pointer to **string** |  | [optional] 
**SupportedPromptValues** | Pointer to **[]string** |  | [optional] 
**VerifiedClaimsValidationSchemaSet** | Pointer to **string** |  | [optional] 
**PreAuthorizedGrantAnonymousAccessSupported** | Pointer to **bool** |  | [optional] 
**CredentialTransactionDuration** | Pointer to **int64** |  | [optional] 
**CredentialDuration** | Pointer to **int64** |  | [optional] 
**CredentialJwks** | Pointer to **string** |  | [optional] 
**IdTokenReissuable** | Pointer to **bool** |  | [optional] 
**CnonceDuration** | Pointer to **int64** |  | [optional] 

## Methods

### NewService

`func NewService() *Service`

NewService instantiates a new Service object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceWithDefaults

`func NewServiceWithDefaults() *Service`

NewServiceWithDefaults instantiates a new Service object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *Service) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Service) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Service) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *Service) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetServiceOwnerNumber

`func (o *Service) GetServiceOwnerNumber() int32`

GetServiceOwnerNumber returns the ServiceOwnerNumber field if non-nil, zero value otherwise.

### GetServiceOwnerNumberOk

`func (o *Service) GetServiceOwnerNumberOk() (*int32, bool)`

GetServiceOwnerNumberOk returns a tuple with the ServiceOwnerNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceOwnerNumber

`func (o *Service) SetServiceOwnerNumber(v int32)`

SetServiceOwnerNumber sets ServiceOwnerNumber field to given value.

### HasServiceOwnerNumber

`func (o *Service) HasServiceOwnerNumber() bool`

HasServiceOwnerNumber returns a boolean if a field has been set.

### GetServiceName

`func (o *Service) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *Service) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *Service) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *Service) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetApiKey

`func (o *Service) GetApiKey() int64`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *Service) GetApiKeyOk() (*int64, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *Service) SetApiKey(v int64)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *Service) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetApiSecret

`func (o *Service) GetApiSecret() string`

GetApiSecret returns the ApiSecret field if non-nil, zero value otherwise.

### GetApiSecretOk

`func (o *Service) GetApiSecretOk() (*string, bool)`

GetApiSecretOk returns a tuple with the ApiSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiSecret

`func (o *Service) SetApiSecret(v string)`

SetApiSecret sets ApiSecret field to given value.

### HasApiSecret

`func (o *Service) HasApiSecret() bool`

HasApiSecret returns a boolean if a field has been set.

### GetIssuer

`func (o *Service) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *Service) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *Service) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *Service) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetAuthorizationEndpoint

`func (o *Service) GetAuthorizationEndpoint() string`

GetAuthorizationEndpoint returns the AuthorizationEndpoint field if non-nil, zero value otherwise.

### GetAuthorizationEndpointOk

`func (o *Service) GetAuthorizationEndpointOk() (*string, bool)`

GetAuthorizationEndpointOk returns a tuple with the AuthorizationEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationEndpoint

`func (o *Service) SetAuthorizationEndpoint(v string)`

SetAuthorizationEndpoint sets AuthorizationEndpoint field to given value.

### HasAuthorizationEndpoint

`func (o *Service) HasAuthorizationEndpoint() bool`

HasAuthorizationEndpoint returns a boolean if a field has been set.

### GetTokenEndpoint

`func (o *Service) GetTokenEndpoint() string`

GetTokenEndpoint returns the TokenEndpoint field if non-nil, zero value otherwise.

### GetTokenEndpointOk

`func (o *Service) GetTokenEndpointOk() (*string, bool)`

GetTokenEndpointOk returns a tuple with the TokenEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpoint

`func (o *Service) SetTokenEndpoint(v string)`

SetTokenEndpoint sets TokenEndpoint field to given value.

### HasTokenEndpoint

`func (o *Service) HasTokenEndpoint() bool`

HasTokenEndpoint returns a boolean if a field has been set.

### GetRevocationEndpoint

`func (o *Service) GetRevocationEndpoint() string`

GetRevocationEndpoint returns the RevocationEndpoint field if non-nil, zero value otherwise.

### GetRevocationEndpointOk

`func (o *Service) GetRevocationEndpointOk() (*string, bool)`

GetRevocationEndpointOk returns a tuple with the RevocationEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationEndpoint

`func (o *Service) SetRevocationEndpoint(v string)`

SetRevocationEndpoint sets RevocationEndpoint field to given value.

### HasRevocationEndpoint

`func (o *Service) HasRevocationEndpoint() bool`

HasRevocationEndpoint returns a boolean if a field has been set.

### GetSupportedRevocationAuthMethods

`func (o *Service) GetSupportedRevocationAuthMethods() []string`

GetSupportedRevocationAuthMethods returns the SupportedRevocationAuthMethods field if non-nil, zero value otherwise.

### GetSupportedRevocationAuthMethodsOk

`func (o *Service) GetSupportedRevocationAuthMethodsOk() (*[]string, bool)`

GetSupportedRevocationAuthMethodsOk returns a tuple with the SupportedRevocationAuthMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedRevocationAuthMethods

`func (o *Service) SetSupportedRevocationAuthMethods(v []string)`

SetSupportedRevocationAuthMethods sets SupportedRevocationAuthMethods field to given value.

### HasSupportedRevocationAuthMethods

`func (o *Service) HasSupportedRevocationAuthMethods() bool`

HasSupportedRevocationAuthMethods returns a boolean if a field has been set.

### GetUserInfoEndpoint

`func (o *Service) GetUserInfoEndpoint() string`

GetUserInfoEndpoint returns the UserInfoEndpoint field if non-nil, zero value otherwise.

### GetUserInfoEndpointOk

`func (o *Service) GetUserInfoEndpointOk() (*string, bool)`

GetUserInfoEndpointOk returns a tuple with the UserInfoEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInfoEndpoint

`func (o *Service) SetUserInfoEndpoint(v string)`

SetUserInfoEndpoint sets UserInfoEndpoint field to given value.

### HasUserInfoEndpoint

`func (o *Service) HasUserInfoEndpoint() bool`

HasUserInfoEndpoint returns a boolean if a field has been set.

### GetJwksUri

`func (o *Service) GetJwksUri() string`

GetJwksUri returns the JwksUri field if non-nil, zero value otherwise.

### GetJwksUriOk

`func (o *Service) GetJwksUriOk() (*string, bool)`

GetJwksUriOk returns a tuple with the JwksUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksUri

`func (o *Service) SetJwksUri(v string)`

SetJwksUri sets JwksUri field to given value.

### HasJwksUri

`func (o *Service) HasJwksUri() bool`

HasJwksUri returns a boolean if a field has been set.

### GetJwks

`func (o *Service) GetJwks() string`

GetJwks returns the Jwks field if non-nil, zero value otherwise.

### GetJwksOk

`func (o *Service) GetJwksOk() (*string, bool)`

GetJwksOk returns a tuple with the Jwks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwks

`func (o *Service) SetJwks(v string)`

SetJwks sets Jwks field to given value.

### HasJwks

`func (o *Service) HasJwks() bool`

HasJwks returns a boolean if a field has been set.

### GetRegistrationEndpoint

`func (o *Service) GetRegistrationEndpoint() string`

GetRegistrationEndpoint returns the RegistrationEndpoint field if non-nil, zero value otherwise.

### GetRegistrationEndpointOk

`func (o *Service) GetRegistrationEndpointOk() (*string, bool)`

GetRegistrationEndpointOk returns a tuple with the RegistrationEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationEndpoint

`func (o *Service) SetRegistrationEndpoint(v string)`

SetRegistrationEndpoint sets RegistrationEndpoint field to given value.

### HasRegistrationEndpoint

`func (o *Service) HasRegistrationEndpoint() bool`

HasRegistrationEndpoint returns a boolean if a field has been set.

### GetRegistrationManagementEndpoint

`func (o *Service) GetRegistrationManagementEndpoint() string`

GetRegistrationManagementEndpoint returns the RegistrationManagementEndpoint field if non-nil, zero value otherwise.

### GetRegistrationManagementEndpointOk

`func (o *Service) GetRegistrationManagementEndpointOk() (*string, bool)`

GetRegistrationManagementEndpointOk returns a tuple with the RegistrationManagementEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationManagementEndpoint

`func (o *Service) SetRegistrationManagementEndpoint(v string)`

SetRegistrationManagementEndpoint sets RegistrationManagementEndpoint field to given value.

### HasRegistrationManagementEndpoint

`func (o *Service) HasRegistrationManagementEndpoint() bool`

HasRegistrationManagementEndpoint returns a boolean if a field has been set.

### GetSupportedScopes

`func (o *Service) GetSupportedScopes() []Scope`

GetSupportedScopes returns the SupportedScopes field if non-nil, zero value otherwise.

### GetSupportedScopesOk

`func (o *Service) GetSupportedScopesOk() (*[]Scope, bool)`

GetSupportedScopesOk returns a tuple with the SupportedScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedScopes

`func (o *Service) SetSupportedScopes(v []Scope)`

SetSupportedScopes sets SupportedScopes field to given value.

### HasSupportedScopes

`func (o *Service) HasSupportedScopes() bool`

HasSupportedScopes returns a boolean if a field has been set.

### GetSupportedResponseTypes

`func (o *Service) GetSupportedResponseTypes() []string`

GetSupportedResponseTypes returns the SupportedResponseTypes field if non-nil, zero value otherwise.

### GetSupportedResponseTypesOk

`func (o *Service) GetSupportedResponseTypesOk() (*[]string, bool)`

GetSupportedResponseTypesOk returns a tuple with the SupportedResponseTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedResponseTypes

`func (o *Service) SetSupportedResponseTypes(v []string)`

SetSupportedResponseTypes sets SupportedResponseTypes field to given value.

### HasSupportedResponseTypes

`func (o *Service) HasSupportedResponseTypes() bool`

HasSupportedResponseTypes returns a boolean if a field has been set.

### GetSupportedGrantTypes

`func (o *Service) GetSupportedGrantTypes() []string`

GetSupportedGrantTypes returns the SupportedGrantTypes field if non-nil, zero value otherwise.

### GetSupportedGrantTypesOk

`func (o *Service) GetSupportedGrantTypesOk() (*[]string, bool)`

GetSupportedGrantTypesOk returns a tuple with the SupportedGrantTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedGrantTypes

`func (o *Service) SetSupportedGrantTypes(v []string)`

SetSupportedGrantTypes sets SupportedGrantTypes field to given value.

### HasSupportedGrantTypes

`func (o *Service) HasSupportedGrantTypes() bool`

HasSupportedGrantTypes returns a boolean if a field has been set.

### GetSupportedAcrs

`func (o *Service) GetSupportedAcrs() []string`

GetSupportedAcrs returns the SupportedAcrs field if non-nil, zero value otherwise.

### GetSupportedAcrsOk

`func (o *Service) GetSupportedAcrsOk() (*[]string, bool)`

GetSupportedAcrsOk returns a tuple with the SupportedAcrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedAcrs

`func (o *Service) SetSupportedAcrs(v []string)`

SetSupportedAcrs sets SupportedAcrs field to given value.

### HasSupportedAcrs

`func (o *Service) HasSupportedAcrs() bool`

HasSupportedAcrs returns a boolean if a field has been set.

### GetSupportedTokenAuthMethods

`func (o *Service) GetSupportedTokenAuthMethods() []string`

GetSupportedTokenAuthMethods returns the SupportedTokenAuthMethods field if non-nil, zero value otherwise.

### GetSupportedTokenAuthMethodsOk

`func (o *Service) GetSupportedTokenAuthMethodsOk() (*[]string, bool)`

GetSupportedTokenAuthMethodsOk returns a tuple with the SupportedTokenAuthMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedTokenAuthMethods

`func (o *Service) SetSupportedTokenAuthMethods(v []string)`

SetSupportedTokenAuthMethods sets SupportedTokenAuthMethods field to given value.

### HasSupportedTokenAuthMethods

`func (o *Service) HasSupportedTokenAuthMethods() bool`

HasSupportedTokenAuthMethods returns a boolean if a field has been set.

### GetSupportedDisplays

`func (o *Service) GetSupportedDisplays() []string`

GetSupportedDisplays returns the SupportedDisplays field if non-nil, zero value otherwise.

### GetSupportedDisplaysOk

`func (o *Service) GetSupportedDisplaysOk() (*[]string, bool)`

GetSupportedDisplaysOk returns a tuple with the SupportedDisplays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedDisplays

`func (o *Service) SetSupportedDisplays(v []string)`

SetSupportedDisplays sets SupportedDisplays field to given value.

### HasSupportedDisplays

`func (o *Service) HasSupportedDisplays() bool`

HasSupportedDisplays returns a boolean if a field has been set.

### GetSupportedClaimTypes

`func (o *Service) GetSupportedClaimTypes() []string`

GetSupportedClaimTypes returns the SupportedClaimTypes field if non-nil, zero value otherwise.

### GetSupportedClaimTypesOk

`func (o *Service) GetSupportedClaimTypesOk() (*[]string, bool)`

GetSupportedClaimTypesOk returns a tuple with the SupportedClaimTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedClaimTypes

`func (o *Service) SetSupportedClaimTypes(v []string)`

SetSupportedClaimTypes sets SupportedClaimTypes field to given value.

### HasSupportedClaimTypes

`func (o *Service) HasSupportedClaimTypes() bool`

HasSupportedClaimTypes returns a boolean if a field has been set.

### GetSupportedClaims

`func (o *Service) GetSupportedClaims() []string`

GetSupportedClaims returns the SupportedClaims field if non-nil, zero value otherwise.

### GetSupportedClaimsOk

`func (o *Service) GetSupportedClaimsOk() (*[]string, bool)`

GetSupportedClaimsOk returns a tuple with the SupportedClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedClaims

`func (o *Service) SetSupportedClaims(v []string)`

SetSupportedClaims sets SupportedClaims field to given value.

### HasSupportedClaims

`func (o *Service) HasSupportedClaims() bool`

HasSupportedClaims returns a boolean if a field has been set.

### GetServiceDocumentation

`func (o *Service) GetServiceDocumentation() string`

GetServiceDocumentation returns the ServiceDocumentation field if non-nil, zero value otherwise.

### GetServiceDocumentationOk

`func (o *Service) GetServiceDocumentationOk() (*string, bool)`

GetServiceDocumentationOk returns a tuple with the ServiceDocumentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDocumentation

`func (o *Service) SetServiceDocumentation(v string)`

SetServiceDocumentation sets ServiceDocumentation field to given value.

### HasServiceDocumentation

`func (o *Service) HasServiceDocumentation() bool`

HasServiceDocumentation returns a boolean if a field has been set.

### GetSupportedClaimLocales

`func (o *Service) GetSupportedClaimLocales() []string`

GetSupportedClaimLocales returns the SupportedClaimLocales field if non-nil, zero value otherwise.

### GetSupportedClaimLocalesOk

`func (o *Service) GetSupportedClaimLocalesOk() (*[]string, bool)`

GetSupportedClaimLocalesOk returns a tuple with the SupportedClaimLocales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedClaimLocales

`func (o *Service) SetSupportedClaimLocales(v []string)`

SetSupportedClaimLocales sets SupportedClaimLocales field to given value.

### HasSupportedClaimLocales

`func (o *Service) HasSupportedClaimLocales() bool`

HasSupportedClaimLocales returns a boolean if a field has been set.

### GetSupportedUiLocales

`func (o *Service) GetSupportedUiLocales() []string`

GetSupportedUiLocales returns the SupportedUiLocales field if non-nil, zero value otherwise.

### GetSupportedUiLocalesOk

`func (o *Service) GetSupportedUiLocalesOk() (*[]string, bool)`

GetSupportedUiLocalesOk returns a tuple with the SupportedUiLocales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedUiLocales

`func (o *Service) SetSupportedUiLocales(v []string)`

SetSupportedUiLocales sets SupportedUiLocales field to given value.

### HasSupportedUiLocales

`func (o *Service) HasSupportedUiLocales() bool`

HasSupportedUiLocales returns a boolean if a field has been set.

### GetPolicyUri

`func (o *Service) GetPolicyUri() string`

GetPolicyUri returns the PolicyUri field if non-nil, zero value otherwise.

### GetPolicyUriOk

`func (o *Service) GetPolicyUriOk() (*string, bool)`

GetPolicyUriOk returns a tuple with the PolicyUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyUri

`func (o *Service) SetPolicyUri(v string)`

SetPolicyUri sets PolicyUri field to given value.

### HasPolicyUri

`func (o *Service) HasPolicyUri() bool`

HasPolicyUri returns a boolean if a field has been set.

### GetTosUri

`func (o *Service) GetTosUri() string`

GetTosUri returns the TosUri field if non-nil, zero value otherwise.

### GetTosUriOk

`func (o *Service) GetTosUriOk() (*string, bool)`

GetTosUriOk returns a tuple with the TosUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTosUri

`func (o *Service) SetTosUri(v string)`

SetTosUri sets TosUri field to given value.

### HasTosUri

`func (o *Service) HasTosUri() bool`

HasTosUri returns a boolean if a field has been set.

### GetAuthenticationCallbackEndpoint

`func (o *Service) GetAuthenticationCallbackEndpoint() string`

GetAuthenticationCallbackEndpoint returns the AuthenticationCallbackEndpoint field if non-nil, zero value otherwise.

### GetAuthenticationCallbackEndpointOk

`func (o *Service) GetAuthenticationCallbackEndpointOk() (*string, bool)`

GetAuthenticationCallbackEndpointOk returns a tuple with the AuthenticationCallbackEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationCallbackEndpoint

`func (o *Service) SetAuthenticationCallbackEndpoint(v string)`

SetAuthenticationCallbackEndpoint sets AuthenticationCallbackEndpoint field to given value.

### HasAuthenticationCallbackEndpoint

`func (o *Service) HasAuthenticationCallbackEndpoint() bool`

HasAuthenticationCallbackEndpoint returns a boolean if a field has been set.

### GetAuthenticationCallbackApiKey

`func (o *Service) GetAuthenticationCallbackApiKey() string`

GetAuthenticationCallbackApiKey returns the AuthenticationCallbackApiKey field if non-nil, zero value otherwise.

### GetAuthenticationCallbackApiKeyOk

`func (o *Service) GetAuthenticationCallbackApiKeyOk() (*string, bool)`

GetAuthenticationCallbackApiKeyOk returns a tuple with the AuthenticationCallbackApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationCallbackApiKey

`func (o *Service) SetAuthenticationCallbackApiKey(v string)`

SetAuthenticationCallbackApiKey sets AuthenticationCallbackApiKey field to given value.

### HasAuthenticationCallbackApiKey

`func (o *Service) HasAuthenticationCallbackApiKey() bool`

HasAuthenticationCallbackApiKey returns a boolean if a field has been set.

### GetAuthenticationCallbackApiSecret

`func (o *Service) GetAuthenticationCallbackApiSecret() string`

GetAuthenticationCallbackApiSecret returns the AuthenticationCallbackApiSecret field if non-nil, zero value otherwise.

### GetAuthenticationCallbackApiSecretOk

`func (o *Service) GetAuthenticationCallbackApiSecretOk() (*string, bool)`

GetAuthenticationCallbackApiSecretOk returns a tuple with the AuthenticationCallbackApiSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationCallbackApiSecret

`func (o *Service) SetAuthenticationCallbackApiSecret(v string)`

SetAuthenticationCallbackApiSecret sets AuthenticationCallbackApiSecret field to given value.

### HasAuthenticationCallbackApiSecret

`func (o *Service) HasAuthenticationCallbackApiSecret() bool`

HasAuthenticationCallbackApiSecret returns a boolean if a field has been set.

### GetSupportedSnses

`func (o *Service) GetSupportedSnses() []string`

GetSupportedSnses returns the SupportedSnses field if non-nil, zero value otherwise.

### GetSupportedSnsesOk

`func (o *Service) GetSupportedSnsesOk() (*[]string, bool)`

GetSupportedSnsesOk returns a tuple with the SupportedSnses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedSnses

`func (o *Service) SetSupportedSnses(v []string)`

SetSupportedSnses sets SupportedSnses field to given value.

### HasSupportedSnses

`func (o *Service) HasSupportedSnses() bool`

HasSupportedSnses returns a boolean if a field has been set.

### GetSnsCredentials

`func (o *Service) GetSnsCredentials() []SnsCredentials`

GetSnsCredentials returns the SnsCredentials field if non-nil, zero value otherwise.

### GetSnsCredentialsOk

`func (o *Service) GetSnsCredentialsOk() (*[]SnsCredentials, bool)`

GetSnsCredentialsOk returns a tuple with the SnsCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnsCredentials

`func (o *Service) SetSnsCredentials(v []SnsCredentials)`

SetSnsCredentials sets SnsCredentials field to given value.

### HasSnsCredentials

`func (o *Service) HasSnsCredentials() bool`

HasSnsCredentials returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Service) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Service) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Service) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Service) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *Service) GetModifiedAt() int64`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Service) GetModifiedAtOk() (*int64, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Service) SetModifiedAt(v int64)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *Service) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetDeveloperAuthenticationCallbackEndpoint

`func (o *Service) GetDeveloperAuthenticationCallbackEndpoint() string`

GetDeveloperAuthenticationCallbackEndpoint returns the DeveloperAuthenticationCallbackEndpoint field if non-nil, zero value otherwise.

### GetDeveloperAuthenticationCallbackEndpointOk

`func (o *Service) GetDeveloperAuthenticationCallbackEndpointOk() (*string, bool)`

GetDeveloperAuthenticationCallbackEndpointOk returns a tuple with the DeveloperAuthenticationCallbackEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperAuthenticationCallbackEndpoint

`func (o *Service) SetDeveloperAuthenticationCallbackEndpoint(v string)`

SetDeveloperAuthenticationCallbackEndpoint sets DeveloperAuthenticationCallbackEndpoint field to given value.

### HasDeveloperAuthenticationCallbackEndpoint

`func (o *Service) HasDeveloperAuthenticationCallbackEndpoint() bool`

HasDeveloperAuthenticationCallbackEndpoint returns a boolean if a field has been set.

### GetDeveloperAuthenticationCallbackApiKey

`func (o *Service) GetDeveloperAuthenticationCallbackApiKey() string`

GetDeveloperAuthenticationCallbackApiKey returns the DeveloperAuthenticationCallbackApiKey field if non-nil, zero value otherwise.

### GetDeveloperAuthenticationCallbackApiKeyOk

`func (o *Service) GetDeveloperAuthenticationCallbackApiKeyOk() (*string, bool)`

GetDeveloperAuthenticationCallbackApiKeyOk returns a tuple with the DeveloperAuthenticationCallbackApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperAuthenticationCallbackApiKey

`func (o *Service) SetDeveloperAuthenticationCallbackApiKey(v string)`

SetDeveloperAuthenticationCallbackApiKey sets DeveloperAuthenticationCallbackApiKey field to given value.

### HasDeveloperAuthenticationCallbackApiKey

`func (o *Service) HasDeveloperAuthenticationCallbackApiKey() bool`

HasDeveloperAuthenticationCallbackApiKey returns a boolean if a field has been set.

### GetDeveloperAuthenticationCallbackApiSecret

`func (o *Service) GetDeveloperAuthenticationCallbackApiSecret() string`

GetDeveloperAuthenticationCallbackApiSecret returns the DeveloperAuthenticationCallbackApiSecret field if non-nil, zero value otherwise.

### GetDeveloperAuthenticationCallbackApiSecretOk

`func (o *Service) GetDeveloperAuthenticationCallbackApiSecretOk() (*string, bool)`

GetDeveloperAuthenticationCallbackApiSecretOk returns a tuple with the DeveloperAuthenticationCallbackApiSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperAuthenticationCallbackApiSecret

`func (o *Service) SetDeveloperAuthenticationCallbackApiSecret(v string)`

SetDeveloperAuthenticationCallbackApiSecret sets DeveloperAuthenticationCallbackApiSecret field to given value.

### HasDeveloperAuthenticationCallbackApiSecret

`func (o *Service) HasDeveloperAuthenticationCallbackApiSecret() bool`

HasDeveloperAuthenticationCallbackApiSecret returns a boolean if a field has been set.

### GetSupportedDeveloperSnses

`func (o *Service) GetSupportedDeveloperSnses() []string`

GetSupportedDeveloperSnses returns the SupportedDeveloperSnses field if non-nil, zero value otherwise.

### GetSupportedDeveloperSnsesOk

`func (o *Service) GetSupportedDeveloperSnsesOk() (*[]string, bool)`

GetSupportedDeveloperSnsesOk returns a tuple with the SupportedDeveloperSnses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedDeveloperSnses

`func (o *Service) SetSupportedDeveloperSnses(v []string)`

SetSupportedDeveloperSnses sets SupportedDeveloperSnses field to given value.

### HasSupportedDeveloperSnses

`func (o *Service) HasSupportedDeveloperSnses() bool`

HasSupportedDeveloperSnses returns a boolean if a field has been set.

### GetDeveloperSnsCredentials

`func (o *Service) GetDeveloperSnsCredentials() []SnsCredentials`

GetDeveloperSnsCredentials returns the DeveloperSnsCredentials field if non-nil, zero value otherwise.

### GetDeveloperSnsCredentialsOk

`func (o *Service) GetDeveloperSnsCredentialsOk() (*[]SnsCredentials, bool)`

GetDeveloperSnsCredentialsOk returns a tuple with the DeveloperSnsCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperSnsCredentials

`func (o *Service) SetDeveloperSnsCredentials(v []SnsCredentials)`

SetDeveloperSnsCredentials sets DeveloperSnsCredentials field to given value.

### HasDeveloperSnsCredentials

`func (o *Service) HasDeveloperSnsCredentials() bool`

HasDeveloperSnsCredentials returns a boolean if a field has been set.

### GetClientsPerDeveloper

`func (o *Service) GetClientsPerDeveloper() int32`

GetClientsPerDeveloper returns the ClientsPerDeveloper field if non-nil, zero value otherwise.

### GetClientsPerDeveloperOk

`func (o *Service) GetClientsPerDeveloperOk() (*int32, bool)`

GetClientsPerDeveloperOk returns a tuple with the ClientsPerDeveloper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientsPerDeveloper

`func (o *Service) SetClientsPerDeveloper(v int32)`

SetClientsPerDeveloper sets ClientsPerDeveloper field to given value.

### HasClientsPerDeveloper

`func (o *Service) HasClientsPerDeveloper() bool`

HasClientsPerDeveloper returns a boolean if a field has been set.

### GetDirectAuthorizationEndpointEnabled

`func (o *Service) GetDirectAuthorizationEndpointEnabled() bool`

GetDirectAuthorizationEndpointEnabled returns the DirectAuthorizationEndpointEnabled field if non-nil, zero value otherwise.

### GetDirectAuthorizationEndpointEnabledOk

`func (o *Service) GetDirectAuthorizationEndpointEnabledOk() (*bool, bool)`

GetDirectAuthorizationEndpointEnabledOk returns a tuple with the DirectAuthorizationEndpointEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectAuthorizationEndpointEnabled

`func (o *Service) SetDirectAuthorizationEndpointEnabled(v bool)`

SetDirectAuthorizationEndpointEnabled sets DirectAuthorizationEndpointEnabled field to given value.

### HasDirectAuthorizationEndpointEnabled

`func (o *Service) HasDirectAuthorizationEndpointEnabled() bool`

HasDirectAuthorizationEndpointEnabled returns a boolean if a field has been set.

### GetDirectTokenEndpointEnabled

`func (o *Service) GetDirectTokenEndpointEnabled() bool`

GetDirectTokenEndpointEnabled returns the DirectTokenEndpointEnabled field if non-nil, zero value otherwise.

### GetDirectTokenEndpointEnabledOk

`func (o *Service) GetDirectTokenEndpointEnabledOk() (*bool, bool)`

GetDirectTokenEndpointEnabledOk returns a tuple with the DirectTokenEndpointEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectTokenEndpointEnabled

`func (o *Service) SetDirectTokenEndpointEnabled(v bool)`

SetDirectTokenEndpointEnabled sets DirectTokenEndpointEnabled field to given value.

### HasDirectTokenEndpointEnabled

`func (o *Service) HasDirectTokenEndpointEnabled() bool`

HasDirectTokenEndpointEnabled returns a boolean if a field has been set.

### GetDirectRevocationEndpointEnabled

`func (o *Service) GetDirectRevocationEndpointEnabled() bool`

GetDirectRevocationEndpointEnabled returns the DirectRevocationEndpointEnabled field if non-nil, zero value otherwise.

### GetDirectRevocationEndpointEnabledOk

`func (o *Service) GetDirectRevocationEndpointEnabledOk() (*bool, bool)`

GetDirectRevocationEndpointEnabledOk returns a tuple with the DirectRevocationEndpointEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectRevocationEndpointEnabled

`func (o *Service) SetDirectRevocationEndpointEnabled(v bool)`

SetDirectRevocationEndpointEnabled sets DirectRevocationEndpointEnabled field to given value.

### HasDirectRevocationEndpointEnabled

`func (o *Service) HasDirectRevocationEndpointEnabled() bool`

HasDirectRevocationEndpointEnabled returns a boolean if a field has been set.

### GetDirectUserInfoEndpointEnabled

`func (o *Service) GetDirectUserInfoEndpointEnabled() bool`

GetDirectUserInfoEndpointEnabled returns the DirectUserInfoEndpointEnabled field if non-nil, zero value otherwise.

### GetDirectUserInfoEndpointEnabledOk

`func (o *Service) GetDirectUserInfoEndpointEnabledOk() (*bool, bool)`

GetDirectUserInfoEndpointEnabledOk returns a tuple with the DirectUserInfoEndpointEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectUserInfoEndpointEnabled

`func (o *Service) SetDirectUserInfoEndpointEnabled(v bool)`

SetDirectUserInfoEndpointEnabled sets DirectUserInfoEndpointEnabled field to given value.

### HasDirectUserInfoEndpointEnabled

`func (o *Service) HasDirectUserInfoEndpointEnabled() bool`

HasDirectUserInfoEndpointEnabled returns a boolean if a field has been set.

### GetDirectJwksEndpointEnabled

`func (o *Service) GetDirectJwksEndpointEnabled() bool`

GetDirectJwksEndpointEnabled returns the DirectJwksEndpointEnabled field if non-nil, zero value otherwise.

### GetDirectJwksEndpointEnabledOk

`func (o *Service) GetDirectJwksEndpointEnabledOk() (*bool, bool)`

GetDirectJwksEndpointEnabledOk returns a tuple with the DirectJwksEndpointEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectJwksEndpointEnabled

`func (o *Service) SetDirectJwksEndpointEnabled(v bool)`

SetDirectJwksEndpointEnabled sets DirectJwksEndpointEnabled field to given value.

### HasDirectJwksEndpointEnabled

`func (o *Service) HasDirectJwksEndpointEnabled() bool`

HasDirectJwksEndpointEnabled returns a boolean if a field has been set.

### GetDirectIntrospectionEndpointEnabled

`func (o *Service) GetDirectIntrospectionEndpointEnabled() bool`

GetDirectIntrospectionEndpointEnabled returns the DirectIntrospectionEndpointEnabled field if non-nil, zero value otherwise.

### GetDirectIntrospectionEndpointEnabledOk

`func (o *Service) GetDirectIntrospectionEndpointEnabledOk() (*bool, bool)`

GetDirectIntrospectionEndpointEnabledOk returns a tuple with the DirectIntrospectionEndpointEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectIntrospectionEndpointEnabled

`func (o *Service) SetDirectIntrospectionEndpointEnabled(v bool)`

SetDirectIntrospectionEndpointEnabled sets DirectIntrospectionEndpointEnabled field to given value.

### HasDirectIntrospectionEndpointEnabled

`func (o *Service) HasDirectIntrospectionEndpointEnabled() bool`

HasDirectIntrospectionEndpointEnabled returns a boolean if a field has been set.

### GetSingleAccessTokenPerSubject

`func (o *Service) GetSingleAccessTokenPerSubject() bool`

GetSingleAccessTokenPerSubject returns the SingleAccessTokenPerSubject field if non-nil, zero value otherwise.

### GetSingleAccessTokenPerSubjectOk

`func (o *Service) GetSingleAccessTokenPerSubjectOk() (*bool, bool)`

GetSingleAccessTokenPerSubjectOk returns a tuple with the SingleAccessTokenPerSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleAccessTokenPerSubject

`func (o *Service) SetSingleAccessTokenPerSubject(v bool)`

SetSingleAccessTokenPerSubject sets SingleAccessTokenPerSubject field to given value.

### HasSingleAccessTokenPerSubject

`func (o *Service) HasSingleAccessTokenPerSubject() bool`

HasSingleAccessTokenPerSubject returns a boolean if a field has been set.

### GetPkceRequired

`func (o *Service) GetPkceRequired() bool`

GetPkceRequired returns the PkceRequired field if non-nil, zero value otherwise.

### GetPkceRequiredOk

`func (o *Service) GetPkceRequiredOk() (*bool, bool)`

GetPkceRequiredOk returns a tuple with the PkceRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkceRequired

`func (o *Service) SetPkceRequired(v bool)`

SetPkceRequired sets PkceRequired field to given value.

### HasPkceRequired

`func (o *Service) HasPkceRequired() bool`

HasPkceRequired returns a boolean if a field has been set.

### GetPkceS256Required

`func (o *Service) GetPkceS256Required() bool`

GetPkceS256Required returns the PkceS256Required field if non-nil, zero value otherwise.

### GetPkceS256RequiredOk

`func (o *Service) GetPkceS256RequiredOk() (*bool, bool)`

GetPkceS256RequiredOk returns a tuple with the PkceS256Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkceS256Required

`func (o *Service) SetPkceS256Required(v bool)`

SetPkceS256Required sets PkceS256Required field to given value.

### HasPkceS256Required

`func (o *Service) HasPkceS256Required() bool`

HasPkceS256Required returns a boolean if a field has been set.

### GetRefreshTokenKept

`func (o *Service) GetRefreshTokenKept() bool`

GetRefreshTokenKept returns the RefreshTokenKept field if non-nil, zero value otherwise.

### GetRefreshTokenKeptOk

`func (o *Service) GetRefreshTokenKeptOk() (*bool, bool)`

GetRefreshTokenKeptOk returns a tuple with the RefreshTokenKept field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenKept

`func (o *Service) SetRefreshTokenKept(v bool)`

SetRefreshTokenKept sets RefreshTokenKept field to given value.

### HasRefreshTokenKept

`func (o *Service) HasRefreshTokenKept() bool`

HasRefreshTokenKept returns a boolean if a field has been set.

### GetRefreshTokenDurationKept

`func (o *Service) GetRefreshTokenDurationKept() bool`

GetRefreshTokenDurationKept returns the RefreshTokenDurationKept field if non-nil, zero value otherwise.

### GetRefreshTokenDurationKeptOk

`func (o *Service) GetRefreshTokenDurationKeptOk() (*bool, bool)`

GetRefreshTokenDurationKeptOk returns a tuple with the RefreshTokenDurationKept field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenDurationKept

`func (o *Service) SetRefreshTokenDurationKept(v bool)`

SetRefreshTokenDurationKept sets RefreshTokenDurationKept field to given value.

### HasRefreshTokenDurationKept

`func (o *Service) HasRefreshTokenDurationKept() bool`

HasRefreshTokenDurationKept returns a boolean if a field has been set.

### GetRefreshTokenDurationReset

`func (o *Service) GetRefreshTokenDurationReset() bool`

GetRefreshTokenDurationReset returns the RefreshTokenDurationReset field if non-nil, zero value otherwise.

### GetRefreshTokenDurationResetOk

`func (o *Service) GetRefreshTokenDurationResetOk() (*bool, bool)`

GetRefreshTokenDurationResetOk returns a tuple with the RefreshTokenDurationReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenDurationReset

`func (o *Service) SetRefreshTokenDurationReset(v bool)`

SetRefreshTokenDurationReset sets RefreshTokenDurationReset field to given value.

### HasRefreshTokenDurationReset

`func (o *Service) HasRefreshTokenDurationReset() bool`

HasRefreshTokenDurationReset returns a boolean if a field has been set.

### GetErrorDescriptionOmitted

`func (o *Service) GetErrorDescriptionOmitted() bool`

GetErrorDescriptionOmitted returns the ErrorDescriptionOmitted field if non-nil, zero value otherwise.

### GetErrorDescriptionOmittedOk

`func (o *Service) GetErrorDescriptionOmittedOk() (*bool, bool)`

GetErrorDescriptionOmittedOk returns a tuple with the ErrorDescriptionOmitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescriptionOmitted

`func (o *Service) SetErrorDescriptionOmitted(v bool)`

SetErrorDescriptionOmitted sets ErrorDescriptionOmitted field to given value.

### HasErrorDescriptionOmitted

`func (o *Service) HasErrorDescriptionOmitted() bool`

HasErrorDescriptionOmitted returns a boolean if a field has been set.

### GetErrorUriOmitted

`func (o *Service) GetErrorUriOmitted() bool`

GetErrorUriOmitted returns the ErrorUriOmitted field if non-nil, zero value otherwise.

### GetErrorUriOmittedOk

`func (o *Service) GetErrorUriOmittedOk() (*bool, bool)`

GetErrorUriOmittedOk returns a tuple with the ErrorUriOmitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorUriOmitted

`func (o *Service) SetErrorUriOmitted(v bool)`

SetErrorUriOmitted sets ErrorUriOmitted field to given value.

### HasErrorUriOmitted

`func (o *Service) HasErrorUriOmitted() bool`

HasErrorUriOmitted returns a boolean if a field has been set.

### GetClientIdAliasEnabled

`func (o *Service) GetClientIdAliasEnabled() bool`

GetClientIdAliasEnabled returns the ClientIdAliasEnabled field if non-nil, zero value otherwise.

### GetClientIdAliasEnabledOk

`func (o *Service) GetClientIdAliasEnabledOk() (*bool, bool)`

GetClientIdAliasEnabledOk returns a tuple with the ClientIdAliasEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIdAliasEnabled

`func (o *Service) SetClientIdAliasEnabled(v bool)`

SetClientIdAliasEnabled sets ClientIdAliasEnabled field to given value.

### HasClientIdAliasEnabled

`func (o *Service) HasClientIdAliasEnabled() bool`

HasClientIdAliasEnabled returns a boolean if a field has been set.

### GetSupportedServiceProfiles

`func (o *Service) GetSupportedServiceProfiles() []string`

GetSupportedServiceProfiles returns the SupportedServiceProfiles field if non-nil, zero value otherwise.

### GetSupportedServiceProfilesOk

`func (o *Service) GetSupportedServiceProfilesOk() (*[]string, bool)`

GetSupportedServiceProfilesOk returns a tuple with the SupportedServiceProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedServiceProfiles

`func (o *Service) SetSupportedServiceProfiles(v []string)`

SetSupportedServiceProfiles sets SupportedServiceProfiles field to given value.

### HasSupportedServiceProfiles

`func (o *Service) HasSupportedServiceProfiles() bool`

HasSupportedServiceProfiles returns a boolean if a field has been set.

### GetTlsClientCertificateBoundAccessTokens

`func (o *Service) GetTlsClientCertificateBoundAccessTokens() bool`

GetTlsClientCertificateBoundAccessTokens returns the TlsClientCertificateBoundAccessTokens field if non-nil, zero value otherwise.

### GetTlsClientCertificateBoundAccessTokensOk

`func (o *Service) GetTlsClientCertificateBoundAccessTokensOk() (*bool, bool)`

GetTlsClientCertificateBoundAccessTokensOk returns a tuple with the TlsClientCertificateBoundAccessTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsClientCertificateBoundAccessTokens

`func (o *Service) SetTlsClientCertificateBoundAccessTokens(v bool)`

SetTlsClientCertificateBoundAccessTokens sets TlsClientCertificateBoundAccessTokens field to given value.

### HasTlsClientCertificateBoundAccessTokens

`func (o *Service) HasTlsClientCertificateBoundAccessTokens() bool`

HasTlsClientCertificateBoundAccessTokens returns a boolean if a field has been set.

### GetIntrospectionEndpoint

`func (o *Service) GetIntrospectionEndpoint() string`

GetIntrospectionEndpoint returns the IntrospectionEndpoint field if non-nil, zero value otherwise.

### GetIntrospectionEndpointOk

`func (o *Service) GetIntrospectionEndpointOk() (*string, bool)`

GetIntrospectionEndpointOk returns a tuple with the IntrospectionEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntrospectionEndpoint

`func (o *Service) SetIntrospectionEndpoint(v string)`

SetIntrospectionEndpoint sets IntrospectionEndpoint field to given value.

### HasIntrospectionEndpoint

`func (o *Service) HasIntrospectionEndpoint() bool`

HasIntrospectionEndpoint returns a boolean if a field has been set.

### GetSupportedIntrospectionAuthMethods

`func (o *Service) GetSupportedIntrospectionAuthMethods() []string`

GetSupportedIntrospectionAuthMethods returns the SupportedIntrospectionAuthMethods field if non-nil, zero value otherwise.

### GetSupportedIntrospectionAuthMethodsOk

`func (o *Service) GetSupportedIntrospectionAuthMethodsOk() (*[]string, bool)`

GetSupportedIntrospectionAuthMethodsOk returns a tuple with the SupportedIntrospectionAuthMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedIntrospectionAuthMethods

`func (o *Service) SetSupportedIntrospectionAuthMethods(v []string)`

SetSupportedIntrospectionAuthMethods sets SupportedIntrospectionAuthMethods field to given value.

### HasSupportedIntrospectionAuthMethods

`func (o *Service) HasSupportedIntrospectionAuthMethods() bool`

HasSupportedIntrospectionAuthMethods returns a boolean if a field has been set.

### GetMutualTlsValidatePkiCertChain

`func (o *Service) GetMutualTlsValidatePkiCertChain() bool`

GetMutualTlsValidatePkiCertChain returns the MutualTlsValidatePkiCertChain field if non-nil, zero value otherwise.

### GetMutualTlsValidatePkiCertChainOk

`func (o *Service) GetMutualTlsValidatePkiCertChainOk() (*bool, bool)`

GetMutualTlsValidatePkiCertChainOk returns a tuple with the MutualTlsValidatePkiCertChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutualTlsValidatePkiCertChain

`func (o *Service) SetMutualTlsValidatePkiCertChain(v bool)`

SetMutualTlsValidatePkiCertChain sets MutualTlsValidatePkiCertChain field to given value.

### HasMutualTlsValidatePkiCertChain

`func (o *Service) HasMutualTlsValidatePkiCertChain() bool`

HasMutualTlsValidatePkiCertChain returns a boolean if a field has been set.

### GetTrustedRootCertificates

`func (o *Service) GetTrustedRootCertificates() []string`

GetTrustedRootCertificates returns the TrustedRootCertificates field if non-nil, zero value otherwise.

### GetTrustedRootCertificatesOk

`func (o *Service) GetTrustedRootCertificatesOk() (*[]string, bool)`

GetTrustedRootCertificatesOk returns a tuple with the TrustedRootCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedRootCertificates

`func (o *Service) SetTrustedRootCertificates(v []string)`

SetTrustedRootCertificates sets TrustedRootCertificates field to given value.

### HasTrustedRootCertificates

`func (o *Service) HasTrustedRootCertificates() bool`

HasTrustedRootCertificates returns a boolean if a field has been set.

### GetDynamicRegistrationSupported

`func (o *Service) GetDynamicRegistrationSupported() bool`

GetDynamicRegistrationSupported returns the DynamicRegistrationSupported field if non-nil, zero value otherwise.

### GetDynamicRegistrationSupportedOk

`func (o *Service) GetDynamicRegistrationSupportedOk() (*bool, bool)`

GetDynamicRegistrationSupportedOk returns a tuple with the DynamicRegistrationSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicRegistrationSupported

`func (o *Service) SetDynamicRegistrationSupported(v bool)`

SetDynamicRegistrationSupported sets DynamicRegistrationSupported field to given value.

### HasDynamicRegistrationSupported

`func (o *Service) HasDynamicRegistrationSupported() bool`

HasDynamicRegistrationSupported returns a boolean if a field has been set.

### GetEndSessionEndpoint

`func (o *Service) GetEndSessionEndpoint() string`

GetEndSessionEndpoint returns the EndSessionEndpoint field if non-nil, zero value otherwise.

### GetEndSessionEndpointOk

`func (o *Service) GetEndSessionEndpointOk() (*string, bool)`

GetEndSessionEndpointOk returns a tuple with the EndSessionEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndSessionEndpoint

`func (o *Service) SetEndSessionEndpoint(v string)`

SetEndSessionEndpoint sets EndSessionEndpoint field to given value.

### HasEndSessionEndpoint

`func (o *Service) HasEndSessionEndpoint() bool`

HasEndSessionEndpoint returns a boolean if a field has been set.

### GetDescription

`func (o *Service) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Service) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Service) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Service) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAccessTokenType

`func (o *Service) GetAccessTokenType() string`

GetAccessTokenType returns the AccessTokenType field if non-nil, zero value otherwise.

### GetAccessTokenTypeOk

`func (o *Service) GetAccessTokenTypeOk() (*string, bool)`

GetAccessTokenTypeOk returns a tuple with the AccessTokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenType

`func (o *Service) SetAccessTokenType(v string)`

SetAccessTokenType sets AccessTokenType field to given value.

### HasAccessTokenType

`func (o *Service) HasAccessTokenType() bool`

HasAccessTokenType returns a boolean if a field has been set.

### GetAccessTokenSignAlg

`func (o *Service) GetAccessTokenSignAlg() string`

GetAccessTokenSignAlg returns the AccessTokenSignAlg field if non-nil, zero value otherwise.

### GetAccessTokenSignAlgOk

`func (o *Service) GetAccessTokenSignAlgOk() (*string, bool)`

GetAccessTokenSignAlgOk returns a tuple with the AccessTokenSignAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenSignAlg

`func (o *Service) SetAccessTokenSignAlg(v string)`

SetAccessTokenSignAlg sets AccessTokenSignAlg field to given value.

### HasAccessTokenSignAlg

`func (o *Service) HasAccessTokenSignAlg() bool`

HasAccessTokenSignAlg returns a boolean if a field has been set.

### GetAccessTokenDuration

`func (o *Service) GetAccessTokenDuration() int64`

GetAccessTokenDuration returns the AccessTokenDuration field if non-nil, zero value otherwise.

### GetAccessTokenDurationOk

`func (o *Service) GetAccessTokenDurationOk() (*int64, bool)`

GetAccessTokenDurationOk returns a tuple with the AccessTokenDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenDuration

`func (o *Service) SetAccessTokenDuration(v int64)`

SetAccessTokenDuration sets AccessTokenDuration field to given value.

### HasAccessTokenDuration

`func (o *Service) HasAccessTokenDuration() bool`

HasAccessTokenDuration returns a boolean if a field has been set.

### GetRefreshTokenDuration

`func (o *Service) GetRefreshTokenDuration() int64`

GetRefreshTokenDuration returns the RefreshTokenDuration field if non-nil, zero value otherwise.

### GetRefreshTokenDurationOk

`func (o *Service) GetRefreshTokenDurationOk() (*int64, bool)`

GetRefreshTokenDurationOk returns a tuple with the RefreshTokenDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenDuration

`func (o *Service) SetRefreshTokenDuration(v int64)`

SetRefreshTokenDuration sets RefreshTokenDuration field to given value.

### HasRefreshTokenDuration

`func (o *Service) HasRefreshTokenDuration() bool`

HasRefreshTokenDuration returns a boolean if a field has been set.

### GetIdTokenDuration

`func (o *Service) GetIdTokenDuration() int64`

GetIdTokenDuration returns the IdTokenDuration field if non-nil, zero value otherwise.

### GetIdTokenDurationOk

`func (o *Service) GetIdTokenDurationOk() (*int64, bool)`

GetIdTokenDurationOk returns a tuple with the IdTokenDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTokenDuration

`func (o *Service) SetIdTokenDuration(v int64)`

SetIdTokenDuration sets IdTokenDuration field to given value.

### HasIdTokenDuration

`func (o *Service) HasIdTokenDuration() bool`

HasIdTokenDuration returns a boolean if a field has been set.

### GetAuthorizationResponseDuration

`func (o *Service) GetAuthorizationResponseDuration() int64`

GetAuthorizationResponseDuration returns the AuthorizationResponseDuration field if non-nil, zero value otherwise.

### GetAuthorizationResponseDurationOk

`func (o *Service) GetAuthorizationResponseDurationOk() (*int64, bool)`

GetAuthorizationResponseDurationOk returns a tuple with the AuthorizationResponseDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationResponseDuration

`func (o *Service) SetAuthorizationResponseDuration(v int64)`

SetAuthorizationResponseDuration sets AuthorizationResponseDuration field to given value.

### HasAuthorizationResponseDuration

`func (o *Service) HasAuthorizationResponseDuration() bool`

HasAuthorizationResponseDuration returns a boolean if a field has been set.

### GetPushedAuthReqDuration

`func (o *Service) GetPushedAuthReqDuration() int64`

GetPushedAuthReqDuration returns the PushedAuthReqDuration field if non-nil, zero value otherwise.

### GetPushedAuthReqDurationOk

`func (o *Service) GetPushedAuthReqDurationOk() (*int64, bool)`

GetPushedAuthReqDurationOk returns a tuple with the PushedAuthReqDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushedAuthReqDuration

`func (o *Service) SetPushedAuthReqDuration(v int64)`

SetPushedAuthReqDuration sets PushedAuthReqDuration field to given value.

### HasPushedAuthReqDuration

`func (o *Service) HasPushedAuthReqDuration() bool`

HasPushedAuthReqDuration returns a boolean if a field has been set.

### GetMetadata

`func (o *Service) GetMetadata() []Pair`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Service) GetMetadataOk() (*[]Pair, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Service) SetMetadata(v []Pair)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Service) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetAccessTokenSignatureKeyId

`func (o *Service) GetAccessTokenSignatureKeyId() string`

GetAccessTokenSignatureKeyId returns the AccessTokenSignatureKeyId field if non-nil, zero value otherwise.

### GetAccessTokenSignatureKeyIdOk

`func (o *Service) GetAccessTokenSignatureKeyIdOk() (*string, bool)`

GetAccessTokenSignatureKeyIdOk returns a tuple with the AccessTokenSignatureKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenSignatureKeyId

`func (o *Service) SetAccessTokenSignatureKeyId(v string)`

SetAccessTokenSignatureKeyId sets AccessTokenSignatureKeyId field to given value.

### HasAccessTokenSignatureKeyId

`func (o *Service) HasAccessTokenSignatureKeyId() bool`

HasAccessTokenSignatureKeyId returns a boolean if a field has been set.

### GetAuthorizationSignatureKeyId

`func (o *Service) GetAuthorizationSignatureKeyId() string`

GetAuthorizationSignatureKeyId returns the AuthorizationSignatureKeyId field if non-nil, zero value otherwise.

### GetAuthorizationSignatureKeyIdOk

`func (o *Service) GetAuthorizationSignatureKeyIdOk() (*string, bool)`

GetAuthorizationSignatureKeyIdOk returns a tuple with the AuthorizationSignatureKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationSignatureKeyId

`func (o *Service) SetAuthorizationSignatureKeyId(v string)`

SetAuthorizationSignatureKeyId sets AuthorizationSignatureKeyId field to given value.

### HasAuthorizationSignatureKeyId

`func (o *Service) HasAuthorizationSignatureKeyId() bool`

HasAuthorizationSignatureKeyId returns a boolean if a field has been set.

### GetIdTokenSignatureKeyId

`func (o *Service) GetIdTokenSignatureKeyId() string`

GetIdTokenSignatureKeyId returns the IdTokenSignatureKeyId field if non-nil, zero value otherwise.

### GetIdTokenSignatureKeyIdOk

`func (o *Service) GetIdTokenSignatureKeyIdOk() (*string, bool)`

GetIdTokenSignatureKeyIdOk returns a tuple with the IdTokenSignatureKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTokenSignatureKeyId

`func (o *Service) SetIdTokenSignatureKeyId(v string)`

SetIdTokenSignatureKeyId sets IdTokenSignatureKeyId field to given value.

### HasIdTokenSignatureKeyId

`func (o *Service) HasIdTokenSignatureKeyId() bool`

HasIdTokenSignatureKeyId returns a boolean if a field has been set.

### GetUserInfoSignatureKeyId

`func (o *Service) GetUserInfoSignatureKeyId() string`

GetUserInfoSignatureKeyId returns the UserInfoSignatureKeyId field if non-nil, zero value otherwise.

### GetUserInfoSignatureKeyIdOk

`func (o *Service) GetUserInfoSignatureKeyIdOk() (*string, bool)`

GetUserInfoSignatureKeyIdOk returns a tuple with the UserInfoSignatureKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInfoSignatureKeyId

`func (o *Service) SetUserInfoSignatureKeyId(v string)`

SetUserInfoSignatureKeyId sets UserInfoSignatureKeyId field to given value.

### HasUserInfoSignatureKeyId

`func (o *Service) HasUserInfoSignatureKeyId() bool`

HasUserInfoSignatureKeyId returns a boolean if a field has been set.

### GetSupportedBackchannelTokenDeliveryModes

`func (o *Service) GetSupportedBackchannelTokenDeliveryModes() []string`

GetSupportedBackchannelTokenDeliveryModes returns the SupportedBackchannelTokenDeliveryModes field if non-nil, zero value otherwise.

### GetSupportedBackchannelTokenDeliveryModesOk

`func (o *Service) GetSupportedBackchannelTokenDeliveryModesOk() (*[]string, bool)`

GetSupportedBackchannelTokenDeliveryModesOk returns a tuple with the SupportedBackchannelTokenDeliveryModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedBackchannelTokenDeliveryModes

`func (o *Service) SetSupportedBackchannelTokenDeliveryModes(v []string)`

SetSupportedBackchannelTokenDeliveryModes sets SupportedBackchannelTokenDeliveryModes field to given value.

### HasSupportedBackchannelTokenDeliveryModes

`func (o *Service) HasSupportedBackchannelTokenDeliveryModes() bool`

HasSupportedBackchannelTokenDeliveryModes returns a boolean if a field has been set.

### GetBackchannelAuthenticationEndpoint

`func (o *Service) GetBackchannelAuthenticationEndpoint() string`

GetBackchannelAuthenticationEndpoint returns the BackchannelAuthenticationEndpoint field if non-nil, zero value otherwise.

### GetBackchannelAuthenticationEndpointOk

`func (o *Service) GetBackchannelAuthenticationEndpointOk() (*string, bool)`

GetBackchannelAuthenticationEndpointOk returns a tuple with the BackchannelAuthenticationEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackchannelAuthenticationEndpoint

`func (o *Service) SetBackchannelAuthenticationEndpoint(v string)`

SetBackchannelAuthenticationEndpoint sets BackchannelAuthenticationEndpoint field to given value.

### HasBackchannelAuthenticationEndpoint

`func (o *Service) HasBackchannelAuthenticationEndpoint() bool`

HasBackchannelAuthenticationEndpoint returns a boolean if a field has been set.

### GetBackchannelUserCodeParameterSupported

`func (o *Service) GetBackchannelUserCodeParameterSupported() bool`

GetBackchannelUserCodeParameterSupported returns the BackchannelUserCodeParameterSupported field if non-nil, zero value otherwise.

### GetBackchannelUserCodeParameterSupportedOk

`func (o *Service) GetBackchannelUserCodeParameterSupportedOk() (*bool, bool)`

GetBackchannelUserCodeParameterSupportedOk returns a tuple with the BackchannelUserCodeParameterSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackchannelUserCodeParameterSupported

`func (o *Service) SetBackchannelUserCodeParameterSupported(v bool)`

SetBackchannelUserCodeParameterSupported sets BackchannelUserCodeParameterSupported field to given value.

### HasBackchannelUserCodeParameterSupported

`func (o *Service) HasBackchannelUserCodeParameterSupported() bool`

HasBackchannelUserCodeParameterSupported returns a boolean if a field has been set.

### GetBackchannelAuthReqIdDuration

`func (o *Service) GetBackchannelAuthReqIdDuration() int32`

GetBackchannelAuthReqIdDuration returns the BackchannelAuthReqIdDuration field if non-nil, zero value otherwise.

### GetBackchannelAuthReqIdDurationOk

`func (o *Service) GetBackchannelAuthReqIdDurationOk() (*int32, bool)`

GetBackchannelAuthReqIdDurationOk returns a tuple with the BackchannelAuthReqIdDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackchannelAuthReqIdDuration

`func (o *Service) SetBackchannelAuthReqIdDuration(v int32)`

SetBackchannelAuthReqIdDuration sets BackchannelAuthReqIdDuration field to given value.

### HasBackchannelAuthReqIdDuration

`func (o *Service) HasBackchannelAuthReqIdDuration() bool`

HasBackchannelAuthReqIdDuration returns a boolean if a field has been set.

### GetBackchannelPollingInterval

`func (o *Service) GetBackchannelPollingInterval() int32`

GetBackchannelPollingInterval returns the BackchannelPollingInterval field if non-nil, zero value otherwise.

### GetBackchannelPollingIntervalOk

`func (o *Service) GetBackchannelPollingIntervalOk() (*int32, bool)`

GetBackchannelPollingIntervalOk returns a tuple with the BackchannelPollingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackchannelPollingInterval

`func (o *Service) SetBackchannelPollingInterval(v int32)`

SetBackchannelPollingInterval sets BackchannelPollingInterval field to given value.

### HasBackchannelPollingInterval

`func (o *Service) HasBackchannelPollingInterval() bool`

HasBackchannelPollingInterval returns a boolean if a field has been set.

### GetBackchannelBindingMessageRequiredInFapi

`func (o *Service) GetBackchannelBindingMessageRequiredInFapi() bool`

GetBackchannelBindingMessageRequiredInFapi returns the BackchannelBindingMessageRequiredInFapi field if non-nil, zero value otherwise.

### GetBackchannelBindingMessageRequiredInFapiOk

`func (o *Service) GetBackchannelBindingMessageRequiredInFapiOk() (*bool, bool)`

GetBackchannelBindingMessageRequiredInFapiOk returns a tuple with the BackchannelBindingMessageRequiredInFapi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackchannelBindingMessageRequiredInFapi

`func (o *Service) SetBackchannelBindingMessageRequiredInFapi(v bool)`

SetBackchannelBindingMessageRequiredInFapi sets BackchannelBindingMessageRequiredInFapi field to given value.

### HasBackchannelBindingMessageRequiredInFapi

`func (o *Service) HasBackchannelBindingMessageRequiredInFapi() bool`

HasBackchannelBindingMessageRequiredInFapi returns a boolean if a field has been set.

### GetAllowableClockSkew

`func (o *Service) GetAllowableClockSkew() int32`

GetAllowableClockSkew returns the AllowableClockSkew field if non-nil, zero value otherwise.

### GetAllowableClockSkewOk

`func (o *Service) GetAllowableClockSkewOk() (*int32, bool)`

GetAllowableClockSkewOk returns a tuple with the AllowableClockSkew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowableClockSkew

`func (o *Service) SetAllowableClockSkew(v int32)`

SetAllowableClockSkew sets AllowableClockSkew field to given value.

### HasAllowableClockSkew

`func (o *Service) HasAllowableClockSkew() bool`

HasAllowableClockSkew returns a boolean if a field has been set.

### GetDeviceAuthorizationEndpoint

`func (o *Service) GetDeviceAuthorizationEndpoint() string`

GetDeviceAuthorizationEndpoint returns the DeviceAuthorizationEndpoint field if non-nil, zero value otherwise.

### GetDeviceAuthorizationEndpointOk

`func (o *Service) GetDeviceAuthorizationEndpointOk() (*string, bool)`

GetDeviceAuthorizationEndpointOk returns a tuple with the DeviceAuthorizationEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceAuthorizationEndpoint

`func (o *Service) SetDeviceAuthorizationEndpoint(v string)`

SetDeviceAuthorizationEndpoint sets DeviceAuthorizationEndpoint field to given value.

### HasDeviceAuthorizationEndpoint

`func (o *Service) HasDeviceAuthorizationEndpoint() bool`

HasDeviceAuthorizationEndpoint returns a boolean if a field has been set.

### GetDeviceVerificationUri

`func (o *Service) GetDeviceVerificationUri() string`

GetDeviceVerificationUri returns the DeviceVerificationUri field if non-nil, zero value otherwise.

### GetDeviceVerificationUriOk

`func (o *Service) GetDeviceVerificationUriOk() (*string, bool)`

GetDeviceVerificationUriOk returns a tuple with the DeviceVerificationUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceVerificationUri

`func (o *Service) SetDeviceVerificationUri(v string)`

SetDeviceVerificationUri sets DeviceVerificationUri field to given value.

### HasDeviceVerificationUri

`func (o *Service) HasDeviceVerificationUri() bool`

HasDeviceVerificationUri returns a boolean if a field has been set.

### GetDeviceVerificationUriComplete

`func (o *Service) GetDeviceVerificationUriComplete() string`

GetDeviceVerificationUriComplete returns the DeviceVerificationUriComplete field if non-nil, zero value otherwise.

### GetDeviceVerificationUriCompleteOk

`func (o *Service) GetDeviceVerificationUriCompleteOk() (*string, bool)`

GetDeviceVerificationUriCompleteOk returns a tuple with the DeviceVerificationUriComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceVerificationUriComplete

`func (o *Service) SetDeviceVerificationUriComplete(v string)`

SetDeviceVerificationUriComplete sets DeviceVerificationUriComplete field to given value.

### HasDeviceVerificationUriComplete

`func (o *Service) HasDeviceVerificationUriComplete() bool`

HasDeviceVerificationUriComplete returns a boolean if a field has been set.

### GetDeviceFlowCodeDuration

`func (o *Service) GetDeviceFlowCodeDuration() int32`

GetDeviceFlowCodeDuration returns the DeviceFlowCodeDuration field if non-nil, zero value otherwise.

### GetDeviceFlowCodeDurationOk

`func (o *Service) GetDeviceFlowCodeDurationOk() (*int32, bool)`

GetDeviceFlowCodeDurationOk returns a tuple with the DeviceFlowCodeDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceFlowCodeDuration

`func (o *Service) SetDeviceFlowCodeDuration(v int32)`

SetDeviceFlowCodeDuration sets DeviceFlowCodeDuration field to given value.

### HasDeviceFlowCodeDuration

`func (o *Service) HasDeviceFlowCodeDuration() bool`

HasDeviceFlowCodeDuration returns a boolean if a field has been set.

### GetDeviceFlowPollingInterval

`func (o *Service) GetDeviceFlowPollingInterval() int32`

GetDeviceFlowPollingInterval returns the DeviceFlowPollingInterval field if non-nil, zero value otherwise.

### GetDeviceFlowPollingIntervalOk

`func (o *Service) GetDeviceFlowPollingIntervalOk() (*int32, bool)`

GetDeviceFlowPollingIntervalOk returns a tuple with the DeviceFlowPollingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceFlowPollingInterval

`func (o *Service) SetDeviceFlowPollingInterval(v int32)`

SetDeviceFlowPollingInterval sets DeviceFlowPollingInterval field to given value.

### HasDeviceFlowPollingInterval

`func (o *Service) HasDeviceFlowPollingInterval() bool`

HasDeviceFlowPollingInterval returns a boolean if a field has been set.

### GetUserCodeCharset

`func (o *Service) GetUserCodeCharset() string`

GetUserCodeCharset returns the UserCodeCharset field if non-nil, zero value otherwise.

### GetUserCodeCharsetOk

`func (o *Service) GetUserCodeCharsetOk() (*string, bool)`

GetUserCodeCharsetOk returns a tuple with the UserCodeCharset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCodeCharset

`func (o *Service) SetUserCodeCharset(v string)`

SetUserCodeCharset sets UserCodeCharset field to given value.

### HasUserCodeCharset

`func (o *Service) HasUserCodeCharset() bool`

HasUserCodeCharset returns a boolean if a field has been set.

### GetUserCodeLength

`func (o *Service) GetUserCodeLength() int32`

GetUserCodeLength returns the UserCodeLength field if non-nil, zero value otherwise.

### GetUserCodeLengthOk

`func (o *Service) GetUserCodeLengthOk() (*int32, bool)`

GetUserCodeLengthOk returns a tuple with the UserCodeLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCodeLength

`func (o *Service) SetUserCodeLength(v int32)`

SetUserCodeLength sets UserCodeLength field to given value.

### HasUserCodeLength

`func (o *Service) HasUserCodeLength() bool`

HasUserCodeLength returns a boolean if a field has been set.

### GetPushedAuthReqEndpoint

`func (o *Service) GetPushedAuthReqEndpoint() string`

GetPushedAuthReqEndpoint returns the PushedAuthReqEndpoint field if non-nil, zero value otherwise.

### GetPushedAuthReqEndpointOk

`func (o *Service) GetPushedAuthReqEndpointOk() (*string, bool)`

GetPushedAuthReqEndpointOk returns a tuple with the PushedAuthReqEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushedAuthReqEndpoint

`func (o *Service) SetPushedAuthReqEndpoint(v string)`

SetPushedAuthReqEndpoint sets PushedAuthReqEndpoint field to given value.

### HasPushedAuthReqEndpoint

`func (o *Service) HasPushedAuthReqEndpoint() bool`

HasPushedAuthReqEndpoint returns a boolean if a field has been set.

### GetMtlsEndpointAliases

`func (o *Service) GetMtlsEndpointAliases() []NamedUri`

GetMtlsEndpointAliases returns the MtlsEndpointAliases field if non-nil, zero value otherwise.

### GetMtlsEndpointAliasesOk

`func (o *Service) GetMtlsEndpointAliasesOk() (*[]NamedUri, bool)`

GetMtlsEndpointAliasesOk returns a tuple with the MtlsEndpointAliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtlsEndpointAliases

`func (o *Service) SetMtlsEndpointAliases(v []NamedUri)`

SetMtlsEndpointAliases sets MtlsEndpointAliases field to given value.

### HasMtlsEndpointAliases

`func (o *Service) HasMtlsEndpointAliases() bool`

HasMtlsEndpointAliases returns a boolean if a field has been set.

### GetSupportedAuthorizationDetailsTypes

`func (o *Service) GetSupportedAuthorizationDetailsTypes() []string`

GetSupportedAuthorizationDetailsTypes returns the SupportedAuthorizationDetailsTypes field if non-nil, zero value otherwise.

### GetSupportedAuthorizationDetailsTypesOk

`func (o *Service) GetSupportedAuthorizationDetailsTypesOk() (*[]string, bool)`

GetSupportedAuthorizationDetailsTypesOk returns a tuple with the SupportedAuthorizationDetailsTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedAuthorizationDetailsTypes

`func (o *Service) SetSupportedAuthorizationDetailsTypes(v []string)`

SetSupportedAuthorizationDetailsTypes sets SupportedAuthorizationDetailsTypes field to given value.

### HasSupportedAuthorizationDetailsTypes

`func (o *Service) HasSupportedAuthorizationDetailsTypes() bool`

HasSupportedAuthorizationDetailsTypes returns a boolean if a field has been set.

### GetSupportedTrustFrameworks

`func (o *Service) GetSupportedTrustFrameworks() []string`

GetSupportedTrustFrameworks returns the SupportedTrustFrameworks field if non-nil, zero value otherwise.

### GetSupportedTrustFrameworksOk

`func (o *Service) GetSupportedTrustFrameworksOk() (*[]string, bool)`

GetSupportedTrustFrameworksOk returns a tuple with the SupportedTrustFrameworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedTrustFrameworks

`func (o *Service) SetSupportedTrustFrameworks(v []string)`

SetSupportedTrustFrameworks sets SupportedTrustFrameworks field to given value.

### HasSupportedTrustFrameworks

`func (o *Service) HasSupportedTrustFrameworks() bool`

HasSupportedTrustFrameworks returns a boolean if a field has been set.

### GetSupportedEvidence

`func (o *Service) GetSupportedEvidence() []string`

GetSupportedEvidence returns the SupportedEvidence field if non-nil, zero value otherwise.

### GetSupportedEvidenceOk

`func (o *Service) GetSupportedEvidenceOk() (*[]string, bool)`

GetSupportedEvidenceOk returns a tuple with the SupportedEvidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedEvidence

`func (o *Service) SetSupportedEvidence(v []string)`

SetSupportedEvidence sets SupportedEvidence field to given value.

### HasSupportedEvidence

`func (o *Service) HasSupportedEvidence() bool`

HasSupportedEvidence returns a boolean if a field has been set.

### GetSupportedIdentityDocuments

`func (o *Service) GetSupportedIdentityDocuments() []string`

GetSupportedIdentityDocuments returns the SupportedIdentityDocuments field if non-nil, zero value otherwise.

### GetSupportedIdentityDocumentsOk

`func (o *Service) GetSupportedIdentityDocumentsOk() (*[]string, bool)`

GetSupportedIdentityDocumentsOk returns a tuple with the SupportedIdentityDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedIdentityDocuments

`func (o *Service) SetSupportedIdentityDocuments(v []string)`

SetSupportedIdentityDocuments sets SupportedIdentityDocuments field to given value.

### HasSupportedIdentityDocuments

`func (o *Service) HasSupportedIdentityDocuments() bool`

HasSupportedIdentityDocuments returns a boolean if a field has been set.

### GetSupportedDocuments

`func (o *Service) GetSupportedDocuments() []string`

GetSupportedDocuments returns the SupportedDocuments field if non-nil, zero value otherwise.

### GetSupportedDocumentsOk

`func (o *Service) GetSupportedDocumentsOk() (*[]string, bool)`

GetSupportedDocumentsOk returns a tuple with the SupportedDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedDocuments

`func (o *Service) SetSupportedDocuments(v []string)`

SetSupportedDocuments sets SupportedDocuments field to given value.

### HasSupportedDocuments

`func (o *Service) HasSupportedDocuments() bool`

HasSupportedDocuments returns a boolean if a field has been set.

### GetSupportedVerificationMethods

`func (o *Service) GetSupportedVerificationMethods() []string`

GetSupportedVerificationMethods returns the SupportedVerificationMethods field if non-nil, zero value otherwise.

### GetSupportedVerificationMethodsOk

`func (o *Service) GetSupportedVerificationMethodsOk() (*[]string, bool)`

GetSupportedVerificationMethodsOk returns a tuple with the SupportedVerificationMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedVerificationMethods

`func (o *Service) SetSupportedVerificationMethods(v []string)`

SetSupportedVerificationMethods sets SupportedVerificationMethods field to given value.

### HasSupportedVerificationMethods

`func (o *Service) HasSupportedVerificationMethods() bool`

HasSupportedVerificationMethods returns a boolean if a field has been set.

### GetSupportedDocumentsMethods

`func (o *Service) GetSupportedDocumentsMethods() []string`

GetSupportedDocumentsMethods returns the SupportedDocumentsMethods field if non-nil, zero value otherwise.

### GetSupportedDocumentsMethodsOk

`func (o *Service) GetSupportedDocumentsMethodsOk() (*[]string, bool)`

GetSupportedDocumentsMethodsOk returns a tuple with the SupportedDocumentsMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedDocumentsMethods

`func (o *Service) SetSupportedDocumentsMethods(v []string)`

SetSupportedDocumentsMethods sets SupportedDocumentsMethods field to given value.

### HasSupportedDocumentsMethods

`func (o *Service) HasSupportedDocumentsMethods() bool`

HasSupportedDocumentsMethods returns a boolean if a field has been set.

### GetSupportedDocumentsValidationMethods

`func (o *Service) GetSupportedDocumentsValidationMethods() []string`

GetSupportedDocumentsValidationMethods returns the SupportedDocumentsValidationMethods field if non-nil, zero value otherwise.

### GetSupportedDocumentsValidationMethodsOk

`func (o *Service) GetSupportedDocumentsValidationMethodsOk() (*[]string, bool)`

GetSupportedDocumentsValidationMethodsOk returns a tuple with the SupportedDocumentsValidationMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedDocumentsValidationMethods

`func (o *Service) SetSupportedDocumentsValidationMethods(v []string)`

SetSupportedDocumentsValidationMethods sets SupportedDocumentsValidationMethods field to given value.

### HasSupportedDocumentsValidationMethods

`func (o *Service) HasSupportedDocumentsValidationMethods() bool`

HasSupportedDocumentsValidationMethods returns a boolean if a field has been set.

### GetSupportedDocumentsVerificationMethods

`func (o *Service) GetSupportedDocumentsVerificationMethods() []string`

GetSupportedDocumentsVerificationMethods returns the SupportedDocumentsVerificationMethods field if non-nil, zero value otherwise.

### GetSupportedDocumentsVerificationMethodsOk

`func (o *Service) GetSupportedDocumentsVerificationMethodsOk() (*[]string, bool)`

GetSupportedDocumentsVerificationMethodsOk returns a tuple with the SupportedDocumentsVerificationMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedDocumentsVerificationMethods

`func (o *Service) SetSupportedDocumentsVerificationMethods(v []string)`

SetSupportedDocumentsVerificationMethods sets SupportedDocumentsVerificationMethods field to given value.

### HasSupportedDocumentsVerificationMethods

`func (o *Service) HasSupportedDocumentsVerificationMethods() bool`

HasSupportedDocumentsVerificationMethods returns a boolean if a field has been set.

### GetSupportedDocumentsCheckMethods

`func (o *Service) GetSupportedDocumentsCheckMethods() []string`

GetSupportedDocumentsCheckMethods returns the SupportedDocumentsCheckMethods field if non-nil, zero value otherwise.

### GetSupportedDocumentsCheckMethodsOk

`func (o *Service) GetSupportedDocumentsCheckMethodsOk() (*[]string, bool)`

GetSupportedDocumentsCheckMethodsOk returns a tuple with the SupportedDocumentsCheckMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedDocumentsCheckMethods

`func (o *Service) SetSupportedDocumentsCheckMethods(v []string)`

SetSupportedDocumentsCheckMethods sets SupportedDocumentsCheckMethods field to given value.

### HasSupportedDocumentsCheckMethods

`func (o *Service) HasSupportedDocumentsCheckMethods() bool`

HasSupportedDocumentsCheckMethods returns a boolean if a field has been set.

### GetSupportedElectronicRecords

`func (o *Service) GetSupportedElectronicRecords() []string`

GetSupportedElectronicRecords returns the SupportedElectronicRecords field if non-nil, zero value otherwise.

### GetSupportedElectronicRecordsOk

`func (o *Service) GetSupportedElectronicRecordsOk() (*[]string, bool)`

GetSupportedElectronicRecordsOk returns a tuple with the SupportedElectronicRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedElectronicRecords

`func (o *Service) SetSupportedElectronicRecords(v []string)`

SetSupportedElectronicRecords sets SupportedElectronicRecords field to given value.

### HasSupportedElectronicRecords

`func (o *Service) HasSupportedElectronicRecords() bool`

HasSupportedElectronicRecords returns a boolean if a field has been set.

### GetSupportedVerifiedClaims

`func (o *Service) GetSupportedVerifiedClaims() []string`

GetSupportedVerifiedClaims returns the SupportedVerifiedClaims field if non-nil, zero value otherwise.

### GetSupportedVerifiedClaimsOk

`func (o *Service) GetSupportedVerifiedClaimsOk() (*[]string, bool)`

GetSupportedVerifiedClaimsOk returns a tuple with the SupportedVerifiedClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedVerifiedClaims

`func (o *Service) SetSupportedVerifiedClaims(v []string)`

SetSupportedVerifiedClaims sets SupportedVerifiedClaims field to given value.

### HasSupportedVerifiedClaims

`func (o *Service) HasSupportedVerifiedClaims() bool`

HasSupportedVerifiedClaims returns a boolean if a field has been set.

### GetSupportedAttachments

`func (o *Service) GetSupportedAttachments() []string`

GetSupportedAttachments returns the SupportedAttachments field if non-nil, zero value otherwise.

### GetSupportedAttachmentsOk

`func (o *Service) GetSupportedAttachmentsOk() (*[]string, bool)`

GetSupportedAttachmentsOk returns a tuple with the SupportedAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedAttachments

`func (o *Service) SetSupportedAttachments(v []string)`

SetSupportedAttachments sets SupportedAttachments field to given value.

### HasSupportedAttachments

`func (o *Service) HasSupportedAttachments() bool`

HasSupportedAttachments returns a boolean if a field has been set.

### GetSupportedDigestAlgorithms

`func (o *Service) GetSupportedDigestAlgorithms() []string`

GetSupportedDigestAlgorithms returns the SupportedDigestAlgorithms field if non-nil, zero value otherwise.

### GetSupportedDigestAlgorithmsOk

`func (o *Service) GetSupportedDigestAlgorithmsOk() (*[]string, bool)`

GetSupportedDigestAlgorithmsOk returns a tuple with the SupportedDigestAlgorithms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedDigestAlgorithms

`func (o *Service) SetSupportedDigestAlgorithms(v []string)`

SetSupportedDigestAlgorithms sets SupportedDigestAlgorithms field to given value.

### HasSupportedDigestAlgorithms

`func (o *Service) HasSupportedDigestAlgorithms() bool`

HasSupportedDigestAlgorithms returns a boolean if a field has been set.

### GetMissingClientIdAllowed

`func (o *Service) GetMissingClientIdAllowed() bool`

GetMissingClientIdAllowed returns the MissingClientIdAllowed field if non-nil, zero value otherwise.

### GetMissingClientIdAllowedOk

`func (o *Service) GetMissingClientIdAllowedOk() (*bool, bool)`

GetMissingClientIdAllowedOk returns a tuple with the MissingClientIdAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingClientIdAllowed

`func (o *Service) SetMissingClientIdAllowed(v bool)`

SetMissingClientIdAllowed sets MissingClientIdAllowed field to given value.

### HasMissingClientIdAllowed

`func (o *Service) HasMissingClientIdAllowed() bool`

HasMissingClientIdAllowed returns a boolean if a field has been set.

### GetParRequired

`func (o *Service) GetParRequired() bool`

GetParRequired returns the ParRequired field if non-nil, zero value otherwise.

### GetParRequiredOk

`func (o *Service) GetParRequiredOk() (*bool, bool)`

GetParRequiredOk returns a tuple with the ParRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParRequired

`func (o *Service) SetParRequired(v bool)`

SetParRequired sets ParRequired field to given value.

### HasParRequired

`func (o *Service) HasParRequired() bool`

HasParRequired returns a boolean if a field has been set.

### GetRequestObjectRequired

`func (o *Service) GetRequestObjectRequired() bool`

GetRequestObjectRequired returns the RequestObjectRequired field if non-nil, zero value otherwise.

### GetRequestObjectRequiredOk

`func (o *Service) GetRequestObjectRequiredOk() (*bool, bool)`

GetRequestObjectRequiredOk returns a tuple with the RequestObjectRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestObjectRequired

`func (o *Service) SetRequestObjectRequired(v bool)`

SetRequestObjectRequired sets RequestObjectRequired field to given value.

### HasRequestObjectRequired

`func (o *Service) HasRequestObjectRequired() bool`

HasRequestObjectRequired returns a boolean if a field has been set.

### GetTraditionalRequestObjectProcessingApplied

`func (o *Service) GetTraditionalRequestObjectProcessingApplied() bool`

GetTraditionalRequestObjectProcessingApplied returns the TraditionalRequestObjectProcessingApplied field if non-nil, zero value otherwise.

### GetTraditionalRequestObjectProcessingAppliedOk

`func (o *Service) GetTraditionalRequestObjectProcessingAppliedOk() (*bool, bool)`

GetTraditionalRequestObjectProcessingAppliedOk returns a tuple with the TraditionalRequestObjectProcessingApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraditionalRequestObjectProcessingApplied

`func (o *Service) SetTraditionalRequestObjectProcessingApplied(v bool)`

SetTraditionalRequestObjectProcessingApplied sets TraditionalRequestObjectProcessingApplied field to given value.

### HasTraditionalRequestObjectProcessingApplied

`func (o *Service) HasTraditionalRequestObjectProcessingApplied() bool`

HasTraditionalRequestObjectProcessingApplied returns a boolean if a field has been set.

### GetClaimShortcutRestrictive

`func (o *Service) GetClaimShortcutRestrictive() bool`

GetClaimShortcutRestrictive returns the ClaimShortcutRestrictive field if non-nil, zero value otherwise.

### GetClaimShortcutRestrictiveOk

`func (o *Service) GetClaimShortcutRestrictiveOk() (*bool, bool)`

GetClaimShortcutRestrictiveOk returns a tuple with the ClaimShortcutRestrictive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimShortcutRestrictive

`func (o *Service) SetClaimShortcutRestrictive(v bool)`

SetClaimShortcutRestrictive sets ClaimShortcutRestrictive field to given value.

### HasClaimShortcutRestrictive

`func (o *Service) HasClaimShortcutRestrictive() bool`

HasClaimShortcutRestrictive returns a boolean if a field has been set.

### GetScopeRequired

`func (o *Service) GetScopeRequired() bool`

GetScopeRequired returns the ScopeRequired field if non-nil, zero value otherwise.

### GetScopeRequiredOk

`func (o *Service) GetScopeRequiredOk() (*bool, bool)`

GetScopeRequiredOk returns a tuple with the ScopeRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeRequired

`func (o *Service) SetScopeRequired(v bool)`

SetScopeRequired sets ScopeRequired field to given value.

### HasScopeRequired

`func (o *Service) HasScopeRequired() bool`

HasScopeRequired returns a boolean if a field has been set.

### GetNbfOptional

`func (o *Service) GetNbfOptional() bool`

GetNbfOptional returns the NbfOptional field if non-nil, zero value otherwise.

### GetNbfOptionalOk

`func (o *Service) GetNbfOptionalOk() (*bool, bool)`

GetNbfOptionalOk returns a tuple with the NbfOptional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbfOptional

`func (o *Service) SetNbfOptional(v bool)`

SetNbfOptional sets NbfOptional field to given value.

### HasNbfOptional

`func (o *Service) HasNbfOptional() bool`

HasNbfOptional returns a boolean if a field has been set.

### GetIssSuppressed

`func (o *Service) GetIssSuppressed() bool`

GetIssSuppressed returns the IssSuppressed field if non-nil, zero value otherwise.

### GetIssSuppressedOk

`func (o *Service) GetIssSuppressedOk() (*bool, bool)`

GetIssSuppressedOk returns a tuple with the IssSuppressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssSuppressed

`func (o *Service) SetIssSuppressed(v bool)`

SetIssSuppressed sets IssSuppressed field to given value.

### HasIssSuppressed

`func (o *Service) HasIssSuppressed() bool`

HasIssSuppressed returns a boolean if a field has been set.

### GetAttributes

`func (o *Service) GetAttributes() []Pair`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Service) GetAttributesOk() (*[]Pair, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Service) SetAttributes(v []Pair)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Service) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetSupportedCustomClientMetadata

`func (o *Service) GetSupportedCustomClientMetadata() []string`

GetSupportedCustomClientMetadata returns the SupportedCustomClientMetadata field if non-nil, zero value otherwise.

### GetSupportedCustomClientMetadataOk

`func (o *Service) GetSupportedCustomClientMetadataOk() (*[]string, bool)`

GetSupportedCustomClientMetadataOk returns a tuple with the SupportedCustomClientMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedCustomClientMetadata

`func (o *Service) SetSupportedCustomClientMetadata(v []string)`

SetSupportedCustomClientMetadata sets SupportedCustomClientMetadata field to given value.

### HasSupportedCustomClientMetadata

`func (o *Service) HasSupportedCustomClientMetadata() bool`

HasSupportedCustomClientMetadata returns a boolean if a field has been set.

### GetTokenExpirationLinked

`func (o *Service) GetTokenExpirationLinked() bool`

GetTokenExpirationLinked returns the TokenExpirationLinked field if non-nil, zero value otherwise.

### GetTokenExpirationLinkedOk

`func (o *Service) GetTokenExpirationLinkedOk() (*bool, bool)`

GetTokenExpirationLinkedOk returns a tuple with the TokenExpirationLinked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExpirationLinked

`func (o *Service) SetTokenExpirationLinked(v bool)`

SetTokenExpirationLinked sets TokenExpirationLinked field to given value.

### HasTokenExpirationLinked

`func (o *Service) HasTokenExpirationLinked() bool`

HasTokenExpirationLinked returns a boolean if a field has been set.

### GetFrontChannelRequestObjectEncryptionRequired

`func (o *Service) GetFrontChannelRequestObjectEncryptionRequired() bool`

GetFrontChannelRequestObjectEncryptionRequired returns the FrontChannelRequestObjectEncryptionRequired field if non-nil, zero value otherwise.

### GetFrontChannelRequestObjectEncryptionRequiredOk

`func (o *Service) GetFrontChannelRequestObjectEncryptionRequiredOk() (*bool, bool)`

GetFrontChannelRequestObjectEncryptionRequiredOk returns a tuple with the FrontChannelRequestObjectEncryptionRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontChannelRequestObjectEncryptionRequired

`func (o *Service) SetFrontChannelRequestObjectEncryptionRequired(v bool)`

SetFrontChannelRequestObjectEncryptionRequired sets FrontChannelRequestObjectEncryptionRequired field to given value.

### HasFrontChannelRequestObjectEncryptionRequired

`func (o *Service) HasFrontChannelRequestObjectEncryptionRequired() bool`

HasFrontChannelRequestObjectEncryptionRequired returns a boolean if a field has been set.

### GetRequestObjectEncryptionAlgMatchRequired

`func (o *Service) GetRequestObjectEncryptionAlgMatchRequired() bool`

GetRequestObjectEncryptionAlgMatchRequired returns the RequestObjectEncryptionAlgMatchRequired field if non-nil, zero value otherwise.

### GetRequestObjectEncryptionAlgMatchRequiredOk

`func (o *Service) GetRequestObjectEncryptionAlgMatchRequiredOk() (*bool, bool)`

GetRequestObjectEncryptionAlgMatchRequiredOk returns a tuple with the RequestObjectEncryptionAlgMatchRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestObjectEncryptionAlgMatchRequired

`func (o *Service) SetRequestObjectEncryptionAlgMatchRequired(v bool)`

SetRequestObjectEncryptionAlgMatchRequired sets RequestObjectEncryptionAlgMatchRequired field to given value.

### HasRequestObjectEncryptionAlgMatchRequired

`func (o *Service) HasRequestObjectEncryptionAlgMatchRequired() bool`

HasRequestObjectEncryptionAlgMatchRequired returns a boolean if a field has been set.

### GetRequestObjectEncryptionEncMatchRequired

`func (o *Service) GetRequestObjectEncryptionEncMatchRequired() bool`

GetRequestObjectEncryptionEncMatchRequired returns the RequestObjectEncryptionEncMatchRequired field if non-nil, zero value otherwise.

### GetRequestObjectEncryptionEncMatchRequiredOk

`func (o *Service) GetRequestObjectEncryptionEncMatchRequiredOk() (*bool, bool)`

GetRequestObjectEncryptionEncMatchRequiredOk returns a tuple with the RequestObjectEncryptionEncMatchRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestObjectEncryptionEncMatchRequired

`func (o *Service) SetRequestObjectEncryptionEncMatchRequired(v bool)`

SetRequestObjectEncryptionEncMatchRequired sets RequestObjectEncryptionEncMatchRequired field to given value.

### HasRequestObjectEncryptionEncMatchRequired

`func (o *Service) HasRequestObjectEncryptionEncMatchRequired() bool`

HasRequestObjectEncryptionEncMatchRequired returns a boolean if a field has been set.

### GetHsmEnabled

`func (o *Service) GetHsmEnabled() bool`

GetHsmEnabled returns the HsmEnabled field if non-nil, zero value otherwise.

### GetHsmEnabledOk

`func (o *Service) GetHsmEnabledOk() (*bool, bool)`

GetHsmEnabledOk returns a tuple with the HsmEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsmEnabled

`func (o *Service) SetHsmEnabled(v bool)`

SetHsmEnabled sets HsmEnabled field to given value.

### HasHsmEnabled

`func (o *Service) HasHsmEnabled() bool`

HasHsmEnabled returns a boolean if a field has been set.

### GetHsks

`func (o *Service) GetHsks() []Hsk`

GetHsks returns the Hsks field if non-nil, zero value otherwise.

### GetHsksOk

`func (o *Service) GetHsksOk() (*[]Hsk, bool)`

GetHsksOk returns a tuple with the Hsks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsks

`func (o *Service) SetHsks(v []Hsk)`

SetHsks sets Hsks field to given value.

### HasHsks

`func (o *Service) HasHsks() bool`

HasHsks returns a boolean if a field has been set.

### GetGrantManagementEndpoint

`func (o *Service) GetGrantManagementEndpoint() string`

GetGrantManagementEndpoint returns the GrantManagementEndpoint field if non-nil, zero value otherwise.

### GetGrantManagementEndpointOk

`func (o *Service) GetGrantManagementEndpointOk() (*string, bool)`

GetGrantManagementEndpointOk returns a tuple with the GrantManagementEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantManagementEndpoint

`func (o *Service) SetGrantManagementEndpoint(v string)`

SetGrantManagementEndpoint sets GrantManagementEndpoint field to given value.

### HasGrantManagementEndpoint

`func (o *Service) HasGrantManagementEndpoint() bool`

HasGrantManagementEndpoint returns a boolean if a field has been set.

### GetGrantManagementActionRequired

`func (o *Service) GetGrantManagementActionRequired() bool`

GetGrantManagementActionRequired returns the GrantManagementActionRequired field if non-nil, zero value otherwise.

### GetGrantManagementActionRequiredOk

`func (o *Service) GetGrantManagementActionRequiredOk() (*bool, bool)`

GetGrantManagementActionRequiredOk returns a tuple with the GrantManagementActionRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantManagementActionRequired

`func (o *Service) SetGrantManagementActionRequired(v bool)`

SetGrantManagementActionRequired sets GrantManagementActionRequired field to given value.

### HasGrantManagementActionRequired

`func (o *Service) HasGrantManagementActionRequired() bool`

HasGrantManagementActionRequired returns a boolean if a field has been set.

### GetUnauthorizedOnClientConfigSupported

`func (o *Service) GetUnauthorizedOnClientConfigSupported() bool`

GetUnauthorizedOnClientConfigSupported returns the UnauthorizedOnClientConfigSupported field if non-nil, zero value otherwise.

### GetUnauthorizedOnClientConfigSupportedOk

`func (o *Service) GetUnauthorizedOnClientConfigSupportedOk() (*bool, bool)`

GetUnauthorizedOnClientConfigSupportedOk returns a tuple with the UnauthorizedOnClientConfigSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnauthorizedOnClientConfigSupported

`func (o *Service) SetUnauthorizedOnClientConfigSupported(v bool)`

SetUnauthorizedOnClientConfigSupported sets UnauthorizedOnClientConfigSupported field to given value.

### HasUnauthorizedOnClientConfigSupported

`func (o *Service) HasUnauthorizedOnClientConfigSupported() bool`

HasUnauthorizedOnClientConfigSupported returns a boolean if a field has been set.

### GetDcrScopeUsedAsRequestable

`func (o *Service) GetDcrScopeUsedAsRequestable() bool`

GetDcrScopeUsedAsRequestable returns the DcrScopeUsedAsRequestable field if non-nil, zero value otherwise.

### GetDcrScopeUsedAsRequestableOk

`func (o *Service) GetDcrScopeUsedAsRequestableOk() (*bool, bool)`

GetDcrScopeUsedAsRequestableOk returns a tuple with the DcrScopeUsedAsRequestable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcrScopeUsedAsRequestable

`func (o *Service) SetDcrScopeUsedAsRequestable(v bool)`

SetDcrScopeUsedAsRequestable sets DcrScopeUsedAsRequestable field to given value.

### HasDcrScopeUsedAsRequestable

`func (o *Service) HasDcrScopeUsedAsRequestable() bool`

HasDcrScopeUsedAsRequestable returns a boolean if a field has been set.

### GetPredefinedTransformedClaims

`func (o *Service) GetPredefinedTransformedClaims() string`

GetPredefinedTransformedClaims returns the PredefinedTransformedClaims field if non-nil, zero value otherwise.

### GetPredefinedTransformedClaimsOk

`func (o *Service) GetPredefinedTransformedClaimsOk() (*string, bool)`

GetPredefinedTransformedClaimsOk returns a tuple with the PredefinedTransformedClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredefinedTransformedClaims

`func (o *Service) SetPredefinedTransformedClaims(v string)`

SetPredefinedTransformedClaims sets PredefinedTransformedClaims field to given value.

### HasPredefinedTransformedClaims

`func (o *Service) HasPredefinedTransformedClaims() bool`

HasPredefinedTransformedClaims returns a boolean if a field has been set.

### GetLoopbackRedirectionUriVariable

`func (o *Service) GetLoopbackRedirectionUriVariable() bool`

GetLoopbackRedirectionUriVariable returns the LoopbackRedirectionUriVariable field if non-nil, zero value otherwise.

### GetLoopbackRedirectionUriVariableOk

`func (o *Service) GetLoopbackRedirectionUriVariableOk() (*bool, bool)`

GetLoopbackRedirectionUriVariableOk returns a tuple with the LoopbackRedirectionUriVariable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopbackRedirectionUriVariable

`func (o *Service) SetLoopbackRedirectionUriVariable(v bool)`

SetLoopbackRedirectionUriVariable sets LoopbackRedirectionUriVariable field to given value.

### HasLoopbackRedirectionUriVariable

`func (o *Service) HasLoopbackRedirectionUriVariable() bool`

HasLoopbackRedirectionUriVariable returns a boolean if a field has been set.

### GetRequestObjectAudienceChecked

`func (o *Service) GetRequestObjectAudienceChecked() bool`

GetRequestObjectAudienceChecked returns the RequestObjectAudienceChecked field if non-nil, zero value otherwise.

### GetRequestObjectAudienceCheckedOk

`func (o *Service) GetRequestObjectAudienceCheckedOk() (*bool, bool)`

GetRequestObjectAudienceCheckedOk returns a tuple with the RequestObjectAudienceChecked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestObjectAudienceChecked

`func (o *Service) SetRequestObjectAudienceChecked(v bool)`

SetRequestObjectAudienceChecked sets RequestObjectAudienceChecked field to given value.

### HasRequestObjectAudienceChecked

`func (o *Service) HasRequestObjectAudienceChecked() bool`

HasRequestObjectAudienceChecked returns a boolean if a field has been set.

### GetAccessTokenForExternalAttachmentEmbedded

`func (o *Service) GetAccessTokenForExternalAttachmentEmbedded() bool`

GetAccessTokenForExternalAttachmentEmbedded returns the AccessTokenForExternalAttachmentEmbedded field if non-nil, zero value otherwise.

### GetAccessTokenForExternalAttachmentEmbeddedOk

`func (o *Service) GetAccessTokenForExternalAttachmentEmbeddedOk() (*bool, bool)`

GetAccessTokenForExternalAttachmentEmbeddedOk returns a tuple with the AccessTokenForExternalAttachmentEmbedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenForExternalAttachmentEmbedded

`func (o *Service) SetAccessTokenForExternalAttachmentEmbedded(v bool)`

SetAccessTokenForExternalAttachmentEmbedded sets AccessTokenForExternalAttachmentEmbedded field to given value.

### HasAccessTokenForExternalAttachmentEmbedded

`func (o *Service) HasAccessTokenForExternalAttachmentEmbedded() bool`

HasAccessTokenForExternalAttachmentEmbedded returns a boolean if a field has been set.

### GetRefreshTokenIdempotent

`func (o *Service) GetRefreshTokenIdempotent() bool`

GetRefreshTokenIdempotent returns the RefreshTokenIdempotent field if non-nil, zero value otherwise.

### GetRefreshTokenIdempotentOk

`func (o *Service) GetRefreshTokenIdempotentOk() (*bool, bool)`

GetRefreshTokenIdempotentOk returns a tuple with the RefreshTokenIdempotent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenIdempotent

`func (o *Service) SetRefreshTokenIdempotent(v bool)`

SetRefreshTokenIdempotent sets RefreshTokenIdempotent field to given value.

### HasRefreshTokenIdempotent

`func (o *Service) HasRefreshTokenIdempotent() bool`

HasRefreshTokenIdempotent returns a boolean if a field has been set.

### GetFederationEnabled

`func (o *Service) GetFederationEnabled() bool`

GetFederationEnabled returns the FederationEnabled field if non-nil, zero value otherwise.

### GetFederationEnabledOk

`func (o *Service) GetFederationEnabledOk() (*bool, bool)`

GetFederationEnabledOk returns a tuple with the FederationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederationEnabled

`func (o *Service) SetFederationEnabled(v bool)`

SetFederationEnabled sets FederationEnabled field to given value.

### HasFederationEnabled

`func (o *Service) HasFederationEnabled() bool`

HasFederationEnabled returns a boolean if a field has been set.

### GetOrganizationName

`func (o *Service) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *Service) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *Service) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *Service) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### GetAuthorityHints

`func (o *Service) GetAuthorityHints() []string`

GetAuthorityHints returns the AuthorityHints field if non-nil, zero value otherwise.

### GetAuthorityHintsOk

`func (o *Service) GetAuthorityHintsOk() (*[]string, bool)`

GetAuthorityHintsOk returns a tuple with the AuthorityHints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorityHints

`func (o *Service) SetAuthorityHints(v []string)`

SetAuthorityHints sets AuthorityHints field to given value.

### HasAuthorityHints

`func (o *Service) HasAuthorityHints() bool`

HasAuthorityHints returns a boolean if a field has been set.

### GetTrustAnchors

`func (o *Service) GetTrustAnchors() []TrustAnchor`

GetTrustAnchors returns the TrustAnchors field if non-nil, zero value otherwise.

### GetTrustAnchorsOk

`func (o *Service) GetTrustAnchorsOk() (*[]TrustAnchor, bool)`

GetTrustAnchorsOk returns a tuple with the TrustAnchors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustAnchors

`func (o *Service) SetTrustAnchors(v []TrustAnchor)`

SetTrustAnchors sets TrustAnchors field to given value.

### HasTrustAnchors

`func (o *Service) HasTrustAnchors() bool`

HasTrustAnchors returns a boolean if a field has been set.

### GetFederationJwks

`func (o *Service) GetFederationJwks() string`

GetFederationJwks returns the FederationJwks field if non-nil, zero value otherwise.

### GetFederationJwksOk

`func (o *Service) GetFederationJwksOk() (*string, bool)`

GetFederationJwksOk returns a tuple with the FederationJwks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederationJwks

`func (o *Service) SetFederationJwks(v string)`

SetFederationJwks sets FederationJwks field to given value.

### HasFederationJwks

`func (o *Service) HasFederationJwks() bool`

HasFederationJwks returns a boolean if a field has been set.

### GetFederationSignatureKeyId

`func (o *Service) GetFederationSignatureKeyId() string`

GetFederationSignatureKeyId returns the FederationSignatureKeyId field if non-nil, zero value otherwise.

### GetFederationSignatureKeyIdOk

`func (o *Service) GetFederationSignatureKeyIdOk() (*string, bool)`

GetFederationSignatureKeyIdOk returns a tuple with the FederationSignatureKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederationSignatureKeyId

`func (o *Service) SetFederationSignatureKeyId(v string)`

SetFederationSignatureKeyId sets FederationSignatureKeyId field to given value.

### HasFederationSignatureKeyId

`func (o *Service) HasFederationSignatureKeyId() bool`

HasFederationSignatureKeyId returns a boolean if a field has been set.

### GetFederationConfigurationDuration

`func (o *Service) GetFederationConfigurationDuration() int32`

GetFederationConfigurationDuration returns the FederationConfigurationDuration field if non-nil, zero value otherwise.

### GetFederationConfigurationDurationOk

`func (o *Service) GetFederationConfigurationDurationOk() (*int32, bool)`

GetFederationConfigurationDurationOk returns a tuple with the FederationConfigurationDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederationConfigurationDuration

`func (o *Service) SetFederationConfigurationDuration(v int32)`

SetFederationConfigurationDuration sets FederationConfigurationDuration field to given value.

### HasFederationConfigurationDuration

`func (o *Service) HasFederationConfigurationDuration() bool`

HasFederationConfigurationDuration returns a boolean if a field has been set.

### GetSignedJwksUri

`func (o *Service) GetSignedJwksUri() string`

GetSignedJwksUri returns the SignedJwksUri field if non-nil, zero value otherwise.

### GetSignedJwksUriOk

`func (o *Service) GetSignedJwksUriOk() (*string, bool)`

GetSignedJwksUriOk returns a tuple with the SignedJwksUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedJwksUri

`func (o *Service) SetSignedJwksUri(v string)`

SetSignedJwksUri sets SignedJwksUri field to given value.

### HasSignedJwksUri

`func (o *Service) HasSignedJwksUri() bool`

HasSignedJwksUri returns a boolean if a field has been set.

### GetFederationRegistrationEndpoint

`func (o *Service) GetFederationRegistrationEndpoint() string`

GetFederationRegistrationEndpoint returns the FederationRegistrationEndpoint field if non-nil, zero value otherwise.

### GetFederationRegistrationEndpointOk

`func (o *Service) GetFederationRegistrationEndpointOk() (*string, bool)`

GetFederationRegistrationEndpointOk returns a tuple with the FederationRegistrationEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederationRegistrationEndpoint

`func (o *Service) SetFederationRegistrationEndpoint(v string)`

SetFederationRegistrationEndpoint sets FederationRegistrationEndpoint field to given value.

### HasFederationRegistrationEndpoint

`func (o *Service) HasFederationRegistrationEndpoint() bool`

HasFederationRegistrationEndpoint returns a boolean if a field has been set.

### GetSupportedClientRegistrationTypes

`func (o *Service) GetSupportedClientRegistrationTypes() []string`

GetSupportedClientRegistrationTypes returns the SupportedClientRegistrationTypes field if non-nil, zero value otherwise.

### GetSupportedClientRegistrationTypesOk

`func (o *Service) GetSupportedClientRegistrationTypesOk() (*[]string, bool)`

GetSupportedClientRegistrationTypesOk returns a tuple with the SupportedClientRegistrationTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedClientRegistrationTypes

`func (o *Service) SetSupportedClientRegistrationTypes(v []string)`

SetSupportedClientRegistrationTypes sets SupportedClientRegistrationTypes field to given value.

### HasSupportedClientRegistrationTypes

`func (o *Service) HasSupportedClientRegistrationTypes() bool`

HasSupportedClientRegistrationTypes returns a boolean if a field has been set.

### GetTokenExchangeByIdentifiableClientsOnly

`func (o *Service) GetTokenExchangeByIdentifiableClientsOnly() bool`

GetTokenExchangeByIdentifiableClientsOnly returns the TokenExchangeByIdentifiableClientsOnly field if non-nil, zero value otherwise.

### GetTokenExchangeByIdentifiableClientsOnlyOk

`func (o *Service) GetTokenExchangeByIdentifiableClientsOnlyOk() (*bool, bool)`

GetTokenExchangeByIdentifiableClientsOnlyOk returns a tuple with the TokenExchangeByIdentifiableClientsOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExchangeByIdentifiableClientsOnly

`func (o *Service) SetTokenExchangeByIdentifiableClientsOnly(v bool)`

SetTokenExchangeByIdentifiableClientsOnly sets TokenExchangeByIdentifiableClientsOnly field to given value.

### HasTokenExchangeByIdentifiableClientsOnly

`func (o *Service) HasTokenExchangeByIdentifiableClientsOnly() bool`

HasTokenExchangeByIdentifiableClientsOnly returns a boolean if a field has been set.

### GetTokenExchangeByConfidentialClientsOnly

`func (o *Service) GetTokenExchangeByConfidentialClientsOnly() bool`

GetTokenExchangeByConfidentialClientsOnly returns the TokenExchangeByConfidentialClientsOnly field if non-nil, zero value otherwise.

### GetTokenExchangeByConfidentialClientsOnlyOk

`func (o *Service) GetTokenExchangeByConfidentialClientsOnlyOk() (*bool, bool)`

GetTokenExchangeByConfidentialClientsOnlyOk returns a tuple with the TokenExchangeByConfidentialClientsOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExchangeByConfidentialClientsOnly

`func (o *Service) SetTokenExchangeByConfidentialClientsOnly(v bool)`

SetTokenExchangeByConfidentialClientsOnly sets TokenExchangeByConfidentialClientsOnly field to given value.

### HasTokenExchangeByConfidentialClientsOnly

`func (o *Service) HasTokenExchangeByConfidentialClientsOnly() bool`

HasTokenExchangeByConfidentialClientsOnly returns a boolean if a field has been set.

### GetTokenExchangeByPermittedClientsOnly

`func (o *Service) GetTokenExchangeByPermittedClientsOnly() bool`

GetTokenExchangeByPermittedClientsOnly returns the TokenExchangeByPermittedClientsOnly field if non-nil, zero value otherwise.

### GetTokenExchangeByPermittedClientsOnlyOk

`func (o *Service) GetTokenExchangeByPermittedClientsOnlyOk() (*bool, bool)`

GetTokenExchangeByPermittedClientsOnlyOk returns a tuple with the TokenExchangeByPermittedClientsOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExchangeByPermittedClientsOnly

`func (o *Service) SetTokenExchangeByPermittedClientsOnly(v bool)`

SetTokenExchangeByPermittedClientsOnly sets TokenExchangeByPermittedClientsOnly field to given value.

### HasTokenExchangeByPermittedClientsOnly

`func (o *Service) HasTokenExchangeByPermittedClientsOnly() bool`

HasTokenExchangeByPermittedClientsOnly returns a boolean if a field has been set.

### GetTokenExchangeEncryptedJwtRejected

`func (o *Service) GetTokenExchangeEncryptedJwtRejected() bool`

GetTokenExchangeEncryptedJwtRejected returns the TokenExchangeEncryptedJwtRejected field if non-nil, zero value otherwise.

### GetTokenExchangeEncryptedJwtRejectedOk

`func (o *Service) GetTokenExchangeEncryptedJwtRejectedOk() (*bool, bool)`

GetTokenExchangeEncryptedJwtRejectedOk returns a tuple with the TokenExchangeEncryptedJwtRejected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExchangeEncryptedJwtRejected

`func (o *Service) SetTokenExchangeEncryptedJwtRejected(v bool)`

SetTokenExchangeEncryptedJwtRejected sets TokenExchangeEncryptedJwtRejected field to given value.

### HasTokenExchangeEncryptedJwtRejected

`func (o *Service) HasTokenExchangeEncryptedJwtRejected() bool`

HasTokenExchangeEncryptedJwtRejected returns a boolean if a field has been set.

### GetTokenExchangeUnsignedJwtRejected

`func (o *Service) GetTokenExchangeUnsignedJwtRejected() bool`

GetTokenExchangeUnsignedJwtRejected returns the TokenExchangeUnsignedJwtRejected field if non-nil, zero value otherwise.

### GetTokenExchangeUnsignedJwtRejectedOk

`func (o *Service) GetTokenExchangeUnsignedJwtRejectedOk() (*bool, bool)`

GetTokenExchangeUnsignedJwtRejectedOk returns a tuple with the TokenExchangeUnsignedJwtRejected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExchangeUnsignedJwtRejected

`func (o *Service) SetTokenExchangeUnsignedJwtRejected(v bool)`

SetTokenExchangeUnsignedJwtRejected sets TokenExchangeUnsignedJwtRejected field to given value.

### HasTokenExchangeUnsignedJwtRejected

`func (o *Service) HasTokenExchangeUnsignedJwtRejected() bool`

HasTokenExchangeUnsignedJwtRejected returns a boolean if a field has been set.

### GetJwtGrantByIdentifiableClientsOnly

`func (o *Service) GetJwtGrantByIdentifiableClientsOnly() bool`

GetJwtGrantByIdentifiableClientsOnly returns the JwtGrantByIdentifiableClientsOnly field if non-nil, zero value otherwise.

### GetJwtGrantByIdentifiableClientsOnlyOk

`func (o *Service) GetJwtGrantByIdentifiableClientsOnlyOk() (*bool, bool)`

GetJwtGrantByIdentifiableClientsOnlyOk returns a tuple with the JwtGrantByIdentifiableClientsOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtGrantByIdentifiableClientsOnly

`func (o *Service) SetJwtGrantByIdentifiableClientsOnly(v bool)`

SetJwtGrantByIdentifiableClientsOnly sets JwtGrantByIdentifiableClientsOnly field to given value.

### HasJwtGrantByIdentifiableClientsOnly

`func (o *Service) HasJwtGrantByIdentifiableClientsOnly() bool`

HasJwtGrantByIdentifiableClientsOnly returns a boolean if a field has been set.

### GetJwtGrantEncryptedJwtRejected

`func (o *Service) GetJwtGrantEncryptedJwtRejected() bool`

GetJwtGrantEncryptedJwtRejected returns the JwtGrantEncryptedJwtRejected field if non-nil, zero value otherwise.

### GetJwtGrantEncryptedJwtRejectedOk

`func (o *Service) GetJwtGrantEncryptedJwtRejectedOk() (*bool, bool)`

GetJwtGrantEncryptedJwtRejectedOk returns a tuple with the JwtGrantEncryptedJwtRejected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtGrantEncryptedJwtRejected

`func (o *Service) SetJwtGrantEncryptedJwtRejected(v bool)`

SetJwtGrantEncryptedJwtRejected sets JwtGrantEncryptedJwtRejected field to given value.

### HasJwtGrantEncryptedJwtRejected

`func (o *Service) HasJwtGrantEncryptedJwtRejected() bool`

HasJwtGrantEncryptedJwtRejected returns a boolean if a field has been set.

### GetJwtGrantUnsignedJwtRejected

`func (o *Service) GetJwtGrantUnsignedJwtRejected() bool`

GetJwtGrantUnsignedJwtRejected returns the JwtGrantUnsignedJwtRejected field if non-nil, zero value otherwise.

### GetJwtGrantUnsignedJwtRejectedOk

`func (o *Service) GetJwtGrantUnsignedJwtRejectedOk() (*bool, bool)`

GetJwtGrantUnsignedJwtRejectedOk returns a tuple with the JwtGrantUnsignedJwtRejected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtGrantUnsignedJwtRejected

`func (o *Service) SetJwtGrantUnsignedJwtRejected(v bool)`

SetJwtGrantUnsignedJwtRejected sets JwtGrantUnsignedJwtRejected field to given value.

### HasJwtGrantUnsignedJwtRejected

`func (o *Service) HasJwtGrantUnsignedJwtRejected() bool`

HasJwtGrantUnsignedJwtRejected returns a boolean if a field has been set.

### GetDcrDuplicateSoftwareIdBlocked

`func (o *Service) GetDcrDuplicateSoftwareIdBlocked() bool`

GetDcrDuplicateSoftwareIdBlocked returns the DcrDuplicateSoftwareIdBlocked field if non-nil, zero value otherwise.

### GetDcrDuplicateSoftwareIdBlockedOk

`func (o *Service) GetDcrDuplicateSoftwareIdBlockedOk() (*bool, bool)`

GetDcrDuplicateSoftwareIdBlockedOk returns a tuple with the DcrDuplicateSoftwareIdBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcrDuplicateSoftwareIdBlocked

`func (o *Service) SetDcrDuplicateSoftwareIdBlocked(v bool)`

SetDcrDuplicateSoftwareIdBlocked sets DcrDuplicateSoftwareIdBlocked field to given value.

### HasDcrDuplicateSoftwareIdBlocked

`func (o *Service) HasDcrDuplicateSoftwareIdBlocked() bool`

HasDcrDuplicateSoftwareIdBlocked returns a boolean if a field has been set.

### GetResourceSignatureKeyId

`func (o *Service) GetResourceSignatureKeyId() string`

GetResourceSignatureKeyId returns the ResourceSignatureKeyId field if non-nil, zero value otherwise.

### GetResourceSignatureKeyIdOk

`func (o *Service) GetResourceSignatureKeyIdOk() (*string, bool)`

GetResourceSignatureKeyIdOk returns a tuple with the ResourceSignatureKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSignatureKeyId

`func (o *Service) SetResourceSignatureKeyId(v string)`

SetResourceSignatureKeyId sets ResourceSignatureKeyId field to given value.

### HasResourceSignatureKeyId

`func (o *Service) HasResourceSignatureKeyId() bool`

HasResourceSignatureKeyId returns a boolean if a field has been set.

### GetRsResponseSigned

`func (o *Service) GetRsResponseSigned() bool`

GetRsResponseSigned returns the RsResponseSigned field if non-nil, zero value otherwise.

### GetRsResponseSignedOk

`func (o *Service) GetRsResponseSignedOk() (*bool, bool)`

GetRsResponseSignedOk returns a tuple with the RsResponseSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRsResponseSigned

`func (o *Service) SetRsResponseSigned(v bool)`

SetRsResponseSigned sets RsResponseSigned field to given value.

### HasRsResponseSigned

`func (o *Service) HasRsResponseSigned() bool`

HasRsResponseSigned returns a boolean if a field has been set.

### GetOpenidDroppedOnRefreshWithoutOfflineAccess

`func (o *Service) GetOpenidDroppedOnRefreshWithoutOfflineAccess() bool`

GetOpenidDroppedOnRefreshWithoutOfflineAccess returns the OpenidDroppedOnRefreshWithoutOfflineAccess field if non-nil, zero value otherwise.

### GetOpenidDroppedOnRefreshWithoutOfflineAccessOk

`func (o *Service) GetOpenidDroppedOnRefreshWithoutOfflineAccessOk() (*bool, bool)`

GetOpenidDroppedOnRefreshWithoutOfflineAccessOk returns a tuple with the OpenidDroppedOnRefreshWithoutOfflineAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenidDroppedOnRefreshWithoutOfflineAccess

`func (o *Service) SetOpenidDroppedOnRefreshWithoutOfflineAccess(v bool)`

SetOpenidDroppedOnRefreshWithoutOfflineAccess sets OpenidDroppedOnRefreshWithoutOfflineAccess field to given value.

### HasOpenidDroppedOnRefreshWithoutOfflineAccess

`func (o *Service) HasOpenidDroppedOnRefreshWithoutOfflineAccess() bool`

HasOpenidDroppedOnRefreshWithoutOfflineAccess returns a boolean if a field has been set.

### GetVerifiableCredentialsEnabled

`func (o *Service) GetVerifiableCredentialsEnabled() bool`

GetVerifiableCredentialsEnabled returns the VerifiableCredentialsEnabled field if non-nil, zero value otherwise.

### GetVerifiableCredentialsEnabledOk

`func (o *Service) GetVerifiableCredentialsEnabledOk() (*bool, bool)`

GetVerifiableCredentialsEnabledOk returns a tuple with the VerifiableCredentialsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiableCredentialsEnabled

`func (o *Service) SetVerifiableCredentialsEnabled(v bool)`

SetVerifiableCredentialsEnabled sets VerifiableCredentialsEnabled field to given value.

### HasVerifiableCredentialsEnabled

`func (o *Service) HasVerifiableCredentialsEnabled() bool`

HasVerifiableCredentialsEnabled returns a boolean if a field has been set.

### GetCredentialIssuerMetadata

`func (o *Service) GetCredentialIssuerMetadata() CredentialIssuerMetadata`

GetCredentialIssuerMetadata returns the CredentialIssuerMetadata field if non-nil, zero value otherwise.

### GetCredentialIssuerMetadataOk

`func (o *Service) GetCredentialIssuerMetadataOk() (*CredentialIssuerMetadata, bool)`

GetCredentialIssuerMetadataOk returns a tuple with the CredentialIssuerMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialIssuerMetadata

`func (o *Service) SetCredentialIssuerMetadata(v CredentialIssuerMetadata)`

SetCredentialIssuerMetadata sets CredentialIssuerMetadata field to given value.

### HasCredentialIssuerMetadata

`func (o *Service) HasCredentialIssuerMetadata() bool`

HasCredentialIssuerMetadata returns a boolean if a field has been set.

### GetCredentialOfferDuration

`func (o *Service) GetCredentialOfferDuration() int64`

GetCredentialOfferDuration returns the CredentialOfferDuration field if non-nil, zero value otherwise.

### GetCredentialOfferDurationOk

`func (o *Service) GetCredentialOfferDurationOk() (*int64, bool)`

GetCredentialOfferDurationOk returns a tuple with the CredentialOfferDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialOfferDuration

`func (o *Service) SetCredentialOfferDuration(v int64)`

SetCredentialOfferDuration sets CredentialOfferDuration field to given value.

### HasCredentialOfferDuration

`func (o *Service) HasCredentialOfferDuration() bool`

HasCredentialOfferDuration returns a boolean if a field has been set.

### GetUserPinLength

`func (o *Service) GetUserPinLength() int32`

GetUserPinLength returns the UserPinLength field if non-nil, zero value otherwise.

### GetUserPinLengthOk

`func (o *Service) GetUserPinLengthOk() (*int32, bool)`

GetUserPinLengthOk returns a tuple with the UserPinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPinLength

`func (o *Service) SetUserPinLength(v int32)`

SetUserPinLength sets UserPinLength field to given value.

### HasUserPinLength

`func (o *Service) HasUserPinLength() bool`

HasUserPinLength returns a boolean if a field has been set.

### GetIdTokenAudType

`func (o *Service) GetIdTokenAudType() string`

GetIdTokenAudType returns the IdTokenAudType field if non-nil, zero value otherwise.

### GetIdTokenAudTypeOk

`func (o *Service) GetIdTokenAudTypeOk() (*string, bool)`

GetIdTokenAudTypeOk returns a tuple with the IdTokenAudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTokenAudType

`func (o *Service) SetIdTokenAudType(v string)`

SetIdTokenAudType sets IdTokenAudType field to given value.

### HasIdTokenAudType

`func (o *Service) HasIdTokenAudType() bool`

HasIdTokenAudType returns a boolean if a field has been set.

### GetSupportedPromptValues

`func (o *Service) GetSupportedPromptValues() []string`

GetSupportedPromptValues returns the SupportedPromptValues field if non-nil, zero value otherwise.

### GetSupportedPromptValuesOk

`func (o *Service) GetSupportedPromptValuesOk() (*[]string, bool)`

GetSupportedPromptValuesOk returns a tuple with the SupportedPromptValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedPromptValues

`func (o *Service) SetSupportedPromptValues(v []string)`

SetSupportedPromptValues sets SupportedPromptValues field to given value.

### HasSupportedPromptValues

`func (o *Service) HasSupportedPromptValues() bool`

HasSupportedPromptValues returns a boolean if a field has been set.

### GetVerifiedClaimsValidationSchemaSet

`func (o *Service) GetVerifiedClaimsValidationSchemaSet() string`

GetVerifiedClaimsValidationSchemaSet returns the VerifiedClaimsValidationSchemaSet field if non-nil, zero value otherwise.

### GetVerifiedClaimsValidationSchemaSetOk

`func (o *Service) GetVerifiedClaimsValidationSchemaSetOk() (*string, bool)`

GetVerifiedClaimsValidationSchemaSetOk returns a tuple with the VerifiedClaimsValidationSchemaSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedClaimsValidationSchemaSet

`func (o *Service) SetVerifiedClaimsValidationSchemaSet(v string)`

SetVerifiedClaimsValidationSchemaSet sets VerifiedClaimsValidationSchemaSet field to given value.

### HasVerifiedClaimsValidationSchemaSet

`func (o *Service) HasVerifiedClaimsValidationSchemaSet() bool`

HasVerifiedClaimsValidationSchemaSet returns a boolean if a field has been set.

### GetPreAuthorizedGrantAnonymousAccessSupported

`func (o *Service) GetPreAuthorizedGrantAnonymousAccessSupported() bool`

GetPreAuthorizedGrantAnonymousAccessSupported returns the PreAuthorizedGrantAnonymousAccessSupported field if non-nil, zero value otherwise.

### GetPreAuthorizedGrantAnonymousAccessSupportedOk

`func (o *Service) GetPreAuthorizedGrantAnonymousAccessSupportedOk() (*bool, bool)`

GetPreAuthorizedGrantAnonymousAccessSupportedOk returns a tuple with the PreAuthorizedGrantAnonymousAccessSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreAuthorizedGrantAnonymousAccessSupported

`func (o *Service) SetPreAuthorizedGrantAnonymousAccessSupported(v bool)`

SetPreAuthorizedGrantAnonymousAccessSupported sets PreAuthorizedGrantAnonymousAccessSupported field to given value.

### HasPreAuthorizedGrantAnonymousAccessSupported

`func (o *Service) HasPreAuthorizedGrantAnonymousAccessSupported() bool`

HasPreAuthorizedGrantAnonymousAccessSupported returns a boolean if a field has been set.

### GetCredentialTransactionDuration

`func (o *Service) GetCredentialTransactionDuration() int64`

GetCredentialTransactionDuration returns the CredentialTransactionDuration field if non-nil, zero value otherwise.

### GetCredentialTransactionDurationOk

`func (o *Service) GetCredentialTransactionDurationOk() (*int64, bool)`

GetCredentialTransactionDurationOk returns a tuple with the CredentialTransactionDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialTransactionDuration

`func (o *Service) SetCredentialTransactionDuration(v int64)`

SetCredentialTransactionDuration sets CredentialTransactionDuration field to given value.

### HasCredentialTransactionDuration

`func (o *Service) HasCredentialTransactionDuration() bool`

HasCredentialTransactionDuration returns a boolean if a field has been set.

### GetCredentialDuration

`func (o *Service) GetCredentialDuration() int64`

GetCredentialDuration returns the CredentialDuration field if non-nil, zero value otherwise.

### GetCredentialDurationOk

`func (o *Service) GetCredentialDurationOk() (*int64, bool)`

GetCredentialDurationOk returns a tuple with the CredentialDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialDuration

`func (o *Service) SetCredentialDuration(v int64)`

SetCredentialDuration sets CredentialDuration field to given value.

### HasCredentialDuration

`func (o *Service) HasCredentialDuration() bool`

HasCredentialDuration returns a boolean if a field has been set.

### GetCredentialJwks

`func (o *Service) GetCredentialJwks() string`

GetCredentialJwks returns the CredentialJwks field if non-nil, zero value otherwise.

### GetCredentialJwksOk

`func (o *Service) GetCredentialJwksOk() (*string, bool)`

GetCredentialJwksOk returns a tuple with the CredentialJwks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialJwks

`func (o *Service) SetCredentialJwks(v string)`

SetCredentialJwks sets CredentialJwks field to given value.

### HasCredentialJwks

`func (o *Service) HasCredentialJwks() bool`

HasCredentialJwks returns a boolean if a field has been set.

### GetIdTokenReissuable

`func (o *Service) GetIdTokenReissuable() bool`

GetIdTokenReissuable returns the IdTokenReissuable field if non-nil, zero value otherwise.

### GetIdTokenReissuableOk

`func (o *Service) GetIdTokenReissuableOk() (*bool, bool)`

GetIdTokenReissuableOk returns a tuple with the IdTokenReissuable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTokenReissuable

`func (o *Service) SetIdTokenReissuable(v bool)`

SetIdTokenReissuable sets IdTokenReissuable field to given value.

### HasIdTokenReissuable

`func (o *Service) HasIdTokenReissuable() bool`

HasIdTokenReissuable returns a boolean if a field has been set.

### GetCnonceDuration

`func (o *Service) GetCnonceDuration() int64`

GetCnonceDuration returns the CnonceDuration field if non-nil, zero value otherwise.

### GetCnonceDurationOk

`func (o *Service) GetCnonceDurationOk() (*int64, bool)`

GetCnonceDurationOk returns a tuple with the CnonceDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnonceDuration

`func (o *Service) SetCnonceDuration(v int64)`

SetCnonceDuration sets CnonceDuration field to given value.

### HasCnonceDuration

`func (o *Service) HasCnonceDuration() bool`

HasCnonceDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


