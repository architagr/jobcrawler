build:
	make -C ../scrapper lambda-build
	make -C ../database-lambda lambda-build

clear:
	make -C ../scrapper clear
	make -C ../database-lambda clear

bootstrap:
	cdk bootstrap

deploy:
	make build
	rm -rf cdk.out
	clear
	cdk deploy --all --require-approval never
	make clear

destroy:
	cdk destroy --all --require-approval never