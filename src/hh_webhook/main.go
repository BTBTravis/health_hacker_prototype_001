package main

import (
	"context"
        "main/sender"
        "main/store"
	"main/chatbot"
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
    Name string
    ChannelId string
    UserId string
}

var Travis = SlackUser{
    Name: "travis",
    ChannelId: "DN7L5L8AW",
    UserId: "UMYDK4Q66",
}

var Alex = SlackUser{
    Name: "alex",
    ChannelId: "DNLGGKBGD",
    UserId: "UND2UQZB9",
}

var Peter = SlackUser{
    Name: "peter",
    ChannelId: "DN95RPK6F",
    UserId: "UNNRQT6BY",
}

var Patient = Peter
var Volunteer = Travis

const (
  MODE_AI = "ai"
  MODE_HUMAN = "human"
)

func isPatientByUserId(userId string) bool {
    return userId == Patient.UserId
}

func isVolunteerByUserId(userId string) bool {
    return userId == Volunteer.UserId
}

func getPatient() SlackUser {
    return Patient
}

func getVolunteer() SlackUser {
    return Volunteer
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

	r, _ := json.Marshal(v)
	log.Printf("Request: ", string(r))

	userId := v.Event.User

	isPatient := isPatientByUserId(userId)
	isVolunteer := isVolunteerByUserId(userId)

	if isVolunteer {
		sender.SendMessage(getPatient().ChannelId, v.Event.Text)
	} else if isPatient {
	    currentMode := store.GetMode(userId)
	    ogMsg := v.Event.Text
	    if currentMode == MODE_HUMAN {
		sender.SendMessage(getVolunteer().ChannelId, ogMsg)
	    } else {
		//ogMsg := v.Event.Text
		//jumper, _ := regexp.MatchString("jump", ogMsg)
			    //
		//if jumper {
		//    sender.SendMessage(getPatient().ChannelId, "don't do it call: XXX-XXXX-XXXX")
		//    sender.SendMessage(getVolunteer().ChannelId, "WARNING: alex is on the edge")
		//} else {
		//    sender.SendMessage(getVolunteer().ChannelId, ogMsg)
		//}
		replyMessage, err := chatbot.SendToLex(v.Event.Text, getPatient().UserId)
		sender.SendMessage(getPatient().ChannelId, replyMessage)
		if err != nil {
		    // TODO: send chat recap to volunteer
		    store.PutMode(userId, MODE_HUMAN)
		    sender.SendMessage(getVolunteer().ChannelId, v.Event.Text)
		}
	    }
	}
	//if v.Type == "url_verification" {
	//	return events.APIGatewayProxyResponse{
	//		StatusCode: 200,
	//		Body:       v.Challenge,
	//	}, nil
	//} else if v.Type == "event_callback" {
	//	sender.SendMessage("Hello there cowboy")
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       v.Challenge,
	}, nil
}

func init() {
	log.Print("Runnning initialization...")
}

func main() {
	log.Print("main function...")
	lambda.Start(Handler)
}
