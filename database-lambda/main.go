package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	sqs_message "github.com/architagr/common-models/sqs-message"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, sqsEvent events.SQSEvent) error {
	log.Printf("lambda handler start")
	for _, message := range sqsEvent.Records {
		data := new(sqs_message.MessageBody)
		json.Unmarshal([]byte(message.Body), data)
		fmt.Printf("The message %s for event source %s, data: %+v \n", message.MessageId, message.EventSource, data)
	}

	return nil
}

func main() {
	log.Printf("lambda start")
	lambda.Start(handler)
}
