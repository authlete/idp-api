# \ServiceApiAPI

All URIs are relative to *https://devidp.authlete.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdoptService**](ServiceApiAPI.md#AdoptService) | **Post** /api/service/adopt | 
[**CreateService**](ServiceApiAPI.md#CreateService) | **Post** /api/service | 
[**DeleteService**](ServiceApiAPI.md#DeleteService) | **Post** /api/service/remove | 
[**FindService**](ServiceApiAPI.md#FindService) | **Get** /api/service/find/{id} | 
[**GetOrphans**](ServiceApiAPI.md#GetOrphans) | **Get** /api/service/orphans | 
[**MoveService**](ServiceApiAPI.md#MoveService) | **Post** /api/service/move | 
[**RemoveOrphanService**](ServiceApiAPI.md#RemoveOrphanService) | **Post** /api/service/orphans/remove | 



## AdoptService

> ServiceInstanceManagementResponse AdoptService(ctx).AdoptServiceRequest(adoptServiceRequest).Authorization(authorization).DPoP(dPoP).Execute()



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
	adoptServiceRequest := *openapiclient.NewAdoptServiceRequest(int64(123), int64(123), int64(123)) // AdoptServiceRequest | 
	authorization := "authorization_example" // string |  (optional)
	dPoP := "dPoP_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceApiAPI.AdoptService(context.Background()).AdoptServiceRequest(adoptServiceRequest).Authorization(authorization).DPoP(dPoP).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceApiAPI.AdoptService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdoptService`: ServiceInstanceManagementResponse
	fmt.Fprintf(os.Stdout, "Response from `ServiceApiAPI.AdoptService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdoptServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adoptServiceRequest** | [**AdoptServiceRequest**](AdoptServiceRequest.md) |  | 
 **authorization** | **string** |  | 
 **dPoP** | **string** |  | 

### Return type

[**ServiceInstanceManagementResponse**](ServiceInstanceManagementResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateService

> Service CreateService(ctx).CreateServiceRequest(createServiceRequest).Authorization(authorization).DPoP(dPoP).Execute()



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
	createServiceRequest := *openapiclient.NewCreateServiceRequest(int64(123), int64(123)) // CreateServiceRequest | 
	authorization := "authorization_example" // string |  (optional)
	dPoP := "dPoP_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceApiAPI.CreateService(context.Background()).CreateServiceRequest(createServiceRequest).Authorization(authorization).DPoP(dPoP).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceApiAPI.CreateService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateService`: Service
	fmt.Fprintf(os.Stdout, "Response from `ServiceApiAPI.CreateService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createServiceRequest** | [**CreateServiceRequest**](CreateServiceRequest.md) |  | 
 **authorization** | **string** |  | 
 **dPoP** | **string** |  | 

### Return type

[**Service**](Service.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteService

> DeleteService(ctx).DeleteServiceRequest(deleteServiceRequest).Authorization(authorization).DPoP(dPoP).Execute()



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
	deleteServiceRequest := *openapiclient.NewDeleteServiceRequest(int64(123), int64(123), int64(123)) // DeleteServiceRequest | 
	authorization := "authorization_example" // string |  (optional)
	dPoP := "dPoP_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServiceApiAPI.DeleteService(context.Background()).DeleteServiceRequest(deleteServiceRequest).Authorization(authorization).DPoP(dPoP).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceApiAPI.DeleteService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteServiceRequest** | [**DeleteServiceRequest**](DeleteServiceRequest.md) |  | 
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


## FindService

> []ServiceInstanceManagementResponse FindService(ctx, id).Authorization(authorization).DPoP(dPoP).Execute()



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
	resp, r, err := apiClient.ServiceApiAPI.FindService(context.Background(), id).Authorization(authorization).DPoP(dPoP).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceApiAPI.FindService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindService`: []ServiceInstanceManagementResponse
	fmt.Fprintf(os.Stdout, "Response from `ServiceApiAPI.FindService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** |  | 
 **dPoP** | **string** |  | 

### Return type

[**[]ServiceInstanceManagementResponse**](ServiceInstanceManagementResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrphans

> OrphanServiceResponse GetOrphans(ctx).Authorization(authorization).DPoP(dPoP).Execute()



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
	resp, r, err := apiClient.ServiceApiAPI.GetOrphans(context.Background()).Authorization(authorization).DPoP(dPoP).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceApiAPI.GetOrphans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrphans`: OrphanServiceResponse
	fmt.Fprintf(os.Stdout, "Response from `ServiceApiAPI.GetOrphans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrphansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **dPoP** | **string** |  | 

### Return type

[**OrphanServiceResponse**](OrphanServiceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoveService

> ServiceInstanceManagementResponse MoveService(ctx).MoveServiceRequest(moveServiceRequest).Authorization(authorization).DPoP(dPoP).Execute()



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
	moveServiceRequest := *openapiclient.NewMoveServiceRequest(int64(123), int64(123), int64(123), int64(123)) // MoveServiceRequest | 
	authorization := "authorization_example" // string |  (optional)
	dPoP := "dPoP_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceApiAPI.MoveService(context.Background()).MoveServiceRequest(moveServiceRequest).Authorization(authorization).DPoP(dPoP).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceApiAPI.MoveService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MoveService`: ServiceInstanceManagementResponse
	fmt.Fprintf(os.Stdout, "Response from `ServiceApiAPI.MoveService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMoveServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **moveServiceRequest** | [**MoveServiceRequest**](MoveServiceRequest.md) |  | 
 **authorization** | **string** |  | 
 **dPoP** | **string** |  | 

### Return type

[**ServiceInstanceManagementResponse**](ServiceInstanceManagementResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveOrphanService

> RemoveOrphanService(ctx).RemoveOrphanRequest(removeOrphanRequest).Authorization(authorization).DPoP(dPoP).Execute()



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
	removeOrphanRequest := *openapiclient.NewRemoveOrphanRequest(int64(123), int64(123)) // RemoveOrphanRequest | 
	authorization := "authorization_example" // string |  (optional)
	dPoP := "dPoP_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServiceApiAPI.RemoveOrphanService(context.Background()).RemoveOrphanRequest(removeOrphanRequest).Authorization(authorization).DPoP(dPoP).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceApiAPI.RemoveOrphanService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveOrphanServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **removeOrphanRequest** | [**RemoveOrphanRequest**](RemoveOrphanRequest.md) |  | 
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

