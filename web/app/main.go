package main

import (
	"fmt"
	"os"
	"otus_final_project/src/component/app"

	_ "github.com/jackc/pgx/v4/stdlib"
)

const (
	DefaultConfigPath = "./configs"
)

func main() {
	appName := "Web"
	confApp := app.InitConfig(DefaultConfigPath)
	application := app.NewWeb(confApp, app.DefaultWebAppPort)

	fmt.Printf("\nStarting %s app...\n", appName)

	go app.WaitStopSignal()

	application.Run()

	fmt.Println("App is stopped")
	os.Exit(0)
}
