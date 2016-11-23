package source

import (
	"github.com/jung-kurt/gofpdf"
	"io"
	"math"
	"log"
	"github.com/edustor/gen/bindata"
	"github.com/satori/go.uuid"
	"fmt"
	"strings"
)

const LR_MARGIN float64 = 5.0
const TOP_MARGIN float64 = 7.5
const BOTTOM_MARGIN float64 = 10.0

func GenPdf(writter io.Writer, pageCount int) (err error) {
	pdf := gofpdf.New("P", "mm", "A4", "")

	for i:=0; i < pageCount; i++ {
		pdf.AddPage()

		pageId := uuid.NewV4().String()
		uri := fmt.Sprintf("edustor://d/%s", pageId)
		qr, err := GenQR(uri)
		if err != nil {
			panic(err)
		}

		_idParts := strings.Split(pageId, "-")
		_idTail := _idParts[len(_idParts) - 1]

		topId := fmt.Sprintf("#%s  #__________", _idTail[8:12])
		bottomId := fmt.Sprintf("#%s-%s-%s", _idTail[0:4], _idTail[4:8], _idTail[8:12])

		drawPage(pdf, topId, bottomId, qr)
	}

	err = pdf.Output(writter)
	return err
}


func drawPage(pdf *gofpdf.Fpdf, topId string, bottomId string, qr []byte) {
	fontJson, err := bindata.Asset("fonts/Proxima Nova Thin.json")
	if err != nil {
		log.Panic(err)
	}
	fontZ, err := bindata.Asset("fonts/Proxima Nova Thin.z")
	if err != nil {
		log.Panic(err)
	}
	pdf.AddFontFromBytes("Proxima Nova", "Thin", fontJson, fontZ)
	pdf.SetFont("Proxima Nova", "Thin", 11)
	pdf.SetMargins(0, 0, 0)
	pdf.SetAutoPageBreak(false, 0)

	pdf.SetLineWidth(0.2)
	pdf.SetDrawColor(128, 128, 128)

	pageWidth, pageHeight := pdf.GetPageSize()
	xMin, xMax := calculateMinMaxPoints(pageWidth, 5.0, LR_MARGIN, LR_MARGIN)
	yMin, yMax := calculateMinMaxPoints(pageHeight, 5.0, TOP_MARGIN, BOTTOM_MARGIN)

	// Grid
	for x := xMin; x <= xMax; x += 5 {
		pdf.Line(x, yMin, x, yMax)
	}

	for y := yMin; y <= yMax; y += 5 {
		pdf.Line(xMin, y, xMax, y)
	}

	// Header
	pdf.SetY(yMin - 2.7)
	pdf.SetX(xMin - 1.4)
	pdf.CellFormat(1, 1, "Edustor Alpha", "", 0, "", false, 0, "")

	pdf.SetX(xMax - 1 - pdf.GetStringWidth(topId))
	pdf.CellFormat(1, 1, topId, "", 0, "", false, 0, "")
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
