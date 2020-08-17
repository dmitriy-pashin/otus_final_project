package main

import (
	"fmt"
	"os"
	"otus_final_project/src/component/app"
)

const (
	DefaultConfigPath = "./configs"
)

func main() {
	var appName = "Web"
	confApp := app.InitConfig(DefaultConfigPath)
	application := app.NewWeb(confApp, app.DefaultWebAppPort)

	fmt.Printf("\nStarting %s app...\n", appName)
	fmt.Printf("+%v", application)

	fmt.Printf("+%v", application.Config())

	fmt.Printf("%+v", confApp)

	go app.WaitStopSignal()

	application.Run()

	fmt.Println("App is stopped")
	os.Exit(0)
}
