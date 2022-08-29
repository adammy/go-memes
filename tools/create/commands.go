package create

import (
	"strings"
)

func ToCommandType(str string) (CommandType, bool) {
	tmplTypeMap := map[string]CommandType{
		"api":    APICommand,
		"server": ServerCommand,
	}
	t, ok := tmplTypeMap[strings.ToLower(str)]
	return t, ok
}
