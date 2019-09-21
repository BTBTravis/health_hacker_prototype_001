package chatbot

import (
	"encoding/json"
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lexruntimeservice"
	"github.com/google/uuid"
	"log"
)

func SendToLex(message string, userID string) (string, error)  {
	botAlias, botName := "ZurichBot", "ZurichBot"

	randomId, _ := uuid.NewRandom()
	lexResponse := getLexResponse(botAlias, botName, message, randomId.String())

	if lexResponse.IntentName != nil && *lexResponse.IntentName == "volunteer" {
		return *lexResponse.Message, errors.New("Call volunteer")
	}
	lexResponse2 := getLexResponse(botAlias, botName, message, uuid.New().String())

	return *lexResponse2.Message, nil
}

func getSession() (*session.Session) {
	sess, err := session.NewSessionWithOptions(session.Options{
		Config:            aws.Config{Region: aws.String("eu-west-1")},
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		log.Print("Error while getting session: ", err)
		panic(err)
	}
	return sess
}

func getLexResponse(botAlias string, botName string, inputText string, userID string)(*lexruntimeservice.PostTextOutput){
	lex := lexruntimeservice.New(getSession(), aws.NewConfig().WithRegion("eu-west-1"))
	lexResponse, err := lex.PostText(&lexruntimeservice.PostTextInput{
		BotAlias:  &botAlias,
		BotName:   &botName,
		InputText: &inputText,
		UserId:    &userID,
	})
	if err != nil {
		log.Fatal("Error: Failed to get response from lex", err)
	}

	lexResp, _ := json.Marshal(*lexResponse)
	log.Print("Lex response: ", string(lexResp))

	return lexResponse
}