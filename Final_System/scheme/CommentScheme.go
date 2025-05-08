package scheme

// 评论发布
type PublishCommentReq struct {
	ArticleId int    `json:"article_id" binding:"required"`
	Content   string `json:"content" binding:"required"`
	UserId    int    `json:"user_id" binding:"required"`
}

// 评论查询
type GetCommentReq struct {
	Id int `json:"id" binding:"required"`
}

// 评论删除
type DelCommentReq struct {
	Id int `json:"id" binding:"required"`
}
