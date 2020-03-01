package config

import (
	"github.com/deweppro/core/filedecoder"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Env       string    `yaml:"env"`
	LogFile   string    `yaml:"log"`
	Http      Http      `yaml:"http"`
	Providers Providers `yaml:"providers"`
}

type Providers struct {
	DB SQlite `yaml:"db"`
}

type SQlite struct {
	SQL  string `yaml:"init"`
	Path string `yaml:"path"`
}
type Http struct {
	Port   string `yaml:"port"`
	Prefix string `yaml:"prefix"`
}

func MustNew(path string) (c *Config) {
	if err := filedecoder.Yaml(path, &c); err != nil {
		logrus.Fatal(err)
	}
	return
}

func (c *Config) IsDebug() bool {
	return c.Env == "dev"
}
