package routers

import (
	"elasticsearchservice/config"
	"elasticsearchservice/controller"
	"elasticsearchservice/logger"
	"elasticsearchservice/models"

	"context"
	middlewarePkg "elasticsearchservice/middleware"
	"errors"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginMiddleware middlewarePkg.IMiddleware[gin.HandlerFunc]

type ginRouter struct {
	ginEngine *gin.Engine
	env       config.IConfig
	logObj    logger.ILogger
}

var ginLambda *ginadapter.GinLambda

// Handler is the function that executes for every Request passed into the Lambda
func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}
func (router *ginRouter) StartApp(port int) {
	if router.env.IsLambda() {
		ginLambda = ginadapter.New(router.ginEngine)
		lambda.Start(Handler)
	} else {
		router.ginEngine.Run(fmt.Sprintf(":%d", port))
	}
}

func InitGinRouters(indexController controller.IIndexController, logObj logger.ILogger) IRouter {
	ginMiddleware = getMiddlewares()
	ginEngine := gin.Default()
	registerInitialCommonMiddleware(ginEngine, logObj)
	routerGroup := getInitialRouteGroup(ginEngine)

	routerGroup.GET("/healthCheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "elasticsearch api is working",
		})
	})
	createIndexRoutes(routerGroup, indexController, logObj)
	return &ginRouter{
		ginEngine: ginEngine,
		env:       config.GetConfig(),
		logObj:    logObj,
	}
}

func createIndexRoutes(group *gin.RouterGroup, indexController controller.IIndexController, logObj logger.ILogger) {
	indexGroup := group.Group("index")

	indexGroup.POST("/", func(ginContext *gin.Context) {
		var req models.CreateIndexRequest
		if err := ginContext.ShouldBind(&req); err != nil {
			logObj.Printf("wrong request body %+v", err)
			ginContext.Errors = append(ginContext.Errors, &gin.Error{
				Err:  err,
				Type: gin.ErrorTypeBind,
			})
			return
		}
		err := indexController.Create(req.Name, &req.Mapping)
		if err != nil {
			ginContext.Errors = append(ginContext.Errors, &gin.Error{
				Err:  err,
				Type: gin.ErrorTypePrivate,
			})
			return
		}
		ginContext.JSON(http.StatusCreated, map[string]string{
			"status": fmt.Sprintf("index: %s created", req.Name),
		})
	})
	indexGroup.DELETE("/:name", func(ginContext *gin.Context) {
		name, ok := ginContext.Params.Get("name")
		if !ok {
			logObj.Printf("name is required")
			ginContext.Errors = append(ginContext.Errors, &gin.Error{
				Err:  errors.New("name is required"),
				Type: gin.ErrorTypeBind,
			})
			return
		}
		err := indexController.Delete(name)
		if err != nil {
			ginContext.Errors = append(ginContext.Errors, &gin.Error{
				Err:  err,
				Type: gin.ErrorTypePrivate,
			})
			return
		}
		ginContext.JSON(http.StatusOK, map[string]string{
			"status": fmt.Sprintf("index: %s deleted", name),
		})
	})

	indexGroup.GET("/validate/:name", func(ginContext *gin.Context) {
		name, ok := ginContext.Params.Get("name")
		if !ok {
			logObj.Printf("name is required")
			ginContext.Errors = append(ginContext.Errors, &gin.Error{
				Err:  errors.New("name is required"),
				Type: gin.ErrorTypeBind,
			})
			return
		}
		status, err := indexController.Exists(name)
		if err != nil {
			ginContext.Errors = append(ginContext.Errors, &gin.Error{
				Err:  err,
				Type: gin.ErrorTypePrivate,
			})
			return
		}
		ginContext.JSON(http.StatusOK, map[string]bool{
			"status": status,
		})
	})

}
func getInitialRouteGroup(ginEngine *gin.Engine) *gin.RouterGroup {
	return ginEngine.Group("/es")
}

func registerInitialCommonMiddleware(ginEngine *gin.Engine, logObj logger.ILogger) {
	ginEngine.Use(ginMiddleware.GetErrorHandler(logObj))
	// ginEngine.Use(ginMiddleware.GetCorsMiddelware())
}
func getMiddlewares() middlewarePkg.IMiddleware[gin.HandlerFunc] {
	return middlewarePkg.InitGinMiddelware()
}
