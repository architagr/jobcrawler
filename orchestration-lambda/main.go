package main

import (
	"context"

	"log"
	localAws "orchestration-lambda/aws"
	"orchestration-lambda/config"

	"orchestration-lambda/notification"
	oschestrationService "orchestration-lambda/service"

	"github.com/architagr/repository/connection"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	log.Printf("lambda start")

	lambda.Start(handler)
	// to be used for local
	// handle()
}
func handler(ctx context.Context, sqsEvent interface{}) error {
	return handle()
}
func handle() error {
	env := config.GetConfig()
	notify := initNotification(env)
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
	notify.SendNotificationToMonitoring(0)
	return nil
}

func initNotification(env config.IConfig) notification.INotification {
	snsSvc := localAws.GetSnsService()
	notify := notification.InitNotificationService(snsSvc, env)
	return notify
}
