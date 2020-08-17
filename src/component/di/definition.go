package di

import (
	"github.com/sarulabs/di"
	"otus_final_project/src/component/config"
	"otus_final_project/src/component/di/definitions"
)

func InitDefinitions(configApp *config.Config) di.Container {
	var definitionsGetters = []func(conf *config.Config) []di.Def{
		definitions.Components,
		definitions.Controllers,
		definitions.Services,
		definitions.Repositories,
	}

	builder, err := di.NewBuilder()

	if err != nil {
		panic(err)
	}

	for _, getter := range definitionsGetters {
		var listDefinitions = getter(configApp)
		for _, d := range listDefinitions {
			err := builder.Add(d)

			if err != nil {
				panic(err)
			}
		}
	}

	return builder.Build()
}
