package server

import (
	"context"
	"s3_multiclient/config"
	_ "s3_multiclient/docs"
	"testing"
)

func TestServer_MustStart(t *testing.T) {
	type fields struct {
		ctx         context.Context
		loadManager LoadManager
	}
	type args struct {
		srvcfg config.ServerConfig
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				ctx:         tt.fields.ctx,
				loadManager: tt.fields.loadManager,
			}
			if err := s.MustStart(tt.args.srvcfg); (err != nil) != tt.wantErr {
				t.Errorf("Server.MustStart() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
