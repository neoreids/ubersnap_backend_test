package tests

import (
	"github.com/stretchr/testify/assert"
	"image"
	"os"
	"testing"
	"ubersnap/modules/image_tools/logic"
)

var ImageLogic = new(logic.ImageRouteLogic)
var Image2MB = "./Sample-jpg-image-2mb.jpg"
func TestResize(t *testing.T) {
	file, err := os.Open(Image2MB)
	if err != nil {
		t.Fatalf("error open image file test : %v", err)
	}

	defer file.Close()

	imagePoint := image.Point{
		X: 256,
		Y: 256,
	}

	originalMat, resized, errResizing := ImageLogic.ResizingImage(file, imagePoint)
	if errResizing != nil {
		t.Fatalf("error resizing image : %v", errResizing)
	}
	defer originalMat.Close()
	defer resized.Close()
	assert.Equal(t, resized.Size()[0], imagePoint.X)
	assert.Equal(t, resized.Size()[1], imagePoint.Y)
}
