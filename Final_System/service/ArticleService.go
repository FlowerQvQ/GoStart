package service

import (
	"Final_System/biz"
	"Final_System/model"
	"Final_System/scheme"
	"github.com/gin-gonic/gin"
)

type ArticleService struct {
	ArticleBiz *biz.ArticleBiz
}

func NewArticleService(articleBiz *biz.ArticleBiz) *ArticleService {
	return &ArticleService{
		ArticleBiz: articleBiz,
	}

}

// 文章发布
func (s *ArticleService) PublishArticleService(c *gin.Context) {
	var publishReq scheme.PublishArticleReq
	if err := c.ShouldBindJSON(&publishReq); err != nil {
		c.JSON(400, gin.H{
			"error": "参数错误",
		})
		return
	}
	newArticle := model.Article{
		Title:   publishReq.Title,
		Content: publishReq.Content,
		UserID:  publishReq.UserId,
	}

	if !CheckPermission(c, publishReq.UserId) {
		c.JSON(403, gin.H{"error": "权限不足，不能操作他人文章"})
		return
	}
	reArticle, err := s.ArticleBiz.PublishArticleBiz(newArticle)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "发布失败",
		})
		return
	}
	reArticle = &model.Article{
		Id:      reArticle.Id,
		Title:   reArticle.Title,
		Content: reArticle.Content,
		UserID:  reArticle.UserID,
	}
	c.JSON(200, gin.H{
		"data":    reArticle,
		"message": "发布成功",
	})
}

// 文章查询
func (s *ArticleService) GetArticleService(c *gin.Context) {
	var getReq scheme.GetArticleReq
	if err := c.ShouldBindJSON(&getReq); err != nil {
		c.JSON(400, gin.H{
			"error": "参数错误",
		})
		return
	}
	article, err := s.ArticleBiz.GetArticleBiz(getReq.Id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "查询失败",
		})
		return
	}
	getArticle := &model.Article{
		Id:      article.Id,
		Title:   article.Title,
		Content: article.Content,
		UserID:  article.UserID,
	}

	c.JSON(200, gin.H{
		"message": "查询成功",
		"data":    getArticle,
	})
}

// 文章修改
func (s *ArticleService) UpdateArticleService(c *gin.Context) {
	var updateReq scheme.UpdateArticleReq
	if err := c.ShouldBindJSON(&updateReq); err != nil {
		c.JSON(400, gin.H{
			"error": "参数错误",
		})
		return
	}

	updateArticle := &model.Article{
		Id:      updateReq.Id,
		Title:   updateReq.Title,
		Content: updateReq.Content,
		UserID:  updateReq.UserId,
	}
	if !CheckPermission(c, updateReq.UserId) {
		c.JSON(403, gin.H{"error": "无权修改他人文章"})
		return
	}
	reArticle, err := s.ArticleBiz.UpdateArticleBiz(*updateArticle)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "修改失败",
		})
		return
	}
	updateArticle = &model.Article{
		Id:      reArticle.Id,
		Title:   reArticle.Title,
		Content: reArticle.Content,
		UserID:  reArticle.UserID,
	}

	c.JSON(200, gin.H{
		"data":    updateArticle,
		"message": "修改成功",
	})
}

// 文章删除
func (s *ArticleService) DeleteArticleService(c *gin.Context) {
	var deleteReq scheme.DeleteReq
	if err := c.ShouldBindJSON(&deleteReq); err != nil {
		c.JSON(400, gin.H{
			"error": "参数错误",
		})
		return
	}
	if !CheckPermission(c, deleteReq.Id) {
		c.JSON(403, gin.H{
			"error": "无权删除他人文章",
		})
		return
	}
	if err := s.ArticleBiz.DeleteArticleBiz(deleteReq.Id); err != nil {
		c.JSON(400, gin.H{
			"error": "删除失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "删除成功",
		"data":    deleteReq.Id,
	})
}
