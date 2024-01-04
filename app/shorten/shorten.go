/*
 *  Copyright (c) 2020-2024 Mikhail Knyazhev <markus621@yandex.ru>. All rights reserved.
 *  Use of this source code is governed by a GPL-3.0 license that can be found in the LICENSE file.
 */

package shorten

//go:generate easyjson

import (
	"crypto/sha256"
	"fmt"
	"net/url"

	"github.com/osspkg/uri-one/app/mainapp"
	goshorten "go.osspkg.com/algorithms/shorten"
	"go.osspkg.com/goppy/orm"
	"go.osspkg.com/goppy/ormmysql"
	"go.osspkg.com/goppy/plugins"
	"go.osspkg.com/goppy/web"
	"go.osspkg.com/goppy/xlog"
)

var Plugin = plugins.Plugin{
	Config: &Config{},
	Inject: New,
}

type Shorten struct {
	route   web.Router
	codec   *goshorten.Shorten
	db      orm.Stmt
	address mainapp.Address
}

func New(r web.RouterPool, c *Config, db ormmysql.MySQL, d mainapp.Address) *Shorten {
	return &Shorten{
		route:   r.Main(),
		codec:   goshorten.New(c.Shorten),
		db:      db.Pool("main"),
		address: d,
	}
}

func (v *Shorten) Up() error {
	v.route.Get("/{code}", v.Index)
	v.route.Post("/add", v.Add)

	return nil
}

func (v *Shorten) Down() error {
	return nil
}

func (v *Shorten) Index(ctx web.Context) {
	code, err := ctx.Param("code").String()
	if err != nil {
		ctx.String(404, page404HTML)
		ctx.Log().WithFields(xlog.Fields{
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
	URL    string `json:"url"`
	Source string `json:"source"`
}

func (v *Shorten) Add(ctx web.Context) {
	result := &AddModel{}
	if err := ctx.BindJSON(result); err != nil {
		ctx.ErrorJSON(400, fmt.Errorf("invalid body: %s", err.Error()), nil)
		return
	}
	if len(result.Source) == 0 {
		ctx.ErrorJSON(400, fmt.Errorf("invalid `source`: is empty"), nil)
		return
	}
	u, err := url.Parse(result.Source)
	if err != nil {
		ctx.ErrorJSON(400, fmt.Errorf("invalid `source`: %s", err.Error()), nil)
		return
	}
	if len(u.Scheme) == 0 {
		ctx.ErrorJSON(400, fmt.Errorf("invalid `source`: scheme is empty"), nil)
		return
	}
	if u.Hostname() == string(v.address) {
		ctx.ErrorJSON(400, fmt.Errorf("invalid `source`: unsupported hostname"), nil)
		return
	}

	h := sha256.New()
	//nolint:staticcheck
	if _, err = fmt.Fprintf(h, result.Source); err != nil {
		ctx.ErrorJSON(400, fmt.Errorf("invalid `source`: hashing"), nil)
		return
	}
	hash := fmt.Sprintf("%x", h.Sum(nil))

	var id uint64
	err = v.db.ExecContext("insert_new_shorten", ctx.Context(), func(q orm.Executor) {
		q.SQL("INSERT INTO `shorten` (`data`, `hash`, `lock`, `created_at`) VALUES (?, ?, 0, now());")
		q.Params(result.Source, hash)
		q.Bind(func(rowsAffected, lastInsertId int64) error {
			id = uint64(lastInsertId)
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
			ctx.ErrorJSON(400, fmt.Errorf("invalid `source`: cant save"), nil)
			return
		}
	}

	code := v.codec.Encode(id)
	result.URL = fmt.Sprintf("%s/%s", string(v.address), code)
	ctx.JSON(200, result)
}
