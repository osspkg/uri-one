/*
 *  Copyright (c) 2020-2024 Mikhail Knyazhev <markus621@yandex.ru>. All rights reserved.
 *  Use of this source code is governed by a GPL-3.0 license that can be found in the LICENSE file.
 */

package mainapp

type Config struct {
	Address string `yaml:"address"`
}

func (v *Config) Default() {
	v.Address = "http://localhost:8080"
}
