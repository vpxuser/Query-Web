package model

import (
	"gorm.io/gorm"
	"query/utils/errmsg"
)

type Swindler struct {
	gorm.Model
	Tieba string `gorm:"type:varchar(13);not null" json:"tieba" label:"贴吧号"`
	Wechat string `gorm:"type:varchar(28)" json:"wechat" validate:"required,min=6,max=28" label:"微信号"`
	Phone string `gorm:"type:char(11)" json:"phone" label:"手机号"`
}

func AddSwindler(data *Swindler) int {
	err = db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func DeleteSwindlerById(id int) int {
	err = db.Where("id = ? ",id).Delete(&Swindler{}).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func EditSwindlerById(id int,data *Swindler) int {
	data.Tieba = ""
	err = db.Model(&Swindler{}).
		Where("id = ? ",id).
		Updates(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func FindSwindlerById(id int) (Swindler,int) {
	var swindler Swindler
	err = db.Where("ID = ? ",id).First(&swindler).Error
	if err != nil {
		return swindler,errmsg.ERROR
	}
	return swindler,errmsg.SUCCESS
}

func GetSwindlerList(pageSize, pageNum int) ([]Swindler,int64) {
	var swindlerList []Swindler
	var total int64

	db.Model(&Swindler{}).Count(&total).
		Limit(pageSize).
		Offset((pageNum-1)*pageSize).
		Find(&swindlerList)

	return swindlerList,total
}

func GetSwindlerListByTieba(tieba string,pageSize, pageNum int) ([]Swindler,int64) {
	var swindlerList []Swindler
	var total int64

	db.Model(&Swindler{}).
		Where("tieba = ? ",tieba).
		Count(&total).
		Limit(pageSize).
		Offset((pageNum-1)*pageSize).
		Find(&swindlerList)

	return swindlerList,total
}