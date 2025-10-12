package config

import (
	"fmt"
	"log/slog"

	"github.com/joho/godotenv"
)

type StorageType string

const (
	StorageMinIO StorageType = "minio"
	StorageCeph  StorageType = "ceph"
)

type AppConfig interface {
	Load(map[string]string) error
	Validate() error
}

type ServerConfig struct {
	Host string
	Port int
}

type S3Config struct {
	// Общие поля для MinIO и Ceph
	Type            StorageType `env:"S3_TYPE"`              // Тип хранилища (например, "minio" или "ceph")
	UseSSL          bool        `env:"S3_USE_SSL"`           // Использовать SSL (true/false)
	Endpoint        string      `env:"S3_ENDPOINT"`          // Адрес эндпоинта (например, "s3.example.com:9000")
	AccessKeyID     string      `env:"S3_ACCESS_KEY_ID"`     // Ключ доступа
	SecretAccessKey string      `env:"S3_SECRET_ACCESS_KEY"` // Секретный ключ доступа
	BucketName      string      `env:"S3_BUCKET_NAME"`       // Имя бакета
	Region          string      `env:"S3_REGION"`            // Регион (если требуется)
	PathStyle       bool        `env:"S3_PATH_STYLE"`        // Использовать path style (true для Ceph, false для MinIO)

	// Специфичное поле для MinIO
	StorageClass string `env:"S3_STORAGE_NAME"` // Класс хранения (например, "FORTRESS")
}

type Config struct {
	SrvCfg ServerConfig
	S3Cfg  S3Config
}

func readEnv() (map[string]string, error) {
	myEnv, err := godotenv.Read()
	if err != nil {
		// slog.Error("Ошибка при чтении переменных окружения", "error", err)
		return nil, fmt.Errorf("ошибка при чтении переменных окружения: %w", err)
	}
	return myEnv, nil
}

func MustLoad() (Config, error) {
	if err := godotenv.Load(); err != nil {
		// slog.Error("Ошибка при загрузке .env файла", "error", err)
		return Config{}, fmt.Errorf("ошибка при загрузке .env файла: %w", err)
	}

	envMap, err := readEnv()
	if err != nil {
		return Config{}, fmt.Errorf("ошибка при чтении переменных окружения: %w", err)
	}

	srvCfg := &ServerConfig{}
	s3Cfg := &S3Config{}

	configs := []AppConfig{srvCfg, s3Cfg}
	for _, cfg := range configs {
		if err := cfg.Load(envMap); err != nil {
			// slog.Error("Ошибка при загрузке конфигурации", "error", err)
			return Config{}, err
		}
		if err := cfg.Validate(); err != nil {
			// slog.Error("Ошибка валидации конфигурации", "error", err)
			return Config{}, err
		}
	}

	cfg := Config{
		SrvCfg: *srvCfg,
		S3Cfg:  *s3Cfg,
	}

	slog.Info("Все конфигурации успешно загружены")

	return cfg, nil
}

