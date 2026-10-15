package config

import "github.com/caarlos0/env/v11"

type Config struct {
	Telegram Telegram
}

func Load() (Config, error) {
	return env.ParseAs[Config]()
}
