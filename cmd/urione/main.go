/*
 *  Copyright (c) 2020-2024 Mikhail Knyazhev <markus621@yandex.ru>. All rights reserved.
 *  Use of this source code is governed by a GPL-3.0 license that can be found in the LICENSE file.
 */

package main

import (
	"github.com/osspkg/uri-one/app/badges"
	"github.com/osspkg/uri-one/app/common"
	"github.com/osspkg/uri-one/app/mainapp"
	"github.com/osspkg/uri-one/app/shorten"
	"go.osspkg.com/goppy"
	"go.osspkg.com/goppy/ormmysql"
	"go.osspkg.com/goppy/web"
)

func main() {
	app := goppy.New()
	app.Plugins(
		web.WithHTTP(),
		ormmysql.WithMySQL(),
	)
	app.Plugins(
		mainapp.Plugin,
		badges.Plugin,
		shorten.Plugin,
		common.Plugin,
	)
	app.Run()
}
