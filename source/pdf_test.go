package source

import (
	"testing"
	"os"
	"log"
	"github.com/stretchr/testify/assert"
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

func TestCalculateMinMax(t *testing.T) {
	max, min := calculateMinMaxPoints(210, 5, 5, 5)
	assert.EqualValues(t, 5, max)
	assert.EqualValues(t, 205, min)

	max, min = calculateMinMaxPoints(297, 5, 5, 5)
	assert.EqualValues(t, 6, max)

	assert.EqualValues(t, 291, min)
}