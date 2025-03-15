package config

type ProgressAPI struct {
	Token string `yaml:"token" env:"APP_PROGRESS_TOKEN"`
	URL   string `yaml:"url" env:"APP_PROGRESS_URL"`
}
