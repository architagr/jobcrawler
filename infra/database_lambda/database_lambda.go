package databaselambda

import (
	"fmt"

	"github.com/architagr/common-constants/constants"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"

	lambdaEvent "github.com/aws/aws-cdk-go/awscdk/v2/awslambdaeventsources"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type DatabaseLambdaStackProps struct {
	awscdk.StackProps
	Queues map[constants.HostName]awssqs.Queue
}

func NewDatabaseLambdaStack(scope constructs.Construct, id string, props *DatabaseLambdaStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	for hostName, scrapperQueue := range props.Queues {
		env := make(map[string]*string)
		env["DatabaseSqsUrl"] = scrapperQueue.QueueUrl()
		env["DbConnectionString"] = jsii.String("mongodb+srv://webscrapper:WebScrapper123@cluster0.xzvihm7.mongodb.net/?retryWrites=true&w=majority")

		lambdaFunction := awslambda.NewFunction(stack, jsii.String(fmt.Sprintf("%sDatabaseLambda", hostName)), &awslambda.FunctionProps{
			Environment:  &env,
			Runtime:      awslambda.Runtime_GO_1_X(),
			Handler:      jsii.String("database-lambda"),
			Code:         awslambda.Code_FromAsset(jsii.String("./../database-lambda/main.zip"), &awss3assets.AssetOptions{}),
			FunctionName: jsii.String(fmt.Sprintf("%s-database-lambda-fn", hostName)),
		})
		scrapperQueue.GrantConsumeMessages(lambdaFunction)
		triggerEvent := lambdaEvent.NewSqsEventSource(scrapperQueue, &lambdaEvent.SqsEventSourceProps{
			BatchSize:         jsii.Number(10),
			MaxBatchingWindow: awscdk.Duration_Seconds(jsii.Number(1)),
			MaxConcurrency:    jsii.Number(5),
			Enabled:           jsii.Bool(true),
		})

		lambdaFunction.AddEventSource(triggerEvent)
	}
	return stack
}
