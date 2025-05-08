package biz

import (
	"Final_System/data"
	"Final_System/model"
)

type ArticleBiz struct {
	ArticleData *data.ArticleData
}

func NewArticleBiz(articleData *data.ArticleData) *ArticleBiz {
	return &ArticleBiz{
		ArticleData: articleData,
	}
}

// 文章发表
func (b *ArticleBiz) PublishArticleBiz(article model.Article) (*model.Article, error) {
	return b.ArticleData.PublishArticleData(article)
}

// 文章查询
func (b *ArticleBiz) GetArticleBiz(id int) (model.Article, error) {
	article, err := b.ArticleData.GetArticleData(id)
	if err != nil {
		return model.Article{}, err
	}
	return article, nil
}

// 文章修改
func (b *ArticleBiz) UpdateArticleBiz(article model.Article) (*model.Article, error) {

	return b.ArticleData.UpdateArticleData(article)
}

// 文章删除
func (b *ArticleBiz) DeleteArticleBiz(id int) error {
	return b.ArticleData.DeleteArticleData(id)
}
