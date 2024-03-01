package utils

import (
	"reflect"
	"testing"
)

func TestGetMonthDays(t *testing.T) {
	type args struct {
		month string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMonthDays(tt.args.month); got != tt.want {
				t.Errorf("GetMonthDays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMonthColumn(t *testing.T) {
	type args struct {
		gasto bool
		month string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMonthColumn(tt.args.gasto, tt.args.month); got != tt.want {
				t.Errorf("GetMonthColumn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLlenarGastosIngresos(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LlenarGastosIngresos()
		})
	}
}

func TestGetCells(t *testing.T) {
	type args struct {
		gastos bool
		month  string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCells(tt.args.gastos, tt.args.month); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCells() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateCell(t *testing.T) {
	type args struct {
		column string
		row    int
		gasto  bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateCell(tt.args.column, tt.args.row, tt.args.gasto); got != tt.want {
				t.Errorf("GenerateCell() = %v, want %v", got, tt.want)
			}
		})
	}
}
