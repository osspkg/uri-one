package main

import (
	"github.com/dewep-online/uri-one/internal/api"
	"github.com/dewep-online/uri-one/pkg"
	"github.com/deweppro/go-app/application"
	"github.com/deweppro/go-app/console"
	"github.com/deweppro/go-logger"
)

func main() {
	root := console.New("uri-one", "help uri-one")
	root.AddCommand(appRun())
	root.Exec()
}

func appRun() console.CommandGetter {
	return console.NewCommand(func(setter console.CommandSetter) {
		setter.Setup("run", "run application")
		setter.Example("run --config=./config.yaml")
		setter.Flag(func(f console.FlagsSetter) {
			f.StringVar("config", "./config.yaml", "path to config file")
		})
		setter.ExecFunc(func(_ []string, config string) {
			application.New().
				Logger(logger.Default()).
				ConfigFile(
					config,
					pkg.Config,
					api.Config,
				).
				Modules(
					pkg.Module,
					api.Module,
				).
				Run()
		})
	})
}
