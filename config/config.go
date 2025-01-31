package config

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type ProjectConfig struct {
	Team string `env:"PROJECT_TEAM" env-default:"ngikut"`
	Name string `env:"PROJECT_NAME" env-default:"tutuplapak"`
}

type DataBaseConfig struct {
	Username string `env:"DB_USER"`
	Password string `env:"DB_PASS"`
	Host     string `env:"DB_HOST"`
	Port     uint16 `env:"DB_PORT"`
	SSLMode  string `env:"DB_SSLMODE" env-default:"require"`
}

func (db *DataBaseConfig) GetDSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		db.Username,
		db.Password,
		db.Host,
		db.Port,
		db.Username,
		db.SSLMode,
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
	FileMaxSize      int64 `env:"FILE_MAX_SIZE" env-default:"102400"`     // In bytes
	ThumbnailMaxSize int64 `env:"THUMBNAIL_MAX_SIZE" env-default:"10240"` // In bytes
}

type Config struct {
	ServerPort     string `env:"SERVER_PORT" env-default:"8080"`
	RequestTimeout int    `env:"REQUEST_TIMEOUT" env-default:"2"` // In seconds
	Project        ProjectConfig
	Database       DataBaseConfig
	AWS            AWSConfig
	File           FileConfig
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
