package models

import (
	"path"
	"runtime"

	"github.com/BurntSushi/toml"
)

type EnvConfig struct {
	Environment int

	Mysql struct {
		Dsn string
	}

	Redis struct {
		Addr     string
		Password string
	}

	Supervise struct {
		Ip   string
		Port int
		Flag bool
	}

	Log struct {
		Dir   bool `toml:"dir"`
		Alert struct {
			Open bool     `toml:"open"`
			At   []string `toml:"at"`
		} `toml:"alert"`
	} `toml:"log"`

	Wechat struct {
		MPAccessToken string `toml:"MPToken"`
	} `toml:"wechat"`
}

type EnvType int8

const (
	EnvDefault EnvType = 0
	EnvLocal   EnvType = 1
	EnvDevelop EnvType = 2
	EnvOnline  EnvType = 3
)

func (env EnvType) ToString() string {
	switch env {
	case EnvLocal:
		return "EnvLocal"
	case EnvDevelop:
		return "EnvDevelop"
	case EnvOnline:
		return "EnvOnline"
	default:
		return "EnvNil"
	}
}

var (
	envConf EnvConfig
)

// 仅供单元测试使用
func GetEnvPath() string {
	_, filename, _, _ := runtime.Caller(0)
	return path.Dir(filename)
}

func InitEnvConf(path string) {
	if _, err := toml.DecodeFile(path, &envConf); err != nil {
		log.Fatal(err)
	}
}

// local, develop, online
func (e *EnvConfig) GetEnvironment() EnvType {
	switch envConf.Environment {
	case 1:
		return EnvLocal
	case 2:
		return EnvDevelop
	case 3:
		return EnvOnline
	default:
		return EnvDefault
	}
}
