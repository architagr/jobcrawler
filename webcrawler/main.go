package main

import (
	"context"
	"fmt"
	"jobcrawler/notification"
	"jobcrawler/urlfrontier"
	"jobcrawler/urlseeding"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	// setupDB()
	crawlLinkedIn()
}
func crawlLinkedIn() {
	notification := new(notification.Notification)
	search := &urlseeding.SearchCondition{
		JobTitle: urlseeding.JobTitle_SoftwareEngineer,
		LocationInfo: urlseeding.Location{
			Country: "United States",
			City:    "New York",
		},
		RoleName:   urlseeding.Role_Engineering,
		JobType:    urlseeding.JobType_FullTime,
		JobModel:   urlseeding.JobModel_OnSite,
		Experience: urlseeding.ExperienceLevel_EntryLevel,
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
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://webscrapper:WebScrapper123@cluster0.xzvihm7.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)
}
