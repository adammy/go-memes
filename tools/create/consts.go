package create

import (
	"strings"
	"text/template"
)

const (
	APIConfig  ConfigType = "api"
	MainConfig ConfigType = "main"

	APICommand    CommandType = "api"
	ServerCommand CommandType = "server"
)

var (
	FuncMap = template.FuncMap{
		"Title": strings.Title,
	}
)
