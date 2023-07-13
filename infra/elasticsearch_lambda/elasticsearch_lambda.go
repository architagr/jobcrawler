package elasticsearchlambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type ElasticSearchLambdaStackProps struct {
	awscdk.StackProps
	DeadLetterQueue awssqs.IQueue
}

func NewElasticSearchLambdaStack(scope constructs.Construct, id string, props *ElasticSearchLambdaStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	env := make(map[string]*string)

	awslambda.NewFunction(stack, jsii.String("ElasticSearchLambda"), &awslambda.FunctionProps{
		Environment:  &env,
		Runtime:      awslambda.Runtime_GO_1_X(),
		Handler:      jsii.String("elasticsearch-lambda"),
		Code:         awslambda.Code_FromAsset(jsii.String("./../elasticsearchservice/main.zip"), &awss3assets.AssetOptions{}),
		FunctionName: jsii.String("elasticsearch-lambda-fn"),
		Timeout:      awscdk.Duration_Minutes(jsii.Number(5)),
	})
	return stack
}
