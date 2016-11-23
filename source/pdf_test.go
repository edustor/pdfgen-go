package source

import (
	"testing"
	"os"
	"log"
)

func TestGenPdf(t *testing.T) {
	os.Mkdir("../test", 0777)
	file, err := os.Create("../test/pdf.pdf")
	if err != nil {
		log.Panic(err)
	}

	err = GenPdf(file)
	if err != nil {
		log.Panic(err)
	}
}