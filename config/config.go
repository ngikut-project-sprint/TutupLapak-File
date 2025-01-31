package config

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type ProjectConfig struct {
	Team string `env:"PROJECT_TEAM"`
	Name string `env:"PROJECT_NAME"`
}

type DataBaseConfig struct {
	Username string `env:"DB_USER"`
	Password string `env:"DB_PASS"`
	Host     string `env:"DB_HOST"`
	Port     uint16 `env:"DB_PORT"`
}

func (db *DataBaseConfig) GetDSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		db.Username,
		db.Password,
		db.Host,
		db.Port,
		db.Username,
		"require",
	)
}

type AWSConfig struct {
	AccessKey    string `env:"AWS_ACCESS_KEY_ID"`
	SecretKey    string `env:"AWS_SECRET_ACCESS_KEY"`
	BucketName   string `env:"AWS_S3_BUCKET_NAME"`
	Region       string `env:"AWS_REGION"`
	SessionToken string `env:"AWS_SESSION_TOKEN"`
}

func (c *AWSConfig) GetCredential() credentials.StaticCredentialsProvider {
	return credentials.NewStaticCredentialsProvider(
		c.AccessKey,
		c.SecretKey,
		c.SessionToken,
	)
}

type FileConfig struct {
	// In bytes
	FileMaxSize      int64 `env:"FILE_MAX_SIZE"`
	ThumbnailMaxSize int64 `env:"THUMBNAIL_MAX_SIZE"`
}

type Config struct {
	Project  ProjectConfig
	Database DataBaseConfig
	AWS      AWSConfig
	File     FileConfig
}

func Load() (*Config, error) {
	cfg := &Config{}

	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	if err := cleanenv.ReadEnv(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
