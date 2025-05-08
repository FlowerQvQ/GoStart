package service

import (
	"Final_System/biz"
	"Final_System/model"
	"Final_System/scheme"
	"github.com/gin-gonic/gin"
)

type AdminService struct {
	AdminBiz *biz.AdminBiz
}

func NewAdminService(adminBiz *biz.AdminBiz) *AdminService {
	return &AdminService{
		AdminBiz: adminBiz,
	}
}

// 通过Id查询用户信息
func (s *AdminService) GetInfosService(c *gin.Context) {
	var infos scheme.GetReq
	if err := c.ShouldBindJSON(&infos); err != nil {
		c.JSON(400, gin.H{
			"error": "参数错误",
		})
		return
	}
	//验证用户凭证
	//用user接收biz返回的信息
	user, err := s.AdminBiz.GetInfosBiz(infos.Id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "查询失败",
		})
		return
	}
	//从user里获取除密码外的其他信息
	getWithoutPassword := &model.User{
		Id:       user.Id,
		Username: user.Username,
		Phone:    user.Phone,
		Sex:      user.Sex,
		IsAdmin:  user.IsAdmin,
	}
	c.JSON(200, gin.H{
		"message": "查询成功",
		"data":    getWithoutPassword,
	})
}

// 添加

func (s *AdminService) AddInfosService(c *gin.Context) {
	var addInfos scheme.AddReq
	if err := c.ShouldBindJSON(&addInfos); err != nil {
		c.JSON(400, gin.H{
			"error": "参数错误",
		})
		return
	}
	//AddInfosBiz需要的参数是model.User类型，这里把addInfos转换成model.User类型
	addNewUser := &model.User{
		Username: addInfos.Username,
		Password: addInfos.Password,
		Phone:    addInfos.Phone,
		Sex:      addInfos.Sex,
	}
	//权限验证，只有管理员能添加用户
	if !UserAuth(c) {
		c.JSON(403, gin.H{
			"error": "权限不足，不能添加用户",
		})
		return
	}
	newUser, err := s.AdminBiz.AddInfosBiz(*addNewUser)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "添加失败",
		})
		return
	}

	addWithoutPassword := &model.User{
		Id:       newUser.Id,
		Username: newUser.Username,
		Phone:    newUser.Phone,
		Sex:      newUser.Sex,
		IsAdmin:  newUser.IsAdmin,
	}
	c.JSON(200, gin.H{
		"message": "添加成功",
		"data":    addWithoutPassword,
	})
}

// 修改

func (s *AdminService) UpdateInfosService(c *gin.Context) {
	var updateInfos scheme.UpdateReq
	if err := c.ShouldBindJSON(&updateInfos); err != nil {
		c.JSON(400, gin.H{
			"error": "参数错误",
		})
		return
	}

	Infos := &model.User{
		Id:       updateInfos.Id,
		Password: updateInfos.Password,
		Username: updateInfos.Username,
		Phone:    updateInfos.Phone,
		Sex:      updateInfos.Sex,
	}

	//权限验证
	if !CheckPermission(c, updateInfos.Id) {
		c.JSON(403, gin.H{
			"error": "权限不足，不能修改他人信息",
		})
		return
	}

	user, err := s.AdminBiz.UpdateInfosBiz(*Infos)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "修改失败",
		})
		return
	}
	updatedInfos := &model.User{
		Id:       user.Id,
		Username: user.Username,
		Phone:    user.Phone,
		Sex:      user.Sex,
		IsAdmin:  user.IsAdmin,
	}
	if updateInfos.Id == 0 {
		c.JSON(400, gin.H{
			"error": "无效的ID",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "修改成功",
		"data":    updatedInfos,
	})
	return

}

//删除

func (s *AdminService) DeleteInfosService(c *gin.Context) {
	var delInfos scheme.DeleteReq
	if err := c.ShouldBindJSON(&delInfos); err != nil {
		c.JSON(400, gin.H{
			"error": "参数错误",
		})
		return
	}
	if delInfos.Id == 0 {
		c.JSON(400, gin.H{
			"error": "无效的ID",
		})
		return
	}

	//权限验证
	if !CheckPermission(c, delInfos.Id) {
		c.JSON(403, gin.H{
			"error": "权限不足，不能删除他人信息",
		})
		return
	}

	if err := s.AdminBiz.DeleteInfosBiz(delInfos.Id); err != nil {
		c.JSON(400, gin.H{
			"error": "删除失败",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "删除成功",
		"data":    delInfos.Id,
	})

}
