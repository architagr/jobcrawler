package main

import (
	"log"

	"github.com/architagr/repository/collection"
	"github.com/architagr/repository/connection"
)

type Jobs struct {
	JobId    int
	Title    string
	Company  string
	Location string
}

func main() {
	conn := connection.InitConnection("mongodb+srv://webscrapper:WebScrapper123@cluster0.xzvihm7.mongodb.net/?retryWrites=true&w=majority", 10)
	err := conn.ValidateConnection()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("connection validated")

	doc, err := collection.InitCollection[Jobs](conn, "webscrapper", "jobs")
	if err != nil {
		log.Fatal(err)
	}
	// id, err := doc.AddSingle(Jobs{JobId: 4, Title: "Workplace Services Manager, Google", Company: "Google", Location: "New York, NY"})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("new id %+v", id)

	// newDoc, err := doc.GetById(id)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("new doc retrived %+v", newDoc)

	filter := map[string]interface{}{"title": "Workplace Services Manager, Google"}

	result, err := doc.Get(filter, 10, 0)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("filtered doc retrived %+v", result)

}
