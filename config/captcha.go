package config

type Captcha struct {
	KeyLong   int `json:"key_long" yaml:"key-long" mapstructure:"key-long"`
	ImgWidth  int `json:"img_width" yaml:"img-width" mapstructure:"img-width"`
	ImgHeight int `json:"img_height" yaml:"img-height" mapstructure:"img-height"`
}
