package definitions

import (
	"otus_final_project/src/component/config"
	"otus_final_project/src/component/db"

	"github.com/sarulabs/di"
)

func Components(configApp *config.Config) []di.Def {
	definitions := []di.Def{
		{
			Name:  "app.config",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				return configApp, nil
			},
		},
		{
			Name:  "app.component.db",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				return db.NewDB(configApp.DB), nil
			},
		},
	}

	return definitions
}
