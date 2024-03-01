package main

import (
	"net/http"
	"slices"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/hbourgeot/goxcel/excel"
	"github.com/hbourgeot/goxcel/utils"
)

func (app *App) initGoxcel(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

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
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Goxcel initialized successfully!"))
}

func (app *App) appendDay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	requestUser := chi.URLParam(r, "user")
	if !slices.Contains(app.users, requestUser) {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	} else if requestUser != app.CurrentUser && slices.Contains(app.users, app.CurrentUser) {
		app.CurrentUser = requestUser
		app.g.FileName = "gastos_ingresos_" + app.CurrentUser + ".xlsx"
	}

	// Get the current user
	user := app.CurrentUser
	if user == "" {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// gets gasto and ingreso from the request
	gasto := chi.URLParam(r, "gasto")
	ingreso := chi.URLParam(r, "ingreso")
	if gasto == "" || ingreso == "" {
		http.Error(w, "Missing parameters", http.StatusBadRequest)
		return
	}

	// Get the current day
	day := time.Now().Day()
	month := time.Now().Month().String()

	// Set gasto in Meses sheet
	cell := utils.GenerateCell(month, day, true)

	cellValue, err := app.g.GetCellValue("Meses", cell)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if cellValue == "" {
		cellValue = gasto
	} else {
		gastoInt, err := strconv.Atoi(gasto)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		cellValueInt, err := strconv.Atoi(cellValue)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		cellValue = strconv.Itoa(gastoInt + cellValueInt)
	}

	err = app.g.SetCellValue("Meses", cell, cellValue)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set ingreso in Meses sheet
	cell = utils.GenerateCell(month, day, false)

	cellValue, err = app.g.GetCellValue("Meses", cell)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if cellValue == "" {
		cellValue = ingreso
	} else {
		ingresoInt, err := strconv.Atoi(ingreso)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		cellValueInt, err := strconv.Atoi(cellValue)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		cellValue = strconv.Itoa(ingresoInt + cellValueInt)
	}

	err = app.g.SetCellValue("Meses", cell, cellValue)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// send a success message with status 200
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Day appended successfully!"))
}

func (app *App) getGastosIngresos(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	requestUser := chi.URLParam(r, "user")
	if !slices.Contains(app.users, requestUser) {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	} else if requestUser != app.CurrentUser && slices.Contains(app.users, app.CurrentUser) {
		app.CurrentUser = requestUser
		app.g.FileName = "gastos_ingresos_" + app.CurrentUser + ".xlsx"
	}

	// Get the current user
	user := app.CurrentUser
	if user == "" {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Get the current month
	month := time.Now().Month().String()
	monthDays := utils.GetMonthDays(month)

	// Get gastos and ingresos from the Meses sheet
	gastos := 0
	ingresos := 0
	for i := 1; i <= monthDays; i++ {
		cell := utils.GenerateCell(month, i, true)
		cellValue, err := app.g.GetCellValue("Meses", cell)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if cellValue != "" {
			gasto, err := strconv.Atoi(cellValue)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			gastos += gasto
		}

		cell = utils.GenerateCell(month, i, false)
		cellValue, err = app.g.GetCellValue("Meses", cell)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if cellValue != "" {
			ingreso, err := strconv.Atoi(cellValue)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			ingresos += ingreso
		}
	}

	// send a success message with status 200 with a json
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"gastos": ` + strconv.Itoa(gastos) + `, "ingresos": ` + strconv.Itoa(ingresos) + `}`))
}
