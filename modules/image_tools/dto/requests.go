package dto

import (
	"github.com/gofiber/fiber/v2"
	"io"
	"mime/multipart"
	"ubersnap/static"
)

type (
	RequestObject interface {
		AdditionalParser(ctx *fiber.Ctx) error
	}

	ImageResizeRequest struct {
		Image  *multipart.FileHeader `json:"image" form:"image"`
		Domain string
		File   io.Reader
		Width  int `json:"width" form:"width" validate:"required"`
		Height int `json:"height" form:"height" validate:"required"`
	}

	ImageCompressRequest struct {
		Image   *multipart.FileHeader `json:"image" form:"image"`
		Domain  string
		File    io.Reader
		Quality int `json:"quality" form:"quality" validate:"required"`
	}

	ImageConvertRequest struct {
		Image   *multipart.FileHeader `json:"image" form:"image"`
		Domain  string
		File    io.Reader
	}
)


func (i *ImageResizeRequest) AdditionalParser(ctx *fiber.Ctx) error {
	// open the file from multipart.FileHeaader pointer of fiber request form file
	file, errOpenFile := i.Image.Open()
	if errOpenFile != nil {
		return static.FAILED_OPEN_FILE
	}
	defer func(file multipart.File) {
		_ = file.Close()
	}(file)

	// parse io.Reader file into request object
	i.File = file
	i.Domain = ctx.Hostname()
	return nil
}

func (i *ImageCompressRequest) AdditionalParser(ctx *fiber.Ctx) error {
	// open the file from multipart.FileHeaader pointer of fiber request form file
	file, errOpenFile := i.Image.Open()
	if errOpenFile != nil {
		return static.FAILED_OPEN_FILE
	}
	defer func(file multipart.File) {
		_ = file.Close()
	}(file)

	// parse io.Reader file into request object
	i.File = file
	i.Domain = ctx.Hostname()
	return nil
}

func (i *ImageConvertRequest) AdditionalParser(ctx *fiber.Ctx) error {
	// open the file from multipart.FileHeaader pointer of fiber request form file
	file, errOpenFile := i.Image.Open()
	if errOpenFile != nil {
		return static.FAILED_OPEN_FILE
	}
	defer func(file multipart.File) {
		_ = file.Close()
	}(file)

	// parse io.Reader file into request object
	i.File = file
	i.Domain = ctx.Hostname()
	return nil
}


