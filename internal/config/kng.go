package config

type KngAPI struct {
	Token string `yaml:"token" env:"APP_KNG_TOKEN" env-required:"true"`
	URL   string `yaml:"url" env:"APP_KNG_URL" env-required:"true"`
}
