package controller

import (
	"github.com/HSczy/CLBackendGO/global"
	"github.com/HSczy/CLBackendGO/middlewares"
	"github.com/HSczy/CLBackendGO/model/request"
	"github.com/HSczy/CLBackendGO/model/response"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func LoginHandler(c *gin.Context) {
	login := request.LoginRequest{}
	err := c.ShouldBindJSON(login)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.WithMessage("传入参数有问题，请检查传入参数"))
		return
	}
	if login.Username == "123" && login.Password == "456" {
		if store.Verify(login.CaptchaId, login.CaptchaString, true) {
			tokenNext(c, login.Username)
		} else {
			c.JSON(http.StatusBadRequest, response.WithMessage("验证码错误"))
			return
		}
	}

}

func tokenNext(ctx *gin.Context, username string) {
	j := middlewares.NewJWT()

	claims := request.MyClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                         // 签名生效时间
			ExpiresAt: time.Now().Unix() + global.CONFIG.JWT.ExpireTime, // 过期时间 7天  配置文件
			Issuer:    global.CONFIG.JWT.SignatureName,                  // 签名的发行者
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.WithMessage("生成JWT token错误"))
		return
	}
	ctx.JSON(http.StatusBadRequest, response.WithData("生成JWT token成功", response.LoginResponse{
		User:      username,
		Token:     token,
		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	}))
	return
}
