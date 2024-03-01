package excel

import (
	"io"
	"os"

	"github.com/xuri/excelize/v2"
)

type Goxcel struct {
	FileName string
	Template string
	File     *excelize.File
}

func (g *Goxcel) Open() (*excelize.File, error) {
	file, err := excelize.OpenFile(g.FileName)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (g *Goxcel) CheckFile() error {
	file, err := g.Open()
	if err != nil {
		return err
	}

	g.File = file
	return nil
}

func (g *Goxcel) Save() error {
	if g.File == nil {
		if err := g.CheckFile(); err != nil {
			return err
		}
	}
	return g.File.Save()
}

func (g *Goxcel) SaveAs(filename string) error {
	if g.File == nil {
		if err := g.CheckFile(); err != nil {
			return err
		}
	}
	return g.File.SaveAs(filename)
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
	if g.File == nil {
		if err := g.CheckFile(); err != nil {
			return err
		}
	}

	return g.File.SetCellValue(sheet, cell, value)
}

func (g *Goxcel) GetCellValue(sheet, cell string) (string, error) {
	if g.File == nil {
		if err := g.CheckFile(); err != nil {
			return "", err
		}
	}

	return g.File.GetCellValue(sheet, cell)
}
