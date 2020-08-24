package definitions

import (
	"otus_final_project/src/component/config"
	"otus_final_project/src/controller"

	"github.com/sarulabs/di"
)

func Controllers(configApp *config.Config) []di.Def {
	definitions := []di.Def{
		{
			Name:  "app.controller.blacklist",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				add := ctn.Get("app.service.blacklist").(controller.AddHandler)
				del := ctn.Get("app.service.blacklist").(controller.DeleteHandler)

				return controller.NewBlacklist(add, del), nil
			},
		},
		{
			Name:  "app.controller.whitelist",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				add := ctn.Get("app.service.whitelist").(controller.AddHandler)
				del := ctn.Get("app.service.whitelist").(controller.DeleteHandler)

				return controller.NewWhitelist(add, del), nil
			},
		},
		{
			Name:  "app.controller.login",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				login := ctn.Get("app.service.rate_limiter").(controller.LoginHandler)

				return controller.NewLogin(login), nil
			},
		},
	}

	return definitions
}
