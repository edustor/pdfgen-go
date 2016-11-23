package source

import (
	"github.com/jung-kurt/gofpdf"
	"io"
	"math"
	"log"
)

const LR_MARGIN float64 = 5.0
const TOP_MARGIN float64 = 5.0
const BOTTOM_MARGIN float64 = 5.0

func GenPdf(writter io.Writer, pageCount int) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Helvetica", "B", 16)
	pdf.SetMargins(5, 5, 5)
	pdf.SetAutoPageBreak(false, 0)

	pdf.SetLineWidth(0.2)
	pdf.SetDrawColor(128, 128, 128)

	pageWidth, pageHeight := pdf.GetPageSize()

	xMin, xMax := calculateMinMaxPoints(pageWidth, 5.0, LR_MARGIN, LR_MARGIN)
	yMin, yMax := calculateMinMaxPoints(pageHeight, 5.0, TOP_MARGIN, BOTTOM_MARGIN)

	log.Printf("%v %v", xMin, xMax)
	log.Printf("%v %v", yMin, yMax)

	for x := xMin; x <= xMax; x += 5 {
		pdf.Line(x, yMin, x, yMax)
	}

	for y := yMin; y <= yMax; y += 5 {
		pdf.Line(xMin, y, xMax, y)
	}

	err := pdf.Output(writter)
	return err
}


func calculateMinMaxPoints(total float64, step float64, marginStart float64, marginEnd float64) (min float64, max float64) {
	allowedUsage := total - (marginStart + marginEnd)
	maxRoundUsage := allowedUsage - math.Mod(total, step)

	additionalMargin := (allowedUsage - maxRoundUsage) / 2

	min = marginStart + additionalMargin
	max = total - (marginEnd + additionalMargin)

	e := math.Mod(max - min, step)
	//max = max - e

	log.Printf("allowedUsage %v maxRound: %v margin %v e %v", allowedUsage, maxRoundUsage, additionalMargin, e)

	return
}
