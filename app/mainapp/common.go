/*
 *  Copyright (c) 2020-2024 Mikhail Knyazhev <markus621@yandex.ru>. All rights reserved.
 *  Use of this source code is governed by a GPL-3.0 license that can be found in the LICENSE file.
 */

package mainapp

import "go.osspkg.com/goppy/plugins"

type Address string

var Plugin = plugins.Plugin{
	Config: &Config{},
	Inject: func(c *Config) Address {
		return Address(c.Address)
	},
}
