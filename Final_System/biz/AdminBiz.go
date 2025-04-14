package biz

import (
	"Final_System/data"
	"Final_System/model"
)

type AdminBiz struct {
	AdminData *data.AdminData
}

func NewAdminBiz(adminData *data.AdminData) *AdminBiz {
	return &AdminBiz{
		AdminData: adminData,
	}
}

// GetInfosBiz 根据id获取管理员信息
func (b *AdminBiz) GetInfosBiz(id int) (*model.User, error) {
	result, err := b.AdminData.GetInfosData(id)
	return result, err
}

// 添加用户信息

func (b *AdminBiz) AddInfosBiz(user model.User) (*model.User, error) {
	result, err := b.AdminData.AddInfosData(user)
	return result, err
}

// 修改信息

func (b *AdminBiz) UpdateInfosBiz(user model.User) (*model.User, error) {
	result, err := b.AdminData.UpdateInfosData(user)
	return result, err
}

// 删除信息
func (b *AdminBiz) DeleteInfosBiz(id int) error {
	err := b.AdminData.DeleteInfosData(id)
	return err
}
