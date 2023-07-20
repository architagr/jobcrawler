lambda-build:
	make clear
	GOOS=linux GOARCH=amd64 go build -o elasticsearch-lambda api_main.go 
	zip main.zip elasticsearch-lambda
	rm -f elasticsearch-lambda

clear:
	rm -f elasticsearch-lambda main.zip
	
test:
	go test `go list ./... | grep -v examples` -coverprofile=../elastic_coverage.txt -covermode=atomic