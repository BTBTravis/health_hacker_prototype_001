package store

import (
    "log"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type UserData struct {
    UserId   string
    CommonWords  string
}

func AddCommonWords(userId string, words string) {
    log.Println("====== Saving common words to db ======")
    sess := session.Must(session.NewSessionWithOptions(session.Options{
        SharedConfigState: session.SharedConfigEnable,
    }))

    svc := dynamodb.New(sess)

    rawNewRecord := UserData{
        UserId: userId,
        CommonWords: words,
    }

    newRecord, err := dynamodbattribute.MarshalMap(rawNewRecord)
    if err != nil {
        log.Println("Error with dynamo marshal map")
    }

    input := &dynamodb.PutItemInput{
        Item:      newRecord,
        TableName: aws.String("patientData"),
    }

    _, err = svc.PutItem(input)
    if err != nil {
        log.Println("Got error calling PutItem:")
        log.Println(err.Error())
    }
}


