package adddocument

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AddDocument struct {
}

type Jobs struct {
	JobId    int
	Title    string
	Company  string
	Location string
}

func (addDoc *AddDocument) InsertJobDetails() {
	//Create client connection
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

	//fetch DB
	coll := client.Database("webscrapper").Collection("jobs")
	docs := []interface{}{
		Jobs{JobId: 1, Title: "Workplace Services Manager, Google", Company: "Google", Location: "New York, NY"},
		Jobs{JobId: 2, Title: "Technical Accounting Specialist", Company: "Google", Location: "New York, NY"},
	}
	result, err := coll.InsertMany(context.TODO(), docs)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Fatal(result.InsertedIDs...)
	}
}
