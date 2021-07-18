package request

type LoginRequest struct {
	Username      string `json:"username,omitempty"`
	Password      string `json:"password,omitempty"`
	CaptchaId     string `json:"captcha_id,omitempty"`
	CaptchaString string `json:"captcha_string,omitempty"`
}
