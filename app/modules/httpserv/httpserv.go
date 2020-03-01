package httpserv

import (
	"strings"
	"sync"

	"uri-one/app/config"
	"uri-one/providers/db"

	"github.com/deweppro/core/servers/http"
)

type HttpSrv struct {
	cfg   *config.Config
	db    *db.SQLite
	srv   *http.Server
	cache map[string]string
	sync.RWMutex
}

func MustNew(cfg *config.Config, db *db.SQLite) *HttpSrv {
	return &HttpSrv{
		cfg:   cfg,
		db:    db,
		srv:   http.New(),
		cache: make(map[string]string),
	}
}

func (h *HttpSrv) Start() error {

	h.srv.SetAddr(strings.Join([]string{"127.0.0.1", h.cfg.Http.Port}, ":"))

	h.srv.Route("GET", "/new", h.New)
	h.srv.Route("*", "", h.Get)

	return h.srv.Start()
}

func (h *HttpSrv) Stop() error {
	return h.srv.Stop()
}
