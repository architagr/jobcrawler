package main

import (
	"context"
	"fmt"
	"log"
	"monitorqueuelambda/config"
	"monitorqueuelambda/notification"
	"monitorqueuelambda/service"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	log.Printf("lambda start")
	config.InitConfig()
	notification.InitNotificationService()
	service.InitMonitoringService()
	lambda.Start(handler)
	//handle()
}
func handler(ctx context.Context, sqsEvent events.SQSEvent) error {
	return handle()
}

func handle() error {
	svc := service.GetMonitoringService()
	svc.StartMonitoring()
	if svc.GetCountOfMessages() == 0 {
		return fmt.Errorf("no messages to process")
	}
	notify := notification.GetNotificationObj()
	notify.SendUrlNotificationToMonioring(svc.GetCountOfMessages())
	return nil
}
