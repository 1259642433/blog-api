package models

import (
	"blog-api/pkg/util"
)

type Banner struct {
	Model
	Title string `json:"title"`
	Url string `json:"url"`
	Description string `json:"description"`
	Link string `json:"link"`
	Type int `json:"type"`
}

//获取多个banner
func GetBanners(page util.PageVar, maps interface {})(data []Banner,err error){
	err = db.Where(maps).Offset(page.page).Limit(page.size).Find(&data).Error
	return
}

// 查询指定id的banner
func FindBanner(id uint)(data []*Banner,err error){
	err =db.Where("id = ?",id).Find(&data).Error
	return
}

// 创建banner记录
func CreateBanner(data Banner)(err error){
	err = db.Create(&data).Error
	return
}

//修改指定banner数据
func UpdateBanner(data Banner)(err error){
	err=db.Model(&Banner{}).Updates(&data).Error
	return
}

//删除指定banner
func DeleteBanner(data Banner)(err error){
	err = db.Delete(&data).Error
	return
}

func GetBannerTotal(maps interface {}) (count int,err error){
	err = db.Model(&Banner{}).Where(maps).Count(&count).Error
	return
}