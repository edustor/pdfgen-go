package source

import (
	"testing"
	"os"
	"log"
)

func TestGenPdf(t *testing.T) {
	os.Mkdir("../test", 0777)
	file, err := os.Create("../test/A4.pdf")
	if err != nil {
		log.Panic(err)
	}

	err = GenPdf(file, 1)
	if err != nil {
		log.Panic(err)
	}
}