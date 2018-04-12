// filename: cmd/main.go
package main
import (
  "io"
  "net/http"
  "os"
  sparta "github.com/mweagle/Sparta"
)
// AWS Lambda Go style
func helloWorld() (string, error) {
  return "Hello World", nil
}
////////////////////////////////////////////////////////////////////////////////
// Main
func main() {
  lambdaFn := sparta.HandleAWSLambda('Hello World',
    helloWorld,
    sparta.IAMRoleDefinition{})
  var lambdaFunctions []*sparta.LambdaAWSInfo
  lambdaFunctions = append(lambdaFunctions, lambdaFn)
    err := sparta.Main('Hello World',
    'Hello World',
    lambdaFunctions,
    nil,
    nil)
  if err != nil {
    os.Exit(1)
  }
}
