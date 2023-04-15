package main

import (
	"context"

	"log"
	"orchestration/config"
	"orchestration/notification"
	oschestrationService "orchestration/service"

	"github.com/architagr/repository/connection"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	log.Printf("lambda start")
	lambda.Start(handler)
	// to be used for local
	//handle()
}
func handler(ctx context.Context, sqsEvent interface{}) error {
	return handle()
}
func handle() error {
	env := config.GetConfig()
	notify := notification.GetNotificationObj()
	conn := connection.InitConnection(env.GetDatabaseConnectionString(), 10)
	err := conn.ValidateConnection()
	if err != nil {
		log.Println("error", err)
		return err
	}
	defer conn.Disconnect()
	svc, err := oschestrationService.InitService(conn, notify, env)
	if err != nil {
		log.Println("error", err)
		return err
	}
	svc.Start()
	return nil
}
