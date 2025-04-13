package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"sync"
)

type AppConfig struct {
	Planet       PlanetAPI   `yaml:"planet"`
	Kng          KngAPI      `yaml:"kng"`
	Progress     ProgressAPI `yaml:"progress"`
	Essco        EsscoAPI    `yaml:"essco"`
	MockAdapters string      `yaml:"mock_adapters" env:"APP_MOCK_ADAPTERS"`
}

var conf AppConfig

var once sync.Once

func GetConfig() AppConfig {
	once.Do(func() {

		path, ok := os.LookupEnv("CONF_PATH")
		log.Println("Config path ", path, ok)

		if !ok {
			err := cleanenv.ReadEnv(&conf)
			if err != nil {
				panic(err)
			}
		} else {
			err := cleanenv.ReadConfig(path, &conf)
			if err != nil {
				panic(err)
			}
		}

		log.Print(conf)
	})
	return conf
}
