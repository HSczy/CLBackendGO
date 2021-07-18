package controller

import (
	"github.com/HSczy/CLBackendGO/global"
	"github.com/HSczy/CLBackendGO/model/response"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"net/http"
)

var store = base64Captcha.DefaultMemStore

func CaptchaHandler(ctx *gin.Context) {
	// 字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(
		global.CONFIG.CAPTCHA.ImgHeight,
		global.CONFIG.CAPTCHA.ImgWidth,
		global.CONFIG.CAPTCHA.KeyLong,
		0.7,
		80,
	)
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := cp.Generate()
	if err != nil {
		ctx.JSON(http.StatusExpectationFailed, response.WithMessage(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.WithData("生成成功", response.CaptchaResponse{
		CaptchaId: id,
		PicPath:   b64s,
	}))
	return

}
