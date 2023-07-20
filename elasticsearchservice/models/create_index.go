package models

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type CreateIndexRequest struct {
	Mapping types.TypeMapping `json:"mapping"` // TODO: create a custom type and create a way to create map
	Name    string            `json:"name"`
}
