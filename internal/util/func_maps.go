package util

import "html/template"

func GetFuncMaps() template.FuncMap {
	return template.FuncMap{
		"inc": func(i int) int {
			return i + 1
		},
	}
}
