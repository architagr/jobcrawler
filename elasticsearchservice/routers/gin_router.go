package routers

import (
	"UserAPI/config"
	"UserAPI/controller"
	"UserAPI/logger"
	middlewarePkg "UserAPI/middleware"
	"UserAPI/models"
	"context"
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

func InitGinRouters(userController controller.IUserController, logObj logger.ILogger) IRouter {
	ginMiddleware = getMiddlewares()
	ginEngine := gin.Default()
	registerInitialCommonMiddleware(ginEngine, logObj)
	routerGroup := getInitialRouteGroup(ginEngine)

	routerGroup.GET("/healthCheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "User api is working",
		})
	})
	createUserProfileRoutes(routerGroup, userController, logObj)
	return &ginRouter{
		ginEngine: ginEngine,
		env:       config.GetConfig(),
		logObj:    logObj,
	}
}

func createUserProfileRoutes(group *gin.RouterGroup, userController controller.IUserController, logObj logger.ILogger) {
	userProfileGroup := group.Group("profile")

	userProfileGroup.POST("/", func(ginContext *gin.Context) {
		var userDetails models.UserDetail
		if err := ginContext.ShouldBind(&userDetails); err != nil {
			logObj.Printf("wrong request body %+v", err)
			ginContext.Errors = append(ginContext.Errors, &gin.Error{
				Err:  err,
				Type: gin.ErrorTypeBind,
			})
			return
		}
		userData, err := userController.SaveUserProfile(&userDetails)
		if err != nil {
			ginContext.Errors = append(ginContext.Errors, &gin.Error{
				Err:  err,
				Type: gin.ErrorTypePrivate,
			})
			return
		}
		ginContext.JSON(http.StatusOK, userData)
	})
	userProfileGroup.PUT("/image/:id", func(ginContext *gin.Context) {

		var updateAvatarRequest models.UpdateAvatarRequest
		if err := ginContext.ShouldBind(&updateAvatarRequest); err != nil {
			logObj.Printf("wrong request body %+v", err)
			ginContext.Errors = append(ginContext.Errors, &gin.Error{
				Err:  err,
				Type: gin.ErrorTypeBind,
			})
			return
		}
		if err := ginContext.ShouldBindUri(&updateAvatarRequest); err != nil {
			logObj.Printf("wrong request param %+v", err)
			ginContext.Errors = append(ginContext.Errors, &gin.Error{
				Err:  err,
				Type: gin.ErrorTypeBind,
			})
			return
		}
		err := userController.SaveUserImage(&updateAvatarRequest)
		if err != nil {
			ginContext.Errors = append(ginContext.Errors, &gin.Error{
				Err:  err,
				Type: gin.ErrorTypePrivate,
			})
			return
		}
		ginContext.JSON(http.StatusOK, map[string]string{
			"message": "image uploaded",
		})
	})

	userProfileGroup.GET("/:username", func(ginContext *gin.Context) {
		email, ok := ginContext.Params.Get("username")
		if !ok {
			logObj.Printf("username is required")
			ginContext.Errors = append(ginContext.Errors, &gin.Error{
				Err:  errors.New("username is required"),
				Type: gin.ErrorTypeBind,
			})
			return
		}
		userData, err := userController.GetUserProfile(email)
		if err != nil {
			ginContext.Errors = append(ginContext.Errors, &gin.Error{
				Err:  err,
				Type: gin.ErrorTypePrivate,
			})
			return
		}
		ginContext.JSON(http.StatusOK, userData)
	})

}
func getInitialRouteGroup(ginEngine *gin.Engine) *gin.RouterGroup {
	return ginEngine.Group("/user")
}
func registerInitialCommonMiddleware(ginEngine *gin.Engine, logObj logger.ILogger) {
	ginEngine.Use(ginMiddleware.GetErrorHandler(logObj))
	ginEngine.Use(ginMiddleware.GetCorsMiddelware())
}
func getMiddlewares() middlewarePkg.IMiddleware[gin.HandlerFunc] {
	return middlewarePkg.InitGinMiddelware()
}
