package config

import (
	"github.com/Devil666face/telecho/pkg/lib"

	"github.com/ilyakaznacheev/cleanenv"
)

var defaultConfigPath = []string{
	lib.MustPath("~/.config/telecho/telecho.yaml"),
	lib.MustPath("~/.config/telecho/telecho.yml"),
	lib.MustPath("~/.config/telecho/.telecho.env"),
	lib.MustPath("~/telecho.yaml"),
	lib.MustPath("~/telecho.yml"),
	lib.MustPath("~/.telecho.env"),
	lib.MustPath("telecho.yaml"),
	lib.MustPath("telecho.yml"),
	lib.MustPath(".telecho.env"),
}

type Config struct {
	BotToken string   `env:"BOT_TOKEN" yaml:"token" env-required:"true"`
	GroupsID []string `env:"GROUPS_ID" yaml:"groups_id" env-required:"true"`
}

func New(path string) (*Config, error) {
	cfg := Config{}
	if path != "" {
		defaultConfigPath = append(defaultConfigPath, path)
	}
	for i := len(defaultConfigPath) - 1; i >= 0; i-- {
		if err := cleanenv.ReadConfig(defaultConfigPath[i], &cfg); err == nil {
			return &cfg, nil
		}
	}
	return &cfg, cleanenv.ReadEnv(&cfg)
}
