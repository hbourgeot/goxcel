package utils

import (
	"github.com/hbourgeot/goxcel/excel"
	"strconv"
)

var Meses = []string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre"}

var Semanas = 52

func getMonthDays(month string) int {
	switch month {
	case "Enero", "Marzo", "Mayo", "Julio", "Agosto", "Octubre", "Diciembre":
		return 31
	case "Febrero":
		return 28
	default:
		return 30
	}
}

func getMonthColumn(gasto bool, month string) string {
	if gasto {
		switch month {
		case "Enero":
			return "A"
		case "Febrero":
			return "C"
		case "Marzo":
			return "E"
		case "Abril":
			return "G"
		case "Mayo":
			return "I"
		case "Junio":
			return "K"
		case "Julio":
			return "M"
		case "Agosto":
			return "O"
		case "Septiembre":
			return "Q"
		case "Octubre":
			return "S"
		case "Noviembre":
			return "U"
		case "Diciembre":
			return "W"
		}
	}

	switch month {
	case "Enero":
		return "B"
	case "Febrero":
		return "D"
	case "Marzo":
		return "F"
	case "Abril":
		return "H"
	case "Mayo":
		return "J"
	case "Junio":
		return "L"
	case "Julio":
		return "N"
	case "Agosto":
		return "P"
	case "Septiembre":
		return "R"
	case "Octubre":
		return "T"
	case "Noviembre":
		return "V"
	case "Diciembre":
		return "X"
	}

	return ""
}

func LlenarGastosIngresos() {
	meses := excel.GastosIngresos{}
	for key, value := range excel.Meses {
		for secondKey := range value {
			cells := GetCells(secondKey == "Gastos", key)
			meses[key][secondKey] = cells
		}
	}

	excel.Meses = meses
}

func LlenarGastosIngresosSem() {
	semanas := excel.GastosIngresosSem{}
	for i := 1; i <= Semanas; i++ {
		semana := map[string]any{}
		str := strconv.Itoa(i)
		gastos := []string{"A" + str, "C" + str, "E" + str, "G" + str, "I" + str, "K" + str, "M" + str}
		ingresos := []string{"B" + str, "D" + str, "F" + str, "H" + str, "J" + str, "L" + str, "N" + str}

		semana["Gastos"] = gastos
		semana["Ingresos"] = ingresos
		semanas[i] = semana
	}

	print(len(semanas))
	excel.Semanas = semanas
}

func GetCells(gastos bool, month string) []string {
	cells := []string{}
	for i := 0; i < getMonthDays(month); i++ {
		cells = append(cells, getMonthColumn(gastos, month)+strconv.Itoa(i+2))
	}

	return cells
}
