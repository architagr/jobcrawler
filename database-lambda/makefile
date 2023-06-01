lambda-build:
	make clear
	GOOS=linux GOARCH=amd64 go build -o database-lambda main.go 
	zip main.zip database-lambda
	rm -f database-lambda

clear:
	rm -f scrapper main.zip
	
test:
	go test `go list ./... | grep -v examples` -coverprofile=../database_coverage.txt -covermode=atomic