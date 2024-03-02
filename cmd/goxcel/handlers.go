package main

import (
	"fmt"
	"net/http"
	"slices"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/hbourgeot/goxcel/excel"
	"github.com/hbourgeot/goxcel/utils"
)

type Month struct {
	Month string `json:"month"`
	Days  []Day  `json:"days"`
}

type Day struct {
	Day      int `json:"day"`
	Gastos   int `json:"gastos"`
	Ingresos int `json:"ingresos"`
}

func (app *App) fillStructs(w http.ResponseWriter) {
	app.g = &excel.Goxcel{
		FileName: "gastos_ingresos_" + app.CurrentUser + ".xlsx",
		Template: "gastos_ingresos_template.xlsx",
	}

	// Open the file
	err := app.g.Open()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (app *App) initGoxcel(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	user := chi.URLParam(r, "user")
	if !slices.Contains(app.users, user) {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Crea una nueva instancia de Goxcel
	app.g = &excel.Goxcel{
		FileName: "gastos_ingresos_" + user + ".xlsx",
		Template: "gastos_ingresos_template.xlsx",
	}

	// Copia la plantilla al nuevo archivo si es necesario
	if err := app.g.CopyTemplate(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Abre el archivo
	if err := app.g.Open(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Envía un mensaje de éxito
	w.WriteHeader(http.StatusCreated)
	render.PlainText(w, r, "Goxcel initialized successfully!")
}

func (app *App) appendDay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		render.Status(r, 405)
		render.PlainText(w, r, "Method not allowed")
		return
	}

	requestUser := chi.URLParam(r, "user")

	if !slices.Contains(app.users, requestUser) {
		render.Status(r, 404)
		render.PlainText(w, r, "User not found")
		return
	} else if requestUser != app.CurrentUser {
		fmt.Println("Changing user")
		app.CurrentUser = requestUser
	}

	if app.g == nil || app.g.File == nil {
		app.fillStructs(w)
	}

	// gets gasto and ingreso from the request
	gasto := chi.URLParam(r, "gasto")
	ingreso := chi.URLParam(r, "ingreso")
	if gasto == "" || ingreso == "" {
		render.Status(r, 400)
		render.PlainText(w, r, "Bad request: gasto and ingreso are required")
		return
	}

	// Get the current day
	day := time.Now().Day()
	month := time.Now().Month().String()

	// Set gasto in Meses sheet
	cell := utils.GenerateCell(month, day, true)

	cellValue, err := app.g.GetCellValue("Meses", cell)
	if err != nil {
		render.Status(r, 500)
		render.PlainText(w, r, err.Error())
		return
	}

	if cellValue == "" {
		cellValue = gasto
	} else {
		gastoInt, err := strconv.Atoi(gasto)
		if err != nil {
			render.Status(r, 500)
			render.PlainText(w, r, err.Error())
			return
		}

		cellValueInt, err := strconv.Atoi(cellValue)
		if err != nil {
			render.Status(r, 500)
			render.PlainText(w, r, err.Error())
			return
		}

		cellValue = strconv.Itoa(gastoInt + cellValueInt)
	}

	err = app.g.SetCellValue("Meses", cell, cellValue)
	if err != nil {
		render.Status(r, 500)
		render.PlainText(w, r, err.Error())
		return
	}

	// Set ingreso in Meses sheet
	cell = utils.GenerateCell(month, day, false)

	cellValue, err = app.g.GetCellValue("Meses", cell)
	if err != nil {
		render.Status(r, 500)
		render.PlainText(w, r, err.Error())
		return
	}

	if cellValue == "" {
		cellValue = ingreso
	} else {
		ingresoInt, err := strconv.Atoi(ingreso)
		if err != nil {
			render.Status(r, 500)
			render.PlainText(w, r, err.Error())
			return
		}

		cellValueInt, err := strconv.Atoi(cellValue)
		if err != nil {
			render.Status(r, 500)
			render.PlainText(w, r, err.Error())
			return
		}

		cellValue = strconv.Itoa(ingresoInt + cellValueInt)
	}

	err = app.g.SetCellValue("Meses", cell, cellValue)
	if err != nil {
		render.Status(r, 500)
		render.PlainText(w, r, err.Error())
		return
	}

	// send a success message with status 200
	render.Status(r, 200)
	render.PlainText(w, r, "Day appended successfully!")
}

func (app *App) getGastosIngresos(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	requestUser := chi.URLParam(r, "user")
	if !slices.Contains(app.users, requestUser) {
		render.Status(r, 404)
		render.PlainText(w, r, "User not found")
		return
	} else if requestUser != app.CurrentUser {
		app.CurrentUser = requestUser
	}

	if app.g == nil || app.g.File == nil {
		app.fillStructs(w)
	}

	monthsData := []*Month{}
	for _, month := range utils.Meses {
		// Get the current month
		monthDays := utils.GetMonthDays(month)

		monthData := Month{
			Month: month,
			Days:  make([]Day, monthDays),
		}

		// Get gastos and ingresos from the Meses sheet
		gastos := 0
		ingresos := 0
		for i := 1; i <= monthDays; i++ {
			cell := utils.GenerateCell(month, i, true)
			cellValue, err := app.g.GetCellValue("Meses", cell)
			if err != nil {
				render.Status(r, 500)
				render.PlainText(w, r, err.Error())
				return
			}

			monthData.Days[i-1].Day = i
			if cellValue != "" {
				gasto, err := strconv.Atoi(cellValue)
				if err != nil {
					render.Status(r, 500)
					render.PlainText(w, r, err.Error())
					return
				}

				monthData.Days[i-1].Gastos = gasto

				gastos += gasto
			}

			cell = utils.GenerateCell(month, i, false)
			cellValue, err = app.g.GetCellValue("Meses", cell)
			if err != nil {
				render.Status(r, 500)
				render.PlainText(w, r, err.Error())
				return
			}

			if cellValue != "" {
				ingreso, err := strconv.Atoi(cellValue)
				if err != nil {
					render.Status(r, 500)
					render.PlainText(w, r, err.Error())
					return
				}

				monthData.Days[i-1].Ingresos = ingreso
				ingresos += ingreso
			}

		}
		monthsData = append(monthsData, &monthData)
	}
	// send a success message with status 200 with a json
	render.Status(r, 200)
	render.JSON(w, r, monthsData)
}
