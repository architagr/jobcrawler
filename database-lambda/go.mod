module database_lambda

go 1.20

require (
	github.com/aws/aws-lambda-go v1.39.1
	github.com/architagr/common-constants v0.0.0-00010101000000-000000000000
	github.com/architagr/common-models v0.0.0-00010101000000-000000000000

)

replace (
	github.com/architagr/common-constants => /Users/architagarwal/code/jobcrawler/common-constants
	github.com/architagr/common-models => /Users/architagarwal/code/jobcrawler/common-models
	github.com/architagr/repository => /Users/architagarwal/code/jobcrawler/repository
)
