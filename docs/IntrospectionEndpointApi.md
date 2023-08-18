# \IntrospectionEndpointApi

All URIs are relative to *https://devidp.authlete.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Introspect**](IntrospectionEndpointApi.md#Introspect) | **Post** /introspect | 



## Introspect

> string Introspect(ctx).Empty(empty).Execute()



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
    empty := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntrospectionEndpointApi.Introspect(context.Background()).Empty(empty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntrospectionEndpointApi.Introspect``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Introspect`: string
    fmt.Fprintf(os.Stdout, "Response from `IntrospectionEndpointApi.Introspect`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIntrospectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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

