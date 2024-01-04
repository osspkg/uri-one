/*
 *  Copyright (c) 2020-2024 Mikhail Knyazhev <markus621@yandex.ru>. All rights reserved.
 *  Use of this source code is governed by a GPL-3.0 license that can be found in the LICENSE file.
 */

package common

import (
	"go.osspkg.com/goppy/plugins"
	"go.osspkg.com/goppy/web"
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

func (v *Shorten) Up() error {
	v.route.NotFoundHandler(v.Page404)
	return nil
}

func (v *Shorten) Down() error {
	return nil
}

func (v *Shorten) Page404(ctx web.Context) {
	ctx.String(404, page404HTML)
}
