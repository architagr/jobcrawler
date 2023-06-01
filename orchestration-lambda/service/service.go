package service

import (
	"orchestration-lambda/config"
	"orchestration-lambda/notification"

	jobcrawlerlinks "common-models/job-crawler-links"

	"repository/collection"
	"repository/connection"
)

type IService interface {
	Start() error
}

type Service struct {
	notify        notification.INotification
	env           config.IConfig
	collectionObj collection.ICollection[jobcrawlerlinks.JobCrawlerLink]
}

func InitService(conn connection.IConnection, notify notification.INotification, env config.IConfig) (IService, error) {
	jobCrawlDoc, err := collection.InitCollection[jobcrawlerlinks.JobCrawlerLink](conn, env.GetDatabaseName(), env.GetCollectionName())
	if err != nil {
		return nil, err
	}
	return &Service{
		notify:        notify,
		env:           env,
		collectionObj: jobCrawlDoc,
	}, nil
}

func (svc *Service) Start() error {
	for pageSize, startPage := 10, 0; ; startPage++ {
		links, err := svc.collectionObj.Get(nil, int64(pageSize), int64(startPage))
		if err != nil {
			return err
		}
		if len(links) == 0 {
			break
		}
		for _, link := range links {
			svc.notify.SendUrlNotificationToCrawler(&link.SearchCondition, link.HostName, link.Link)
		}
	}
	return nil
}
