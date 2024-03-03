package main

import (
	"net/http"
	"os"
	"path/filepath"

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

func main() {
	app := &App{
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
	app.Router.Use(middleware.Recoverer)
	app.Router.Use(middleware.Logger)

	// workDir:= "/usr/local/bin"
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "frontend", "dist"))

	app.Router.Route("/api", func(r chi.Router) {
		r.Post("/initGoxcel/{user}", app.initGoxcel)
		r.Post("/appendDay/{user}/{month}-{day}/-{gasto}-{ingreso}", app.appendDay)
		r.Get("/getGasIng/{user}", app.getGastosIngresos)
	})
	app.Router.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		if _, err := os.Stat(filepath.Join(string(filesDir), r.URL.Path)); os.IsNotExist(err) {
			// Si no se encuentra el archivo estático, sirve index.html
			http.ServeFile(w, r, filepath.Join(string(filesDir), "index.html"))
		} else {
			// Sirve el archivo estático
			http.FileServer(filesDir).ServeHTTP(w, r)
		}
	})
}
