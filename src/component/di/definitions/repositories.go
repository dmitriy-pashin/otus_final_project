package definitions

import (
	"github.com/sarulabs/di"
	"otus_final_project/src/component/config"
	"otus_final_project/src/component/db"
	"otus_final_project/src/repository"
)

func Repositories(configApp *config.Config) []di.Def {
	var definitions = []di.Def{
		{
			Name:  "app.repository.blacklist",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				componentDB := ctn.Get("app.component.db").(*db.DB)
				repBase := &repository.Base{DB: componentDB}

				return &repository.Blacklist{
					Base: repBase,
				}, nil
			},
		},
		{
			Name:  "app.repository.whitelist",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				componentDB := ctn.Get("app.component.db").(*db.DB)
				repBase := &repository.Base{DB: componentDB}

				return &repository.Whitelist{
					Base: repBase,
				}, nil
			},
		},
	}

	return definitions
}
