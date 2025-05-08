package data

import (
	"Final_System/model"
)

//UserData 结构体封装了与数据相关的操作。
//它持有一个指向 Data 类型的指针，用于执行数据操作

type UserData struct {
	DB *Data
}

// NewUserData 创建并返回一个新的UserData实例。
// 该函数接收一个指向Data类型的指针作为参数，将其用作初始化UserData结构体的DB字段。

//返回值:返回一个指向UserData类型的指针，其DB字段被初始化为传入的Data指针

func NewUserData(data *Data) *UserData {
	return &UserData{
		DB: data,
	}

}

// 注册数据层

func (d *UserData) RegisterData(user model.User) (model.User, error) {

	err := d.DB.DBClient.Model(&model.User{}).Create(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, err //直接用&user返回数据库中拿到新注册的用户信息
}

// 登录数据层
func (d *UserData) LoginData(username, password string) (*model.User, error) {

	var user model.User
	err := d.DB.DBClient.Model(&model.User{}).Where("username = ? AND password = ?", username, password).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
