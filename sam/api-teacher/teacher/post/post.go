package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	headers := make(map[string]string)
	headers["Access-Control-Allow-Origin"] = "*"

	return events.APIGatewayProxyResponse{
		Headers:headers,
		Body: "ok " , StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}

