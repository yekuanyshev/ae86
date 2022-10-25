package config

import (
	"errors"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"os"
)

var Global *Config

type Config struct {
	DB struct {
		Host     string `envconfig:"DB_HOST" required:"true"`
		Port     string `envconfig:"DB_PORT" required:"true"`
		User     string `envconfig:"DB_USER" required:"true"`
		Password string `envconfig:"DB_PASSWORD" required:"true"`
		Name     string `envconfig:"DB_NAME" required:"true"`
		SSLMode  string `envconfig:"DB_SSLMODE" default:"disable"`
	}

	Minio struct {
		Enabled  bool   `envconfig:"MINIO_ENABLED" default:"false"`
		Host     string `envconfig:"MINIO_HOST"`
		Port     string `envconfig:"MINIO_PORT"`
		User     string `envconfig:"MINIO_USER"`
		Password string `envconfig:"MINIO_PASSWORD"`
		UseSSL   bool   `envconfig:"MINIO_USE_SSL"`
		Bucket   string `envconfig:"MINIO_BUCKET"`
	}

	HTTP struct {
		Host      string `envconfig:"HTTP_HOST"`
		Port      string `envconfig:"HTTP_PORT" default:"8000"`
		TLSEnable bool   `envconfig:"HTTP_TLS_ENABLE" default:"false"`
		CertFile  string `envconfig:"HTTP_CERTFILE"`
		KeyFile   string `envconfig:"HTTP_KEYFILE"`
	}
}

func Load(envPrefix string, filenames ...string) error {
	err := godotenv.Load(filenames...)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}

	conf := Config{}
	err = envconfig.Process(envPrefix, &conf)
	if err != nil {
		return err
	}

	Global = &conf
	return nil
}
