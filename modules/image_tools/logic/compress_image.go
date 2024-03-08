package logic

import (
	"gocv.io/x/gocv"
	"io"
	"path/filepath"
	"ubersnap/utilities"
)

var mapExtension = map[string]int{
	".jpg": gocv.IMWriteJpegQuality,
	".png": gocv.IMWritePngCompression,
}

func (a *ImageRouteLogic) CompressImage(file io.Reader, quality int, pathFile string) (*gocv.Mat, *gocv.NativeByteBuffer, error) {
	ext := filepath.Ext(pathFile)
	fileMat, errDecode := a.ConvertFileHeaderToMat(file)
	if errDecode != nil {
		return nil, nil, errDecode
	}

	if ext != ".jpg" {
		quality = utilities.ConvertScale(quality)
	}

	buf, errEncode := gocv.IMEncodeWithParams(gocv.FileExt(ext), *fileMat, []int{mapExtension[ext], quality})

	if errEncode != nil {
		return fileMat, nil, errEncode
	}

	return fileMat, buf, nil
}
