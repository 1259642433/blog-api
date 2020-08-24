package model

import (
	"github.com/jinzhu/gorm"
	"go-template/dao"
)

type Banner struct {
	gorm.Model
	Title string `json:"title"`
	Url string `json:"url"`
	Description string `json:"description"`
	Link string `json:"link"`
	Type int `json:"type"`
}

func GetBanners()(data []*Banner,err error){
	if err:=dao.DB.Find(&data).Error;err != nil{
		return nil,err
	}
	return
}

func FindBanner(id uint)(data []*Banner,err error){
	if err:=dao.DB.Where("id = ?",id).Find(&data).Error;err != nil{
		return nil,err
	}
	return
}

func CreateBanner(data Banner)(err error){
	if err:=dao.DB.Create(&data).Error;err != nil{
		return err
	}
	return
}

func UpdateBanner(data Banner)(err error){
	if err:=dao.DB.Model(&Banner{}).Updates(&data).Error;err != nil{
		return err
	}
	return
}

func DeleteBanner(data Banner)(err error){
	if err:=dao.DB.Delete(&data).Error;err != nil{
		return err
	}
	return
}