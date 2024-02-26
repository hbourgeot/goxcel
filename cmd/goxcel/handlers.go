package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/hbourgeot/goxcel/excel"
	"net/http"
	"slices"
)

func (app *App) initGoxcel(w http.ResponseWriter, r *http.Request) {
	user := chi.URLParam(r, "user")
	if slices.Contains(app.users, user) {
		app.CurrentUser = user
	} else {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Create a new instance of Goxcel
	app.g = &excel.Goxcel{
		FileName: "gastos_ingresos_" + user + ".xlsx",
		Template: "gastos_ingresos_template.xlsx",
	}

	if err := app.g.CopyTemplate(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// send a success message with status 200
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Goxcel initialized successfully!"))
}
