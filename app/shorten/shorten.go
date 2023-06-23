/*
 *  Copyright (c) 2020-2023 Mikhail Knyazhev <markus621@yandex.ru>. All rights reserved.
 *  Use of this source code is governed by a GPL-3.0 license that can be found in the LICENSE file.
 */

package shorten

//go:generate easyjson

import (
	"crypto/sha256"
	"fmt"
	"net/url"

	goshorten "github.com/osspkg/go-algorithms/shorten"
	"github.com/osspkg/go-sdk/app"
	"github.com/osspkg/go-sdk/log"
	"github.com/osspkg/go-sdk/orm"
	"github.com/osspkg/goppy/plugins"
	"github.com/osspkg/goppy/plugins/database"
	"github.com/osspkg/goppy/plugins/web"
)

var Plugin = plugins.Plugin{
	Config: &Config{},
	Inject: New,
}

type Shorten struct {
	route web.Router
	codec *goshorten.Shorten
	db    orm.Stmt
}

func New(r web.RouterPool, c *Config, db database.MySQL) *Shorten {
	return &Shorten{
		route: r.Main(),
		codec: goshorten.New(c.Shorten),
		db:    db.Pool("main"),
	}
}

func (v *Shorten) Up(ctx app.Context) (err error) {
	v.route.Get("/~{code}", v.Index)
	v.route.Get("/+", v.Add)

	return nil
}

func (v *Shorten) Down() error {
	return nil
}

func (v *Shorten) Index(ctx web.Context) {
	code, err := ctx.Param("code").String()
	if err != nil {
		ctx.String(404, page404HTML)
		ctx.Log().WithFields(log.Fields{
			"err": err.Error(),
			"key": "id",
		}).Errorf("Invalid shorten key")
		return
	}

	id := v.codec.Decode(code)
	var uri string
	err = v.db.QueryContext("select_shorten", ctx.Context(), func(q orm.Querier) {
		q.SQL("SELECT `data` FROM `shorten` WHERE `id` = ? AND `lock` = 0 LIMIT 1;", id)
		q.Bind(func(bind orm.Scanner) error {
			return bind.Scan(&uri)
		})
	})
	if err != nil || len(uri) == 0 {
		ctx.String(404, page404HTML)
		return
	}

	ctx.Redirect(uri)
}

//easyjson:json
type AddModel struct {
	URI    string `json:"uri"`
	Source string `json:"source"`
}

func (v *Shorten) Add(ctx web.Context) {
	uri := ctx.URL().Query().Get("uri")
	if len(uri) == 0 {
		ctx.ErrorJSON(400, fmt.Errorf("invalid `uri`: is empty"), nil)
		return
	}
	u, err := url.Parse(uri)
	if err != nil {
		ctx.ErrorJSON(400, fmt.Errorf("invalid `uri`: %s", err.Error()), nil)
		return
	}
	if len(u.Scheme) == 0 {
		ctx.ErrorJSON(400, fmt.Errorf("invalid `uri`: scheme is empty"), nil)
		return
	}
	if u.Hostname() == ctx.URL().Hostname() {
		ctx.ErrorJSON(400, fmt.Errorf("invalid `uri`: unsupported hostname"), nil)
		return
	}

	h := sha256.New()
	//nolint:staticcheck
	if _, err = fmt.Fprintf(h, uri); err != nil {
		ctx.ErrorJSON(400, fmt.Errorf("invalid `uri`: hashing"), nil)
		return
	}
	hash := fmt.Sprintf("%x", h.Sum(nil))

	var id uint64
	err = v.db.ExecContext("insert_new_shorten", ctx.Context(), func(q orm.Executor) {
		q.SQL("INSERT INTO `shorten` (`data`, `hash`, `lock`, `created_at`) VALUES (?, ?, 0, now());")
		q.Params(uri, hash)
		q.Bind(func(result orm.Result) error {
			id = uint64(result.LastInsertId)
			if id == 0 {
				return fmt.Errorf("invalid insert")
			}
			return nil
		})
	})
	if err != nil {
		err = v.db.QueryContext("select_shorten", ctx.Context(), func(q orm.Querier) {
			q.SQL("SELECT `id` FROM `shorten` WHERE `hash` = ? LIMIT 1;", hash)
			q.Bind(func(bind orm.Scanner) error {
				return bind.Scan(&id)
			})
		})
		if err != nil {
			ctx.ErrorJSON(400, fmt.Errorf("invalid `uri`: cant save"), nil)
			return
		}
	}

	code := v.codec.Encode(id)

	model := &AddModel{
		URI:    fmt.Sprintf("https://%s/~%s", ctx.URL().Host, code),
		Source: uri,
	}

	ctx.JSON(200, model)
}
