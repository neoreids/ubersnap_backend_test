package handler

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"ubersnap/modules/image_tools/dto"
	"ubersnap/modules/image_tools/logic"
	"ubersnap/utilities"
)

type (
	ImageRouteHandler struct {
		fx.In
		ImageLogic logic.ImageRouteLogic
	}
)

// Resize handler for resizing image endpoint
// form field image: Binary
// width: int
// height: int
func (a *ImageRouteHandler) Resize(ctx *fiber.Ctx) error {
	requestObject := new(dto.ImageResizeRequest)
	// parse form into object
	if err := utilities.RequestBodyParser(requestObject, ctx); err != nil {
		return utilities.Response(ctx, &utilities.ResponseRequest{
			Error: err,
		})
	}

	// validate the object instead form validation
	if err := utilities.ValidateFormFile(requestObject, ctx); err != nil {
		return utilities.Response(ctx, &utilities.ResponseRequest{
			Error: err,
		})
	}

	// parse FileHeader pointer into IO.Reader and set domain field struct to generate image URL in the next process
	if err := requestObject.AdditionalParser(ctx); err != nil {
		return utilities.Response(ctx, &utilities.ResponseRequest{
			Error: err,
		})
	}

	resizeLogic := a.ImageLogic.Resize(requestObject)
	return utilities.Response(ctx, resizeLogic)
}

// Compress handler for compressing image file size
// form field image: Binary
// quality: int with percent
func (a *ImageRouteHandler) Compress(ctx *fiber.Ctx) error {
	requestObject := new(dto.ImageCompressRequest)
	// parse form into object
	if err := utilities.RequestBodyParser(requestObject, ctx); err != nil {
		return utilities.Response(ctx, &utilities.ResponseRequest{
			Error: err,
		})
	}

	// validate the object instead form validation
	if err := utilities.ValidateFormFile(requestObject, ctx); err != nil {
		return utilities.Response(ctx, &utilities.ResponseRequest{
			Error: err,
		})
	}

	// parse FileHeader pointer into IO.Reader and set domain field struct to generate image URL in the next process
	if err := requestObject.AdditionalParser(ctx); err != nil {
		return utilities.Response(ctx, &utilities.ResponseRequest{
			Error: err,
		})
	}

	resizeLogic := a.ImageLogic.Compress(requestObject)
	return utilities.Response(ctx, resizeLogic)
}

// Convert handler for convert image PNG to JPEG file
// form field image: Binary
func (a *ImageRouteHandler) Convert(ctx *fiber.Ctx) error {
	requestObject := new(dto.ImageConvertRequest)
	// parse form into object
	if err := utilities.RequestBodyParser(requestObject, ctx); err != nil {
		return utilities.Response(ctx, &utilities.ResponseRequest{
			Error: err,
		})
	}

	// validate the object instead form validation
	if err := utilities.ValidateFormFile(requestObject, ctx); err != nil {
		return utilities.Response(ctx, &utilities.ResponseRequest{
			Error: err,
		})
	}

	// parse FileHeader pointer into IO.Reader and set domain field struct to generate image URL in the next process
	if err := requestObject.AdditionalParser(ctx); err != nil {
		return utilities.Response(ctx, &utilities.ResponseRequest{
			Error: err,
		})
	}

	resizeLogic := a.ImageLogic.Convert(requestObject)
	return utilities.Response(ctx, resizeLogic)
}