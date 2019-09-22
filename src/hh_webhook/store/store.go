package store

import (
    "log"
    //"time"
    //"strconv"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type UserData struct {
    UserId  string `type:"string"`
    Mode string `type:"string"`
}

var transcriptTable = aws.String("patientTranscripts")
var modeTable = aws.String("patientMode")

type Item struct {
    Mode string
}

func GetMode(userId string) string {
    record := dbGet(modeTable, dynamodb.GetItemInput{
	TableName: modeTable,
	Key: map[string]*dynamodb.AttributeValue{
	    "UserId": {
		S: aws.String(userId),
	    },
	},
    })


    return record.Mode
}

func dbGet(table *string, getItemInput dynamodb.GetItemInput) Item {
    sess := session.Must(session.NewSessionWithOptions(session.Options{
	SharedConfigState: session.SharedConfigEnable,
    }))

    svc := dynamodb.New(sess)
    result, err := svc.GetItem(&getItemInput)
    if err != nil {
	log.Println("Error with dynamo get")
    }

    item := Item{}
    err = dynamodbattribute.UnmarshalMap(result.Item, &item)
    if err != nil {
        log.Println("Failed to unmarshal Record while getting record")
    }
    return item
}

func dbPut(table string, record map[string]*dynamodb.AttributeValue) {
    sess := session.Must(session.NewSessionWithOptions(session.Options{
	SharedConfigState: session.SharedConfigEnable,
    }))

    svc := dynamodb.New(sess)

    input := &dynamodb.PutItemInput{
	Item:      record,
	TableName: &table,
    }

    var _, err = svc.PutItem(input)
    if err != nil {
	log.Println("Got error calling PutItem:")
	log.Println(err.Error())
    }
}

func PutMode(userId string, mode string) {
    rawRecord := UserData{
        UserId: userId,
        Mode: mode,
    }
    newRecord, err := dynamodbattribute.MarshalMap(rawRecord)
    if err != nil {
	log.Println("Error with dynamo marshal map")
    }
    dbPut(*modeTable, newRecord)
}

//func PutTranscript(userId string, message string) {
    //nowInt := time.Now().Unix()
    //nowStr := strconv.FormatInt(nowInt, 10)
    //rawNewRecord := UserData{
        //UserId: userId + "-" + nowStr,
        //Message: msg,
        //TimeStamp: nowInt,
    //}

    //newRecord, err := dynamodbattribute.MarshalMap(rawRecord)
    //dbPut(transcriptTable, newRecord)
//}
