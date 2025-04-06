package config

import "time"

type Config struct {
	App AppConfig
	DB  DatabaseConfig
	JWT JWTConfig
}

type AppConfig struct {
	Name     string
	HttpPort int
	AppURL   string
}

type DatabaseConfig struct {
	Driver              string
	URL                 string
	MaxIdleConnections  int
	MaxOpenConnections  int
	MaxIdleDuration     time.Duration
	MaxLifeTimeDuration time.Duration
	MigrationPath       string
}

type JWTConfig struct {
	JWTSecretKey              string
	LoginExpirationDuration   time.Duration
	RefreshExpirationDuration time.Duration
}
