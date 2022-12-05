package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"query/utils"
	"time"
)

var (
	db  *gorm.DB
	err error
)

func InitDb() {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	)
	db, err = gorm.Open(mysql.Open(dns),&gorm.Config{
		// gorm日志模式
		Logger: logger.Default.LogMode(logger.Silent),
		// 外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
		// 禁用默认事务（提高运行速度）
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			// 使用单数表名，启用该选项，此时，`User`的表名应该是`user`
			SingularTable: true,
		},
	})
	if err != nil {
		log.Println("Connect Database Fail :", err)
		os.Exit(1)
	}

	// 迁移数据表，在没有数据表结构变更时候，建议注释不执行
	db.AutoMigrate(&User{}, &Swindler{})

	sqlDB,_ := db.DB()
	// 最大闲置连接数
	sqlDB.SetMaxIdleConns(10)

	// 数据库最大连接数
	sqlDB.SetMaxOpenConns(100)

	// 连接最大可复用时间
	sqlDB.SetConnMaxLifetime(18 * time.Second)

	//db.Close()
}