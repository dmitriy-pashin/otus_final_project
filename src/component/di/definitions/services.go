package definitions

import (
	"otus_final_project/src/component/config"
	"otus_final_project/src/repository"
	"otus_final_project/src/service"

	"github.com/sarulabs/di"
)

func Services(configApp *config.Config) []di.Def {
	definitions := []di.Def{
		{
			Name:  "app.service.blacklist",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				rep := ctn.Get("app.repository.blacklist").(*repository.Blacklist)

				return service.NewBlacklist(rep), nil
			},
		},
		{
			Name:  "app.service.whitelist",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				rep := ctn.Get("app.repository.whitelist").(*repository.Whitelist)

				return service.NewWhitelist(rep), nil
			},
		},
		{
			Name:  "app.service.login",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				repWhitelist := ctn.Get("app.repository.whitelist").(*repository.Whitelist)
				repBlacklist := ctn.Get("app.repository.blacklist").(*repository.Blacklist)

				return service.NewLogin(repWhitelist, repBlacklist, configApp.Limits), nil
			},
		},
	}

	return definitions
}
