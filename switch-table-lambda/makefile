lambda-build:
	make clear
	GOOS=linux GOARCH=amd64 go build -o switchtable main.go 
	zip main.zip switchtable
	rm -f switchtable

clear:
	rm -f switchtable main.zip