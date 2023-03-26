lambda-build:
	rm -f scrapper main.zip
	GOOS=linux GOARCH=amd64 go build -o scrapper main.go 
	zip main.zip scrapper
	rm -f scrapper