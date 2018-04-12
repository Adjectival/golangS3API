// Deploy it
stackName := spartaCF.UserScopedStackName("SpartaHTML")
sparta.Main(stackName,
  fmt.Sprintf("SpartaHTML provisions a static S3 hosted website with an API Gateway resource backed by a custom Lambda function"),
  spartaHTMLLambdaFunctions(apiGateway),
  apiGateway,
  s3Site)