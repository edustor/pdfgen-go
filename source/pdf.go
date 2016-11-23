package source

import (
	"github.com/jung-kurt/gofpdf"
	"io"
)

const LR_MARGIN float64 = 5.0
const TOP_MARGIN float64 = 5.0
const BOTTOM_MARGIN float64 = 5.0

func GenPdf(writter io.Writer) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Helvetica", "B", 16)
	pdf.SetMargins(5, 5, 5)
	pdf.SetAutoPageBreak(false, 0)

	pdf.SetLineWidth(0.2)

	pageWidth, pageHeight := pdf.GetPageSize()

	xMax := float64(int(pageWidth) - int(pageWidth) % 5) - 5
	yMax := float64(int(pageHeight) - int(pageHeight) % 5) - 5

	//rightBorder := w - 5
	//bottomBorder := h - 5


	for x := 5.0; x <= xMax; x += 5 {
		pdf.Line(x, 5, x, yMax)
	}


	for y := 5.0; y <= yMax; y += 5 {
		pdf.Line(5, y, xMax, y)
	}


	err := pdf.Output(writter)
	return err
}
