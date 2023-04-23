package databasesns

import (
	"github.com/architagr/common-constants/constants"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssnssubscriptions"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type DatabaseSNSStackProps struct {
	awscdk.StackProps
	Queues map[constants.HostName]awssqs.IQueue
}

func NewDatabaseSNSStack(scope constructs.Construct, id string, props *DatabaseSNSStackProps) (awscdk.Stack, awssns.ITopic) {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)
	topic := awssns.NewTopic(stack, aws.String("DatabaseSns"), &awssns.TopicProps{
		DisplayName: aws.String("Database SNS Topic"),
		TopicName:   aws.String("database-sns-topic"),
	})
	for hostName, scrapperQueue := range props.Queues {
		filter := map[string]awssns.SubscriptionFilter{
			"hostName": awssns.SubscriptionFilter_StringFilter(&awssns.StringConditions{Allowlist: jsii.Strings(string(hostName))}),
		}
		topic.AddSubscription(awssnssubscriptions.NewSqsSubscription(scrapperQueue, &awssnssubscriptions.SqsSubscriptionProps{
			FilterPolicy: &filter,
		}))
	}
	return stack, topic
}
