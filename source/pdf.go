package source

import (
	"github.com/jung-kurt/gofpdf"
	"io"
	"log"
)

func GenPdf(writter io.Writer) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Helvetica", "B", 16)
	pdf.SetMargins(5, 5, 5)
	pdf.SetAutoPageBreak(false, 0)

	w, h := pdf.GetPageSize()

	rightBorder := w - 5
	bottomBorder := h - 5

	pdf.SetX(5)
	pdf.SetY(5)

	for pdf.GetY() + 5 < bottomBorder {
		log.Printf("%v %v", pdf.GetX(), pdf.GetY())
		for pdf.GetX() + 5 < rightBorder {
			pdf.CellFormat(5, 5, "", "LT", 0, "", false, 0, "")
		}
		pdf.Ln(-1)
	}


	err := pdf.Output(writter)
	return err
}
