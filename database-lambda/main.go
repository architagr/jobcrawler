package main

import (
	"context"
	"database_lambda/config"
	"encoding/json"
	"fmt"
	"log"

	jobdetails "github.com/architagr/common-models/job-details"
	notificationModel "github.com/architagr/common-models/sns-notification"
	"github.com/architagr/repository/collection"
	"github.com/architagr/repository/connection"
	"gorgonia.org/tensor"

	sqs_message "github.com/architagr/common-models/sqs-message"
	"github.com/aws/aws-lambda-go/events"

	"github.com/nlpodyssey/spago/pkg/nlp/tokenizer/whitespace"
	"github.com/nlpodyssey/spago/pkg/nlp/transformers/bert"
	"gorgonia.org/tensor"
)

var conn connection.IConnection
var env *config.Config

func handler(ctx context.Context, sqsEvent events.SQSEvent) error {
	log.Printf("lambda handler start")
	jobs := make([]jobdetails.JobDetails, 0)
	for _, message := range sqsEvent.Records {
		data := new(sqs_message.MessageBody)
		json.Unmarshal([]byte(message.Body), data)
		messageContent := new(notificationModel.Notification[jobdetails.JobDetails])
		json.Unmarshal([]byte(data.Message), messageContent)
		jobs = append(jobs, messageContent.Data)
		fmt.Printf("The message %s for event source %s, messageContent: %+v \n", message.MessageId, message.EventSource, messageContent)
	}

	doc, err := collection.InitCollection[jobdetails.JobDetails](conn, env.GetDatabaseName(), env.GetCollectionName())
	if err != nil {
		log.Panic(err)
	}
	_, err = doc.AddMany(jobs)
	if err != nil {
		log.Panic(err)
	}
	return nil
}

// func main() {
// 	log.Printf("lambda start")
// 	env = config.GetConfig()
// 	setupDB()
// 	defer conn.Disconnect()
// 	lambda.Start(handler)
// }

func setupDB() {
	conn = connection.InitConnection(env.GetDatabaseConnectionString(), 10)
	err := conn.ValidateConnection()
	if err != nil {
		log.Fatalf("error in conncting to mongo %+v", err)
	}
}

func main() {
	// Define the job description
	jobDescription := "Put your job description here"

	// Load a pre-trained BERT model
	model, err := bert.NewDefaultBertForSequenceClassification()
	if err != nil {
		panic(err)
	}

	// Tokenize the job description
	tokenizer := whitespace.NewTokenizer()
	tokens := tokenizer.Tokenize(jobDescription)

	// Convert the tokens to BERT input format
	inputIds, _, _, err := model.BERTTokenize(tokens)
	if err != nil {
		panic(err)
	}

	// Run the BERT model to obtain the embeddings
	embeddings, err := model.BERTForward(tensor.New(tensor.Of(tensor.Int), tensor.WithShape(1, len(inputIds))).SetData(inputIds))
	if err != nil {
		panic(err)
	}

	// Obtain the top 10 keywords
	topKeywords := getTopKeywords(embeddings, tokens, 10)

	// Print the extracted keywords
	for _, keyword := range topKeywords {
		fmt.Println(keyword)
	}
}

func getTopKeywords(embeddings *bert.Output, tokens []string, k int) []string {
	// Calculate the embeddings' mean and standard deviation
	embeddingsMean, embeddingsStddev := embeddings.PooledOutputMeanAndStd()

	// Calculate the cosine similarity between each token embedding and the embeddings' mean
	similarities := make(map[string]float32)
	for i, token := range tokens {
		embedding := embeddings.Outputs[0][i+1]
		similarity := bert.CosineSimilarity(embedding, embeddingsMean, embeddingsStddev)
		similarities[token] = similarity
	}

	// Sort the tokens by similarity score and extract the top k keywords
	topKeywords := make([]string, 0, k)
	for len(topKeywords) < k && len(similarities) > 0 {
		maxSimilarity := float32(-1)
		maxToken := ""
		for token, similarity := range similarities {
			if similarity > maxSimilarity {
				maxSimilarity = similarity
				maxToken = token
			}
		}
		topKeywords = append(topKeywords, maxToken)
		delete(similarities, maxToken)
	}

	return topKeywords
}
