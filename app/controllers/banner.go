package controllers

import (
	"github.com/gin-gonic/gin"
	"go-template/app/model"
	"go-template/app/utils"
	"net/http"
)

// @Summary 获取banner列表
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success {"data": } *model.Banner{}
// @Router /controllers/v1/GetBanners [get]
func GetBanners(c *gin.Context){
	data,err := model.GetBanners()
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"msg": "查询出错",
		})
	} else {
		c.JSON(http.StatusOK,gin.H{
			"data": data,
		})
	}

}

func CreateBanners(c *gin.Context){
	var data model.Banner
	if err :=c.BindJSON(&data);err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"msg":"绑定出错",
			"err":err.Error(),
		})
		return
	}
	if data.Url != "" {
		if urlResult,err := utils.MoveFileToS(data.Url);err != nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"msg":"文件移动出错",
				"err":err.Error(),
			})
		} else{
			data.Url = urlResult
		}
	}
	if err :=model.CreateBanner(data);err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"msg":"创建出错",
			"err":err.Error(),
		})
	} else {
		c.JSON(http.StatusOK,gin.H{
			"msg": "创建成功",
		})
	}
}

func UpdateBanner(c *gin.Context){
	var data model.Banner
	if err :=c.BindJSON(&data);err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"msg":"绑定出错",
			"err":err.Error(),
		})
		return
	}
	if queryResult,err :=model.FindBanner(data.ID);err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"msg":"根据id索引查询失败",
			"err":err.Error(),
		})
		return
	} else if len(queryResult) == 0 {
		c.JSON(http.StatusOK,gin.H{
			"msg":"没有找到该条记录",
		})
		return
	}
	if data.ID == 0 {
		c.JSON(http.StatusBadRequest,gin.H{
			"msg":"id不能为空",
		})
		return
	}
	if data.Url != "" {
		if urlResult,err := utils.MoveFileToS(data.Url);err != nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"msg":"文件移动出错",
				"err":err.Error(),
			})
		} else{
			data.Url = urlResult
		}
	}
	if err :=model.UpdateBanner(data);err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"msg":"更新出错",
			"err":err.Error(),
		})
	} else {
		c.JSON(http.StatusOK,gin.H{
			"msg": "更新成功",
		})
	}
}

func DeleteBanner(c *gin.Context) {
	var data model.Banner
	if err :=c.BindJSON(&data);err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"msg":"绑定出错",
			"err":err.Error(),
		})
		return
	}
	if queryResult,err :=model.FindBanner(data.ID);err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"msg":"根据id索引查询失败",
			"err":err.Error(),
		})
		return
	} else if len(queryResult) == 0 {
		c.JSON(http.StatusOK,gin.H{
			"msg":"没有找到该条记录",
		})
		return
	}
	if data.ID ==0 {
		c.JSON(http.StatusBadRequest,gin.H{
			"msg":"id不能为空",
		})
		return
	}
	if err :=model.DeleteBanner(data);err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"msg":"删除成功",
			"err":err.Error(),
		})
	} else {
		c.JSON(http.StatusOK,gin.H{
			"msg": "删除成功",
		})
	}
	if data.Url != "" {
		if err := utils.RemoveFile(data.Url);err != nil{
			// 列入消息队列
			// 删除达到最大次数记录到日志当中
			c.JSON(http.StatusOK,gin.H{
				"msg":"文件删除失败",
			})
		}
	}
}