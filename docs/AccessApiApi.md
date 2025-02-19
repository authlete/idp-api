# \AccessApiAPI

All URIs are relative to *https://devidp.authlete.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcceptInvitation**](AccessApiAPI.md#AcceptInvitation) | **Post** /api/access/invite/{id} | 
[**DeclineInvitation**](AccessApiAPI.md#DeclineInvitation) | **Delete** /api/access/invite/{id} | 
[**GetInvitation**](AccessApiAPI.md#GetInvitation) | **Get** /api/access/invite/{id} | 
[**GetReceivedInvitations**](AccessApiAPI.md#GetReceivedInvitations) | **Get** /api/access/invite/received | 
[**GetSentInvitations**](AccessApiAPI.md#GetSentInvitations) | **Get** /api/access/invite/sent | 
[**InviteByEmail**](AccessApiAPI.md#InviteByEmail) | **Post** /api/access/invite | 
[**RemoveApiServerPrivileges**](AccessApiAPI.md#RemoveApiServerPrivileges) | **Post** /api/access/apiserver/remove | 
[**RemoveClientPrivileges**](AccessApiAPI.md#RemoveClientPrivileges) | **Post** /api/access/client/remove | 
[**RemoveOrganizationPrivileges**](AccessApiAPI.md#RemoveOrganizationPrivileges) | **Post** /api/access/organization/remove | 
[**RemoveServicePrivileges**](AccessApiAPI.md#RemoveServicePrivileges) | **Post** /api/access/service/remove | 
[**UpdateApiServerPrivileges**](AccessApiAPI.md#UpdateApiServerPrivileges) | **Post** /api/access/apiserver | 
[**UpdateClientPrivileges**](AccessApiAPI.md#UpdateClientPrivileges) | **Post** /api/access/client | 
[**UpdateOrganizationPrivileges**](AccessApiAPI.md#UpdateOrganizationPrivileges) | **Post** /api/access/organization | 
[**UpdateServicePrivileges**](AccessApiAPI.md#UpdateServicePrivileges) | **Post** /api/access/service | 



## AcceptInvitation

> AcceptInvitation(ctx, id).Authorization(authorization).DPoP(dPoP).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/authlete/idp-api"
)

func main() {
	id := int64(789) // int64 | 
	authorization := "authorization_example" // string |  (optional)
	dPoP := "dPoP_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccessApiAPI.AcceptInvitation(context.Background(), id).Authorization(authorization).DPoP(dPoP).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessApiAPI.AcceptInvitation``: %v\n", err)
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
	openapiclient "github.com/authlete/idp-api"
)

func main() {
	id := int64(789) // int64 | 
	authorization := "authorization_example" // string |  (optional)
	dPoP := "dPoP_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccessApiAPI.DeclineInvitation(context.Background(), id).Authorization(authorization).DPoP(dPoP).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessApiAPI.DeclineInvitation``: %v\n", err)
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
	openapiclient "github.com/authlete/idp-api"
)

func main() {
	id := int64(789) // int64 | 
	authorization := "authorization_example" // string |  (optional)
	dPoP := "dPoP_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessApiAPI.GetInvitation(context.Background(), id).Authorization(authorization).DPoP(dPoP).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessApiAPI.GetInvitation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInvitation`: InvitationResponse
	fmt.Fprintf(os.Stdout, "Response from `AccessApiAPI.GetInvitation`: %v\n", resp)
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
	openapiclient "github.com/authlete/idp-api"
)

func main() {
	authorization := "authorization_example" // string |  (optional)
	dPoP := "dPoP_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessApiAPI.GetReceivedInvitations(context.Background()).Authorization(authorization).DPoP(dPoP).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessApiAPI.GetReceivedInvitations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReceivedInvitations`: []InvitationResponse
	fmt.Fprintf(os.Stdout, "Response from `AccessApiAPI.GetReceivedInvitations`: %v\n", resp)
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
	openapiclient "github.com/authlete/idp-api"
)

func main() {
	authorization := "authorization_example" // string |  (optional)
	dPoP := "dPoP_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessApiAPI.GetSentInvitations(context.Background()).Authorization(authorization).DPoP(dPoP).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessApiAPI.GetSentInvitations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSentInvitations`: []InvitationResponse
	fmt.Fprintf(os.Stdout, "Response from `AccessApiAPI.GetSentInvitations`: %v\n", resp)
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
	openapiclient "github.com/authlete/idp-api"
)

func main() {
	inviteRequest := *openapiclient.NewInviteRequest(int64(123), "Email_example") // InviteRequest | 
	authorization := "authorization_example" // string |  (optional)
	dPoP := "dPoP_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessApiAPI.InviteByEmail(context.Background()).InviteRequest(inviteRequest).Authorization(authorization).DPoP(dPoP).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessApiAPI.InviteByEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InviteByEmail`: InvitationResponse
	fmt.Fprintf(os.Stdout, "Response from `AccessApiAPI.InviteByEmail`: %v\n", resp)
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
	openapiclient "github.com/authlete/idp-api"
)

func main() {
	removeApiServerPrivilegesRequest := *openapiclient.NewRemoveApiServerPrivilegesRequest() // RemoveApiServerPrivilegesRequest | 
	authorization := "authorization_example" // string |  (optional)
	dPoP := "dPoP_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccessApiAPI.RemoveApiServerPrivileges(context.Background()).RemoveApiServerPrivilegesRequest(removeApiServerPrivilegesRequest).Authorization(authorization).DPoP(dPoP).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessApiAPI.RemoveApiServerPrivileges``: %v\n", err)
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
	openapiclient "github.com/authlete/idp-api"
)

func main() {
	removeClientPrivilegesRequest := *openapiclient.NewRemoveClientPrivilegesRequest(int64(123), int64(123), int64(123), int64(123), int64(123)) // RemoveClientPrivilegesRequest | 
	authorization := "authorization_example" // string |  (optional)
	dPoP := "dPoP_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccessApiAPI.RemoveClientPrivileges(context.Background()).RemoveClientPrivilegesRequest(removeClientPrivilegesRequest).Authorization(authorization).DPoP(dPoP).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessApiAPI.RemoveClientPrivileges``: %v\n", err)
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
	openapiclient "github.com/authlete/idp-api"
)

func main() {
	removeOrganizationPrivilegesRequest := *openapiclient.NewRemoveOrganizationPrivilegesRequest(int64(123), int64(123)) // RemoveOrganizationPrivilegesRequest | 
	authorization := "authorization_example" // string |  (optional)
	dPoP := "dPoP_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccessApiAPI.RemoveOrganizationPrivileges(context.Background()).RemoveOrganizationPrivilegesRequest(removeOrganizationPrivilegesRequest).Authorization(authorization).DPoP(dPoP).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessApiAPI.RemoveOrganizationPrivileges``: %v\n", err)
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
	openapiclient "github.com/authlete/idp-api"
)

func main() {
	removeServicePrivilegesRequest := *openapiclient.NewRemoveServicePrivilegesRequest(int64(123), int64(123), int64(123), int64(123)) // RemoveServicePrivilegesRequest | 
	authorization := "authorization_example" // string |  (optional)
	dPoP := "dPoP_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccessApiAPI.RemoveServicePrivileges(context.Background()).RemoveServicePrivilegesRequest(removeServicePrivilegesRequest).Authorization(authorization).DPoP(dPoP).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessApiAPI.RemoveServicePrivileges``: %v\n", err)
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
	openapiclient "github.com/authlete/idp-api"
)

func main() {
	updateApiServerPrivilegesRequest := *openapiclient.NewUpdateApiServerPrivilegesRequest(int64(123), int64(123), int64(123)) // UpdateApiServerPrivilegesRequest | 
	authorization := "authorization_example" // string |  (optional)
	dPoP := "dPoP_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessApiAPI.UpdateApiServerPrivileges(context.Background()).UpdateApiServerPrivilegesRequest(updateApiServerPrivilegesRequest).Authorization(authorization).DPoP(dPoP).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessApiAPI.UpdateApiServerPrivileges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApiServerPrivileges`: ApiServerAccessResponse
	fmt.Fprintf(os.Stdout, "Response from `AccessApiAPI.UpdateApiServerPrivileges`: %v\n", resp)
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
	openapiclient "github.com/authlete/idp-api"
)

func main() {
	updateClientPrivilegesRequest := *openapiclient.NewUpdateClientPrivilegesRequest(int64(123), int64(123), int64(123), int64(123), int64(123)) // UpdateClientPrivilegesRequest | 
	authorization := "authorization_example" // string |  (optional)
	dPoP := "dPoP_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessApiAPI.UpdateClientPrivileges(context.Background()).UpdateClientPrivilegesRequest(updateClientPrivilegesRequest).Authorization(authorization).DPoP(dPoP).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessApiAPI.UpdateClientPrivileges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateClientPrivileges`: ClientAccessResponse
	fmt.Fprintf(os.Stdout, "Response from `AccessApiAPI.UpdateClientPrivileges`: %v\n", resp)
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
	openapiclient "github.com/authlete/idp-api"
)

func main() {
	updateOrganizationPrivilegesRequest := *openapiclient.NewUpdateOrganizationPrivilegesRequest(int64(123), int64(123)) // UpdateOrganizationPrivilegesRequest | 
	authorization := "authorization_example" // string |  (optional)
	dPoP := "dPoP_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessApiAPI.UpdateOrganizationPrivileges(context.Background()).UpdateOrganizationPrivilegesRequest(updateOrganizationPrivilegesRequest).Authorization(authorization).DPoP(dPoP).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessApiAPI.UpdateOrganizationPrivileges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrganizationPrivileges`: OrganizationAccessResponse
	fmt.Fprintf(os.Stdout, "Response from `AccessApiAPI.UpdateOrganizationPrivileges`: %v\n", resp)
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
	openapiclient "github.com/authlete/idp-api"
)

func main() {
	updateServicePrivilegesRequest := *openapiclient.NewUpdateServicePrivilegesRequest(int64(123), int64(123), int64(123), int64(123)) // UpdateServicePrivilegesRequest | 
	authorization := "authorization_example" // string |  (optional)
	dPoP := "dPoP_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessApiAPI.UpdateServicePrivileges(context.Background()).UpdateServicePrivilegesRequest(updateServicePrivilegesRequest).Authorization(authorization).DPoP(dPoP).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessApiAPI.UpdateServicePrivileges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateServicePrivileges`: ServiceAccessResponse
	fmt.Fprintf(os.Stdout, "Response from `AccessApiAPI.UpdateServicePrivileges`: %v\n", resp)
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

