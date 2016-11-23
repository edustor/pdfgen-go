package source

import (
	"github.com/jung-kurt/gofpdf"
	"io"
	"math"
	"log"
	"github.com/edustor/gen/bindata"
)

const LR_MARGIN float64 = 5.0
const TOP_MARGIN float64 = 7.5
const BOTTOM_MARGIN float64 = 10.0

func GenPdf(writter io.Writer, pageCount int) (err error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	template := pdf.CreateTemplate(createTemplate)

	pdf.AddPage()
	pdf.UseTemplate(template)

	err = pdf.Output(writter)
	return err
}

func createTemplate(tpl *gofpdf.Tpl) {
	fontJson, err := bindata.Asset("fonts/Proxima Nova Thin.json")
	if err != nil {
		log.Panic(err)
	}
	fontZ, err := bindata.Asset("fonts/Proxima Nova Thin.z")
	if err != nil {
		log.Panic(err)
	}
	tpl.AddFontFromBytes("Proxima Nova", "Thin", fontJson, fontZ)
	tpl.SetFont("Proxima Nova", "Thin", 11)
	tpl.SetMargins(0, 0, 0)
	tpl.SetAutoPageBreak(false, 0)

	tpl.SetLineWidth(0.2)
	tpl.SetDrawColor(128, 128, 128)

	pageWidth, pageHeight := tpl.GetPageSize()

	xMin, xMax := calculateMinMaxPoints(pageWidth, 5.0, LR_MARGIN, LR_MARGIN)
	yMin, yMax := calculateMinMaxPoints(pageHeight, 5.0, TOP_MARGIN, BOTTOM_MARGIN)

	// Grid
	for x := xMin; x <= xMax; x += 5 {
		tpl.Line(x, yMin, x, yMax)
	}

	for y := yMin; y <= yMax; y += 5 {
		tpl.Line(xMin, y, xMax, y)
	}

	// Header
	tpl.SetY(yMin - 2.7)
	tpl.SetX(xMin - 1.4)
	tpl.CellFormat(1, 1, "Edustor Alpha", "", 0, "", false, 0, "")

	nextIdField := "#__________"
	tpl.SetX(xMax - 1 - tpl.GetStringWidth(nextIdField))
	tpl.CellFormat(1, 1, nextIdField, "", 0, "", false, 0, "")
}

func calculateMinMaxPoints(total float64, step float64, marginStart float64, marginEnd float64) (min float64, max float64) {
	allowedUsage := total - (marginStart + marginEnd)
	maxRoundUsage := allowedUsage - math.Mod(allowedUsage, step)

	additionalMargin := (allowedUsage - maxRoundUsage) / 2

	min = marginStart + additionalMargin
	max = total - (marginEnd + additionalMargin)

	e := math.Mod(max - min, step)
	log.Printf("allowedUsage %v maxRoundUsage: %v additionalMargin %v e %v", allowedUsage, maxRoundUsage, additionalMargin, e)

	return
}
