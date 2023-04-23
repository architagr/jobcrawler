package monitoringsns

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssnssubscriptions"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/constructs-go/constructs/v10"
)

type MonitoringSNSStackProps struct {
	awscdk.StackProps
	Queue awssqs.IQueue
}

func NewMonitoringSNSStack(scope constructs.Construct, id string, props *MonitoringSNSStackProps) (awscdk.Stack, awssns.ITopic) {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)
	topic := awssns.NewTopic(stack, aws.String("MonitoringSns"), &awssns.TopicProps{
		DisplayName: aws.String("Monitoring Queues SNS Topic"),
		TopicName:   aws.String("monitoring-sns-topic"),
	})

	topic.AddSubscription(awssnssubscriptions.NewSqsSubscription(props.Queue, &awssnssubscriptions.SqsSubscriptionProps{}))

	return stack, topic
}
