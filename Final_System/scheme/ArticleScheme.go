package scheme

// 文章发布：
// 用户文章发布结构体
type PublishArticleReq struct {
	Id      int    `json:"id" binding:""`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	UserId  int    `json:"user_id" binding:"required"`
}

// 文章查询
type GetArticleReq struct {
	Id int `json:"id" binding:"required"`
}

// 文章修改
type UpdateArticleReq struct {
	Id      int    `json:"id" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	UserId  int    `json:"user_id" binding:"required"`
}

// 文章删除
type DeleteArticleReq struct {
	Id int `json:"id" binding:"required"`
}
