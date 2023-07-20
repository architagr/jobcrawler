package models

import "elasticsearchservice/enums"

type ErrorResponse struct {
	ErrorCode enums.ErrorCode `json:"errorCode"`
	Message   string          `json:"message"`
}
