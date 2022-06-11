package api

import "github.com/deweppro/go-http/servers"

type (
	//MiddlewareConfig model
	MiddlewareConfig struct {
		Middleware ConfigItem `yaml:"middleware" json:"middleware"`
	}

	//ConfigItem model
	ConfigItem struct {
		Throttling int64 `yaml:"throttling" json:"throttling"`
	}

	//WebConfig model
	WebConfig struct {
		Http  servers.Config `yaml:"http"`
		Debug servers.Config `yaml:"debug"`
	}
)
