package config

import ()

type ServiceConfigItems struct {
	ServiceConfigItems []ServiceConfigItem `json:"services"`
}

type ServiceConfigItem struct {
	Name string `json:"name"`
	Path string `json:"path"`
	IsPlugin bool `json:"isplugin"`
}
