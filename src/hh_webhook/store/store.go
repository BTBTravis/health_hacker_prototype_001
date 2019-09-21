package store

import (
    "log"
    "time"
    "strconv"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type UserData struct {
    UserId  string `type:"string"`
    Message string `type:"string"`
    TimeStamp int64
}

var storeTable = aws.String("patientData")

func AddTranscriptRecord(userId string, msg string) {
    sess := session.Must(session.NewSessionWithOptions(session.Options{
        SharedConfigState: session.SharedConfigEnable,
    }))

    svc := dynamodb.New(sess)

    nowInt := time.Now().Unix()
    nowStr := strconv.FormatInt(nowInt, 10)
    rawNewRecord := UserData{
        UserId: userId + "-" + nowStr,
        Message: msg,
        TimeStamp: nowInt,
    }

    newRecord, err := dynamodbattribute.MarshalMap(rawNewRecord)
    if err != nil {
        log.Println("Error with dynamo marshal map")
    }

    input := &dynamodb.PutItemInput{
        Item:      newRecord,
        TableName: storeTable,
    }

    _, err = svc.PutItem(input)
    if err != nil {
        log.Println("Got error calling PutItem:")
        log.Println(err.Error())
    }
}

