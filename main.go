package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/quanku/pkg/setting"
	"github.com/quanku/router"
	"net/http"
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

	//================ 注册中间件 ===========
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	//================ 注册路由  ============
	router.Route(app)

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        app,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
