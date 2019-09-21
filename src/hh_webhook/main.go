package main

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

type UrlVerification struct{
	Token string	`json:"token"`
	Challenge string `json:"challenge"`
	Type string `json:"type"`
}

func Handler(context context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("1. Processing Lambda request %s\n", request.RequestContext.RequestID)

	// If no query is provided in the HTTP request body, throw an error
	if len(request.Body) < 1 {
		return events.APIGatewayProxyResponse{}, errors.New("Error: no query was provided in the HTTP body")
	}

	var v UrlVerification
	if err := json.Unmarshal([]byte(request.Body), &v); err != nil {
		log.Print("Error: Could not decode body!", err)
	}

	log.Printf("Request: ", request.Body)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body: v.Challenge,
	}, nil
}

func init() {
	log.Print("Runnning initialization...")
}

func main() {
	log.Print("main function...")
	lambda.Start(Handler)
}
