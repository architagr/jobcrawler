build:
	make -C ../scrapper lambda-build

bootstrap:
	cdk bootstrap

deploy:
	make build
	rm -rf cdk.out
	clear
	cdk deploy --all --require-approval never