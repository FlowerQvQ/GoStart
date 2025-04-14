package service

import (
	"Final_System/biz"
	"Final_System/model"
	"Final_System/scheme"
	"github.com/gin-gonic/gin"
)

type UserService struct {
	UserBiz *biz.UserBiz
}

// 通过构造函数注入依赖, *UserService依赖UserBiz

func NewUserService(userBiz *biz.UserBiz) *UserService {
	//返回一个初始化后的 UserService 实例
	return &UserService{
		UserBiz: userBiz,
	}
}

// RegisterService 注册服务
func (s *UserService) RegisterService(c *gin.Context) {
	var registerReq scheme.RegisterReq
	if err := c.ShouldBindJSON(&registerReq); err != nil {
		c.JSON(400, gin.H{
			"error": "参数错误",
		})
		return
	}
	//把接收到的信息封装成User结构体
	receiveUserInfo := &model.User{
		Username: registerReq.Username,
		Password: registerReq.Password,
		Phone:    registerReq.Phone,
		Sex:      registerReq.Sex,
	}
	newUser, err := s.UserBiz.RegisterBiz(*receiveUserInfo)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "注册失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "注册成功",
		"data":    newUser,
	})

}

// LoginService 登录服务
func (s *UserService) LoginService(c *gin.Context) {

	var loginReq scheme.LoginReq
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(400, gin.H{
			"error": "参数错误",
		})
		return
	}

	if _, err := s.UserBiz.LoginBiz(loginReq.Username, loginReq.Password); err != nil {
		c.JSON(400, gin.H{
			"error": "账号或密码错误",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "登录成功",
	})

}
