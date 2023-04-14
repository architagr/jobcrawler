package orchestrationlambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"

	"github.com/aws/aws-cdk-go/awscdk/v2/awseventstargets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type OrchestrationLambdaStackProps struct {
	awscdk.StackProps
	CrawlerSNSTopic *awssns.Topic
}

func NewOrchestrationLambdaStack(scope constructs.Construct, id string, props *OrchestrationLambdaStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)
	env := make(map[string]*string)
	env["CrawlerSNSTopicArn"] = (*props.CrawlerSNSTopic).TopicArn()
	env["DbConnectionString"] = jsii.String("mongodb+srv://webscrapper:WebScrapper123@cluster0.xzvihm7.mongodb.net/?retryWrites=true&w=majority")
	env["DatabaseName"] = jsii.String("webscrapper")
	env["CollectionName"] = jsii.String("jobLinks")

	lambdaFunction := awslambda.NewFunction(stack, jsii.String("OrchestrationLambda"), &awslambda.FunctionProps{
		Environment:  &env,
		Runtime:      awslambda.Runtime_GO_1_X(),
		Handler:      jsii.String("orchestration"),
		Code:         awslambda.Code_FromAsset(jsii.String("./../orchestration-lambda/main.zip"), &awss3assets.AssetOptions{}),
		FunctionName: jsii.String("orchestration-lambda-fn"),
		Timeout:      awscdk.Duration_Seconds(jsii.Number(300)),
	})
	(*props.CrawlerSNSTopic).GrantPublish(lambdaFunction)

	eventRule := awsevents.NewRule(stack, aws.String("TriggerOrchestrationLambdaEvent"), &awsevents.RuleProps{
		Schedule: awsevents.Schedule_Cron(&awsevents.CronOptions{Hour: aws.String("0"), Minute: aws.String("0")}),
	})
	eventRule.AddTarget(awseventstargets.NewLambdaFunction(lambdaFunction, nil))

	return stack
}
