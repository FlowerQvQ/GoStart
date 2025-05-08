package middleWare

import (
	"Final_System/biz"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
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
		//ParseWithClaims()函数会自动判断token是否过期，不需要再写判断语句
		//解析JWT令牌并验证其签名
		claims, err := jwt.ParseWithClaims(tokenStr, &biz.MyClaims{}, func(token *jwt.Token) (interface{}, error) {
			//确保签名方法是HMAC
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			//返回用于签名验证的密钥
			return []byte(biz.SecretKey), nil
		})
		//如果令牌解析出错，返回401错误响应
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			}) //401
			c.Abort()
			return
		}
		//类型断言，将解析后的claims转换为MyClaims类型，
		//如果成功，则将解析后的claims存储在Gin上下文中，并继续处理请求
		//如果类型断言失败，则返回500错误响应
		if claims, ok := claims.Claims.(*biz.MyClaims); ok {
			c.Set("id", claims.Id)
			c.Set("is_admin", claims.IsAdmin)
			c.Next()
		} else {
			c.AbortWithStatusJSON(500, gin.H{
				"error": "claims parsing failed",
			})
		}

	}
}

//func AuthMiddleware() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		//从上下文获取当前用户信息
//		id, _ := c.Get("id")
//		isAdmin, _ := c.Get("is_admin")
//		targetUserID := c.Query("userId") // 从查询参数中获取用户ID
//
//		//管理员直接放行
//		if isAdmin.(int) == 1 {
//			c.Next()
//			return
//		}
//		//普通用户检验ID一致性
//		if fmt.Sprintf("%v", id) != targetUserID {
//			c.JSON(http.StatusForbidden, gin.H{
//				"error": "无权访问他人数据",
//			})
//			c.Abort()
//			return
//		}
//		c.Next()
//	}
//}
