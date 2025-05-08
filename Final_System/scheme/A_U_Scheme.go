package scheme

// 注册请求结构体

type RegisterReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=8,max=16"`
	Phone    string `json:"phone" binding:"required"`
	Sex      string `json:"sex" binding:"required"`
}

// 登录请求结构体

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=8,max=16"`
}

//登录token结构体

type LoginRsp struct {
	UserName string `json:"username"`
	IsAdmin  int    `json:"is_admin"`
	Token    string `json:"token"`
}

// 查询请求结构体

type GetReq struct {
	Id int `json:"id" binding:"required"`
}

// 添加请求结构体

type AddReq struct {
	Id       int    `json:"id" binding:""`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=8,max=16"`
	Phone    string `json:"phone" binding:"required"`
	Sex      string `json:"sex" binding:"required"`
}

// 修改请求结构体

type UpdateReq struct {
	Id       int    `json:"id" binding:"required"`
	Username string `json:"username" `
	Password string `json:"password" binding:"min=8,max=16"`
	Phone    string `json:"phone" `
	Sex      string `json:"sex" `
}

// 删除请求结构体

type DeleteReq struct {
	Id int `json:"id" binding:"required"`
}
