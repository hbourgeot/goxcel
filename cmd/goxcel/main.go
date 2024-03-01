package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/hbourgeot/goxcel/excel"
	"github.com/hbourgeot/goxcel/utils"
)

type App struct {
	CurrentUser string
	users       []string
	Router      *chi.Mux
	g           *excel.Goxcel
}

var users = []string{"henrry", "hilda"}

func main() {
	app := &App{
		users:  users,
		Router: NewRouter(),
	}

	app.Mount()
	utils.LlenarGastosIngresos()

	http.ListenAndServe(":8080", app.Router)
}

func NewRouter() *chi.Mux {
	return chi.NewMux()
}

func (app *App) Mount() {
	app.Router.Use(middleware.Logger)

	app.Router.Post("/initGoxcel/{user}", app.initGoxcel)
	app.Router.Post("/appendDay/{user}-{gasto}-{ingreso}", app.appendDay)
	app.Router.Get("/getGasIng/{user}", app.getGastosIngresos)
}
