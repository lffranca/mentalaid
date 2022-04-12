package util

import (
	"html/template"
	"time"
)

func GetFuncMaps() template.FuncMap {
	return template.FuncMap{
		"inc": func(i int) int {
			return i + 1
		},
		"now": func() int {
			return time.Now().Nanosecond()
		},
	}
}
