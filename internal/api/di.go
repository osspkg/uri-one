package api

import (
	"github.com/deweppro/go-app/application"
	"github.com/deweppro/go-http/web/routes"
	"github.com/deweppro/go-http/web/server"
)

var (
	//Module di injector
	Module = application.Modules{
		server.New,
		routes.NewRouter,
		New,
	}
	//Config di injector
	Config = application.Modules{
		&server.Config{},
		&MiddlewareConfig{},
	}
)
