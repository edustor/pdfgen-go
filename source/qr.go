package source

import (
	"bytes"
	"github.com/oliamb/cutter"
	qr "github.com/skip2/go-qrcode"
	"image/png"
)

func GenQR(content string) ([]byte, error) {

	code, err := qr.New(content, qr.Low)

	if err != nil {
		return nil, err
	}

	image := code.Image(200)

	image, err = cutter.Crop(image, cutter.Config{
		Width:  150,
		Height: 150,
		Mode:   cutter.Centered,
	})

	buffer := bytes.NewBuffer(make([]byte, 0))

	png.Encode(buffer, image)

	return buffer.Bytes(), nil
}