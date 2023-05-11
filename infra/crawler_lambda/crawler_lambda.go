package crawlerlambda

import (
	"fmt"

	"github.com/architagr/common-constants/constants"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"

	lambdaEvent "github.com/aws/aws-cdk-go/awscdk/v2/awslambdaeventsources"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type CrawlerLambdaStackProps struct {
	awscdk.StackProps
	ScrapperSNSTopic awssns.ITopic
	CrawlerQueues    map[constants.HostName]awssqs.IQueue
	DeadLetterQueue  awssqs.IQueue
}

func NewCrawlerLambdaStack(scope constructs.Construct, id string, props *CrawlerLambdaStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)
	for hostName, crawlerQueue := range props.CrawlerQueues {
		env := make(map[string]*string)
		env["ScrapperSnsTopicArn"] = props.ScrapperSNSTopic.TopicArn()

		lambdaFunction := awslambda.NewFunction(stack, jsii.String(fmt.Sprintf("%sCrawlerLambda", hostName)), &awslambda.FunctionProps{
			Environment:  &env,
			Runtime:      awslambda.Runtime_GO_1_X(),
			Handler:      jsii.String("webcrawler"),
			Code:         awslambda.Code_FromAsset(jsii.String("./../jobcrawler/main.zip"), &awss3assets.AssetOptions{}),
			FunctionName: jsii.String(fmt.Sprintf("%s-crawler-lambda-fn", hostName)),
			Timeout:      awscdk.Duration_Seconds(jsii.Number(5)),
		})
		crawlerQueue.GrantConsumeMessages(lambdaFunction)
		props.DeadLetterQueue.GrantSendMessages(lambdaFunction)
		props.ScrapperSNSTopic.GrantPublish(lambdaFunction)

		triggerEvent := lambdaEvent.NewSqsEventSource(crawlerQueue, &lambdaEvent.SqsEventSourceProps{
			BatchSize:         jsii.Number(1),
			MaxBatchingWindow: awscdk.Duration_Millis(jsii.Number(0)),
			MaxConcurrency:    jsii.Number(5),
			Enabled:           jsii.Bool(true),
		})

		lambdaFunction.AddEventSource(triggerEvent)
	}
	return stack
}
