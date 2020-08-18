package di

import (
	"otus_final_project/src/component/config"
	"otus_final_project/src/component/di/definitions"

	"github.com/sarulabs/di"
)

func InitDefinitions(configApp *config.Config) di.Container {
	definitionsGetters := []func(conf *config.Config) []di.Def{
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
		listDefinitions := getter(configApp)
		for _, d := range listDefinitions {
			err := builder.Add(d)
			if err != nil {
				panic(err)
			}
		}
	}

	return builder.Build()
}
