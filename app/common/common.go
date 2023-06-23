/*
 *  Copyright (c) 2020-2023 Mikhail Knyazhev <markus621@yandex.ru>. All rights reserved.
 *  Use of this source code is governed by a GPL-3.0 license that can be found in the LICENSE file.
 */

package common

import (
	"github.com/osspkg/go-sdk/app"
	"github.com/osspkg/goppy/plugins"
	"github.com/osspkg/goppy/plugins/web"
)

var Plugin = plugins.Plugin{
	Inject: New,
}

type Shorten struct {
	route web.Router
}

func New(r web.RouterPool) *Shorten {
	return &Shorten{
		route: r.Main(),
	}
}

func (v *Shorten) Up(ctx app.Context) error {
	v.route.NotFoundHandler(v.Page404)
	return nil
}

func (v *Shorten) Down() error {
	return nil
}

func (v *Shorten) Page404(ctx web.Context) {
	ctx.String(404, page404HTML)
}
