package data

import "Final_System/model"

type ArticleData struct {
	DB *Data
}

func NewArticleData(data *Data) *ArticleData {
	return &ArticleData{
		DB: data,
	}
}

// 用户文章发布（增）

func (d *ArticleData) PublishArticleData(article model.Article) (*model.Article, error) {

	err := d.DB.DBClient.Model(&model.Article{}).Create(&article).Error
	if err != nil {
		return nil, err
	}
	return &article, err
}

// 文章查询（查）
func (d *ArticleData) GetArticleData(id int) (model.Article, error) {
	var article model.Article
	err := d.DB.DBClient.Model(&model.Article{}).Where("id = ?", id).First(&article).Error
	if err != nil {
		return model.Article{}, err
	}
	return article, nil
}

// 文章修改（改）
func (d *ArticleData) UpdateArticleData(article model.Article) (*model.Article, error) {
	err := d.DB.DBClient.Model(&model.Article{}).Where("id = ?", article.Id).Updates(&article).First(&article).Error
	if err != nil {
		return nil, err
	}
	return &article, nil
}

// 文章删除（删）
func (d *ArticleData) DeleteArticleData(id int) error {
	var article model.Article
	// delete不会检查文章是否存在，所以需要写这个↓检查文章是否存在
	err := d.DB.DBClient.Model(&model.Article{}).Where("id = ?", id).First(&article).Error
	if err != nil {
		return err
	}
	err = d.DB.DBClient.Model(&model.Article{}).Where("id = ?", id).Delete(&model.Article{}).Error
	if err != nil {
		return err
	}
	return nil
}
