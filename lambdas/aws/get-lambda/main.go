package main

import (
	"hello-world/shared"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var logger = new(shared.Logger)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	getLambdaHelper := new(getLambdaHelper)

	output, error := getLambdaHelper.procsssRequest(request)

	return output, error

}

func main() {
	lambda.Start(handler)
}
