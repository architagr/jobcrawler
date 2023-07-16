package services

import (
	"fmt"

	"github.com/elastic/go-elasticsearch/v8"
)

type IESConnection interface {
	GetConnection() (*elasticsearch.TypedClient, error)
}
type eSConnection struct {
	esClinet *elasticsearch.TypedClient
}

func (esConnectionObj *eSConnection) GetConnection() (*elasticsearch.TypedClient, error) {
	if esConnectionObj.esClinet == nil {
		return nil, fmt.Errorf("ES client is not initilized")
	}
	return esConnectionObj.esClinet, nil
}

func InitEsConnection(addresses []string, username, password string, certificate []byte) (IESConnection, error) {
	cfg := elasticsearch.Config{
		Addresses: addresses,
		Username:  username,
		Password:  password,
		CACert:    certificate,
	}
	es, err := elasticsearch.NewTypedClient(cfg)
	if err != nil {
		return nil, err
	}
	return &eSConnection{
		esClinet: es,
	}, nil
}
