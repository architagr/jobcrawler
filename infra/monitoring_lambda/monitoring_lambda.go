package monitoringlambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"

	lambdaEvent "github.com/aws/aws-cdk-go/awscdk/v2/awslambdaeventsources"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type MonitoringLambdaStackProps struct {
	awscdk.StackProps
	MonitoringSNSTopic awssns.ITopic
	MonitoringQueue    awssqs.IQueue
	DeadLetterQueue    awssqs.IQueue
}

func NewMonitoringLambdaStack(scope constructs.Construct, id string, props *MonitoringLambdaStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	env := make(map[string]*string)
	env["MonitoringSnsTopicArn"] = props.MonitoringSNSTopic.TopicArn()

	lambdaFunction := awslambda.NewFunction(stack, jsii.String("MonitoringLambda"), &awslambda.FunctionProps{
		Environment:  &env,
		Runtime:      awslambda.Runtime_GO_1_X(),
		Handler:      jsii.String("monitorqueue"),
		Code:         awslambda.Code_FromAsset(jsii.String("./../monitor-queues-lambda/main.zip"), &awss3assets.AssetOptions{}),
		FunctionName: jsii.String("monitor-queues-lambda-fn"),
		Timeout:      awscdk.Duration_Minutes(jsii.Number(5)),
	})
	props.MonitoringQueue.GrantConsumeMessages(lambdaFunction)
	props.DeadLetterQueue.GrantSendMessages(lambdaFunction)
	props.MonitoringSNSTopic.GrantPublish(lambdaFunction)

	triggerEvent := lambdaEvent.NewSqsEventSource(props.MonitoringQueue, &lambdaEvent.SqsEventSourceProps{
		BatchSize:         jsii.Number(1),
		MaxBatchingWindow: awscdk.Duration_Seconds(jsii.Number(1)),
		MaxConcurrency:    jsii.Number(2),
		Enabled:           jsii.Bool(true),
	})

	lambdaFunction.AddEventSource(triggerEvent)
	return stack
}
