type helloWorldResponse struct {
  Message string
  Request spartaAPIG.APIGatewayRequest
}

////////////////////////////////////////////////////////////////////////////////
// Hello world event handler
func helloWorld(ctx context.Context,
  gatewayEvent spartaAPIG.APIGatewayRequest) (helloWorldResponse, error) {

  logger, loggerOk := ctx.Value(sparta.ContextKeyLogger).(*logrus.Logger)
  if loggerOk {
    logger.Info("Hello world structured log message")
  }
  // Return a message, together with the incoming input...
  return helloWorldResponse{
    Message: fmt.Sprintf('Hello world 🌏'),
    Request: gatewayEvent,
  }, nil
}