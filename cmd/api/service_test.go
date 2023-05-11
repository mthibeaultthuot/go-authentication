package main

import (
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"testing"
)

func TestConfig_Delete(t *testing.T) {
	type fields struct {
		Client *mongo.Client
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
			app := &Config{
				Client: tt.fields.Client,
			}
			app.Delete(tt.args.w, tt.args.r)
		})
	}
}

func TestConfig_Login(t *testing.T) {
	type fields struct {
		Client *mongo.Client
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
			app := &Config{
				Client: tt.fields.Client,
			}
			app.Login(tt.args.w, tt.args.r)
		})
	}
}

func TestConfig_Refresh(t *testing.T) {
	type fields struct {
		Client *mongo.Client
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
			app := &Config{
				Client: tt.fields.Client,
			}
			app.Refresh(tt.args.w, tt.args.r)
		})
	}
}

func TestConfig_Registration(t *testing.T) {
	type fields struct {
		Client *mongo.Client
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
			app := &Config{
				Client: tt.fields.Client,
			}
			app.Registration(tt.args.w, tt.args.r)
		})
	}
}

func TestConfig_Token(t *testing.T) {
	type fields struct {
		Client *mongo.Client
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
			app := &Config{
				Client: tt.fields.Client,
			}
			app.Token(tt.args.w, tt.args.r)
		})
	}
}
