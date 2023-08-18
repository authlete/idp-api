# \TokenEndpointApi

All URIs are relative to *https://devidp.authlete.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TokenEndpoint**](TokenEndpointApi.md#TokenEndpoint) | **Post** /token | 



## TokenEndpoint

> string TokenEndpoint(ctx).Authorization(authorization).DPoP(dPoP).Empty(empty).Execute()



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
    empty := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokenEndpointApi.TokenEndpoint(context.Background()).Authorization(authorization).DPoP(dPoP).Empty(empty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenEndpointApi.TokenEndpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TokenEndpoint`: string
    fmt.Fprintf(os.Stdout, "Response from `TokenEndpointApi.TokenEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTokenEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **dPoP** | **string** |  | 
 **empty** | **bool** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

