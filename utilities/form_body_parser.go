package utilities

import (
	"github.com/gofiber/fiber/v2"
	"ubersnap/static"
)

func RequestBodyParser(in interface{}, ctx *fiber.Ctx) error {
	if err := ctx.BodyParser(in); err != nil {
		err = static.FAILED_PARSE_FORM_BODY
		return ErrorRequest(err, fiber.StatusBadRequest)
	}

	return nil
}
