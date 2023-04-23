package switchtablelambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"

	lambdaEvent "github.com/aws/aws-cdk-go/awscdk/v2/awslambdaeventsources"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type SwitchLambdaLambdaStackProps struct {
	awscdk.StackProps
	DeadLetterQueue awssqs.IQueue
}

func NewSwitchTableLambdaStack(scope constructs.Construct, id string, props *SwitchLambdaLambdaStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	env := make(map[string]*string)
	env["DbConnectionString"] = jsii.String("mongodb+srv://webscrapper:WebScrapper123@cluster0.xzvihm7.mongodb.net/?retryWrites=true&w=majority")
	env["DatabaseName"] = jsii.String("webscrapper")
	env["TempCollectionName"] = jsii.String("jobDetailsTemp")
	env["FinalCollectionName"] = jsii.String("jobDetails")

	lambdaFunction := awslambda.NewFunction(stack, jsii.String("SwitchLambdaLambda"), &awslambda.FunctionProps{
		Environment:  &env,
		Runtime:      awslambda.Runtime_GO_1_X(),
		Handler:      jsii.String("switchtable"),
		Code:         awslambda.Code_FromAsset(jsii.String("./../switch-table-lambda/main.zip"), &awss3assets.AssetOptions{}),
		FunctionName: jsii.String("switch-table-lambda-fn"),
		Timeout:      awscdk.Duration_Minutes(jsii.Number(5)),
	})
	props.DeadLetterQueue.GrantSendMessages(lambdaFunction)

	triggerEvent := lambdaEvent.NewSqsEventSource(props.DeadLetterQueue, &lambdaEvent.SqsEventSourceProps{
		BatchSize:         jsii.Number(1),
		MaxBatchingWindow: awscdk.Duration_Seconds(jsii.Number(1)),
		MaxConcurrency:    jsii.Number(2),
		Enabled:           jsii.Bool(true),
	})

	lambdaFunction.AddEventSource(triggerEvent)
	return stack
}
