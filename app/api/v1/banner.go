package v1

import (
	"blog-api/app/models"
	"blog-api/app/utils"
	"blog-api/pkg/e"
	"blog-api/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"log"
	"net/http"
)

// @Summary 获取banner列表
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success {"data": } *modelss.Banner{}
// @Router /api/v1/GetBanners [get]
func GetBanners(c *gin.Context){
	//maps := make(map[string]interface{})
	//data,err := models.GetBanners(10,1,maps)
	//count,err1 := models.GetBannerTotal(maps)
	//if err != nil||err1 != nil {
	//	c.JSON(http.StatusOK,gin.H{
	//		"code": '',
	//		"msg": "查询出错",
	//		"data": err,
	//	})
	//} else {
	//	c.JSON(http.StatusOK,gin.H{
	//		"data": data,
	//	})
	//}
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})



	if name != "" {
		maps["name"] = name
	}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS

	log.Print("%+v",maps)

	data["lists"],_ = models.GetBanners(util.GetPageVar(c), maps)
	data["total"],_ = models.GetBannerTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})
}

func CreateBanner(c *gin.Context){
	var data models.Banner
	if err :=c.BindJSON(&data);err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"msg": "绑定出错",
			"err": err.Error(),
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
	if err :=models.CreateBanner(data);err != nil{
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
	var data models.Banner
	if err :=c.BindJSON(&data);err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"msg":"绑定出错",
			"err":err.Error(),
		})
		return
	}
	if queryResult,err :=models.FindBanner(data.ID);err != nil{
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
	if err :=models.UpdateBanner(data);err != nil{
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
	var data models.Banner
	if err :=c.BindJSON(&data);err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"msg":"绑定出错",
			"err":err.Error(),
		})
		return
	}
	if queryResult,err :=models.FindBanner(data.ID);err != nil{
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
	if err :=models.DeleteBanner(data);err != nil{
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