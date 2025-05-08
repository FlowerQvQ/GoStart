package router

import (
	"Final_System/middleWare"
	"Final_System/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// App 结构体,作用：封装所有服务，作为程序的核心容器，用于组织和管理service层的依赖关系
type App struct {
	UserService    *service.UserService
	AdminService   *service.AdminService
	ArticleService *service.ArticleService
	CommentService *service.CommentService
}

//wire依赖注入（创建并初始化 App 实例）方式：
//1. 使用 wire.NewSet 的方式
//这段代码的功能是使用 wire 库创建一个依赖注入的集合，并将 App 结构体注册到该集合中
//wire.Struct(new(App), "*") 表示将 App 结构体的所有字段自动注入
//wire.NewSet(...)：创建一个提供 *App 类型值的集合，供 Wire 在生成代码时使用。
//特点：
//基于代码生成，没有运行时反射开销。
//编译期检查依赖是否完整。
//更适合大型项目或需要高性能 DI 的场景

var NewApp = wire.NewSet(wire.Struct(new(App), "*"))

//2.手动构造函数的方式
//特点：
//需要开发者显式传入每个依赖项。
//没有使用任何依赖注入框架，是传统的手动依赖注入方式。
//更直观、简单，但在依赖较多或复杂时维护成本高。
//func NewApp(userService *service.UserService, adminService *service.AdminService,) *App {
//	return &App{
//		UserService:  userService,
//		AdminService: adminService,
//		ArticleService *service.ArticleService
//		CommentService *service.CommentService
//	}
//}

func (a *App) SetRouter(engin *gin.Engine) {
	//登录注册
	userRouter := engin.Group("/user")
	userRouter.POST("/register", a.UserService.RegisterService)
	userRouter.POST("/login", a.UserService.LoginService)
	//信息增删改查
	adminRouter := engin.Group("/admin")
	adminRouter.Use(middleWare.JWTAuthMiddleware())
	adminRouter.GET("/getUserInfo", a.AdminService.GetInfosService)
	adminRouter.POST("/addUserInfo", a.AdminService.AddInfosService)
	adminRouter.PUT("/updateUserInfo", a.AdminService.UpdateInfosService)
	adminRouter.DELETE("/deleteUserInfo", a.AdminService.DeleteInfosService)
	//文章增删改查
	articleRouter := engin.Group("/article")
	articleRouter.Use(middleWare.JWTAuthMiddleware())
	articleRouter.POST("/publishArticle", a.ArticleService.PublishArticleService)
	articleRouter.GET("/getArticle", a.ArticleService.GetArticleService)
	articleRouter.PUT("/updateArticle", a.ArticleService.UpdateArticleService)
	articleRouter.DELETE("/deleteArticle", a.ArticleService.DeleteArticleService)
	//评论增删查
	commentRouter := engin.Group("/comment")
	commentRouter.Use(middleWare.JWTAuthMiddleware())
	commentRouter.POST("/publishComment", a.CommentService.PublishCommentService)
	commentRouter.GET("/getComment", a.CommentService.GetCommentService)
	commentRouter.DELETE("/deleteGetComment", a.CommentService.DeleteCommentService)

}

func InitGenEngine(app *App) *gin.Engine {

	engin := gin.Default() // gin.Default()默认包含 Logger 和 Recovery 中间件
	//Logger 中间件用于记录每个请求的基本信息，包括请求路径、请求方法、请求状态码、响应时间等。这对于监控应用和调试问题非常有用。
	//Recovery 中间件主要作用是从任何恐慌（panics）中恢复，并在发生恐慌时写入 HTTP 状态码 500。
	//在 Go 语言里，当使用 panic() 时，程序会崩溃退出。
	//而 gin.Recovery 中间件会捕获这种异常，避免程序因未处理的异常而崩溃退出，保证服务的持续运行。
	//不过，对于链接断开的情况，不会有 HTTP 状态码返回

	app.SetRouter(engin)

	return engin
}
