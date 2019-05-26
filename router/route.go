package router

import (
	"github.com/gin-gonic/gin"
	"github.com/quanku/controller/articler"
	"github.com/quanku/controller/tag"
)

func Route(app *gin.Engine) {

	apiV1 := app.Group("api/v1")
	{
		apiV1.GET("/articles", articler.GetArticles)
		apiV1.POST("/articles", articler.AddArticle)
		apiV1.PUT("/articles/:id", articler.EditArticle)
		apiV1.DELETE("/articles/:id", articler.DeleteArticle)

		//================ 标签路由 ==========================
		apiV1.GET("/tags", tag.GetTags)
		apiV1.POST("/tags", tag.AddTag)
		apiV1.PUT("/tags/:id", tag.EditTag)
		apiV1.DELETE("/tags/:id", tag.DeleteTag)
	}
}
