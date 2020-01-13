package logic

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BusinessLogic struct {
}

func (logic *BusinessLogic) RegisterEndpoints(router *gin.RouterGroup) {
	router.GET("/health", logic.healthCheck)
}

func (logic *BusinessLogic) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "UP"})
}
