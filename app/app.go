package app

import (
	"uri-one/app/config"
	"uri-one/app/modules/httpserv"
	"uri-one/app/modules/logger"
	"uri-one/providers/db"

	"github.com/deweppro/core/event"
	"github.com/deweppro/core/initializer"
	"github.com/sirupsen/logrus"
)

type App struct {
	init []interface{}

	dep *initializer.Dependencies
}

func New(cpath string) *App {
	return &App{
		dep: initializer.New(),
		init: []interface{}{
			config.MustNew(cpath),
			db.MustNew,
			logger.MustNew,
			httpserv.MustNew,
		},
	}
}

func (app *App) Run() {

	if err := app.dep.Register(app.init...); err != nil {
		logrus.Fatal(err)
	}

	if err := app.dep.Start(); err != nil {
		logrus.Fatal(err)
	}

	event.OnSyscallStop(func() {
		if err := app.dep.Stop(); err != nil {
			logrus.Fatal(err)
		}
	})

}
