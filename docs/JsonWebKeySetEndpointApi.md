# \JsonWebKeySetEndpointAPI

All URIs are relative to *https://devidp.authlete.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetJsonWebKeySet**](JsonWebKeySetEndpointAPI.md#GetJsonWebKeySet) | **Get** /jwks | 



## GetJsonWebKeySet

> string GetJsonWebKeySet(ctx).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JsonWebKeySetEndpointAPI.GetJsonWebKeySet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JsonWebKeySetEndpointAPI.GetJsonWebKeySet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJsonWebKeySet`: string
	fmt.Fprintf(os.Stdout, "Response from `JsonWebKeySetEndpointAPI.GetJsonWebKeySet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetJsonWebKeySetRequest struct via the builder pattern


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

