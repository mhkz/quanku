package main

import (
	"github.com/gin-gonic/gin"
	"github.com/quanku/router"
)

func initTemplate(router *gin.Engine)  {
	router.StaticFile("/", "./static/index.html")
	router.StaticFile("/favicon.ico", "./static/favicon.ico")
	router.Static("/css/", "./static/css")
	router.Static("/fonts/", "./static/fonts")
	router.Static("/img/", "./static/img")
	router.Static("/js/", "./static/js")
}

func main() {
	app := gin.New()
	router.Route(app)
	app.Run(":8002")
}
