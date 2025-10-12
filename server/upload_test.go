package server

import (
	"context"
	"net/http"
	"testing"
)

func TestServer_Upload(t *testing.T) {
	type fields struct {
		ctx         context.Context
		loadManager LoadManager
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				ctx:         tt.fields.ctx,
				loadManager: tt.fields.loadManager,
			}
			s.Upload(tt.args.w, tt.args.r)
		})
	}
}
