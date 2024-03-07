package route

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"ubersnap/modules/image_tools/handler"
	"ubersnap/router"
)

type ImageTools struct {
	fx.In
	Router  *router.OwnRouter
	ImageRouteHandler handler.ImageRouteHandler
}

func NewImageRoute(r ImageTools) {
	imageRoute := fiber.New()
	imageRoute.Post("/resize", r.ImageRouteHandler.Resize)
	imageRoute.Post("/convert", r.ImageRouteHandler.Convert)
	imageRoute.Post("/compress", r.ImageRouteHandler.Compress)
	r.Router.Mount("/image", imageRoute)
}