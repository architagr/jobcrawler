package controller

import (
	"elasticsearchservice/logger"
	"elasticsearchservice/services"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

type IIndexController interface {
	Create(name string, mapping *types.TypeMapping) error
	Delete(name string) error
	Exists(name string) (bool, error)
}

type indexController struct {
	indexService services.IIndexService
	loggerObj    logger.ILogger
}

func (ctrl *indexController) Create(name string, mapping *types.TypeMapping) error {
	return ctrl.indexService.Create(name, mapping)
}
func (ctrl *indexController) Delete(name string) error {
	return ctrl.indexService.Delete(name)
}
func (ctrl *indexController) Exists(name string) (bool, error) {
	return ctrl.indexService.Exists(name)
}

func InitIndexController(indexSvc services.IIndexService, logObj logger.ILogger) IIndexController {
	return &indexController{
		indexService: indexSvc,
		loggerObj:    logObj,
	}
}
