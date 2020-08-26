package util

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

type PageVar struct{
	page int
	size int
}

func GetPageVar(c *gin.Context) (result PageVar) {
	result.page, _ = com.StrTo(c.Query("page")).Int()
	result.size,_ = com.StrTo(c.Query("size")).Int()
	return
}

//func GetPageSize(c *gin.Context) int {
//	result := 0
//	page, _ := com.StrTo(c.Query("size")).Int()
//	if page > 0 {
//		result = (page - 1) * setting.PageSize
//	}
//
//	return result
//}