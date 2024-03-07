package tests

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"ubersnap/utilities"
)

func TestConvert(t *testing.T) {
	file, err := os.Open(Image5MB)
	if err != nil {
		t.Fatalf("error open image file test : %v", err)
	}

	defer file.Close()

	originalMat, converted, errResizing := ImageLogic.ConvertImage(file)
	if errResizing != nil {
		t.Fatalf("error resizing image : %v", errResizing)
	}
	defer originalMat.Close()
	defer converted.Close()

	mime, errGetMime := utilities.DetectMimeTypeOfMat(*converted)
	if errGetMime != nil {
		t.Fatalf("error while getting mime : %v", errGetMime)
	}

	assert.Equal(t, "image/jpeg", mime)
}
