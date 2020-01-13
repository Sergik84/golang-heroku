package main

import (
	"github.com/Sergik84/golang-heroku/logic"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./web", true)))
	api := router.Group("/api")
	logic := logic.BusinessLogic{}
	logic.RegisterEndpoints(api)
	router.Run()
}
