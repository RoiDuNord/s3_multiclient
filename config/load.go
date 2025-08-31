package config

import (
	"fmt"
	"log/slog"
	"strconv"
	"strings"
)

func (cfg *S3Config) Load(envMap map[string]string) error {
	var missingVars []string

	getEnv := func(key string) string {
		val, ok := envMap[key]
		if !ok {
			missingVars = append(missingVars, key)
		}
		return val
	}

	// Обязательные поля для всех типов хранилищ
	cfg.Type = StorageType(strings.ToLower(getEnv("S3_TYPE")))
	cfg.Endpoint = getEnv("S3_ENDPOINT")
	cfg.AccessKeyID = getEnv("S3_ACCESS_KEY")
	cfg.SecretAccessKey = getEnv("S3_SECRET_KEY")
	cfg.BucketName = getEnv("S3_BUCKET_NAME")
	cfg.Region = getEnv("S3_REGION")

	// Опциональные булевы поля
	if useSSLStr, ok := envMap["S3_USE_SSL"]; ok {
		cfg.UseSSL = (useSSLStr == "true")
	}

	// Поля, зависящие от типа хранилища
	switch cfg.Type {
	case StorageMinIO:
		cfg.StorageClass = getEnv("S3_STORAGE_NAME")
	case StorageCeph:
		if pathStyleStr, ok := envMap["S3_PATH_STYLE"]; ok {
			cfg.PathStyle = (pathStyleStr == "true")
		}
	default:
		return fmt.Errorf("неизвестный тип хранилища: %s", cfg.Type)
	}

	if len(missingVars) > 0 {
		for _, v := range missingVars {
			slog.Warn(fmt.Sprintf("Переменная %s не определена в .env", v))
		}
		return fmt.Errorf("отсутствуют обязательные переменные окружения для %s: %v", cfg.Type, missingVars)
	}

	return nil
}

func (srvcfg *ServerConfig) Load(envMap map[string]string) error {
	// var ok bool
	// var missingVars []string

	// ap.Host, ok = envMap["APP_HOST"]
	// if !ok {
	// 	missingVars = append(missingVars, "APP_HOST")
	// }

	if err := srvcfg.loadPort(envMap); err != nil {
		return err
	}

	// if len(missingVars) > 0 {
	// 	for _, v := range missingVars {
	// 		slog.Warn(fmt.Sprintf("Переменная %s не определена в .env", v))
	// 	}
	// 	return fmt.Errorf("отсутствуют обязательные переменные окружения APP: %v", missingVars)
	// }

	return nil
}

func (srvcfg *ServerConfig) loadPort(envMap map[string]string) error {
	portStr, ok := envMap["SERVER_PORT"]
	if !ok {
		slog.Warn("Переменная SERVER_PORT не определена в .env")
		return fmt.Errorf("переменная SERVER_PORT не определена в .env")
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return fmt.Errorf("ошибка преобразования SERVER_PORT в число: %w", err)
	}

	srvcfg.Port = port
	return nil
}
