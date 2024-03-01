package excel

import (
	"testing"

	"github.com/xuri/excelize/v2"
)

func TestGoxcel_Save(t *testing.T) {
	type fields struct {
		FileName string
		Template string
		file     *excelize.File
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Goxcel{
				FileName: tt.fields.FileName,
				Template: tt.fields.Template,
				file:     tt.fields.file,
			}
			if err := g.Save(); (err != nil) != tt.wantErr {
				t.Errorf("Goxcel.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGoxcel_SaveAs(t *testing.T) {
	type fields struct {
		FileName string
		Template string
		file     *excelize.File
	}
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Goxcel{
				FileName: tt.fields.FileName,
				Template: tt.fields.Template,
				file:     tt.fields.file,
			}
			if err := g.SaveAs(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("Goxcel.SaveAs() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGoxcel_CopyTemplate(t *testing.T) {
	type fields struct {
		FileName string
		Template string
		file     *excelize.File
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Goxcel{
				FileName: tt.fields.FileName,
				Template: tt.fields.Template,
				file:     tt.fields.file,
			}
			if err := g.CopyTemplate(); (err != nil) != tt.wantErr {
				t.Errorf("Goxcel.CopyTemplate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGoxcel_SetCellValue(t *testing.T) {
	type fields struct {
		FileName string
		Template string
		file     *excelize.File
	}
	type args struct {
		sheet string
		cell  string
		value string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Goxcel{
				FileName: tt.fields.FileName,
				Template: tt.fields.Template,
				file:     tt.fields.file,
			}
			if err := g.SetCellValue(tt.args.sheet, tt.args.cell, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Goxcel.SetCellValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGoxcel_GetCellValue(t *testing.T) {
	type fields struct {
		FileName string
		Template string
		file     *excelize.File
	}
	type args struct {
		sheet string
		cell  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Goxcel{
				FileName: tt.fields.FileName,
				Template: tt.fields.Template,
				file:     tt.fields.file,
			}
			got, err := g.GetCellValue(tt.args.sheet, tt.args.cell)
			if (err != nil) != tt.wantErr {
				t.Errorf("Goxcel.GetCellValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Goxcel.GetCellValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
