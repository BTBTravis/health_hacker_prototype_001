package main

import (
	"context"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

func Handler(context context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("1. Processing Lambda request %s\n", request.RequestContext.RequestID)

	// If no query is provided in the HTTP request body, throw an error
	if len(request.Body) < 1 {
		return events.APIGatewayProxyResponse{}, errors.New("Error: no query was provided in the HTTP body")
	}

	log.Printf("Request: ", request.Body)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}

func init() {
	log.Print("Runnning initialization...")
}

func main() {
	log.Print("main function...")
	lambda.Start(Handler)
}
