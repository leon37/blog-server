package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/leon37/blog-server/internal/routers/api/v1"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	tag := v1.NewTag()
	article := v1.NewArticle()
	api := r.Group("/api/v1")
	{
		api.POST("/tags", tag.Create)
		api.DELETE("/tags/:id", tag.Delete)
		api.PUT("/tags/:id/", tag.Update)
		api.PATCH("/tags/:id/state", tag.Update)
		api.GET("/tags", tag.List)

		api.POST("/articles", article.Create)
		api.DELETE("/articles/:id", article.Delete)
		api.PUT("/articles/:id/", article.Update)
		api.PATCH("/articles/:id/state", article.Update)
		api.GET("/articles", article.List)
		api.GET("/articles/:id", article.Get)
	}
	return r
}