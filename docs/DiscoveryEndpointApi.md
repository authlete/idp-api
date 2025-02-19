# \DiscoveryEndpointAPI

All URIs are relative to *https://devidp.authlete.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConfiguration**](DiscoveryEndpointAPI.md#GetConfiguration) | **Get** /.well-known/openid-configuration | 



## GetConfiguration

> string GetConfiguration(ctx).Execute()



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
	resp, r, err := apiClient.DiscoveryEndpointAPI.GetConfiguration(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoveryEndpointAPI.GetConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfiguration`: string
	fmt.Fprintf(os.Stdout, "Response from `DiscoveryEndpointAPI.GetConfiguration`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigurationRequest struct via the builder pattern


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

