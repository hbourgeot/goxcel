package main

import (
	"reflect"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/hbourgeot/goxcel/excel"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func TestNewRouter(t *testing.T) {
	tests := []struct {
		name string
		want *chi.Mux
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRouter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApp_Mount(t *testing.T) {
	type fields struct {
		CurrentUser string
		users       []string
		Router      *chi.Mux
		g           *excel.Goxcel
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &App{
				CurrentUser: tt.fields.CurrentUser,
				users:       tt.fields.users,
				Router:      tt.fields.Router,
				g:           tt.fields.g,
			}
			app.Mount()
		})
	}
}
