package scrappersns

import (
	"common-constants/constants"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssnssubscriptions"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type ScrapperSnsStackProps struct {
	awscdk.StackProps
	Queues map[constants.HostName]awssqs.IQueue
}

func NewScrapperSnsStack(scope constructs.Construct, id string, props *ScrapperSnsStackProps) (awscdk.Stack, awssns.ITopic) {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)
	topic := awssns.NewTopic(stack, aws.String("ScrapperSns"), &awssns.TopicProps{
		DisplayName: aws.String("Scrapper SNS Topic"),
		TopicName:   aws.String("scrapper-sns-topic"),
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
