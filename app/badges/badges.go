/*
 *  Copyright (c) 2020-2024 Mikhail Knyazhev <markus621@yandex.ru>. All rights reserved.
 *  Use of this source code is governed by a GPL-3.0 license that can be found in the LICENSE file.
 */

package badges

import (
	"html"

	"github.com/osspkg/uri-one/app/mainapp"
	"go.osspkg.com/badges"
	"go.osspkg.com/goppy/plugins"
	"go.osspkg.com/goppy/web"
	"go.osspkg.com/goppy/xlog"
)

var Plugin = plugins.Plugin{
	Inject: New,
}

type Badge struct {
	address mainapp.Address
	route   web.Router
	badge   *badges.Badges
}

func New(r web.RouterPool, d mainapp.Address) *Badge {
	return &Badge{
		address: d,
		route:   r.Main(),
	}
}

func (v *Badge) Up() (err error) {
	if v.badge, err = badges.New(); err != nil {
		return err
	}

	v.route.Get("/badge", v.Index)
	v.route.Get("/badge/{color}/{title}/{data}/image.svg", v.Draw)

	return nil
}

func (v *Badge) Down() error {
	return nil
}

func (v *Badge) Index(ctx web.Context) {
	ctx.String(200, indexHTML, string(v.address), string(v.address))
}

var colors = map[string]badges.Color{
	"primary":   badges.ColorPrimary,
	"secondary": badges.ColorSecondary,
	"success":   badges.ColorSuccess,
	"danger":    badges.ColorDanger,
	"warning":   badges.ColorWarning,
	"info":      badges.ColorInfo,
	"light":     badges.ColorLight,
}

func (v *Badge) Draw(ctx web.Context) {
	title, err := ctx.Param("title").String()
	if err != nil {
		ctx.String(400, "Invalid `title`")
		ctx.Log().WithFields(xlog.Fields{
			"err": err.Error(),
			"key": "title",
		}).Errorf("Invalid badge key")
		return
	}

	data, err := ctx.Param("data").String()
	if err != nil {
		ctx.String(400, "Invalid `data`")
		ctx.Log().WithFields(xlog.Fields{
			"err": err.Error(),
			"key": "data",
		}).Errorf("Invalid badge key")
		return
	}

	color, err := ctx.Param("color").String()
	if err != nil {
		ctx.String(400, "Invalid `color`")
		ctx.Log().WithFields(xlog.Fields{
			"err": err.Error(),
			"key": "color",
		}).Errorf("Invalid badge key")
		return
	}

	colored, ok := colors[color]
	if !ok {
		colored = badges.ColorPrimary
	}

	err = v.badge.WriteResponse(ctx.Response(), colored, html.EscapeString(title), html.EscapeString(data))
	if err != nil {
		ctx.Log().WithFields(xlog.Fields{
			"err": err.Error(),
		}).Errorf("Invalid badge response")
	}
}
