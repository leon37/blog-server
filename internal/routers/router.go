package routers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/leon37/blog-server/docs"
	"github.com/leon37/blog-server/global"
	"github.com/leon37/blog-server/internal/middleware"
	"github.com/leon37/blog-server/internal/routers/api"
	v1 "github.com/leon37/blog-server/internal/routers/api/v1"
	"github.com/leon37/blog-server/pkg/limiter"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"strings"
	"time"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(limiter.LimiterBucketRule{
	Key:          "/auth",
	FillInterval: time.Second,
	Capacity:     10,
	Quantum:      10,
})

func NewRouter() *gin.Engine {
	r := gin.New()
	if strings.ToLower(global.ServerSetting.RunMode) == "debug" {
		r.Use(gin.Logger(), gin.Recovery())
	} else {
		r.Use(middleware.AccessLog(), middleware.Recovery())
	}
	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeout(60 * time.Second))
	r.Use(middleware.Translation())
	r.Use(middleware.Tracing())
	r.GET("/auth", api.GetAuth)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	tag := v1.NewTag()
	article := v1.NewArticle()
	upload := api.NewUpload()
	r.POST("/upload/file", upload.UploadFile)
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))
	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.JWT())
	apiv1.Use(middleware.AppInfo())
	{
		apiv1.POST("/tags/create", tag.Create)
		apiv1.POST("/tags/delete", tag.Delete)
		apiv1.POST("/tags/update", tag.Update)
		apiv1.GET("/tags/list", tag.List)

		apiv1.POST("/articles/create", article.Create)
		apiv1.POST("/articles/delete", article.Delete)
		apiv1.POST("/articles/update", article.Update)
		apiv1.GET("/articles/list", article.List)
		apiv1.GET("/articles/single", article.Get)
	}
	return r
}
