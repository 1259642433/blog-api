package dao

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
)

var (
	DB *gorm.DB
)

func init(){
	var err error
	dsn := "root:WWt1989829@tcp(144.34.165.131:3306)/blog?charset=utf8mb4&parseTime=true"
	DB,err = gorm.Open("mysql",dsn)
	if err!=nil {
		fmt.Printf("mysql连接错误： %v", err)
		panic(err)
	}
	if DB.Error != nil {
		fmt.Printf("数据库错误 %v", DB.Error)
		panic(DB.Error)
	}
}