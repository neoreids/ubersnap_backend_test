package cmd

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"ubersnap/config"
	"ubersnap/modules"
	"ubersnap/router"
)

var httpCommand = &cobra.Command{
	Use:   "start",
	Short: "start HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		fx.New(
			fx.Provide(config.Set),
			fx.Provide(router.NewRouter),
			modules.RouterModule,
			fx.Invoke(hooks),
		).Run()
	},
}

func hooks(ls fx.Lifecycle, router *router.OwnRouter, c *config.Config) {
	hook := fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				err := router.App.Listen(fmt.Sprintf(":%d", c.Server.Port))
				if err != nil {
					log.Fatal(err)
				}
			}()
			return nil
		},
	}

	ls.Append(hook)
}
