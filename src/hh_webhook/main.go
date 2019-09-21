package main

import (
	"context"
        "regexp"
        "main/sender"
	"encoding/json"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

type UrlVerification struct{
    Token string    `json:"token"`
    Challenge string `json:"challenge"`
    Type string `json:"type"`
    Event SlackEvent `json:"event"`
}

type SlackEvent struct {
    ClientMsgId string `json:"client_msg_id"`
    Type string `json:"type"`
    Text string `json:"text"`
    User string `json:"user"`
    Ts string `json:"ts"`
    Team string `json:"team"`
    Channel string `json:"channel"`
    EventTs string `json:"event_ts"`
    ChannelType string `json:"channel_type"`
}

type SlackUser struct{
    ChannelId string
    UserId string
}

var Travis = SlackUser{
    ChannelId: "DN7L5L8AW",
    UserId: "UMYDK4Q66",
}

var Alex = SlackUser{
    ChannelId: "DNL27EDE0",
    UserId: "UND2UQZB9",
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

	if v.Event.User == Travis.UserId {
	    sender.SendMessage(Alex.ChannelId, v.Event.Text)
	} else if v.Event.User == Alex.UserId {
            ogMsg := v.Event.Text
            jumper, _ := regexp.MatchString("jump", ogMsg)

            if jumper {
                sender.SendMessage(Alex.ChannelId, "don't do it call: XXX-XXXX-XXXX")
                sender.SendMessage(Travis.ChannelId, "WARNING: alex is on the edge")
            } else {
                sender.SendMessage(Travis.ChannelId, ogMsg)
            }
	}

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
