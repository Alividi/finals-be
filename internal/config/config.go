package config

type Config struct {
	App AppConfig
}

type AppConfig struct {
	Name     string
	HttpPort int
	AppURL   string
}
