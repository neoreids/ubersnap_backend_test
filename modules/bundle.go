package modules

import (
	"go.uber.org/fx"
	"ubersnap/modules/image_tools/route"
)

var RouterModule = fx.Options(
	fx.Invoke(route.NewImageRoute),
)