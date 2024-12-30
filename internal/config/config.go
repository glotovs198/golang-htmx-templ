package config

type Config struct {
	AppHost   string
	AssetsDir string
}

func NewConfig() *Config {
	return &Config{
		AppHost:   "localhost:8080",
		AssetsDir: "assets",
	}
}
