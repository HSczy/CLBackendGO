package response

type CaptchaResponse struct {
	CaptchaId string `json:"captcha_id,omitempty"`
	PicPath   string `json:"pic_path,omitempty"`
}
