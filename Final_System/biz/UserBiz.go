package biz

import (
	"Final_System/data"
	"Final_System/model"
	"Final_System/scheme"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type UserBiz struct {
	UserData *data.UserData
}

func NewUserBiz(userData *data.UserData) *UserBiz {
	return &UserBiz{
		UserData: userData,
	}

}

// 注册业务逻辑

func (b *UserBiz) RegisterBiz(user model.User) (*model.User, error) {
	newUser := model.User{
		Username: user.Username,
		Password: user.Password,
		Phone:    user.Phone,
		Sex:      user.Sex,
	}
	return b.UserData.RegisterData(newUser)
}

// 登录业务逻辑
type MyClaims struct {
	Id                   int `json:"id"`
	jwt.RegisteredClaims     //嵌入jwt的注册声明
}

// 从登录开始接收到前端的信息，带着信息去数据库查询，查询成功后生成token，
// 校验身份，校验成功生成token，然后把token返回，后续操作都是用token去验证，验证通过后才能操作
func (b *UserBiz) LoginBiz(username, password string) (scheme.LoginRsp, error) {
	var rsp scheme.LoginRsp
	user, err := b.UserData.LoginData(username, password)
	if err != nil {
		return scheme.LoginRsp{}, err //出现错误返回LoginRsp结构体的零实例
	}
	//生成tolken返回给前端
	rsp.Token, err = GenerateToken(user.Id, username)
	rsp.UserName = user.Username
	return rsp, nil
}

// 把秘钥设置为常量，方便以后修改
const SecretKey = "secret" //秘钥名字

// 生成token
func GenerateToken(id int, username string) (string, error) { //返回值是string类型，是因为返回的是生成的token字符串
	//1.定义Claims
	claims := MyClaims{
		Id: id,
		//里边的名字是自定义的，发行者名字不能写死，传进去的人不一样身份也不一样
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    username,                                          //发行者
			Subject:   "test",                                            //主题
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)), //设置过期时间
			NotBefore: jwt.NewNumericDate(time.Now()),                    //设置token生效时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                    //设置发行时间
		},
	}
	//2.创建token对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //NewWithClaims()初始化token结构体header
	//3.使用秘钥签名解析并生成token字符串
	//最好定义为一个常量secret
	tokenString, err := token.SignedString([]byte(SecretKey)) //secret是秘钥名，解析也要用这个名字
	if err != nil {
		return "", err
	}
	return tokenString, nil //返回生成的token字符串

}
