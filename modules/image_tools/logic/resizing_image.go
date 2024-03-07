package logic

import (
	"github.com/gofiber/fiber/v2"
	"gocv.io/x/gocv"
	"image"
	"io"
	"ubersnap/static"
	"ubersnap/utilities"
)

func (a *ImageRouteLogic) ResizingImage(f io.Reader, size image.Point) (*gocv.Mat, *gocv.Mat, error) {
	// Convert image to gocv.Mat
	mat, errConvertMat := a.ConvertFileHeaderToMat(f)
	if errConvertMat != nil {
		return nil, nil, utilities.ErrorRequest(static.UNABLE_CONVERT_IMAGE_TO_MAT, fiber.StatusInternalServerError)
	}

	// Resize image
	resized := gocv.NewMat()
	gocv.Resize(*mat, &resized, size, 0, 0, gocv.InterpolationDefault)
	return mat, &resized, nil
}
