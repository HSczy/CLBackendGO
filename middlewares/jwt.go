package middlewares

import (
	"errors"
	"github.com/HSczy/CLBackendGO/global"
	"github.com/HSczy/CLBackendGO/model/request"
	"github.com/HSczy/CLBackendGO/model/response"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.Request.Header.Get("jwt-token")
		j := NewJWT()
		if tokenString == "" {
			ctx.JSON(http.StatusBadRequest, response.WithMessage("没有找到jwt数据，请重新登陆"))
			ctx.Abort()
			return
		}
		claims, err := j.ParseToken(tokenString)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, response.WithMessage(err.Error()))
			ctx.Abort()
			return
		}
		ctx.Set("claims", claims)
		ctx.Next()
	}
}

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token")
)

func NewJWT() *JWT {
	return &JWT{SigningKey: []byte(global.CONFIG.JWT.Secret)}
}

// CreateToken 创建JWT token
func (j *JWT) CreateToken(claims request.MyClaims) (tokenString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken 解析JWT token
func (j *JWT) ParseToken(tokenString string) (*request.MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &request.MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*request.MyClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	}
	return nil, TokenInvalid
}
