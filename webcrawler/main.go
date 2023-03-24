package main

import (
	"encoding/json"
	"fmt"
	"jobcrawler/config"
	"jobcrawler/notification"
	"jobcrawler/urlfrontier"
	"jobcrawler/urlseeding"
	"log"
	"sync"
	"time"

	"github.com/architagr/common-constants/constants"
	searchcondition "github.com/architagr/common-models/search-condition"
	"github.com/architagr/repository/connection"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

var env *config.Config

func sendSqsMessage() {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := sqs.New(sess)
	search := &searchcondition.SearchCondition{
		JobTitle: constants.JobTitle_SoftwareEngineer,
		LocationInfo: searchcondition.Location{
			Country: "United States",
			City:    "New York",
		},
		RoleName:   constants.Role_Engineering,
		JobType:    constants.JobType_FullTime,
		JobModel:   constants.JobModel_OnSite,
		Experience: constants.ExperienceLevel_EntryLevel,
	}
	bytes, err := json.Marshal(search)
	_, err = svc.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(10),
		MessageBody:  aws.String(string(bytes)),
		QueueUrl:     aws.String(env.GetScrapperQueueUrl()),
	})
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	config.InitConfig()
	env = config.GetConfig()
	fmt.Println(env.GetScrapperQueueUrl())
	// sendSqsMessage()
	// setupDB()
	crawlLinkedIn()
}
func crawlLinkedIn() {
	notification := notification.GetNotificationObj()
	search := &searchcondition.SearchCondition{
		JobTitle: constants.JobTitle_SoftwareEngineer,
		LocationInfo: searchcondition.Location{
			Country: "United States",
			City:    "New York",
		},
		RoleName:   constants.Role_Engineering,
		JobType:    constants.JobType_FullTime,
		JobModel:   constants.JobModel_OnSite,
		Experience: constants.ExperienceLevel_EntryLevel,
	}
	start := time.Now()

	wg := sync.WaitGroup{}
	urlSeeding := urlseeding.InitUrlSeeding()
	linksToCrawl := urlSeeding.GetLinks(search)
	frontier := urlfrontier.InitUrlFrontier(search, linksToCrawl, notification)
	log.Println("*******")
	frontier.Start(&wg)

	log.Println("*******")
	end := time.Now()
	log.Println(end.Sub(start))
}

func setupDB() {
	conn := connection.InitConnection("mongodb+srv://webscrapper:WebScrapper123@cluster0.xzvihm7.mongodb.net/?retryWrites=true&w=majority", 10)
	err := conn.ValidateConnection()
	if err != nil {
		log.Fatalf("error in conncting to mongo %+v", err)
	}

	client, ctx, err := conn.GetConnction()
	if err != nil {
		log.Fatalf("error in conncting to mongo %+v", err)
	}

	defer client.Disconnect(ctx)
}
