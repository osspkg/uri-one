/*
 *  Copyright (c) 2020-2024 Mikhail Knyazhev <markus621@yandex.ru>. All rights reserved.
 *  Use of this source code is governed by a GPL-3.0 license that can be found in the LICENSE file.
 */

package shorten

import (
	"math/rand"
	"time"
)

type Config struct {
	Shorten string `yaml:"shorten_string"`
}

func (v *Config) Default() {
	data := []byte("abcdefghijklmnopqrstuvwxyz0123456789")
	rnd := rand.New(rand.NewSource(time.Now().Unix()))
	rnd.Shuffle(len(data), func(i, j int) {
		data[i], data[j] = data[j], data[i]
	})
	v.Shorten = string(data)
}
