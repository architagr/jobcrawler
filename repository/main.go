package main

import (
	"log"

	"github.com/architagr/repository/connection"
	"github.com/architagr/repository/document"
)

type Jobs struct {
	JobId    int
	Title    string
	Company  string
	Location string
}

func main() {
	connection := connection.InitConnection("mongodb+srv://webscrapper:WebScrapper123@cluster0.xzvihm7.mongodb.net/?retryWrites=true&w=majority", 10)
	err := connection.ValidateConnection()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("connection validated")

	doc, err := document.InitDocument[Jobs](connection, "webscrapper", "jobs")
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

	result, err := doc.Get(filter)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("filtered doc retrived %+v", result)

}
