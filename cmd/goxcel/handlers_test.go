package main

import (
	"net/http"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/hbourgeot/goxcel/excel"
)

func TestApp_initGoxcel(t *testing.T) {
	type fields struct {
		CurrentUser string
		users       []string
		Router      *chi.Mux
		g           *excel.Goxcel
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
			app := &App{
				CurrentUser: tt.fields.CurrentUser,
				users:       tt.fields.users,
				Router:      tt.fields.Router,
				g:           tt.fields.g,
			}
			app.initGoxcel(tt.args.w, tt.args.r)
		})
	}
}

func TestApp_appendDay(t *testing.T) {
	type fields struct {
		CurrentUser string
		users       []string
		Router      *chi.Mux
		g           *excel.Goxcel
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
			app := &App{
				CurrentUser: tt.fields.CurrentUser,
				users:       tt.fields.users,
				Router:      tt.fields.Router,
				g:           tt.fields.g,
			}
			app.appendDay(tt.args.w, tt.args.r)
		})
	}
}

func TestApp_getGastosIngresos(t *testing.T) {
	type fields struct {
		CurrentUser string
		users       []string
		Router      *chi.Mux
		g           *excel.Goxcel
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
			app := &App{
				CurrentUser: tt.fields.CurrentUser,
				users:       tt.fields.users,
				Router:      tt.fields.Router,
				g:           tt.fields.g,
			}
			app.getGastosIngresos(tt.args.w, tt.args.r)
		})
	}
}
