package router

import (
	"go-dummy-project/go-dummy/handler"

	"github.com/gin-gonic/gin"
)

func ApiRouterGroup(apiRouter *gin.RouterGroup) *gin.RouterGroup {
	r := apiRouter.Group("/api")
	{
		r.GET("/", handler.DummyUserHandler)
	}
	return r
}
