package biz

import "github.com/google/wire"

var ProviderSetBiz = wire.NewSet(
	NewAdminBiz,
	NewUserBiz,
	NewArticleBiz,
	NewCommentBiz,
)
