package config

import (
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Log    LogConfig    `yaml:"log"`
	Server ServerConfig `yaml:"server"`
	Minio  MinioConfig  `yaml:"minio"`
	DB     DBConfig     `yaml:"database"`
}

type LogConfig struct {
	Level string `yaml:"level"`
}

type ServerConfig struct {
	Addr string `yaml:"addr"`
}

type MinioConfig struct {
	AccessKey  string `yaml:"access_key" envconfig:"MINIO_ACCESS_KEY"`
	SecretKey  string `yaml:"secret_key" envconfig:"MINIO_SECRET_KEY"`
	BucketName string `yaml:"bucket" envconfig:"MINIO_BUCKET" default:"media_server"`

	EndPoint string `yaml:"endpoint" envconfig:"MINIO_ENDPOINT" default:"localhost:9000"`
	UseSSL   bool   `yaml:"ssl"`
}

type DBConfig struct {
	User     string `yaml:"user" envconfig:"DB_USER"`
	Password string `yaml:"password" envconfig:"DB_PASSWORD"`
	DBName   string `yaml:"dbname" envconfig:"DB_NAME" default:"media_server"`
	Host     string `yaml:"host" envconfig:"DB_HOST" default:"localhost"`
	Port     int    `yaml:"port" envconfig:"DB_PORT" default:"5432"`
}

func New(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var conf Config

	if err := yaml.Unmarshal(data, &conf); err != nil {
		return nil, err
	}

	return &conf, envconfig.Process("", &conf)
}
