package middleware

import "elasticsearchservice/logger"

type IMiddleware[T any] interface {
	GetCorsMiddelware() T
	GetErrorHandler(logObj logger.ILogger) T
}
