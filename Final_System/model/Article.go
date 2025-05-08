package model

type Article struct {
	Id      int    `gorm:"column:id;type:int(11);primary_key" json:"id"`
	Title   string `gorm:"column:title;type:varchar(30);NOT NULL" json:"title"`
	Content string `gorm:"column:content;type:text;NOT NULL" json:"content"`
	UserID  int    `gorm:"column:user_id;type:int(11);NOT NULL" json:"user_id"`
	Status  int    `gorm:"column:status;default:0 ;NOT NULL" json:"status"`
	//Comments []Comment `gorm:"ForeignKey:comment_Id;AssociationForeignKey:Id" json:"comment"` //关联评论
}

func (Article) TableName() string {
	return "articles"
}
