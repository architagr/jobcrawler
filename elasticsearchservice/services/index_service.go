package services

import (
	"context"
	"elasticsearchservice/logger"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/indices/create"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

type IIndexService interface {
	Create(name string, mapping *types.TypeMapping) error
	Delete(name string) error
	Exists(name string) (bool, error)
}

type indexService struct {
	esConn    IESConnection
	loggerObj logger.ILogger
}

func (svc *indexService) Create(name string, mapping *types.TypeMapping) error {
	// TODO: validate name is valid using the naming convention and REGEX
	esSvc, err := svc.esConn.GetConnection()
	if err != nil {
		return err
	}
	svc.loggerObj.Printf("creating index %s with mapping %+v/n", name, mapping)
	createIndexres, err := esSvc.Indices.Create(name).Request(&create.Request{
		Mappings: mapping,
	}).Do(context.Background())

	if err != nil {
		return err
	}
	if !createIndexres.Acknowledged {
		return fmt.Errorf("index not created")
	}
	return nil
}
func (svc *indexService) Delete(name string) error {
	esSvc, err := svc.esConn.GetConnection()
	if err != nil {
		return err
	}
	svc.loggerObj.Printf("deleting index %s/n", name)
	deleteResponse, err := esSvc.Indices.Delete(name).Do(context.Background())

	if err != nil {
		return err
	}
	if !deleteResponse.Acknowledged {
		return fmt.Errorf("index not deleted")
	}
	return nil
}
func (svc *indexService) Exists(name string) (bool, error) {
	esSvc, err := svc.esConn.GetConnection()
	if err != nil {
		return false, err
	}
	return esSvc.Indices.Exists(name).IsSuccess(context.Background())
}

func InitIndexService(conn IESConnection, logObj logger.ILogger) IIndexService {
	return &indexService{
		esConn:    conn,
		loggerObj: logObj,
	}
}
