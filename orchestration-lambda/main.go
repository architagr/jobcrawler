package main

import (
	"log"
	"orchestration/config"
	"orchestration/notification"
	oschestrationService "orchestration/service"

	"github.com/architagr/repository/connection"
)

func main() {
	env := config.GetConfig()
	notify := notification.GetNotificationObj()
	conn := connection.InitConnection(env.GetDatabaseConnectionString(), 10)
	err := conn.ValidateConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Disconnect()
	svc, err := oschestrationService.InitService(conn, notify, env)
	if err != nil {
		log.Fatal(err)
	}
	svc.Start()

}
