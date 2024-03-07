package logic

import (
	"github.com/gofiber/fiber/v2"
	"gocv.io/x/gocv"
	"image"
	"io"
	"ubersnap/static"
	"ubersnap/utilities"
)

// ConvertFileHeaderToMat decode image file and convert to matRGB opencv
func (a *ImageRouteLogic) ConvertFileHeaderToMat(f io.Reader) (*gocv.Mat, error) {
	// Read image from file of opened by multipart.FileHeader before
	img, _, errDecode := image.Decode(f)
	if errDecode != nil {
		return nil, utilities.ErrorRequest(static.UNABLE_DECODE_IMAGE, fiber.StatusInternalServerError)
	}

	// Convert image to gocv.Mat
	mat, errConvertMat := gocv.ImageToMatRGB(img)
	if errConvertMat != nil {
		return nil, utilities.ErrorRequest(static.UNABLE_CONVERT_IMAGE_TO_MAT, fiber.StatusInternalServerError)
	}

	return &mat, nil
}
