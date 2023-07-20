package main

import (
	// "fmt"
	// "log"
	// "strings"

	"flag"

	"github.com/jdkato/prose/v2"

	"context"
	"elasticsearchservice/config"
	"elasticsearchservice/controller"
	"elasticsearchservice/logger"
	"elasticsearchservice/routers"
	"elasticsearchservice/services"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	_ "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/indices/create"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

func main_ML() {
	jobDescription := `Job description<p>The ideal candidate will oversee the online marketing strategy for the organization by planning and executing digital marketing campaigns. This candidate will launch advertisements and create content to increase brand awareness. This candidate will have previous marketing experience and be able to monitor the company's social media presence. </p><p> </p><p>Responsibilities</p><ul><li>Lead Generation</li><li>Design, maintain, and supply content for the organization's website</li><li>Formulate strategies to build lasting digital connection with customers</li><li>Monitor company presence on social media</li><li>Launch advertisements to increase brand awareness</li><li>Hands on Google Ads, Linkedin Ads, Meta Ads.</li><li>Experience of Email Marketing.</li></ul><p><br/></p><p>Qualifications</p><ul><li>Bachelor's degree in Marketing or related field</li><li>Excellent understanding of digital marketing concepts</li><li>Experience with business to customer social media and content generation</li><li>Strong creative and analytical skills</li></ul><p><br/></p>`

	keySkills := extractKeySkills(jobDescription)

	fmt.Printf("Key Skills: %s\n", strings.Join(keySkills, ", "))
}

func extractKeySkills(jobDescription string) []string {
	doc, err := prose.NewDocument(jobDescription, prose.WithSegmentation(true))
	if err != nil {
		log.Fatalf("Error creating document: %v", err)
	}

	var keySkills []string
	data := doc.Entities()
	for _, ent := range data {
		if ent.Label == "PERSON" {
			keySkills = append(keySkills, ent.Text)
		}
	}

	return keySkills
}

var configObj config.IConfig
var esConnection services.IESConnection
var indexSvcObj services.IIndexService
var logObj logger.ILogger
var indexController controller.IIndexController
var (
	port = flag.Int("port", 8080, "this value is used when we run the service on local")
)

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
func initConfig() {
	configObj = config.GetConfig()
}
func initLogger() {
	logObj = logger.InitConsoleLogger()
}
func initEsConnection() {
	cert, _ := ioutil.ReadFile(configObj.GetCertificatePath())
	conn, err := services.InitEsConnection(configObj.GetElasticSearchUrls(), configObj.GetElasticSearchUsername(), configObj.GetElasticSearchPassword(), cert)
	if err != nil {
		panic(err)
	}
	esConnection = conn
}
func initServices() {
	indexSvcObj = services.InitIndexService(esConnection, logObj)
}
func initControllers() {
	indexController = controller.InitIndexController(indexSvcObj, logObj)
}
func main() {
	flag.Parse()
	initConfig()
	initLogger()
	initEsConnection()
	initServices()
	initControllers()
	routers.InitGinRouters(indexController, logObj).StartApp(*port)

}
func maint() {
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

	createIndexres, err := es.Indices.Create(index).Request(&create.Request{
		Mappings: &types.TypeMapping{
			Properties: map[string]types.Property{
				"body": types.NewTextProperty(),
				// "title": types.NewTextProperty(),
			},
		},
	}).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	log.Printf("create index response %+v", createIndexres)

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
