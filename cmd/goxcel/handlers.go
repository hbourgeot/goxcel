package main

import (
	"io"
	"net/http"
	"os"
	"strconv"

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
	if app.CurrentUser != user {
		app.CurrentUser = user

		// Crea una nueva instancia de Goxcel
		app.g = &excel.Goxcel{
			FileName: "gastos_ingresos_" + user + ".xlsx",
			Template: "gastos_ingresos_template.xlsx",
		}
	}


	// Copia la plantilla al nuevo archivo si es necesario
	if err := app.g.CopyTemplate(); err != nil {
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
	
	if requestUser != app.CurrentUser {
		app.CurrentUser = requestUser
		app.fillStructs(w)
	}

	if app.g == nil || app.g.File == nil {
		app.fillStructs(w)
	}
	
	month := chi.URLParam(r, "month")
	day, err := strconv.Atoi(chi.URLParam(r, "day"))
	if err != nil {
		render.Status(r, 400)
		render.PlainText(w, r, "Bad request: day must be a number")
		return
	}

	// gets gasto and ingreso from the request
	gasto := chi.URLParam(r, "gasto")
	ingreso := chi.URLParam(r, "ingreso")
	if gasto == "" || ingreso == "" {
		render.Status(r, 400)
		render.PlainText(w, r, "Bad request: gasto and ingreso are required")
		return
	}

	// Set gasto in Meses sheet
	cell := utils.GenerateCell(month, day, true)

	cellValue, err := app.g.GetCellValue("Meses", cell)
	if err != nil {
		render.Status(r, 500)
		render.PlainText(w, r, err.Error())
		return
	}

	value := 0

	if cellValue == "" {
		value, err = strconv.Atoi(gasto)
		if err != nil {
			render.Status(r, 500)
			render.PlainText(w, r, err.Error())
			return
		}
	} else {
		gastoInt, err := strconv.Atoi(gasto)
		if err != nil {
			render.Status(r, 500)
			render.PlainText(w, r, err.Error())
			return
		}

		value, err = strconv.Atoi(cellValue)
		if err != nil {
			render.Status(r, 500)
			render.PlainText(w, r, err.Error())
			return
		}

		value += gastoInt
	}

	err = app.g.SetCellValue("Meses", cell, value)
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
		value, err = strconv.Atoi(ingreso)
		if err != nil {
			render.Status(r, 500)
			render.PlainText(w, r, err.Error())
			return
		}
	} else {
		ingresoInt, err := strconv.Atoi(ingreso)
		if err != nil {
			render.Status(r, 500)
			render.PlainText(w, r, err.Error())
			return
		}

		value, err = strconv.Atoi(cellValue)
		if err != nil {
			render.Status(r, 500)
			render.PlainText(w, r, err.Error())
			return
		}

		value += ingresoInt
	}

	err = app.g.SetCellValue("Meses", cell, value)
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

	if requestUser != app.CurrentUser {
		app.CurrentUser = requestUser
		app.fillStructs(w)
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

func (app *App) UploadFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the multipart form
	err := r.ParseMultipartForm(10 << 20) // 10MB
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the file from the form
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	user := chi.URLParam(r, "user")
	if app.CurrentUser != user {
		app.CurrentUser = user
		app.g.FileName = "gastos_ingresos_" + user + ".xlsx"
	}

	// Check if the file already exists
	if _, err := os.Stat(app.g.FileName); os.IsNotExist(err) {
		// Create a new file
		f, err := os.OpenFile(app.g.FileName, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0666)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		// Copy the file to the new file
		_, err = io.Copy(f, file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// Send a success message
	w.WriteHeader(http.StatusCreated)
	render.PlainText(w, r, "File uploaded successfully!")
}

func (app *App) DownloadFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	user := chi.URLParam(r, "user")
	if app.CurrentUser != user {
		app.CurrentUser = user
		app.g.FileName = "gastos_ingresos_" + user + ".xlsx"
	}

	// Check if the file exists
	if _, err := os.Stat(app.g.FileName); os.IsNotExist(err) {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	// Open the file
	file, err := os.Open(app.g.FileName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Get the file info
	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the headers
	w.Header().Set("Content-Disposition", "attachment; filename="+app.g.FileName)
	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	w.Header().Set("Content-Length", strconv.FormatInt(fileInfo.Size(), 10))

	// Send the file
	http.ServeContent(w, r, app.g.FileName, fileInfo.ModTime(), file)
}

func (app *App) DownloadTemplate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Check if the file exists
	if _, err := os.Stat(app.g.Template); os.IsNotExist(err) {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	// Open the file
	file, err := os.Open(app.g.Template)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Get the file info
	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the headers
	w.Header().Set("Content-Disposition", "attachment; filename="+app.g.Template)
	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	w.Header().Set("Content-Length", strconv.FormatInt(fileInfo.Size(), 10))

	// Send the file
	http.ServeContent(w, r, app.g.Template, fileInfo.ModTime(), file)
}