package main

import (
	"github.com/go-chi/chi/v5"
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
		users: users,
	}
	utils.LlenarGastosIngresos()
	utils.LlenarGastosIngresosSem()
}
