package main

import (
	"fmt"
	crawlerlambda "infra/crawler_lambda"
	"infra/crawlersns"
	"infra/crawlersqs"
	databaselambda "infra/database_lambda"
	"infra/databasesns"
	"infra/databasesqs"
	orchestrationlambda "infra/orchestration_lambda"
	scrapperlambda "infra/scrapper_lambda"

	"infra/scrappersns"
	"infra/scrappersqs"

	"github.com/architagr/common-constants/constants"
	"github.com/aws/aws-cdk-go/awscdk/v2"

	"github.com/aws/jsii-runtime-go"
)

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	stackProps := env(app)
	//#region create queues
	_, scrapperQueues, scrapperDLQueue := scrappersqs.NewScrapperSqsStack(app, "ScrapperQueues", &scrappersqs.ScrapperSqsStackProps{
		StackProps: stackProps,
		HostNames: map[constants.HostName]float64{
			constants.HostName_Linkedin: float64(1),
			constants.HostName_Indeed:   float64(1),
		},
	})

	_, databaseQueues, databaseDLQueue := databasesqs.NewDatabaseSQSStack(app, "DatabaseQueues", &databasesqs.DatabaseSQSStackProps{
		StackProps: stackProps,
		HostNames: map[constants.HostName]float64{
			constants.HostName_Linkedin: float64(1),
			constants.HostName_Indeed:   float64(1),
		},
	})

	_, crawlerQueues, crawlerDLQueue := crawlersqs.NewCrawlerSQSStack(app, "CrawlerQueues", &crawlersqs.CrawlerSQSStackProps{
		StackProps: stackProps,
		HostNames: map[constants.HostName]float64{
			constants.HostName_Linkedin: float64(1),
			constants.HostName_Indeed:   float64(1),
		},
	})
	//#endregion

	//#region create SNS topics
	_, scrapperTopic := scrappersns.NewScrapperSnsStack(app, "ScrapperTopic", &scrappersns.ScrapperSnsStackProps{
		StackProps: stackProps,
		Queues:     scrapperQueues,
	})

	_, databaseTopic := databasesns.NewDatabaseSNSStack(app, "DatabaseTopic", &databasesns.DatabaseSNSStackProps{
		StackProps: stackProps,
		Queues:     databaseQueues,
	})

	_, crawlerTopic := crawlersns.NewCrawlerSNSStack(app, "CrawlerTopic", &crawlersns.CrawlerSNSStackProps{
		StackProps:    stackProps,
		CrawlerQueues: crawlerQueues,
	})
	//#endregion

	//#region create lambda
	scrapperlambda.NewScrapperLambdaStack(app, "ScrapperLambda", &scrapperlambda.ScrapperLambdaStackProps{
		StackProps:       stackProps,
		Queues:           scrapperQueues,
		DatabaseSNSTopic: databaseTopic,
		DeadLetterQueue:  scrapperDLQueue,
	})

	databaselambda.NewDatabaseLambdaStack(app, "DatabaseLambda", &databaselambda.DatabaseLambdaStackProps{
		StackProps:      stackProps,
		Queues:          databaseQueues,
		DeadLetterQueue: databaseDLQueue,
	})

	crawlerlambda.NewCrawlerLambdaStack(app, "CrawlerLambda", &crawlerlambda.CrawlerLambdaStackProps{
		StackProps:       stackProps,
		ScrapperSNSTopic: scrapperTopic,
		CrawlerQueues:    crawlerQueues,
		DeadLetterQueue:  crawlerDLQueue,
	})

	orchestrationlambda.NewOrchestrationLambdaStack(app, "OrchestrationLambda", &orchestrationlambda.OrchestrationLambdaStackProps{
		StackProps:      stackProps,
		CrawlerSNSTopic: crawlerTopic,
	})
	//#endregion
	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env(app awscdk.App) awscdk.StackProps {
	accountId := fmt.Sprint(app.Node().TryGetContext(jsii.String("ACCOUNT_ID")))
	region := fmt.Sprint(app.Node().TryGetContext(jsii.String("REGION")))
	project := fmt.Sprint(app.Node().TryGetContext(jsii.String("PROJECT")))
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	// return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	return awscdk.StackProps{
		Env: &awscdk.Environment{
			Account: jsii.String(accountId),
			Region:  jsii.String(region),
		},
		Tags: &map[string]*string{
			"project": jsii.String(project),
		},
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
