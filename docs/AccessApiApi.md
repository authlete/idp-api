# \AccessApiApi

All URIs are relative to *https://devidp.authlete.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcceptInvitation**](AccessApiApi.md#AcceptInvitation) | **Post** /api/access/invite/{id} | 
[**DeclineInvitation**](AccessApiApi.md#DeclineInvitation) | **Delete** /api/access/invite/{id} | 
[**GetInvitation**](AccessApiApi.md#GetInvitation) | **Get** /api/access/invite/{id} | 
[**GetReceivedInvitations**](AccessApiApi.md#GetReceivedInvitations) | **Get** /api/access/invite/received | 
[**GetSentInvitations**](AccessApiApi.md#GetSentInvitations) | **Get** /api/access/invite/sent | 
[**InviteByEmail**](AccessApiApi.md#InviteByEmail) | **Post** /api/access/invite | 
[**RemoveApiServerPrivileges**](AccessApiApi.md#RemoveApiServerPrivileges) | **Post** /api/access/apiserver/remove | 
[**RemoveClientPrivileges**](AccessApiApi.md#RemoveClientPrivileges) | **Post** /api/access/client/remove | 
[**RemoveOrganizationPrivileges**](AccessApiApi.md#RemoveOrganizationPrivileges) | **Post** /api/access/organization/remove | 
[**RemoveServicePrivileges**](AccessApiApi.md#RemoveServicePrivileges) | **Post** /api/access/service/remove | 
[**UpdateApiServerPrivileges**](AccessApiApi.md#UpdateApiServerPrivileges) | **Post** /api/access/apiserver | 
[**UpdateClientPrivileges**](AccessApiApi.md#UpdateClientPrivileges) | **Post** /api/access/client | 
[**UpdateOrganizationPrivileges**](AccessApiApi.md#UpdateOrganizationPrivileges) | **Post** /api/access/organization | 
[**UpdateServicePrivileges**](AccessApiApi.md#UpdateServicePrivileges) | **Post** /api/access/service | 



## AcceptInvitation

> AcceptInvitation(ctx, id).Authorization(authorization).DPoP(dPoP).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(789) // int64 | 
    authorization := "authorization_example" // string |  (optional)
    dPoP := "dPoP_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessApiApi.AcceptInvitation(context.Background(), id).Authorization(authorization).DPoP(dPoP).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApiApi.AcceptInvitation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcceptInvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** |  | 
 **dPoP** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeclineInvitation

> DeclineInvitation(ctx, id).Authorization(authorization).DPoP(dPoP).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(789) // int64 | 
    authorization := "authorization_example" // string |  (optional)
    dPoP := "dPoP_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessApiApi.DeclineInvitation(context.Background(), id).Authorization(authorization).DPoP(dPoP).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApiApi.DeclineInvitation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeclineInvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** |  | 
 **dPoP** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvitation

> InvitationResponse GetInvitation(ctx, id).Authorization(authorization).DPoP(dPoP).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(789) // int64 | 
    authorization := "authorization_example" // string |  (optional)
    dPoP := "dPoP_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessApiApi.GetInvitation(context.Background(), id).Authorization(authorization).DPoP(dPoP).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApiApi.GetInvitation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInvitation`: InvitationResponse
    fmt.Fprintf(os.Stdout, "Response from `AccessApiApi.GetInvitation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** |  | 
 **dPoP** | **string** |  | 

### Return type

[**InvitationResponse**](InvitationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReceivedInvitations

> []InvitationResponse GetReceivedInvitations(ctx).Authorization(authorization).DPoP(dPoP).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    authorization := "authorization_example" // string |  (optional)
    dPoP := "dPoP_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessApiApi.GetReceivedInvitations(context.Background()).Authorization(authorization).DPoP(dPoP).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApiApi.GetReceivedInvitations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReceivedInvitations`: []InvitationResponse
    fmt.Fprintf(os.Stdout, "Response from `AccessApiApi.GetReceivedInvitations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetReceivedInvitationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **dPoP** | **string** |  | 

### Return type

[**[]InvitationResponse**](InvitationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSentInvitations

> []InvitationResponse GetSentInvitations(ctx).Authorization(authorization).DPoP(dPoP).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    authorization := "authorization_example" // string |  (optional)
    dPoP := "dPoP_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessApiApi.GetSentInvitations(context.Background()).Authorization(authorization).DPoP(dPoP).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApiApi.GetSentInvitations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSentInvitations`: []InvitationResponse
    fmt.Fprintf(os.Stdout, "Response from `AccessApiApi.GetSentInvitations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSentInvitationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **dPoP** | **string** |  | 

### Return type

[**[]InvitationResponse**](InvitationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InviteByEmail

> InvitationResponse InviteByEmail(ctx).InviteRequest(inviteRequest).Authorization(authorization).DPoP(dPoP).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    inviteRequest := *openapiclient.NewInviteRequest(int64(123), "Email_example") // InviteRequest | 
    authorization := "authorization_example" // string |  (optional)
    dPoP := "dPoP_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessApiApi.InviteByEmail(context.Background()).InviteRequest(inviteRequest).Authorization(authorization).DPoP(dPoP).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApiApi.InviteByEmail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InviteByEmail`: InvitationResponse
    fmt.Fprintf(os.Stdout, "Response from `AccessApiApi.InviteByEmail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInviteByEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inviteRequest** | [**InviteRequest**](InviteRequest.md) |  | 
 **authorization** | **string** |  | 
 **dPoP** | **string** |  | 

### Return type

[**InvitationResponse**](InvitationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveApiServerPrivileges

> RemoveApiServerPrivileges(ctx).RemoveApiServerPrivilegesRequest(removeApiServerPrivilegesRequest).Authorization(authorization).DPoP(dPoP).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    removeApiServerPrivilegesRequest := *openapiclient.NewRemoveApiServerPrivilegesRequest() // RemoveApiServerPrivilegesRequest | 
    authorization := "authorization_example" // string |  (optional)
    dPoP := "dPoP_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessApiApi.RemoveApiServerPrivileges(context.Background()).RemoveApiServerPrivilegesRequest(removeApiServerPrivilegesRequest).Authorization(authorization).DPoP(dPoP).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApiApi.RemoveApiServerPrivileges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveApiServerPrivilegesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **removeApiServerPrivilegesRequest** | [**RemoveApiServerPrivilegesRequest**](RemoveApiServerPrivilegesRequest.md) |  | 
 **authorization** | **string** |  | 
 **dPoP** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveClientPrivileges

> RemoveClientPrivileges(ctx).RemoveClientPrivilegesRequest(removeClientPrivilegesRequest).Authorization(authorization).DPoP(dPoP).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    removeClientPrivilegesRequest := *openapiclient.NewRemoveClientPrivilegesRequest(int64(123), int64(123), int64(123), int64(123), int64(123)) // RemoveClientPrivilegesRequest | 
    authorization := "authorization_example" // string |  (optional)
    dPoP := "dPoP_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessApiApi.RemoveClientPrivileges(context.Background()).RemoveClientPrivilegesRequest(removeClientPrivilegesRequest).Authorization(authorization).DPoP(dPoP).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApiApi.RemoveClientPrivileges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveClientPrivilegesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **removeClientPrivilegesRequest** | [**RemoveClientPrivilegesRequest**](RemoveClientPrivilegesRequest.md) |  | 
 **authorization** | **string** |  | 
 **dPoP** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveOrganizationPrivileges

> RemoveOrganizationPrivileges(ctx).RemoveOrganizationPrivilegesRequest(removeOrganizationPrivilegesRequest).Authorization(authorization).DPoP(dPoP).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    removeOrganizationPrivilegesRequest := *openapiclient.NewRemoveOrganizationPrivilegesRequest(int64(123), int64(123)) // RemoveOrganizationPrivilegesRequest | 
    authorization := "authorization_example" // string |  (optional)
    dPoP := "dPoP_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessApiApi.RemoveOrganizationPrivileges(context.Background()).RemoveOrganizationPrivilegesRequest(removeOrganizationPrivilegesRequest).Authorization(authorization).DPoP(dPoP).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApiApi.RemoveOrganizationPrivileges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveOrganizationPrivilegesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **removeOrganizationPrivilegesRequest** | [**RemoveOrganizationPrivilegesRequest**](RemoveOrganizationPrivilegesRequest.md) |  | 
 **authorization** | **string** |  | 
 **dPoP** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveServicePrivileges

> RemoveServicePrivileges(ctx).RemoveServicePrivilegesRequest(removeServicePrivilegesRequest).Authorization(authorization).DPoP(dPoP).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    removeServicePrivilegesRequest := *openapiclient.NewRemoveServicePrivilegesRequest(int64(123), int64(123), int64(123), int64(123)) // RemoveServicePrivilegesRequest | 
    authorization := "authorization_example" // string |  (optional)
    dPoP := "dPoP_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessApiApi.RemoveServicePrivileges(context.Background()).RemoveServicePrivilegesRequest(removeServicePrivilegesRequest).Authorization(authorization).DPoP(dPoP).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApiApi.RemoveServicePrivileges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveServicePrivilegesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **removeServicePrivilegesRequest** | [**RemoveServicePrivilegesRequest**](RemoveServicePrivilegesRequest.md) |  | 
 **authorization** | **string** |  | 
 **dPoP** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApiServerPrivileges

> ApiServerAccessResponse UpdateApiServerPrivileges(ctx).UpdateApiServerPrivilegesRequest(updateApiServerPrivilegesRequest).Authorization(authorization).DPoP(dPoP).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    updateApiServerPrivilegesRequest := *openapiclient.NewUpdateApiServerPrivilegesRequest(int64(123), int64(123), int64(123)) // UpdateApiServerPrivilegesRequest | 
    authorization := "authorization_example" // string |  (optional)
    dPoP := "dPoP_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessApiApi.UpdateApiServerPrivileges(context.Background()).UpdateApiServerPrivilegesRequest(updateApiServerPrivilegesRequest).Authorization(authorization).DPoP(dPoP).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApiApi.UpdateApiServerPrivileges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApiServerPrivileges`: ApiServerAccessResponse
    fmt.Fprintf(os.Stdout, "Response from `AccessApiApi.UpdateApiServerPrivileges`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiServerPrivilegesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateApiServerPrivilegesRequest** | [**UpdateApiServerPrivilegesRequest**](UpdateApiServerPrivilegesRequest.md) |  | 
 **authorization** | **string** |  | 
 **dPoP** | **string** |  | 

### Return type

[**ApiServerAccessResponse**](ApiServerAccessResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateClientPrivileges

> ClientAccessResponse UpdateClientPrivileges(ctx).UpdateClientPrivilegesRequest(updateClientPrivilegesRequest).Authorization(authorization).DPoP(dPoP).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    updateClientPrivilegesRequest := *openapiclient.NewUpdateClientPrivilegesRequest(int64(123), int64(123), int64(123), int64(123), int64(123)) // UpdateClientPrivilegesRequest | 
    authorization := "authorization_example" // string |  (optional)
    dPoP := "dPoP_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessApiApi.UpdateClientPrivileges(context.Background()).UpdateClientPrivilegesRequest(updateClientPrivilegesRequest).Authorization(authorization).DPoP(dPoP).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApiApi.UpdateClientPrivileges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateClientPrivileges`: ClientAccessResponse
    fmt.Fprintf(os.Stdout, "Response from `AccessApiApi.UpdateClientPrivileges`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClientPrivilegesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateClientPrivilegesRequest** | [**UpdateClientPrivilegesRequest**](UpdateClientPrivilegesRequest.md) |  | 
 **authorization** | **string** |  | 
 **dPoP** | **string** |  | 

### Return type

[**ClientAccessResponse**](ClientAccessResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationPrivileges

> OrganizationAccessResponse UpdateOrganizationPrivileges(ctx).UpdateOrganizationPrivilegesRequest(updateOrganizationPrivilegesRequest).Authorization(authorization).DPoP(dPoP).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    updateOrganizationPrivilegesRequest := *openapiclient.NewUpdateOrganizationPrivilegesRequest(int64(123), int64(123)) // UpdateOrganizationPrivilegesRequest | 
    authorization := "authorization_example" // string |  (optional)
    dPoP := "dPoP_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessApiApi.UpdateOrganizationPrivileges(context.Background()).UpdateOrganizationPrivilegesRequest(updateOrganizationPrivilegesRequest).Authorization(authorization).DPoP(dPoP).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApiApi.UpdateOrganizationPrivileges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationPrivileges`: OrganizationAccessResponse
    fmt.Fprintf(os.Stdout, "Response from `AccessApiApi.UpdateOrganizationPrivileges`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationPrivilegesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateOrganizationPrivilegesRequest** | [**UpdateOrganizationPrivilegesRequest**](UpdateOrganizationPrivilegesRequest.md) |  | 
 **authorization** | **string** |  | 
 **dPoP** | **string** |  | 

### Return type

[**OrganizationAccessResponse**](OrganizationAccessResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServicePrivileges

> ServiceAccessResponse UpdateServicePrivileges(ctx).UpdateServicePrivilegesRequest(updateServicePrivilegesRequest).Authorization(authorization).DPoP(dPoP).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    updateServicePrivilegesRequest := *openapiclient.NewUpdateServicePrivilegesRequest(int64(123), int64(123), int64(123), int64(123)) // UpdateServicePrivilegesRequest | 
    authorization := "authorization_example" // string |  (optional)
    dPoP := "dPoP_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessApiApi.UpdateServicePrivileges(context.Background()).UpdateServicePrivilegesRequest(updateServicePrivilegesRequest).Authorization(authorization).DPoP(dPoP).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApiApi.UpdateServicePrivileges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateServicePrivileges`: ServiceAccessResponse
    fmt.Fprintf(os.Stdout, "Response from `AccessApiApi.UpdateServicePrivileges`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServicePrivilegesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateServicePrivilegesRequest** | [**UpdateServicePrivilegesRequest**](UpdateServicePrivilegesRequest.md) |  | 
 **authorization** | **string** |  | 
 **dPoP** | **string** |  | 

### Return type

[**ServiceAccessResponse**](ServiceAccessResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

