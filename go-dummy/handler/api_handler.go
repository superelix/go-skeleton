package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DummyUserHandler(c *gin.Context) {
	c.JSON(http.StatusOK,
		gin.H{
			"msg": "Dummy User Creation!",
		})
}
