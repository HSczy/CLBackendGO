package config

import "fmt"

type Mysql struct {
	Host          string `yaml:"host" json:"host" mapstructure:"host"`
	Config        string `mapstructure:"config" json:"config" yaml:"config"` // 高级配置
	Port          string `yaml:"port" json:"port" mapstructure:"port"`
	User          string `yaml:"user" json:"user" mapstructure:"user"`
	Password      string `yaml:"password" json:"password" mapstructure:"password"`
	Database      string `yaml:"database" json:"database" mapstructure:"database"`
	MaxIdleConnes int    `yaml:"max-idle-connes" json:"maxIdleConnes" mapstructure:"max-idle-connes"`
	MaxOpenConnes int    `yaml:"max-open-connes" json:"maxOpenConnes" mapstructure:"max-idle-connes"`
}

func (m *Mysql) GetPath() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", m.User, m.Password, m.Host, m.Port, m.Database, m.Config)
}
