package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Route(app *gin.Engine) {
	app.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello gin");
	})
}