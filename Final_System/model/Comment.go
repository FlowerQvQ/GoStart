package model

type Comment struct {
	Id        int    `gorm:"column:id;type:int(11);primary_key" json:"id"`
	Content   string `gorm:"column:content;type:varchar(255);NOT NULL" json:"content"`
	UserId    int    `gorm:"column:user_id;type:int(11);NOT NULL" json:"user_id"`
	ArticleId int    `gorm:"column:article_id;type:int(11);NOT NULL" json:"article_id"`
}

func (Comment) TableName() string {
	return "comments"
}
