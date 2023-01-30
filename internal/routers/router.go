package routers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/leon37/blog-server/docs"
	"github.com/leon37/blog-server/internal/middleware"
	v1 "github.com/leon37/blog-server/internal/routers/api/v1"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	r.Use(middleware.Translation())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	tag := v1.NewTag()
	article := v1.NewArticle()
	api := r.Group("/api/v1")
	{
		api.POST("/tags/create", tag.Create)
		api.POST("/tags/delete", tag.Delete)
		api.POST("/tags/update", tag.Update)
		api.GET("/tags/list", tag.List)

		api.POST("/articles/create", article.Create)
		api.POST("/articles/delete", article.Delete)
		api.POST("/articles/update", article.Update)
		api.GET("/articles/list", article.List)
		api.GET("/articles/single", article.Get)
	}
	return r
}
