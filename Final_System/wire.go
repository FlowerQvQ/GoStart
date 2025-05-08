//go:build wireinject
// +build wireinject

// wire.go不参与运算，只用来生成依赖注入代码
package main

import (
	"Final_System/biz"
	"Final_System/data"
	"Final_System/router"
	"Final_System/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitApp() *gin.Engine {
	panic(wire.Build( //wire.Build用于生成依赖注入代码，会根据给定的参数生成依赖注入代码。(里边的参数是构造函数列表)
		service.ProviderSetService,
		biz.ProviderSetBiz,
		data.ProviderSetData,
		router.NewApp,
		router.InitGenEngine,
	))

	//使用wire命令会自动生成一个wire_gen.go文件
}
