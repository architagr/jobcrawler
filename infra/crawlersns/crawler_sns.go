package crawlersns

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

type CrawlerSNSStackProps struct {
	awscdk.StackProps
	CrawlerQueues map[constants.HostName]awssqs.IQueue
}

func NewCrawlerSNSStack(scope constructs.Construct, id string, props *CrawlerSNSStackProps) (awscdk.Stack, awssns.ITopic) {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)
	topic := awssns.NewTopic(stack, aws.String("CrawlerSns"), &awssns.TopicProps{
		DisplayName: aws.String("Crawler SNS Topic"),
		TopicName:   aws.String("crawler-sns-topic"),
	})
	for hostName, orchestrationQueue := range props.CrawlerQueues {
		filter := map[string]awssns.SubscriptionFilter{
			"hostName": awssns.SubscriptionFilter_StringFilter(&awssns.StringConditions{Allowlist: jsii.Strings(string(hostName))}),
		}
		topic.AddSubscription(awssnssubscriptions.NewSqsSubscription(orchestrationQueue, &awssnssubscriptions.SqsSubscriptionProps{
			FilterPolicy: &filter,
		}))
	}
	return stack, topic
}
