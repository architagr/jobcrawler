package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"scrapper/config"

	"github.com/architagr/repository/connection"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/jsii-runtime-go"
)

var conn connection.IConnection
var env *config.Config

func readQueue() {
	timeout := int64(300)
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := sqs.New(sess)

	msgResult, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl:            jsii.String(env.GetScrapperSqsUrl()),
		MaxNumberOfMessages: aws.Int64(1),
		VisibilityTimeout:   &timeout,
	})
	if err != nil {
		log.Panic(err)
	}
	log.Printf("message from queue %+v", msgResult)
}

type MessageBody struct {
	Type      string `json:"Type"`
	MessageId string `json:"MessageId"`
	TopicArn  string `json:"TopicArn"`
	Message   string `json:"Message"`
}

func handler(ctx context.Context, sqsEvent events.SQSEvent) error {
	log.Printf("lambda handler start")
	for _, message := range sqsEvent.Records {
		data := new(MessageBody)
		json.Unmarshal([]byte(message.Body), data)
		fmt.Printf("The message %s for event source %s = %T %s \n", message.MessageId, message.EventSource, data, data)
	}

	return nil
}

func main() {
	log.Printf("lambda start")
	lambda.Start(handler)
}

// func main() {
// 	//setupDB()
// 	env = config.GetConfig()
// 	readQueue()
// 	// notification := new(notification.Notification)
// 	// search := &searchcondition.SearchCondition{
// 	// 	JobTitle: constants.JobTitle_SoftwareEngineer,
// 	// 	LocationInfo: searchcondition.Location{
// 	// 		Country: "United States",
// 	// 		City:    "New York",
// 	// 	},
// 	// 	RoleName:   constants.Role_Engineering,
// 	// 	JobType:    constants.JobType_FullTime,
// 	// 	JobModel:   constants.JobModel_OnSite,
// 	// 	Experience: constants.ExperienceLevel_EntryLevel,
// 	// }

// 	// extractor := extractor.InitLinkedInExtractor(*search, notification)
// 	// jobdetails, _ := extractor.StartExtraction(models.Link{
// 	// 	Url:        "https://www.linkedin.com/jobs/view/3503498587/?alternateChannel=search&refId=VxnMmgkzzsXQ7%2FXqQ0bAXQ%3D%3D&trackingId=jGNGTkBtW0dzFFGYUhoO8g%3D%3D",
// 	// 	RetryCount: 5,
// 	// })

// 	// defer conn.Disconnect()
// 	// doc, err := document.InitDocument[models.JobDetails](conn, "webscrapper", "jobDetails")
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// id, err := doc.AddSingle(jobdetails)
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// log.Printf("new id %+v", id)
// }

func setupDB() {
	conn = connection.InitConnection("mongodb+srv://webscrapper:WebScrapper123@cluster0.xzvihm7.mongodb.net/?retryWrites=true&w=majority", 10)
	err := conn.ValidateConnection()
	if err != nil {
		log.Fatalf("error in conncting to mongo %+v", err)
	}

}
