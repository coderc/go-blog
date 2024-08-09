package config

import (
	"fmt"

	"time"
)

type EnvParam struct {
	Config   string `env:"CONFIG, required"`   // 配置文件路径
	GIN_MODE string `env:"GIN_MODE, required"` // gin模式
}

func (e *EnvParam) GetConfig() string {
	if e == nil {
		return ""
	}

	return e.Config
}

type Config struct {
	Server Server `mapstructure:"server" validate:"required"`
	Jwt    Jwt    `mapstructure:"jwt" validate:"required"`
	Mysql  Mysql  `mapstructure:"mysql" validate:"required"`
	Redis  Redis  `mapstructure:"redis" validate:"required"`
}

func (c *Config) Print() {
	fmt.Printf("%+v\n", c)
}

type Mysql struct {
	Host     string `mapstructure:"host" validate:"required"`
	Port     int    `mapstructure:"port" validate:"required"`
	User     string `mapstructure:"user" validate:"required"`
	Password string `mapstructure:"password"`
	DbName   string `mapstructure:"db_name" validate:"required"`
}

type Redis struct {
	Host               string        `mapstructure:"host" validate:"required"`
	Port               int           `mapstructure:"port" validate:"required"`
	Password           string        `mapstructure:"password"`
	Db                 int           `mapstructure:"db" validate:"required"`
	PingTimeout        time.Duration `mapstructure:"ping_timeout" validate:"required"`
	PrefixHouseKeeping string        `mapstructure:"prefix_housekeeping" validate:"required"`
}

type Jwt struct {
	Secret string `mapstructure:"secret" validate:"required"` // jwt密钥
}

type Server struct {
	Env      string `mapstructure:"env"  validate:"required"`      // 环境变量
	Port     int    `mapstructure:"port" validate:"required"`      // 服务端口
	Debug    bool   `mapstructure:"debug" validate:"required"`     // 是否开启debug模式
	LogDir   string `mapstructure:"log_dir" validate:"required"`   // 日志目录
	LogLevel string `mapstructure:"log_level" validate:"required"` // 日志级别
}
