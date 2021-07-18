package config

type JWT struct {
	Secret     string `json:"secret" yaml:"secret" mapstructure:"secret"`
	ExpireTime int64    `json:"expire_time" yaml:"expire-time" mapstructure:"expire-time"`
	SignatureName string `json:"signature_name" yaml:"signature-name" mapstructure:"signature-name"`
}
