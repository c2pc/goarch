package config

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"time"
)

const (
	defaultHTTPRWTimeout          = 10 * time.Second
	defaultHTTPMaxHeaderMegabytes = 1
)

type HttpConfig struct {
	Host               string `env:"HTTP_HOST" envDefault:"localhost"`
	Port               string `env:"HTTP_PORT" envDefault:"8089"`
	ReadTimeout        time.Duration
	WriteTimeout       time.Duration
	MaxHeaderMegabytes int
}

func GetHttpConfig() *HttpConfig {
	c := HttpConfig{}
	if err := env.Parse(&c); err != nil {
		fmt.Printf("%+v\n", err)
	}

	c.ReadTimeout = defaultHTTPRWTimeout
	c.WriteTimeout = defaultHTTPRWTimeout
	c.MaxHeaderMegabytes = defaultHTTPMaxHeaderMegabytes

	return &c
}
