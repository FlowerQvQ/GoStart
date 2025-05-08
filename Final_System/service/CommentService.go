package service

import (
	"Final_System/biz"
	"Final_System/model"
	"Final_System/scheme"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CommentService struct {
	CommentBiz *biz.CommentBiz
}

func NewCommentService(commentBiz *biz.CommentBiz) *CommentService {
	return &CommentService{
		CommentBiz: commentBiz,
	}
}

// 发布评论
func (s *CommentService) PublishCommentService(c *gin.Context) {
	var publishComment scheme.PublishCommentReq
	if err := c.ShouldBindJSON(&publishComment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数错误",
		})
		return
	}
	commentReq := model.Comment{
		UserId:    publishComment.UserId,
		ArticleId: publishComment.ArticleId,
		Content:   publishComment.Content,
	}
	if CheckPermission(c, publishComment.UserId) {

		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "权限不足，不能替他人评论",
		})
	}
	comment, err := s.CommentBiz.PublishCommentBiz(commentReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "评论失败",
		})
		return
	}
	comment = &model.Comment{
		Id:        comment.Id,
		UserId:    comment.UserId,
		ArticleId: comment.ArticleId,
		Content:   comment.Content,
	}
	c.JSON(http.StatusOK, gin.H{
		"data": comment,
		"msg":  "评论成功",
	})
}

// 查询评论
func (s *CommentService) GetCommentService(c *gin.Context) {
	var getComment scheme.GetCommentReq
	if err := c.ShouldBindJSON(&getComment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数错误",
		})
		return
	}

	comment, err := s.CommentBiz.GetCommentBiz(getComment.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "查询失败",
		})
		return
	}
	getComments := &model.Comment{
		Id:        comment.Id,
		UserId:    comment.UserId,
		ArticleId: comment.ArticleId,
		Content:   comment.Content,
	}
	c.JSON(http.StatusOK, gin.H{
		"data": getComments,
		"msg":  "查询成功",
	})
}

// 删除评论
func (s *CommentService) DeleteCommentService(c *gin.Context) {
	var deleteComment scheme.DelCommentReq
	if err := c.ShouldBindJSON(&deleteComment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数错误",
		})
		return
	}
	if CheckPermission(c, deleteComment.Id) {

	}
	err := s.CommentBiz.DeleteCommentBiz(deleteComment.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "删除失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": deleteComment.Id,
		"msg":  "删除成功",
	})
}
