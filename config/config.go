package config

type Config struct {
	//子结构体需要增加mapstructure属性后才能解析数字
	MYSQL   Mysql   `yaml:"mysql" json:"mysql" mapstructure:"mysql"`
	JWT     JWT     `yaml:"jwt" json:"jwt" mapstructure:"jwt"`
	CAPTCHA Captcha `yaml:"captcha" json:"captcha" mapstructure:"captcha"`
}
