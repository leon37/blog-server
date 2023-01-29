package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/leon37/blog-server/pkg/app"
	"github.com/leon37/blog-server/pkg/errcode"
)

type Article struct {
}

func NewArticle() Article {
	return Article{}
}

func (a Article) Get(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
}
func (a Article) List(c *gin.Context)   {}
func (a Article) Create(c *gin.Context) {}
func (a Article) Update(c *gin.Context) {}
func (a Article) Delete(c *gin.Context) {}
