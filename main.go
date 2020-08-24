package main

import (
	"go-template/app/model"
	"go-template/dao"
	"go-template/router"
)

func main() {

	// 连接完毕关闭数据库
	defer dao.DB.Close()

	//同步数据库
	dao.DB.AutoMigrate(&model.Banner{},&model.User{})

	//挂载路由
	router.InitRouter()

	// 打印详细日志
	dao.DB.LogMode(true)

	//todo 对接nsq
	//middlewares.Consumer()
	//middlewares.Producer()
}
