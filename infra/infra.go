package main

import (
	"infra/scrappersns"
	"infra/scrappersqs"

	"github.com/architagr/common-constants/constants"
	"github.com/aws/aws-cdk-go/awscdk/v2"

	"github.com/aws/jsii-runtime-go"
)

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	_, queues := scrappersqs.NewScrapperSqsStack(app, "ScrapperQueues", &scrappersqs.ScrapperSqsStackProps{
		StackProps: awscdk.StackProps{
			Env: env(),
		},
		HostNames: map[constants.HostName]float64{
			constants.HostName_Linkedin: float64(1),
			constants.HostName_Indeed:   float64(1),
		},
	})
	scrappersns.NewScrapperSnsStack(app, "ScrapperTopic", &scrappersns.ScrapperSnsStackProps{
		StackProps: awscdk.StackProps{
			Env: env(),
		},
		Queues: queues,
	})
	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	// return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	return &awscdk.Environment{
		Account: jsii.String("638580160310"),
		Region:  jsii.String("ap-southeast-1"),
	}

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}
