package main

import (
	"Final_System/biz"
	"Final_System/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/wire"
	_ "github.com/google/wire"
	"net/http"
	"time"
)

/*
声明结构体， jwt.RegisteredClaims(内嵌标准声明)包含7个字段
type RegisteredClaims struct {
    Issuer    string   // [iss] 签发机构标识
    Subject   string   // [sub] 令牌主题标识（如用户ID）
    Audience  []string // [aud] 接收方标识（可多个）
    ExpiresAt *NumericDate // [exp] 过期时间（Unix时间戳）
    NotBefore *NumericDate // [nbf] 生效时间
    IssuedAt  *NumericDate // [iat] 签发时间
    ID        string   // [jti] 令牌唯一标识
}
*/

// 自定义JWT结构体

//根据请求中携带的信息，生成token并返回给客户端，并由客户端保存
//客户端携带请求头中携带的token向下请求，
//中间件收到请求，先解析token，然后验证token是否正确---→错误返回错误信息
//正确，则向下请求，通过Context.Context获取上下文信息，然后进行业务逻辑

// 验证token
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取 token
		tokenStr := c.GetHeader("Authorization") //Authorization用来接token，前端传也要用这个名字传token
		if tokenStr == "" {
			c.JSON(401, gin.H{
				"error": "Missing token",
			})
			c.Abort()
			return
		}
		// 解析 token
		//ParseWithClaims()函数会自动判断token是否过期，不需要再写判断函数
		token, err := jwt.ParseWithClaims(tokenStr, &biz.MyClaims{}, func(token *jwt.Token) (interface{}, error) {
			//确保签名方法是HMAC
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(biz.SecretKey), nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			}) //401
			c.Abort()
			return
		}
		// 验证 token 是否有效
		if claims, ok := token.Claims.(*biz.MyClaims); ok && token.Valid {
			// 检查 token 是否过期，问题？
			if time.Now().Unix() > claims.ExpiresAt.Unix() {
				c.JSON(http.StatusUnauthorized, gin.H{
					"error": "Token expired",
				})
				c.Abort()
				return
			}
			// 将用户ID存储在上下文中，以便后续使用
			c.Set("user_id", claims.Id)
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})
			c.Abort()
		}
	}
}

// App 结构体,作用：封装所有服务，作为程序的核心容器，用于组织和管理service层的依赖关系
type App struct {
	UserService  *service.UserService
	AdminService *service.AdminService
}

func (a *App) SetRouter(engin *gin.Engine) {
	userRouter := engin.Group("/user")
	userRouter.POST("/register", a.UserService.RegisterService)
	userRouter.POST("/login", a.UserService.LoginService)
	adminRouter := engin.Group("/admin")
	adminRouter.Use(JWTAuthMiddleware())
	adminRouter.GET("/getUserInfo", a.AdminService.GetInfosService)
	adminRouter.POST("/addUserInfo", a.AdminService.AddInfosService)
	adminRouter.PUT("/updateUserInfo", a.AdminService.UpdateInfosService)
	adminRouter.DELETE("/deleteUserInfo", a.AdminService.DeleteInfosService)

}
func initGenEngine(app *App) *gin.Engine {
	//Logger 中间件用于记录每个请求的基本信息，包括请求路径、请求方法、请求状态码、响应时间等。这对于监控应用和调试问题非常有用。
	//Recovery 中间件主要作用是从任何恐慌（panics）中恢复，并在发生恐慌时写入 HTTP 状态码 500。
	//在 Go 语言里，当使用 panic() 时，程序会崩溃退出。
	//而 gin.Recovery 中间件会捕获这种异常，避免程序因未处理的异常而崩溃退出，保证服务的持续运行。
	//不过，对于链接断开的情况，不会有 HTTP 状态码返回
	engin := gin.Default() // gin.Default()默认包含 Logger 和 Recovery 中间件

	app.SetRouter(engin)

	return engin

}

//这段代码的功能是使用 wire 库创建一个依赖注入的集合，并将 App 结构体注册到该集合中
//wire.Struct(new(App), "*")：将 App 结构体及其所有字段标记为需要依赖注入。
//wire.NewSet(...)：创建一个新的依赖注入集合，包含上述结构体的定义。

var NewApp = wire.NewSet(wire.Struct(new(App), "*"))

//func NewApp(userService *service.UserService, adminService *service.AdminService) *App {
//	return &App{
//		UserService:  userService,
//		AdminService: adminService,
//	}
//}

func main() {
	engin := InitApp()

	err := engin.Run()
	if err != nil {
		return
	}

}
