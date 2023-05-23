lambda-build:
	make clear
	GOOS=linux GOARCH=amd64 go build -o scrapper main.go 
	zip main.zip scrapper
	rm -f scrapper

clear:
	rm -f scrapper main.zip

test:
	go test `go list ./... | grep -v examples` -coverprofile=../scrapper_coverage.txt -covermode=atomic