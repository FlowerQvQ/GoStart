package data

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 封装一个结构体，里边包含与数据库交互所有必要信息和功能
type Data struct {
	DBClient *gorm.DB
}

// provider for *Data 没有任何依赖
func NewData() *Data {
	//返回一个初始化后的 Data 实例
	return &Data{
		DBClient: NewMysqlClient(), //通过调用函数NewMysqlClient()生成MySql数据库实例
	}
}

// 数据库客户端

func NewMysqlClient() *gorm.DB {
	//连接数据库
	dsn := "root:123456@tcp(127.0.0.1:3306)/finally?charset=utf8mb4&parseTime=True&loc=Local&timeout=10s"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), //显示sql语句
	})
	if err != nil {
		panic("数据库连接失败: " + err.Error())
	}

	return db
}
