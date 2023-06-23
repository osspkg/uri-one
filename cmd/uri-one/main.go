/*
 *  Copyright (c) 2020-2023 Mikhail Knyazhev <markus621@yandex.ru>. All rights reserved.
 *  Use of this source code is governed by a GPL-3.0 license that can be found in the LICENSE file.
 */

package main

import (
	"github.com/osspkg/goppy"
	"github.com/osspkg/goppy/plugins/database"
	"github.com/osspkg/goppy/plugins/web"
	"github.com/osspkg/uri-one/app/badges"
	"github.com/osspkg/uri-one/app/common"
	"github.com/osspkg/uri-one/app/shorten"
)

func main() {
	app := goppy.New()
	app.WithConfig("./config.yaml") // Reassigned via the `--config` argument when run via the console.
	app.Plugins(
		web.WithHTTP(),
		database.WithMySQL(),
	)
	app.Plugins(
		badges.Plugin,
		shorten.Plugin,
		common.Plugin,
	)
	app.Run()
}
