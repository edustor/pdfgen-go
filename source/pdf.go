package source

import (
	"github.com/jung-kurt/gofpdf"
	"io"
	"math"
	"fmt"
)

const LR_MARGIN float64 = 1.0
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

	maxWidth := pageWidth - 2 * LR_MARGIN
	maxRoundWidth := maxWidth - math.Mod(maxWidth, 5)

	additionalXMargin := (maxWidth - maxRoundWidth) / 2

	xMin := LR_MARGIN + additionalXMargin
	xMax := pageWidth - LR_MARGIN - additionalXMargin

	fmt.Println(xMin)
	fmt.Println(xMax)


	for x := xMin; x <= xMax; x += 5 {
		pdf.Line(x, 5, x, pageHeight)
	}
	//
	//
	//for y := 5.0; y <= yMax; y += 5 {
	//	pdf.Line(5, y, xMax, y)
	//}
	//
	//
	err := pdf.Output(writter)
	return err
}
