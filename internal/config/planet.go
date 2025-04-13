package config

type PlanetAPI struct {
	Token string `yaml:"token" env:"APP_PLANET_TOKEN" env-required:"true"`
	URL   string `yaml:"url" env:"APP_PLANET_URL" env-required:"true"`
}
