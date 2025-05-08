package service

import "github.com/gin-gonic/gin"

// 用户鉴权身份鉴定
func UserAuth(c *gin.Context) bool {
	//判断是否为管理员
	isAdmin, exists := c.Get("is_admin")
	//如果exists不为真或isAdmin不等于1，则返回false
	if !exists || isAdmin.(int) != 1 {
		return false
	}
	return true
}
func CheckPermission(c *gin.Context, reqId int) bool {
	loginId, exists := c.Get("id")
	if !exists {
		return false
	}
	auth := UserAuth(c)
	if auth {
		return true // 管理员有权限
	}
	// 普通用户只能操作自己的数据 //BUG删除时可以删除与用户Id一样的文章
	if loginId.(int) == reqId {
		return true
	}
	return false
}
