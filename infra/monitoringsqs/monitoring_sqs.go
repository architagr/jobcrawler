package monitoringsqs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type MonitoringSQSStackProps struct {
	awscdk.StackProps
}

func NewMonitoringSQSStack(scope constructs.Construct, id string, props *MonitoringSQSStackProps) (awscdk.Stack, awssqs.IQueue, awssqs.IQueue) {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}

	stack := awscdk.NewStack(scope, &id, &sprops)

	deadLetterQueue := createDeadLetterQueue(stack, props)

	queue := awssqs.NewQueue(stack, aws.String("MonitoringQueue"), &awssqs.QueueProps{
		QueueName:              aws.String("monitoring-queue"),
		RetentionPeriod:        awscdk.Duration_Days(jsii.Number(1)),
		MaxMessageSizeBytes:    jsii.Number(262144),
		VisibilityTimeout:      awscdk.Duration_Minutes(jsii.Number(6)),
		DeliveryDelay:          awscdk.Duration_Minutes(jsii.Number(15)),
		ReceiveMessageWaitTime: awscdk.Duration_Seconds(jsii.Number(20)),
		Encryption:             awssqs.QueueEncryption_UNENCRYPTED,
		DeadLetterQueue: &awssqs.DeadLetterQueue{
			MaxReceiveCount: aws.Float64(5),
			Queue:           deadLetterQueue,
		},
	})

	return stack, queue, deadLetterQueue
}

func createDeadLetterQueue(stack awscdk.Stack, props *MonitoringSQSStackProps) awssqs.IQueue {
	return awssqs.NewQueue(stack, aws.String("MonitoringDeadLetterQueue"), &awssqs.QueueProps{
		QueueName:  aws.String("monitoring-dead-letter-queue"),
		Encryption: awssqs.QueueEncryption_UNENCRYPTED,
	})
}
