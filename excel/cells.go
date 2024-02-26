package excel

type GastosIngresos map[string]map[string]any
type GastosIngresosSem map[int]map[string]any

var Meses GastosIngresos = GastosIngresos{
	"Enero":      {"Gastos": []string{}, "Ingresos": []string{}, "AhorroTotal": ""},
	"Febrero":    {"Gastos": []string{}, "Ingresos": []string{}, "AhorroTotal": ""},
	"Marzo":      {"Gastos": []string{}, "Ingresos": []string{}, "AhorroTotal": ""},
	"Abril":      {"Gastos": []string{}, "Ingresos": []string{}, "AhorroTotal": ""},
	"Mayo":       {"Gastos": []string{}, "Ingresos": []string{}, "AhorroTotal": ""},
	"Junio":      {"Gastos": []string{}, "Ingresos": []string{}, "AhorroTotal": ""},
	"Julio":      {"Gastos": []string{}, "Ingresos": []string{}, "AhorroTotal": ""},
	"Agosto":     {"Gastos": []string{}, "Ingresos": []string{}, "AhorroTotal": ""},
	"Septiembre": {"Gastos": []string{}, "Ingresos": []string{}, "AhorroTotal": ""},
	"Octubre":    {"Gastos": []string{}, "Ingresos": []string{}, "AhorroTotal": ""},
	"Noviembre":  {"Gastos": []string{}, "Ingresos": []string{}, "AhorroTotal": ""},
	"Diciembre":  {"Gastos": []string{}, "Ingresos": []string{}, "AhorroTotal": ""},
}

var Semanas GastosIngresosSem = GastosIngresosSem{}

var SemanaAhorroTotal = "O"
