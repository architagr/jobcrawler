package databasesqs

import (
	"fmt"

	"common-constants/constants"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type DatabaseSQSStackProps struct {
	awscdk.StackProps
	HostNames map[constants.HostName]float64
}

func NewDatabaseSQSStack(scope constructs.Construct, id string, props *DatabaseSQSStackProps) (awscdk.Stack, map[constants.HostName]awssqs.IQueue, awssqs.IQueue) {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}

	stack := awscdk.NewStack(scope, &id, &sprops)
	queues := make(map[constants.HostName]awssqs.IQueue)

	deadLetterQueue := createDeadLetterQueue(stack, props)

	for hostName, delay := range props.HostNames {
		queueId := aws.String(fmt.Sprintf("%sDatabaseQueue", hostName))
		queues[hostName] = awssqs.NewQueue(stack, queueId, &awssqs.QueueProps{
			QueueName:              aws.String(fmt.Sprintf("%s-database-queue", hostName)),
			RetentionPeriod:        awscdk.Duration_Days(jsii.Number(1)),
			MaxMessageSizeBytes:    jsii.Number(262144),
			VisibilityTimeout:      awscdk.Duration_Seconds(jsii.Number(10)),
			DeliveryDelay:          awscdk.Duration_Seconds(jsii.Number(delay)),
			ReceiveMessageWaitTime: awscdk.Duration_Seconds(jsii.Number(20)),
			Encryption:             awssqs.QueueEncryption_UNENCRYPTED,
			DeadLetterQueue: &awssqs.DeadLetterQueue{
				MaxReceiveCount: aws.Float64(5),
				Queue:           deadLetterQueue,
			},
		})
	}

	return stack, queues, deadLetterQueue
}

func createDeadLetterQueue(stack awscdk.Stack, props *DatabaseSQSStackProps) awssqs.IQueue {
	return awssqs.NewQueue(stack, aws.String("DatabaseDeadLetterQueue"), &awssqs.QueueProps{
		QueueName:  aws.String("database-dead-letter-queue"),
		Encryption: awssqs.QueueEncryption_UNENCRYPTED,
	})
}
