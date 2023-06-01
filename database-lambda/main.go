package main

import (
	"common-constants/constants"
	"context"
	"database_lambda/config"
	"database_lambda/persistance"
	"encoding/json"
	"fmt"
	"log"
	"mongodbRepo/collection"
	"mongodbRepo/connection"
	"sync"
	"time"

	jobdetails "common-models/job-details"
	notificationModel "common-models/sns-notification"

	sqs_message "common-models/sqs-message"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var persistanceObj map[constants.DatabaseType]persistance.IPersistance
var connObj map[constants.DatabaseType]connection.IConnection

var env config.IConfig

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
	handle(jobs)
	return nil
}
func handle(jobs []jobdetails.JobDetails) {
	wg := sync.WaitGroup{}
	wg.Add(len(persistanceObj))
	for _, c := range persistanceObj {
		go saveMany(c, jobs, &wg)
	}
	wg.Wait()
}
func saveMany(persist persistance.IPersistance, jobs []jobdetails.JobDetails, wg *sync.WaitGroup) {
	defer wg.Done()
	_, err := persist.SaveMany(jobs)
	if err != nil {
		panic(err)
	}
}

func main() {
	log.Printf("lambda start")
	env = config.GetConfig()
	persistanceObj = make(map[constants.DatabaseType]persistance.IPersistance)
	connObj = make(map[constants.DatabaseType]connection.IConnection)

	setupPersistanceConnections()
	for _, c := range connObj {
		defer c.Disconnect()
	}
	if env.IsLocal() {
		time.Sleep(1 * time.Second)
		handle([]jobdetails.JobDetails{
			{
				Title: "test",
			},
		})
	} else {
		lambda.Start(handler)
	}
}

func setupPersistanceConnections() {
	setupMongodbConnection()
}

func setupMongodbConnection() {
	mongodbConnection := setupMongoDbDB()
	doc, err := collection.InitCollection[jobdetails.JobDetails](mongodbConnection, env.GetDatabaseName(), env.GetCollectionName())
	if err != nil {
		log.Panic(err)
	}
	connObj[constants.DatabaseType_Mongodb] = mongodbConnection
	persistanceObj[constants.DatabaseType_Mongodb] = persistance.InitMongoDbPersistance(env, doc)
}

func setupMongoDbDB() connection.IConnection {
	conn := connection.InitConnection(env.GetDatabaseConnectionString(), 10)
	err := conn.ValidateConnection()
	if err != nil {
		log.Panicf("error in conncting to mongo %+v", err)
	}

	return conn
}
