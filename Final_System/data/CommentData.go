package data

import "Final_System/model"

type CommentData struct {
	DB *Data
}

func NewCommentData(data *Data) *CommentData {
	return &CommentData{
		DB: data,
	}
}

// 查询评论
func (d *CommentData) GetCommentData(id int) (model.Comment, error) {
	var comment model.Comment
	err := d.DB.DBClient.Model(&model.Comment{}).Where("id = ?", id).First(&comment).Error
	if err != nil {
		return model.Comment{}, err
	}
	return comment, nil
}

//评论发布

func (d *CommentData) PublishCommentData(comment model.Comment) (*model.Comment, error) {
	err := d.DB.DBClient.Model(&model.Comment{}).Create(&comment).Error
	if err != nil {
		return nil, err
	}
	return &model.Comment{}, err
}

// 删除评论
func (d *CommentData) DeleteCommentData(id int) error {
	var comment model.Comment
	err := d.DB.DBClient.Model(&model.Comment{}).Where("id = ?", id).First(&comment).Error
	if err != nil {
		return err
	}
	err = d.DB.DBClient.Model(&model.Comment{}).Where("id = ?", id).Delete(&model.Comment{}).Error
	if err != nil {
		return err
	}
	return nil
}
