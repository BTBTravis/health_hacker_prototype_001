package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"main/sender"
	"time"
)

var capabilities = []string{
	"I can learn how to do bla bla",
	"I can learn how to do bla bla1",
	"I can learn how to do bla bla2",
	"I can learn how to do bla bla3",
	"I can learn how to do bla bla4",
	"I can learn how to do bla bla5",
	"I can learn how to do bla bla6",
	"I can learn how to do bla bla7",
	"I can learn how to do bla bla8",
	"I can learn how to do bla bla9",
	"I can learn how to do bla bla0",
	"I can learn how to do bla bla1",
	"I can learn how to do bla bla2",
	"I can learn how to do bla bla12",
	"I can learn how to do bla bla12",
	"I can learn how to do bla bla22",
	"I can learn how to do bla bladf",
	"I can learn how to do bla bladf",
	"I can learn how to do bla bla324",
	"I can learn how to do bla bladg",
	"I can learn how to do bla bla234",
	"I can learn how to do bla bla23d",
	"I can learn how to do bla bla234",
	"I can learn how to do bla bla23rf",
	"I can learn how to do bla bla23",
	"I can learn how to do bla bla23r",
}

func Handler(context context.Context, request events.APIGatewayProxyRequest) error {
	if len(request.Body) < 1 {
		log.Fatal("Body is empty")
	}

	for _, v := range capabilities {
		sender.SendMessage(request.Body, v)
		time.Sleep(5 * time.Second)
	}

	return nil
}

func init() {
	log.Print("Runnning initialization...")
}

func main() {
	log.Print("main function...")
	lambda.Start(Handler)
}
