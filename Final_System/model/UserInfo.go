package model

// 用户表对数据库的映射

type User struct {
	Id       int    `gorm:"column:id;type:int(11);primary_key" json:"id"`
	Username string `gorm:"column:username;type:varchar(255);unique;NOT NULL" json:"username"`
	Password string `gorm:"column:password;type:varchar(255);NOT NULL" json:"password"`
	Phone    string `gorm:"column:phone;type:varchar(255);NOT NULL" json:"phone"`
	Sex      string `gorm:"column:sex;type:varchar(255);NOT NULL" json:"sex"`
	IsAdmin  int    `gorm:"column:is_admin;type:tinyint(4);default:0;NOT NULL" json:"is_admin"`
}

func (User) TableName() string {
	return "usersInfo"
}
