package config

type EsscoAPI struct {
	Token string `yaml:"token" env:"APP_ESSCO_TOKEN"`
	URL   string `yaml:"url" env:"APP_ESSCO_URL"`
}
