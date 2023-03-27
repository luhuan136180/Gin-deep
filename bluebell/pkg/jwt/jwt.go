package jwt

//标注：该jwt库已经停止更新，
//若需要新库github.com/golang-jwt/jwt/v4
import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

//定义JWT的过期时间，这里以2小时为例：
const TokenExpireDuration = time.Hour * 24

//定义密钥MySecret
var MySecret = []byte("夏天夏天悄悄过去")

//定义需求
//定制自己的需求来决定JWT中保存哪些数据
// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	UserID   int64  `json:"user-id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

//生成JWT
// GenToken 生成JWT
func GenToken(userID int64, username string) (string, error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		userID,
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), //过期时间
			Issuer:    "bluebell",                                 //签发人

		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(MySecret)
}

//解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	//解析Token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if Claims, ok := token.Claims.(*MyClaims); ok && token.Valid { //校验token
		return Claims, nil
	}
	return nil, errors.New("invalid token")
}
