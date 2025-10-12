package config

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

// import (
// 	"os"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestS3Config_Load_SuccessMinIO(t *testing.T) {
// 	envMap := map[string]string{
// 		"S3_TYPE":              "minio",
// 		"S3_ENDPOINT":          "localhost:9000",
// 		"S3_ACCESS_KEY_ID":     "key",
// 		"S3_SECRET_ACCESS_KEY": "secret",
// 		"S3_BUCKET_NAME":       "my-bucket",
// 		"S3_REGION":            "us-east-1",
// 		"S3_USE_SSL":           "true",
// 		"S3_STORAGE_NAME":      "STANDARD",
// 	}

// 	cfg := &S3Config{}
// 	err := cfg.Load(envMap)
// 	assert.NoError(t, err)
// 	assert.Equal(t, StorageMinIO, cfg.Type)
// 	assert.True(t, cfg.UseSSL)
// 	assert.Equal(t, "STANDARD", cfg.StorageClass)
// }

// func TestS3Config_Load_MissingVar(t *testing.T) {
// 	envMap := map[string]string{}
// 	cfg := &S3Config{}
// 	err := cfg.Load(envMap)
// 	assert.Error(t, err)
// 	assert.Contains(t, err.Error(), "отсутствуют обязательные переменные")
// }

// func TestS3Config_Load_InvalidType(t *testing.T) {
// 	envMap := map[string]string{
// 		"S3_TYPE": "invalid",
// 	}
// 	cfg := &S3Config{}
// 	err := cfg.Load(envMap)
// 	assert.Error(t, err)
// 	assert.Contains(t, err.Error(), "неизвестный тип хранилища")
// }

// func TestS3Config_Validate_Success(t *testing.T) {
// 	cfg := &S3Config{
// 		Type:            StorageMinIO,
// 		Endpoint:        "localhost:9000",
// 		AccessKeyID:     "key",
// 		SecretAccessKey: "secret",
// 		BucketName:      "my-bucket",
// 		Region:          "us-east-1",
// 		StorageClass:    "STANDARD",
// 	}
// 	err := cfg.Validate()
// 	assert.NoError(t, err)
// }

// func TestS3Config_Validate_InvalidBucketName(t *testing.T) {
// 	cfg := &S3Config{
// 		Type:       StorageMinIO,
// 		BucketName: "Invalid_Bucket!",
// 	}
// 	err := cfg.Validate()
// 	assert.Error(t, err)
// 	assert.Contains(t, err.Error(), "содержит недопустимые символы")
// }

// func TestS3Config_Validate_MissingFields(t *testing.T) {
// 	cfg := &S3Config{Type: StorageMinIO}
// 	err := cfg.Validate()
// 	assert.Error(t, err)
// 	assert.Contains(t, err.Error(), "Необходимо задать переменные окружения")
// }

// func TestServerConfig_Load_Success(t *testing.T) {
// 	envMap := map[string]string{
// 		"SERVER_PORT": "8080",
// 	}
// 	cfg := &ServerConfig{}
// 	err := cfg.Load(envMap)
// 	assert.NoError(t, err)
// 	assert.Equal(t, 8080, cfg.Port)
// }

// func TestServerConfig_Load_InvalidPort(t *testing.T) {
// 	envMap := map[string]string{
// 		"SERVER_PORT": "abc",
// 	}
// 	cfg := &ServerConfig{}
// 	err := cfg.Load(envMap)
// 	assert.Error(t, err)
// 	assert.Contains(t, err.Error(), "ошибка преобразования SERVER_PORT")
// }

// func TestServerConfig_Validate_Success(t *testing.T) {
// 	cfg := &ServerConfig{Port: 8080}
// 	err := cfg.Validate()
// 	assert.NoError(t, err)
// }

// func TestServerConfig_Validate_InvalidPort(t *testing.T) {
// 	cfg := &ServerConfig{Port: 70000}
// 	err := cfg.Validate()
// 	assert.Error(t, err)
// 	assert.Contains(t, err.Error(), "SERVER_PORT должен быть в диапазоне")
// }

// func TestMustLoad_Success(t *testing.T) {
// 	// Установите переменные окружения
// 	os.Setenv("S3_TYPE", "minio")
// 	os.Setenv("S3_ENDPOINT", "localhost:9000")
// 	os.Setenv("S3_ACCESS_KEY_ID", "key")
// 	os.Setenv("S3_SECRET_ACCESS_KEY", "secret")
// 	os.Setenv("S3_BUCKET_NAME", "my-bucket")
// 	os.Setenv("S3_REGION", "us-east-1")
// 	os.Setenv("S3_STORAGE_NAME", "STANDARD")
// 	os.Setenv("SERVER_PORT", "8080")

// 	cfg, err := MustLoad()
// 	assert.NoError(t, err)
// 	assert.Equal(t, StorageMinIO, cfg.S3Cfg.Type)
// 	assert.Equal(t, 8080, cfg.SrvCfg.Port)

// 	// Очистка
// 	os.Unsetenv("S3_TYPE")
// 	// ... очистите остальные
// }

//	func TestMustLoad_ValidationError(t *testing.T) {
//		os.Setenv("S3_TYPE", "minio")
//		os.Setenv("SERVER_PORT", "99999") // Неверный порт
//		_, err := MustLoad()
//		assert.Error(t, err)
//		assert.Contains(t, err.Error(), "диапазоне")
//	}
func TestMustLoad(t *testing.T) {
	// Создаём временную директорию
	dir := t.TempDir()

	// Содержимое тестового .env (включая SrvCfg и S3Cfg)
	envContent := `
SERVER_HOST=
SERVER_PORT=8080
S3_TYPE=minio
S3_ENDPOINT=https://example.com
S3_ACCESS_KEY=test-access-key
S3_SECRET_KEY=test-secret-key
S3_BUCKET_NAME=test-bucket
S3_REGION=us-east-1
S3_USE_SSL=true
S3_STORAGE_NAME=standard
S3_PATH_STYLE=false
`
	envPath := filepath.Join(dir, ".env")
	if err := os.WriteFile(envPath, []byte(envContent), 0644); err != nil {
		t.Fatalf("failed to write .env file: %v", err)
	}

	// Сохраняем текущую директорию
	origDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get wd: %v", err)
	}
	// Меняем на временную директорию с тестовым .env
	if err := os.Chdir(dir); err != nil {
		t.Fatalf("failed to chdir: %v", err)
	}
	defer func() {
		_ = os.Chdir(origDir) // Возвращаемся обратно
	}()

	// Ожидаемый результат
	want := Config{
		SrvCfg: ServerConfig{ // Предполагаемая структура, подстройте под ваш код
			Host: "",
			Port: 8080,
		},
		S3Cfg: S3Config{
			Type:            StorageMinIO,
			Endpoint:        "https://example.com",
			AccessKeyID:     "test-access-key",
			SecretAccessKey: "test-secret-key",
			BucketName:      "test-bucket",
			Region:          "us-east-1",
			UseSSL:          true,
			StorageClass:    "standard",
			PathStyle:       false,
		},
	}

	got, err := MustLoad()
	if err != nil {
		t.Fatalf("MustLoad() returned error: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func Test_readEnv(t *testing.T) {
	tests := []struct {
		name    string
		want    map[string]string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readEnv()
			if (err != nil) != tt.wantErr {
				t.Errorf("readEnv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}
