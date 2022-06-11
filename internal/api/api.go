package api

import (
	"net/http"

	"github.com/dewep-online/uri-one/pkg/database"
	"github.com/dewep-online/uri-one/pkg/encode"
	"github.com/deweppro/go-badges"
	"github.com/deweppro/go-http/pkg/routes"
	"github.com/deweppro/go-logger"
	"github.com/deweppro/go-static"
	"github.com/pkg/errors"
)

//go:generate static ./../../web UI

//UI static archive
var UI = "H4sIAAAAAAAA/2IYBaNgFIxYAAgAAP//Lq+17wAEAAA="

//API model
type API struct {
	log    logger.Logger
	cache  *static.Cache
	route  *routes.Router
	conf   *MiddlewareConfig
	db     *database.Database
	enc    *encode.Enc
	badges *badges.Badges
}

func New(l logger.Logger, r *routes.Router, c *MiddlewareConfig, e *encode.Enc, d *database.Database) *API {
	return &API{
		log:   l,
		cache: static.New(),
		route: r,
		conf:  c,
		enc:   e,
		db:    d,
	}
}

//Up startup api service
func (v *API) Up() error {
	var err error

	if v.badges, err = badges.New(); err != nil {
		return errors.Wrap(err, "badges init")
	}

	if err = v.cache.FromBase64TarGZ(UI); err != nil {
		return errors.Wrap(err, "unpack ui")
	}

	v.route.Global(routes.RecoveryMiddleware(v.log))
	v.route.Global(routes.ThrottlingMiddleware(v.conf.Middleware.Throttling))
	v.route.Global(v.BadgesMiddleware())
	v.route.Global(v.DetectLinkMiddleware())

	for _, file := range v.cache.List() {
		logger.Debugf("static: %s", file)
		v.route.Route(file, v.Index, http.MethodGet)
	}
	v.route.Route("/", v.Index, http.MethodGet)
	v.route.Route("/+", v.Add, http.MethodGet)

	return nil
}

//Down shutdown api service
func (v *API) Down() error {
	return nil
}
