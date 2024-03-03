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

const dir = "/usr/local/bin"

func (g *Goxcel) Open() error {
	file, err := excelize.OpenFile(dir + "/" + g.FileName)
	if err != nil {
		return err
	}

	g.File = file
	return nil
}

func (g *Goxcel) CheckFile() error {
	err := g.Open()
	if err != nil {
		return err
	}

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
	if _, err := os.Stat(dir + "/" + g.FileName); err == nil {
		// El archivo ya existe, no hacer nada y no devolver error
		return nil
	} 
	// Abre el archivo de plantilla original
	sourceFile, err := os.Open(dir + "/" + g.Template)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	// Crea un nuevo archivo para la copia
	destinationFile, err := os.Create(dir + "/" + g.FileName) // g.FileName será el nombre del nuevo archivo
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

	err := g.File.SetCellValue(sheet, cell, value)
	if err != nil {
		return err
	}

	return g.Save()
}

func (g *Goxcel) GetCellValue(sheet, cell string) (string, error) {
	if g.File == nil {
		if err := g.CheckFile(); err != nil {
			return "", err
		}
	}

	return g.File.GetCellValue(sheet, cell)
}
