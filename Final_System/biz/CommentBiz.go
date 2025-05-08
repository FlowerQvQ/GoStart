package biz

import (
	"Final_System/data"
	"Final_System/model"
)

type CommentBiz struct {
	CommentData *data.CommentData
}

func NewCommentBiz(commentData *data.CommentData) *CommentBiz {
	return &CommentBiz{
		CommentData: commentData,
	}
}

// 查询评论
func (b *CommentBiz) GetCommentBiz(id int) (model.Comment, error) {
	comment, err := b.CommentData.GetCommentData(id)
	if err != nil {
		return model.Comment{}, err
	}
	return comment, err
}

// 发布评论
func (b *CommentBiz) PublishCommentBiz(comment model.Comment) (*model.Comment, error) {

	return b.CommentData.PublishCommentData(comment)
}

// 删除评论
func (b *CommentBiz) DeleteCommentBiz(id int) error {
	return b.CommentData.DeleteCommentData(id)
}
