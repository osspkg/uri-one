package main

import (
	"flag"

	"uri-one/providers/db"

	"uri-one/app/modules/httpserv"

	"uri-one/app/config"

	"github.com/deweppro/core/pkg/app"
)

var (
	cpath = flag.String("config", "source/config.yaml", "config file path")
)

func main() {
	flag.Parse()

	app.New(
		*cpath,
		app.NewInterfaces().Add(
			//&debug.ConfigDebug{},
			&config.Config{},
		),
		app.NewInterfaces().Add(
			//debug.New,
			db.MustNew,
			httpserv.MustNew,
		),
	).Run()
}
