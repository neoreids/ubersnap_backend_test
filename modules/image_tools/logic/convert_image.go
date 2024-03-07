package logic

import (
	"gocv.io/x/gocv"
	"io"
)

func (i *ImageRouteLogic) ConvertImage(file io.Reader) (*gocv.Mat, *gocv.Mat, error) {

	imgMat, errConvertToMat := i.ConvertFileHeaderToMat(file)
	if errConvertToMat != nil {
		return nil, nil, errConvertToMat
	}

	// Convert the image to BGR format
	bgrImg := gocv.NewMat()
	gocv.CvtColor(*imgMat, &bgrImg, gocv.ColorBGRToRGB)

	return imgMat, &bgrImg, nil
}
