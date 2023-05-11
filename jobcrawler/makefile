lambda-build:
	make clear
	GOOS=linux GOARCH=amd64 go build -o webcrawler main.go 
	zip main.zip webcrawler
	rm -f webcrawler

clear:
	rm -f webcrawler main.zip

test:
	go test `go list ./... | grep -v examples` -coverprofile=../jobcrawler_coverage.txt -covermode=atomic