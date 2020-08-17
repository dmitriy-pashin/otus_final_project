package app

import (
	"github.com/sarulabs/di"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"otus_final_project/src/component/config"
	componentDi "otus_final_project/src/component/di"
	"syscall"
)

type Base struct {
	diContainer di.Container
	conf        *config.Config
}

func NewBase(conf *config.Config) *Base {
	app := &Base{
		conf: conf,
	}
	app.Init()

	return app
}

func (app *Base) Close() {
	if err := app.Container().Delete(); err != nil {
		log.Fatal(err)
	}
}

func (app *Base) Init() {
	app.diContainer = componentDi.InitDefinitions(app.conf)
}

func (app *Base) Container() di.Container {
	return app.diContainer
}

func (app *Base) Config() *config.Config {
	return app.conf
}

func InitConfig(configPath string, fileParams ...string) *config.Config {
	var fileName string
	var fileExtension string
	var count int

	for _, v := range fileParams {
		switch count {
		case 0:
			fileName = v
		case 1:
			fileExtension = v
		}

		count++
	}
	var conf = config.NewConfig(configPath, fileName, fileExtension)

	conf.Fill()

	return conf
}

func WaitStopSignal() os.Signal {
	var takenSignal os.Signal
	var ch = make(chan os.Signal, 1)

	signal.Notify(ch,
		os.Interrupt,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	log.Debugln("Wait stop signal...")
	takenSignal = <-ch

	log.Debugln("got signal:", takenSignal)

	return takenSignal
}
