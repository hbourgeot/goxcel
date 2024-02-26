package excel

import (
	"github.com/xuri/excelize/v2"
	"io"
	"os"
)

type Goxcel struct {
	FileName string
	Template string
	file     *excelize.File
}

func (g *Goxcel) Save() error {
	if g.file != nil {
		return g.file.Save()
	}
	return nil // O considera retornar un error si file es nil
}

func (g *Goxcel) SaveAs(filename string) error {
	if g.file != nil {
		return g.file.SaveAs(filename)
	}
	return nil // O considera retornar un error si file es nil
}

func (g *Goxcel) CopyTemplate() error {
	// Abre el archivo de plantilla original
	sourceFile, err := os.Open(g.Template)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	// Crea un nuevo archivo para la copia
	destinationFile, err := os.Create(g.FileName) // g.FileName será el nombre del nuevo archivo
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	// Copia el contenido del archivo de plantilla al nuevo archivo
	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	// Asegúrate de que el archivo se haya copiado correctamente antes de proceder
	err = destinationFile.Sync()
	if err != nil {
		return err
	}

	return nil
}

func (g *Goxcel) SetCellValue(sheet, cell, value string) error {
	if g.file != nil {
		return g.file.SetCellValue(sheet, cell, value)
	}
	return nil // O considera retornar un error si file es nil
}

func (g *Goxcel) GetCellValue(sheet, cell string) (string, error) {
	if g.file != nil {
		return g.file.GetCellValue(sheet, cell)
	}
	return "", nil // O considera retornar un error si file es nil
}
