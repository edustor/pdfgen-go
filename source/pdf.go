package source

import (
	"github.com/jung-kurt/gofpdf"
	"io"
)

func GenPdf(w io.Writer) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello world")
	err := pdf.Output(w)
	return err
}
