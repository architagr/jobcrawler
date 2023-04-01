package scrapperlambda

import (
	"fmt"

	"github.com/architagr/common-constants/constants"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"

	lambdaEvent "github.com/aws/aws-cdk-go/awscdk/v2/awslambdaeventsources"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type ScrapperLambdaStackProps struct {
	awscdk.StackProps
	Queues           map[constants.HostName]awssqs.Queue
	DatabaseSNSTopic *awssns.Topic
}

func NewScrapperLambdaStack(scope constructs.Construct, id string, props *ScrapperLambdaStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	for hostName, scrapperQueue := range props.Queues {
		env := make(map[string]*string)
		env["DatabaseSNSTopicArn"] = (*props.DatabaseSNSTopic).TopicArn()

		lambdaFunction := awslambda.NewFunction(stack, jsii.String(fmt.Sprintf("%sScrapperLambda", hostName)), &awslambda.FunctionProps{
			Environment:  &env,
			Runtime:      awslambda.Runtime_GO_1_X(),
			Handler:      jsii.String("scrapper"),
			Code:         awslambda.Code_FromAsset(jsii.String("./../scrapper/main.zip"), &awss3assets.AssetOptions{}),
			FunctionName: jsii.String(fmt.Sprintf("%s-scrapper-lambda-fn", hostName)),
		})
		scrapperQueue.GrantConsumeMessages(lambdaFunction)
		(*props.DatabaseSNSTopic).GrantPublish(lambdaFunction)

		triggerEvent := lambdaEvent.NewSqsEventSource(scrapperQueue, &lambdaEvent.SqsEventSourceProps{
			BatchSize:         jsii.Number(1),
			MaxBatchingWindow: awscdk.Duration_Millis(jsii.Number(0)),
			MaxConcurrency:    jsii.Number(5),
			Enabled:           jsii.Bool(true),
		})

		lambdaFunction.AddEventSource(triggerEvent)
	}
	return stack
}
