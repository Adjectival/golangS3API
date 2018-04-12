if nil != api {
  apiGatewayResource, _ := api.NewResource("/hello", lambdaFn)

  // We only return http.StatusOK
  apiMethod, apiMethodErr := apiGatewayResource.NewMethod("GET",
    http.StatusOK,
    http.StatusOK)
  if nil != apiMethodErr {
    panic("Failed to create /hello resource: " + apiMethodErr.Error())
  }
  // The lambda resource only supports application/json Unmarshallable
  // requests.
  apiMethod.SupportedRequestContentTypes = []string{"application/json"}
}
return append(lambdaFunctions, lambdaFn)