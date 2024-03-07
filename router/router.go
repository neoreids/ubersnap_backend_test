package router

import (
	"github.com/aws/smithy-go/ptr"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"ubersnap/config"
	"ubersnap/static"
	"ubersnap/utilities"
)

type OwnRouter struct {
	*fiber.App
}

func NewRouter(
	config *config.Config) *OwnRouter {
	isProduction := config.Production

	router := fiber.New()
	if isProduction {
		router.Use(recover.New())
	}
	router.Use(cors.New())
	router.Use(logger.New())
	router.Static(static.STATIC_ASSET_PATH, static.PUBLIC_DIR)
	router.Get("/health", func(ctx *fiber.Ctx) error {
		return utilities.Response(ctx, &utilities.ResponseRequest{
			Message: ptr.String("it works!"),
		})
	})
	return &OwnRouter{router}
}
