package router

import (
	"github.com/gin-gonic/gin"
	"github.com/quanku/controller/articler"
	"github.com/quanku/controller/auth"
	"github.com/quanku/controller/tag"
	"github.com/quanku/middleware/jwt"
)

func Route(app *gin.Engine) {

	app.GET("/auth", auth.GetAuth)
	apiV1 := app.Group("api/v1")
	apiV1.Use(jwt.JWT())
	{
		apiV1.GET("/articles", articler.GetArticles) // 获取文章列表
		apiV1.GET("/articles/:id", articler.GetArticles) // 获取指定文章
		apiV1.POST("/articles", articler.AddArticle) // 新增文章
		apiV1.PUT("/articles/:id", articler.EditArticle) // 修改文章
		apiV1.DELETE("/articles/:id", articler.DeleteArticle) // 删除文章

		//================ 标签路由 ==========================
		apiV1.GET("/tags", tag.GetTags)
		apiV1.POST("/tags", tag.AddTag)
		apiV1.PUT("/tags/:id", tag.EditTag)
		apiV1.DELETE("/tags/:id", tag.DeleteTag)
	}
}
