package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"hello-world/shared"
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
