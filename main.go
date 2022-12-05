package main

import (
	"query/model"
	"query/routes"
)

func main() {
	// 引用数据库
	model.InitDb()

	//config.Init()

	routes.InitRouter()
}
