package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

func main() {
	//docker ps
	// get container id if the es01
	//docker cp <<containerid>>:/usr/share/elasticsearch/config/certs/ca/ca.crt .
	cert, _ := ioutil.ReadFile("ca.crt")

	cfg := elasticsearch.Config{
		Addresses: []string{
			"https://localhost:9200",
		},
		Username: "elastic",
		Password: "password123",
		CACert:   cert,
	}
	es, err := elasticsearch.NewTypedClient(cfg)
	if err != nil {
		panic(err)
	}

	index := "my_index2"

	// createIndexres, err := es.Indices.Create(index).Request(&create.Request{
	// 	Mappings: &types.TypeMapping{
	// 		Properties: map[string]types.Property{
	// 			"body": types.NewTextProperty(),
	// 			// "title": types.NewTextProperty(),
	// 		},
	// 	},
	// }).
	// 	Do(context.Background())
	// if err != nil {
	// 	panic(err)
	// }
	// log.Printf("create index response %+v", createIndexres)

	// doc := struct {
	// 	Title string `json:"title"`
	// 	Body  string `json:"body"`
	// }{
	// 	Title: "2My first document elasticSearch",
	// 	Body:  "Hello, Elasticsearch!",
	// }
	// addDocRes, err := es.Index(index).Request(doc).Do(context.Background())
	// if err != nil {
	// 	panic(err)
	// }
	// log.Printf("add doc response %+v", addDocRes)

	resSearch, err := es.Search().
		Index(index).
		Request(&search.Request{
			Query: &types.Query{

				Match: map[string]types.MatchQuery{
					"title": {Query: "2 my document", Fuzziness: 0},
				},
			},
		}).Do(context.Background())
	if err != nil {
		panic(err)
	}
	log.Printf("search response %+v", resSearch.Hits.Total)

	// var x []byte
	for i := 0; i < int(resSearch.Hits.Total.Value); i++ {
		x, _ := resSearch.Hits.Hits[i].Source_.MarshalJSON()
		log.Printf("search doc %d , score: %f, %+v", i, float64(resSearch.Hits.Hits[i].Score_), string(x))
	}

	fmt.Println("Document indexed!")
}
