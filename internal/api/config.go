package api

type (
	//MiddlewareConfig model
	MiddlewareConfig struct {
		Middleware ConfigItem `yaml:"middleware" json:"middleware"`
	}
	//ConfigItem model
	ConfigItem struct {
		Throttling int64 `yaml:"throttling" json:"throttling"`
	}
)
