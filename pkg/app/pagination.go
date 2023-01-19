package app

import (
	"github.com/gin-gonic/gin"
	"github.com/leon37/blog-server/global"
	"github.com/leon37/blog-server/pkg/convert"
)

func GetPage(c *gin.Context) int {
	page := convert.StrTo(c.Query("page")).MustInt()
	if page <= 0 {
		page = 1
	}
	return page
}

func GetPageSize(c *gin.Context) int {
	pageSize := convert.StrTo(c.Query("page_size")).MustInt()
	if pageSize <= 0 {
		return global.AppSetting.DefaultPageSize
	}
	if pageSize > global.AppSetting.MaxPageSize {
		return global.AppSetting.MaxPageSize
	}
	return pageSize
}

func GetPageOffset(page, pageSize int) int {
	res := 0
	if page > 0 {
		res = (page - 1) * pageSize
	}
	return res
}
