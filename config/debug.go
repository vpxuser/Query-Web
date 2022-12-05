package config

import (
	"log"
	"query/model"
)

func Init()  {
	var data model.User
	data.Username = "admin"
	data.Password = "1qaz@WSX"
	code := model.AddUser(&data)
	if code == 200 {
		log.Println("init admin success !!!")
	}
}