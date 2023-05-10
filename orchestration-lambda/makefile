lambda-build:
	make clear
	GOOS=linux GOARCH=amd64 go build -o orchestration main.go 
	zip main.zip orchestration
	rm -f orchestration

clear:
	rm -f orchestration main.zip

test:
	go test `go list ./... | grep -v examples` -coverprofile=../coverage.txt -covermode=atomic