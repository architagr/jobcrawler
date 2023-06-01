package main

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strings"
	"switchtablename/config"
	"time"

	"repository/connection"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/aws/aws-lambda-go/lambda"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var dateFormat string = "2006-01-02"
var conn connection.IConnection
var env config.IConfig

func main() {
	log.Printf("lambda start")
	env = config.GetConfig()
	setupDB()
	lambda.Start(handler)
}
func handler(ctx context.Context, sqsEvent events.SQSEvent) error {
	client, _, err := conn.GetConnction()
	if err != nil {
		log.Panic(err)
	}
	deleteOldBackup(client)
	createBackup(client)
	swictCollection(client)
	dropTempCollection(client)
	return nil
}
func swictCollection(client *mongo.Client) {
	log.Printf("renaming %s to %s", env.GetTempCollectionName(), env.GetFinalCollectionName())
	renameCollection(env.GetTempCollectionNameWithDbName(), env.GetFinalCollectionNameWithDbName(), client)
}
func dropTempCollection(client *mongo.Client) {
	db := client.Database(env.GetDatabaseName())
	db.Collection(env.GetTempCollectionName()).Drop(context.TODO())
}
func createBackup(client *mongo.Client) {
	t := time.Now()
	oldFinalTableBackup := fmt.Sprintf("%s_%s", env.GetFinalCollectionNameWithDbName(), t.Format(dateFormat))
	log.Printf("renaming %s to %s", env.GetFinalCollectionName(), oldFinalTableBackup)
	renameCollection(env.GetFinalCollectionNameWithDbName(), oldFinalTableBackup, client)
}
func renameCollection(oldName, newName string, client *mongo.Client) {
	err := client.Database("admin").RunCommand(context.TODO(), bson.D{
		{Key: "renameCollection", Value: oldName},
		{Key: "to", Value: newName},
		{Key: "dropTarget", Value: true},
	}).Err()
	if err != nil {
		log.Panic(err)
	}
}
func deleteOldBackup(client *mongo.Client) {
	names, err := getAllBackupCollectionsNames(client)
	if err != nil {
		log.Panic(err)
	}
	db := client.Database(env.GetDatabaseName())
	log.Printf("collections that will be tried to be droped if older then 30 days %+v", names)
	for _, collectionName := range names {

		splitValues := strings.Split(collectionName, fmt.Sprintf("%s_", env.GetFinalCollectionName()))
		log.Printf("testing %s to be droped if older then 30 days", collectionName)
		date, _ := time.Parse(dateFormat, splitValues[1])
		t := time.Now()
		diff := t.Sub(date)
		days := diff.Hours() / 24
		if days > 30 {
			log.Printf("dropping collection %s", collectionName)
			db.Collection(collectionName).Drop(context.TODO())
		}
	}

}
func getAllBackupCollectionsNames(client *mongo.Client) ([]string, error) {
	x, regexErr := regexp.Compile(fmt.Sprintf("^%s_", env.GetFinalCollectionName()))
	if regexErr != nil {
		log.Panic(regexErr)
	}
	db := client.Database(env.GetDatabaseName())
	return db.ListCollectionNames(context.TODO(), bson.M{
		"name": bson.M{"$regex": x.String()},
	})

}
func setupDB() {
	conn = connection.InitConnection(env.GetDatabaseConnectionString(), 10)
	err := conn.ValidateConnection()
	if err != nil {
		log.Fatalf("error in conncting to mongo %+v", err)
	}
}
