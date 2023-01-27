package config

import (
	"flag"
	"github.com/caarlos0/env"
)

type Config struct {
	URLSiteMap string `env:"START_URL"`
	Debug      bool   `env:"SERVER_DEBUG"`
}

func InitConfig() (*Config, error) {
	var cfg Config

	flag.StringVar(&cfg.URLSiteMap, "url", "https://yandex.ru/support/sitemap.xml", "url=https://***")
	flag.BoolVar(&cfg.Debug, "debug", false, "debug=true")
	//flag.StringVar(&cfg.Path, "path", tPath, "path=C:\\temp\\00")

	flag.Parse()

	err := env.Parse(&cfg)

	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
