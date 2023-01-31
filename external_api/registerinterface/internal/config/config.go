package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	SecretKey string

	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
}
