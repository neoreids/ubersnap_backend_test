package utilities

import (
	"github.com/gofiber/fiber/v2"
	"ubersnap/static"
)

func QueryParamsParser(in interface{}, ctx *fiber.Ctx) error {
	if err := ctx.QueryParser(in); err != nil {
		return ErrorRequest(static.FAILED_PARSE_QUERY_PARAMS, fiber.StatusBadRequest)
	}
	return nil
}
