package config

import (
	"reflect"
	"testing"
)

func Test_isValidBucketName(t *testing.T) {
	type args struct {
		bucketName string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBucketName(tt.args.bucketName); got != tt.want {
				t.Errorf("isValidBucketName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateCommonFields(t *testing.T) {
	type args struct {
		s3Cfg *S3Config
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateCommonFields(tt.args.s3Cfg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validateCommonFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateMinIOFields(t *testing.T) {
	type args struct {
		s3Cfg *S3Config
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateMinIOFields(tt.args.s3Cfg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validateMinIOFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateCephFields(t *testing.T) {
	type args struct {
		s3Cfg *S3Config
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateCephFields(tt.args.s3Cfg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validateCephFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestS3Config_Validate(t *testing.T) {
	tests := []struct {
		name    string
		s3Cfg   *S3Config
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s3Cfg.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("S3Config.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServerConfig_Validate(t *testing.T) {
	tests := []struct {
		name    string
		srvcfg  *ServerConfig
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.srvcfg.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("ServerConfig.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
