package main

import (
	"flag"

	"uri-one/app"
)

var (
	cpath = flag.String("config", "source/config.yaml", "config file path")
)

func main() {
	flag.Parse()

	app.New(*cpath).Run()
}
