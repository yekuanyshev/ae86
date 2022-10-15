package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DBHost     string `envconfig:"DB_HOST" required:"true"`
	DBPort     string `envconfig:"DB_PORT" required:"true"`
	DBUser     string `envconfig:"DB_USER" required:"true"`
	DBPassword string `envconfig:"DB_PASSWORD" required:"true"`
	DBName     string `envconfig:"DB_NAME" required:"true"`
	DBSSLMode  string `envconfig:"DB_SSLMODE" default:"disable"`

	MinioHost     string `envconfig:"MINIO_HOST"`
	MinioPort     string `envconfig:"MINIO_PORT"`
	MinioUser     string `envconfig:"MINIO_USER"`
	MinioPassword string `envconfig:"MINIO_PASSWORD"`
	MinioUseSSL   bool   `envconfig:"MINIO_USE_SSL"`
	MinioBucket   string `envconfig:"MINIO_BUCKET"`

	HTTPHost      string `envconfig:"HTTP_HOST" default:"localhost"`
	HTTPPort      string `envconfig:"HTTP_PORT" default:"8080"`
	HTTPTLSEnable bool   `envconfig:"HTTP_TLS_ENABLE" default:"false"`
	HTTPCertFile  string `envconfig:"HTTP_CERTFILE"`
	HTTPKeyFile   string `envconfig:"HTTP_KEYFILE"`
}

func Load(envPrefix string, filenames ...string) (Config, error) {
	config := Config{}

	err := godotenv.Load(filenames...)
	if err != nil {
		return Config{}, err
	}

	err = envconfig.Process(envPrefix, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
