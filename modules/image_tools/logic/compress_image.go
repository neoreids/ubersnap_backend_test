package logic

import (
	"gocv.io/x/gocv"
	"io"
)

func (a *ImageRouteLogic) CompressImage(file io.Reader, quality int, pathFile string) (*gocv.Mat, bool, error) {
	fileMat, errDecode := a.ConvertFileHeaderToMat(file)
	if errDecode != nil {
		return nil, false, errDecode
	}

	// Compress the image by adjusting quality
	compressionParams := []int{gocv.IMWriteJpegQuality, quality} // Set JPEG quality to 50%

	return fileMat, gocv.IMWriteWithParams(pathFile, *fileMat, compressionParams), nil
}
