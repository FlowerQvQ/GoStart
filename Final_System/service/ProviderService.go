package service

import "github.com/google/wire"

// 使用wire.NewSet方法组合一组provider set
// 依赖分组管理 wire.NewSet()这个函数可以把相关的 provider 组合在一起然后使用。当然也可以单独使用，

var ProviderSetService = wire.NewSet(
	NewAdminService,
	NewUserService,
	NewArticleService,
	NewCommentService,
)
