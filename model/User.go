package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"query/utils/errmsg"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(500);not null" json:"password" validate:"required,min=8,max=20" label:"密码"`
}

// 添加用户
func AddUser(data *User) int {
	err = db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除用户
func DeleteUserById(id int) int {
	err = db.Where("id = ? ",id).Delete(&User{}).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 编辑用户
func EditUserById(id int,data *User) int {
	data.Password = ""
	err = db.Model(&User{}).
		Where("id = ? ",id).
		Updates(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 通过id查询单个用户
func FindUserById(id int) (User,int) {
	var user User
	err = db.Where("ID = ? ",id).First(&user).Error
	if err != nil {
		return user,errmsg.ERROR
	}
	return user,errmsg.SUCCESS
}

// 通过用户名查询用户id
func FindUserByUsername(username string) int {
	var userList []User
	db.Where("username = ? ", username).
		Find(&userList)
	if userList == nil {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

// 获取用户列表
func GetUserListByUsername(username string,pageSize, pageNum int) ([]User,int64) {
	var userList []User
	var total int64

	db.Model(&User{}).
		Where("username LIKE ? ",username+"%").
		Count(&total).
		Limit(pageSize).
		Offset((pageNum-1)*pageSize).
		Find(&userList)

	return userList,total
}

// 修改密码
func RePassword(id int,data *User) int {
	err = db.Select("password").Where("id = ?", id).Updates(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 密码加密
func (user *User)BeforeCreate(_ *gorm.DB) error {
	user.Password = cryptPassword(user.Password)
	return nil
}

func (user *User)BeforeUpdate(_ *gorm.DB) error {
	user.Password = cryptPassword(user.Password)
	return nil
}

// 密码加密
func cryptPassword(password string) string {
	const cost = 10
	hash,err := bcrypt.GenerateFromPassword([]byte(password),cost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hash)
}

// 登录
func Login(username string,password string) (User,int) {
	var user User
	db.Where("username = ? ",username).First(&user)
	if user.ID == 0 {
		return user,errmsg.ERROR_USER_NOT_EXIST
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password))
	if err != nil {
		return user,errmsg.ERROR_PASSWORD_WRONG
	}
	return user,errmsg.SUCCESS
}