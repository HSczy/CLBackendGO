package request

import (
	"github.com/dgrijalva/jwt-go"
)

// MyClaims 自定义JWT载荷中承载的内容，后续可以自行添加
type MyClaims struct {
	Username    string `json:"username"`
	jwt.StandardClaims
}
