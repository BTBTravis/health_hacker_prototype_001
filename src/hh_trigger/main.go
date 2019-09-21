package main

import (
	"context"
	"main/sender"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	"log"
)

func Handler(context context.Context, request events.APIGatewayProxyRequest) error {
	if len(request.Body) < 1 {
		log.Fatal("Body is empty")
	}

	sender.SendMessage(request.Body, "Hey, how are you doing?")

	return nil
}

func init() {
	log.Print("Runnning initialization...")
}

func main() {
	log.Print("main function...")
	lambda.Start(Handler)
}
