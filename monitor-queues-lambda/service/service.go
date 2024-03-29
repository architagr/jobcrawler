package service

import (
	"log"
	"strconv"
	"sync"

	"common-constants/constants"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type IMonitoringService interface {
	StartMonitoring()
	GetCountOfMessages() int64
}

type MonitoringService struct {
	total  int64
	mutex  sync.Mutex
	sqsSvc *sqs.SQS
}

var svc IMonitoringService
var numberOfMessages string = "ApproximateNumberOfMessages"
var numberOfMessagesDelayed string = "ApproximateNumberOfMessagesDelayed"
var numberOfMessagesNotVisible string = "ApproximateNumberOfMessagesNotVisible"

func InitMonitoringService() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	sqsSvc := sqs.New(sess, &aws.Config{})
	svc = &MonitoringService{
		sqsSvc: sqsSvc,
		total:  0,
	}
}
func GetMonitoringService() IMonitoringService {
	if svc == nil {
		InitMonitoringService()
	}
	return svc
}

func (svc *MonitoringService) StartMonitoring() {

	wg := sync.WaitGroup{}
	hostNames := []string{string(constants.HostName_Linkedin), string(constants.HostName_Indeed)}

	for _, hostName := range hostNames {
		list, err := svc.sqsSvc.ListQueues(&sqs.ListQueuesInput{
			MaxResults:      aws.Int64(10),
			QueueNamePrefix: aws.String(hostName),
		})
		if err != nil {
			log.Printf("get list of queue error for host %s, error %+v", hostName, err)
		}
		log.Printf("list of queues for host %s are %+v", hostName, list)
		for _, url := range list.QueueUrls {
			wg.Add(1)
			go svc.countMessage(url, &wg)
		}
	}
	wg.Wait()
	log.Printf("total pending message on queues are %d", svc.total)
}
func (svc *MonitoringService) GetCountOfMessages() int64 {
	return svc.total
}
func (svc *MonitoringService) countMessage(queueUrl *string, wg *sync.WaitGroup) {
	defer wg.Done()
	attr, err := svc.sqsSvc.GetQueueAttributes(&sqs.GetQueueAttributesInput{
		QueueUrl:       queueUrl,
		AttributeNames: aws.StringSlice([]string{numberOfMessages, numberOfMessagesDelayed, numberOfMessagesNotVisible}),
	})
	if err != nil {
		log.Printf("queue: %s, error %+v", *queueUrl, err)
		return
	}
	log.Printf("queue: %s, attr %+v", *queueUrl, attr)
	svc.mutex.Lock()
	defer svc.mutex.Unlock()
	noOfMessage := attr.Attributes[numberOfMessages]
	invisibleMessageCount := attr.Attributes[numberOfMessagesNotVisible]
	delaiedMessageCount := attr.Attributes[numberOfMessagesDelayed]

	count1, _ := strconv.Atoi(*noOfMessage)
	count2, _ := strconv.Atoi(*invisibleMessageCount)
	count3, _ := strconv.Atoi(*delaiedMessageCount)
	svc.total += int64(count1) + int64(count2) + int64(count3)
}
