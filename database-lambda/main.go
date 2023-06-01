package main

import (
	"context"
	"database_lambda/config"
	"encoding/json"
	"fmt"
	"log"

	jobdetails "common-models/job-details"
	notificationModel "common-models/sns-notification"
	"repository/collection"
	"repository/connection"

	sqs_message "common-models/sqs-message"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var conn connection.IConnection
var env *config.Config

func handler(ctx context.Context, sqsEvent events.SQSEvent) error {
	log.Printf("lambda handler start")
	jobs := make([]jobdetails.JobDetails, 0)
	for _, message := range sqsEvent.Records {
		data := new(sqs_message.MessageBody)
		json.Unmarshal([]byte(message.Body), data)
		messageContent := new(notificationModel.Notification[jobdetails.JobDetails])
		json.Unmarshal([]byte(data.Message), messageContent)
		jobs = append(jobs, messageContent.Data)
		fmt.Printf("The message %s for event source %s, messageContent: %+v \n", message.MessageId, message.EventSource, messageContent)
	}

	doc, err := collection.InitCollection[jobdetails.JobDetails](conn, env.GetDatabaseName(), env.GetCollectionName())
	if err != nil {
		log.Panic(err)
	}
	_, err = doc.AddMany(jobs)
	if err != nil {
		log.Panic(err)
	}
	return nil
}

func main() {
	log.Printf("lambda start")
	env = config.GetConfig()
	setupDB()
	defer conn.Disconnect()
	lambda.Start(handler)
}

func setupDB() {
	conn = connection.InitConnection(env.GetDatabaseConnectionString(), 10)
	err := conn.ValidateConnection()
	if err != nil {
		log.Fatalf("error in conncting to mongo %+v", err)
	}
}
