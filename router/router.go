package router

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"go-template/app/controllers"
	"go-template/app/middlewares"
	"net/http"
)

func InitRouter() {
	router := gin.Default()
	// 要在路由组之前全局使用「跨域中间件」, 否则会返回404
	router.Use(middlewares.Cors())

	// var jwt = middlewares.Auth()
	base := router.Group("/app/v1")
	{
		admin := base.Group("/admin")
		{
			banner := admin.Group("/banner")
			{
				banner.GET("", controllers.GetBanners)
				banner.POST("", controllers.CreateBanners)
				banner.PUT("", controllers.UpdateBanner)
				banner.DELETE("", controllers.DeleteBanner)
			}
			user := admin.Group("/user")
			{
				user.GET("", controllers.Verify)
				user.POST("", controllers.UpdateUser)
				user.POST("/login", controllers.Login)
				user.GET("/refresh", controllers.Refresh)
			}
			admin.POST("/file", controllers.UploadFile)
		}

		//cases := v1.Group("/case")
		//{
		//	cases.GET("",controllers.GetCases)
		//	cases.GET("/:regionId",controllers.GetRegionCases)
		//	cases.POST("",controllers.CreateCase)
		//	cases.PUT("",controllers.UpdateCase)
		//	cases.DELETE("",controllers.DeleteCase)
		//}
		//intelligence := v1.Group("/intelligence")
		//{
		//	intelligence.GET("",controllers.GetIntelligences)
		//	intelligence.POST("",controllers.CreateIntelligence)
		//	intelligence.PUT("",controllers.UpdateIntelligence)
		//	intelligence.DELETE("",controllers.DeleteIntelligence)
		//}
		//link := v1.Group("/link")
		//{
		//	link.GET("",controllers.GetLinks)
		//	link.POST("",controllers.CreateLink)
		//	link.PUT("",controllers.UpdateLink)
		//	link.DELETE("",controllers.DeleteLink)
		//}
		//maps := v1.Group("/map")
		//{
		//	maps.GET("",controllers.GetMap)
		//	maps.PUT("",controllers.UpdateMap)
		//}
		//news := v1.Group("/new")
		//{
		//	news.GET("",controllers.GetNews)
		//	news.GET("/:id",controllers.GetNewsDetail)
		//	news.POST("",controllers.CreateNew)
		//	news.PUT("",controllers.UpdateNew)
		//	news.DELETE("",controllers.DeleteNew)
		//}
	}
	router.StaticFS("/static", http.Dir("./static"))
	//router.StaticFile("/static/file", "../static/file/cacheArea/ZQEUANJFFSPVOHLINGOE.jpg")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":12577")
}
