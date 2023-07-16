package main

import (
	"context"
	"elasticsearchservice/logger"
	"elasticsearchservice/services"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

var esConnection services.IESConnection
var indexSvcObj services.IIndexService
var logObj logger.ILogger

type Request struct {
	FunctionName string `json:"functionName"`
	Body         string `json:"body"`
}
type Response struct {
	Data   string `json:"data"`
	Status int    `json:"status"`
}

func ElasticSearch(ctx context.Context, req Request) (Response, error) {
	log.Printf("request %+v", req)
	return Response{Data: "Success", Status: 200}, nil
}
func InitLogger() {
	logObj = logger.InitConsoleLogger()
}
func InitEsConnection() {
	cert, _ := ioutil.ReadFile("ca.crt")
	conn, err := services.InitEsConnection([]string{"https://localhost:9200"}, "elastic", "password123", cert)
	if err != nil {
		panic(err)
	}
	esConnection = conn
}
func InitService() {
	indexSvcObj = services.InitIndexService(esConnection, logObj)
}
func main() {
	// #region create a internal lambda
	// lambda.Start(ElasticSearch)

	// #endregion
	// #region call another lambda function
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	data := Request{
		FunctionName: "CreateIndex",
		Body:         "body",
	}
	payload, _ := json.Marshal(data)
	lambdaSvc := lambda.New(sess)
	res, err := lambdaSvc.InvokeWithContext(context.TODO(), &lambda.InvokeInput{
		FunctionName: aws.String("arn:aws:lambda:ap-southeast-1:638580160310:function:elasticsearch-lambda-fn"),
		Payload:      payload,
	})
	if err != nil {
		log.Panicln(err)
	}
	var responseBody Response
	err = json.Unmarshal(res.Payload, &responseBody)
	log.Printf("%+v", res)
	log.Printf("%+v", responseBody)

	return
	// #endregion
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
