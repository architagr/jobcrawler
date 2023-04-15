lambda-build:
	make clear
	GOOS=linux GOARCH=amd64 go build -o monitorqueue main.go 
	zip main.zip monitorqueue
	rm -f monitorqueue

clear:
	rm -f monitorqueue main.zip