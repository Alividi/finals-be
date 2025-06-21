package config

import (
	"finals-be/internal/lib/utils"
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig
	DB       DatabaseConfig
	JWT      JWTConfig
	Firebase FirebaseConfig
	AWS      AWSConfig
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

type FirebaseConfig struct {
	Type                    string `json:"type"`
	ProjectID               string `json:"project_id"`
	PrivateKeyID            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientID                string `json:"client_id"`
	AuthURI                 string `json:"auth_uri"`
	TokenURI                string `json:"token_uri"`
	AuthProviderX509CertURL string `json:"auth_provider_x509_cert_url"`
	ClientX509CertURL       string `json:"client_x509_cert_url"`
	UniverseDomain          string `json:"universe_domain"`
}

type AWSConfig struct {
	AccessKeyID     string
	SecretAccessKey string
	Region          string
	BucketName      string
	BaseURL         string
}

func getAppConfig() AppConfig {
	return AppConfig{
		Name:     utils.GetStringOrPanic("APP_NAME"),
		HttpPort: utils.GetIntOrDefault("HTTP_PORT", 8080),
		AppURL:   utils.GetStringOrPanic("APP_URL"),
	}
}

func getDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		Driver:              utils.GetStringOrPanic("DB_DRIVER"),
		URL:                 utils.GetStringOrPanic("DB_URL"),
		MaxIdleConnections:  utils.GetIntOrDefault("DB_MAX_IDLE_CONNECTIONS", 20),
		MaxOpenConnections:  utils.GetIntOrDefault("DB_MAX_OPEN_CONNECTIONS", 50),
		MaxIdleDuration:     time.Duration(utils.GetIntOrDefault("DB_MAX_IDLE_DURATION", 60)) * time.Minute,
		MaxLifeTimeDuration: time.Duration(utils.GetIntOrDefault("DB_MAX_CONN_LIFETIME", 30)) * time.Minute,
		MigrationPath:       utils.GetStringOrPanic("DB_MIGRATION_PATH"),
	}
}

func getAPIConfig() JWTConfig {
	return JWTConfig{
		JWTSecretKey:              utils.GetStringOrPanic("JWT_SECRET_KEY"),
		LoginExpirationDuration:   time.Duration(utils.GetIntOrDefault("JWT_LOGIN_EXPIRATION_DURATION_HOUR", 24)) * time.Hour,
		RefreshExpirationDuration: time.Duration(utils.GetIntOrDefault("JWT_REFRESH_EXPIRATION_DURATION_DAYS", 7)) * 24 * time.Hour,
	}
}

func getFirebaseConfig() FirebaseConfig {
	return FirebaseConfig{
		Type:                    utils.GetStringOrPanic("FIREBASE_TYPE"),
		ProjectID:               utils.GetStringOrPanic("FIREBASE_PROJECT_ID"),
		PrivateKeyID:            utils.GetStringOrPanic("FIREBASE_PRIVATE_KEY_ID"),
		PrivateKey:              utils.GetStringOrPanic("FIREBASE_PRIVATE_KEY"),
		ClientEmail:             utils.GetStringOrPanic("FIREBASE_CLIENT_EMAIL"),
		ClientID:                utils.GetStringOrPanic("FIREBASE_CLIENT_ID"),
		AuthURI:                 utils.GetStringOrPanic("FIREBASE_AUTH_URI"),
		TokenURI:                utils.GetStringOrPanic("FIREBASE_TOKEN_URI"),
		AuthProviderX509CertURL: utils.GetStringOrPanic("FIREBASE_AUTH_PROVIDER_X509_CERT_URL"),
		ClientX509CertURL:       utils.GetStringOrPanic("FIREBASE_CLIENT_X509_CERT_URL"),
		UniverseDomain:          utils.GetStringOrPanic("FIREBASE_UNIVERSE_DOMAIN"),
	}
}

func getAWSConfig() AWSConfig {
	return AWSConfig{
		AccessKeyID:     utils.GetStringOrPanic("AWS_ACCESS_KEY_ID"),
		SecretAccessKey: utils.GetStringOrPanic("AWS_SECRET_ACCESS_KEY"),
		Region:          utils.GetStringOrPanic("AWS_REGION"),
		BucketName:      utils.GetStringOrPanic("AWS_BUCKET_NAME"),
		BaseURL:         utils.GetStringOrPanic("AWS_BASE_URL"),
	}
}

func LoadConfigByFile(path, fileName, fileType string) *Config {
	viper.SetConfigName(fileName)
	viper.SetConfigType(fileType)
	viper.AddConfigPath(path)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %v", err)
	}

	return &Config{
		App:      getAppConfig(),
		DB:       getDatabaseConfig(),
		JWT:      getAPIConfig(),
		Firebase: getFirebaseConfig(),
		AWS:      getAWSConfig(),
	}
}
