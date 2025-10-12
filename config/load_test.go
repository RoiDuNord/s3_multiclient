package config

import "testing"

func TestS3Config_Load(t *testing.T) {
	type args struct {
		envMap map[string]string
	}
	tests := []struct {
		name    string
		cfg     *S3Config
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.cfg.Load(tt.args.envMap); (err != nil) != tt.wantErr {
				t.Errorf("S3Config.Load() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServerConfig_Load(t *testing.T) {
	type args struct {
		envMap map[string]string
	}
	tests := []struct {
		name    string
		srvcfg  *ServerConfig
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.srvcfg.Load(tt.args.envMap); (err != nil) != tt.wantErr {
				t.Errorf("ServerConfig.Load() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServerConfig_loadPort(t *testing.T) {
	type args struct {
		envMap map[string]string
	}
	tests := []struct {
		name    string
		srvcfg  *ServerConfig
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.srvcfg.loadPort(tt.args.envMap); (err != nil) != tt.wantErr {
				t.Errorf("ServerConfig.loadPort() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
