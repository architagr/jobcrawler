accountId := $(shell aws sts get-caller-identity --query Account --output text --profile default)
region := $(AWS_REGION)
project := job_crawler
env := $(ENV)
ifndef region
	region = ap-southeast-1
endif

ifndef env
	env = dev
endif

build:
	make -C ../orchestration-lambda lambda-build
	make -C ../jobcrawler lambda-build
	make -C ../scrapper lambda-build
	make -C ../database-lambda lambda-build
	make -C ../monitor-queues-lambda lambda-build
	make -C ../switch-table-lambda lambda-build

clear:
	make -C ../orchestration-lambda clear
	make -C ../jobcrawler clear
	make -C ../scrapper clear
	make -C ../database-lambda clear
	make -C ../monitor-queues-lambda clear
	make -C ../switch-table-lambda clear

bootstrap:
	cdk bootstrap

deploy:
	make build
	rm -rf cdk.out
	clear
	cdk deploy --all --require-approval never \
	--context ACCOUNT_ID=$(accountId) \
	--context REGION=$(region) \
	--context PROJECT=$(project) \
	--context ENV=$(env)
	make clear

destroy:
	cdk destroy --all --require-approval never