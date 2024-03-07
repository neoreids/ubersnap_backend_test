package logic

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"gocv.io/x/gocv"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"path/filepath"
	"strconv"
	"ubersnap/modules/image_tools/dto"
	"ubersnap/static"
	"ubersnap/utilities"
)

// ImageRouteLogic the struct for initiate any requirement for this logic like repository, data cache or anything else
type (
	ImageRouteLogic struct {
		fx.In
	}
)

// Resize resizing image with gocv based from opencv library
// will return response with path of resized image
func (a *ImageRouteLogic) Resize(req *dto.ImageResizeRequest) *utilities.ResponseRequest {
	// get base name of uploaded file and generate new name for resized file
	fileName := req.Image.Filename
	extFile := filepath.Ext(fileName)
	nameWoExt := utilities.FileNameWithoutExtension(fileName)
	newName := fmt.Sprintf("%s-%s-%s%s", nameWoExt, strconv.Itoa(req.Width), strconv.Itoa(req.Height), extFile)

	matOriginalFile, resized, errResizing := a.ResizingImage(req.File, image.Point{
		X: req.Width,
		Y: req.Height,
	})

	if errResizing != nil {
		return &utilities.ResponseRequest{
			Error: errResizing,
		}
	}

	defer func(matOriginalFile *gocv.Mat) {
		_ = matOriginalFile.Close()
	}(matOriginalFile)

	defer func(resized *gocv.Mat) {
		_ = resized.Close()
	}(resized)

	// Save resized image
	gocv.IMWrite(fmt.Sprintf("%s/%s", static.PUBLIC_DIR, newName), *resized)

	return &utilities.ResponseRequest{
		Data: fmt.Sprintf("%s%s/%s", req.Domain, static.STATIC_ASSET_PATH, newName),
	}
}

// Compress reduce image file size
// will return response with path of compressed image
func (a *ImageRouteLogic) Compress(req *dto.ImageCompressRequest) *utilities.ResponseRequest {
	// get base name of uploaded file and generate new name for resized file
	fileName := req.Image.Filename
	extFile := filepath.Ext(fileName)
	nameWoExt := utilities.FileNameWithoutExtension(fileName)
	newName := fmt.Sprintf("%s-compressed%s", nameWoExt, extFile)

	outputFilename := fmt.Sprintf("%s/%s", static.PUBLIC_DIR, newName)
	// compress image
	fileMat, ok, errCompress := a.CompressImage(req.File, req.Quality, outputFilename)
	if errCompress != nil {
		return &utilities.ResponseRequest{
			Error: errCompress,
		}
	}

	defer func(fileMat *gocv.Mat) {
		_ = fileMat.Close()
	}(fileMat)

	if !ok {
		return &utilities.ResponseRequest{
			Error: static.FAILED_COMPRESS_IMAGE,
		}
	}

	return &utilities.ResponseRequest{
		Data: fmt.Sprintf("%s%s/%s", req.Domain, static.STATIC_ASSET_PATH, newName),
	}
}

// Convert image from PNG to JPEG
// will return response with path of compressed image
func (a *ImageRouteLogic) Convert(req *dto.ImageConvertRequest) *utilities.ResponseRequest {
	// get base name of uploaded file and generate new name for resized file
	fileName := req.Image.Filename
	nameWoExt := utilities.FileNameWithoutExtension(fileName)
	newName := fmt.Sprintf("%s-convert%s", nameWoExt, ".jpg")

	outputFilename := fmt.Sprintf("%s/%s", static.PUBLIC_DIR, newName)

	originalFile, converted, errConvert := a.ConvertImage(req.File)
	if errConvert != nil {
		return &utilities.ResponseRequest{
			Error: errConvert,
		}
	}

	defer func(originalFile *gocv.Mat) {
		_ = originalFile.Close()
	}(originalFile)

	defer func(converted *gocv.Mat) {
		_ = converted.Close()
	}(converted)

	if ok := gocv.IMWrite(outputFilename, *converted); !ok {
		return &utilities.ResponseRequest{
			Error: utilities.ErrorRequest(static.FAILED_SAVE_IMAGE, fiber.StatusInternalServerError),
		}
	}

	return &utilities.ResponseRequest{
		Data: fmt.Sprintf("%s%s/%s", req.Domain, static.STATIC_ASSET_PATH, newName),
	}
}
