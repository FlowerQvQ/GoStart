package data

import "Final_System/model"

type AdminData struct {
	DB *Data
}

func NewAdminData(data *Data) *AdminData {
	return &AdminData{
		DB: data,
	}

}

// 获取信息
func (d *AdminData) GetInfosData(id int) (*model.User, error) {
	var result *model.User
	err := d.DB.DBClient.Model(&model.User{}).Where("id = ?", id).First(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 添加信息
func (d *AdminData) AddInfosData(user model.User) (*model.User, error) {
	err := d.DB.DBClient.Model(&model.User{}).Create(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// 修改信息
func (d *AdminData) UpdateInfosData(user model.User) (*model.User, error) {
	err := d.DB.DBClient.Model(&model.User{}).Where("id = ?", user.Id).Updates(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// 删除信息
func (d *AdminData) DeleteInfosData(id int) error {

	var user model.User
	// 检查用户是否存在
	err := d.DB.DBClient.Model(&model.User{}).Where("id = ?", id).First(&user).Error
	if err != nil {
		return err
	}

	err = d.DB.DBClient.Model(&model.User{}).Where("id = ?", id).Delete(&model.User{}).Error
	if err != nil {
		return err
	}
	return nil
}
