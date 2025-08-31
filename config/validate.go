package config

import (
	"errors"
	"fmt"
	"log/slog"
	"regexp"
	"strings"
)

var bucketNameRegex = regexp.MustCompile(`^[a-z0-9\-]+$`)

func isValidBucketName(bucketName string) bool {
	return bucketNameRegex.MatchString(bucketName)
}

func validateCommonFields(s3Cfg *S3Config) []string {
	var missingVars []string
	if s3Cfg.Type == "" {
		missingVars = append(missingVars, "S3_TYPE")
	}
	if s3Cfg.Endpoint == "" {
		missingVars = append(missingVars, "S3_ENDPOINT")
	}
	if s3Cfg.AccessKeyID == "" {
		missingVars = append(missingVars, "S3_ACCESS_KEY")
	}
	if s3Cfg.SecretAccessKey == "" {
		missingVars = append(missingVars, "S3_SECRET_KEY")
	}
	if s3Cfg.BucketName == "" {
		missingVars = append(missingVars, "S3_BUCKET_NAME")
	}
	if s3Cfg.Region == "" {
		missingVars = append(missingVars, "S3_REGION")
	}
	return missingVars
}

func validateMinIOFields(s3Cfg *S3Config) []string {
	var missingVars []string
	if s3Cfg.StorageClass == "" {
		missingVars = append(missingVars, "S3_STORAGE_NAME")
	}
	return missingVars
}

func validateCephFields(s3Cfg *S3Config) []string {
	// PathStyle — булево, обычно true, но проверять необязательно.
	// Если нужна проверка, например, чтобы значение было true или false, можно сделать так:
	// Но в данном случае, если поле булево, оно всегда имеет значение по умолчанию false,
	// так что можно не проверять.
	return nil
}

func (s3Cfg *S3Config) Validate() error {
	missingVars := []string{}

	// Проверяем обязательное поле Type
	if s3Cfg.Type != StorageMinIO && s3Cfg.Type != StorageCeph {
		message := fmt.Sprintf("Неверный тип хранилища: %q. Ожидается '%s' или '%s'", s3Cfg.Type, StorageMinIO, StorageCeph)
		slog.Error(message)
		return errors.New(message)
	}

	// Проверяем общие поля
	missingVars = append(missingVars, validateCommonFields(s3Cfg)...)

	// Проверяем поля в зависимости от типа
	switch s3Cfg.Type {
	case StorageMinIO:
		missingVars = append(missingVars, validateMinIOFields(s3Cfg)...)
	case StorageCeph:
		missingVars = append(missingVars, validateCephFields(s3Cfg)...)
	}

	if len(missingVars) > 0 {
		message := fmt.Sprintf("Необходимо задать переменные окружения: %s", strings.Join(missingVars, ", "))
		slog.Warn(message)
		return errors.New(message)
	}

	if !isValidBucketName(s3Cfg.BucketName) {
		message := fmt.Sprintf("Имя бакета '%s' содержит недопустимые символы. Используйте только строчные буквы, цифры и дефисы.", s3Cfg.BucketName)
		slog.Error(message)
		return errors.New(message)
	}

	return nil
}

func (srvcfg *ServerConfig) Validate() error {
	if srvcfg.Port <= 0 || srvcfg.Port > 65535 {
		return fmt.Errorf("APP_PORT должен быть в диапазоне 1-65535, получено: %d", srvcfg.Port)
	}
	return nil
}
