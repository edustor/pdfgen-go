package source

import (
	"bytes"
	"github.com/oliamb/cutter"
	qr "github.com/skip2/go-qrcode"
	"image/jpeg"
)

func GenQR(content string) ([]byte, error) {

	code, err := qr.New(content, qr.Low)

	if err != nil {
		return nil, err
	}

	image := code.Image(200)

	image, err = cutter.Crop(image, cutter.Config{
		Width:  135,
		Height: 135,
		Mode:   cutter.Centered,
	})

	buffer := bytes.NewBuffer(make([]byte, 0))

	jpeg.Encode(buffer, image, &jpeg.Options{Quality: 100})

	return buffer.Bytes(), nil
}