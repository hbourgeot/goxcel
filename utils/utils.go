package utils

import (
	"strconv"

	"github.com/hbourgeot/goxcel/excel"
)

var Meses = []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

func GetMonthDays(month string) int {
	switch month {
	case "January", "March", "May", "July", "August", "October", "December":
		return 31
	case "February":
		return 28
	default:
		return 30
	}
}

func GetMonthColumn(gasto bool, month string) string {
	if gasto {
		switch month {
		case "January":
			return "A"
		case "February":
			return "C"
		case "March":
			return "E"
		case "April":
			return "G"
		case "May":
			return "I"
		case "June":
			return "K"
		case "July":
			return "M"
		case "August":
			return "O"
		case "September":
			return "Q"
		case "October":
			return "S"
		case "November":
			return "U"
		case "December":
			return "W"
		}
	}

	switch month {
	case "January":
		return "B"
	case "February":
		return "D"
	case "March":
		return "F"
	case "April":
		return "H"
	case "May":
		return "J"
	case "June":
		return "L"
	case "July":
		return "N"
	case "August":
		return "P"
	case "September":
		return "R"
	case "October":
		return "T"
	case "November":
		return "V"
	case "December":
		return "X"
	}

	return ""
}

func LlenarGastosIngresos() {
	meses := make(excel.GastosIngresos) // Initialize the map
	for key, value := range excel.Meses {
		meses[key] = make(map[string]any) // Initialize the inner map with type map[string]any
		for secondKey := range value {
			cells := GetCells(secondKey == "Gastos", key)
			meses[key][secondKey] = cells
		}
	}

	excel.Meses = meses
}

func GetCells(gastos bool, month string) []string {
	cells := []string{}
	for i := 0; i < GetMonthDays(month); i++ {
		cells = append(cells, GetMonthColumn(gastos, month)+strconv.Itoa(i+2))
	}

	return cells
}

func GenerateCell(column string, row int, gasto bool) string {
	return GetMonthColumn(gasto, column) + strconv.Itoa(row+2)
}
