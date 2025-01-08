// 코드 10.2 github.com/caarlos0/env/ 패키지를 사용해 환경 변수 읽기

package config

import (
	"github.com/caarlos0/env/"
)

type Config struct {
	Env  string `env:"TODO_ENV" envDefault:"dev"`
	Port int    `env:"PORT" envDefault:"80"`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}