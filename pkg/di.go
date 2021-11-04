package pkg

import (
	"github.com/dewep-online/uri-one/pkg/database"
	"github.com/dewep-online/uri-one/pkg/encode"
	"github.com/deweppro/go-app/application"
	"github.com/deweppro/go-orm/schema/sqlite"
)

var (
	//Module di injector
	Module = application.Modules{
		database.New,
		encode.New(),
	}
	//Config di injector
	Config = application.Modules{
		&sqlite.Config{},
	}
)
