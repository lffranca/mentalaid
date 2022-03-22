package service

import (
	"bytes"
	"github.com/lffranca/mentalaid/internal/entity"
	"html/template"
	"io/ioutil"
	"path/filepath"

	"github.com/yosssi/gohtml"
)

type PageGenerator interface {
	Generate(key entity.RouteIdentify, item entity.Route) error
}

func NewPageGenerator(
	publicPath string,
	templ *template.Template,
) PageGenerator {
	client := new(pageGenerator)
	client.publicPath = publicPath
	client.templ = templ

	return client
}

type pageGenerator struct {
	publicPath string
	templ      *template.Template
}

func (pkg *pageGenerator) Generate(key entity.RouteIdentify, item entity.Route) error {
	route, routes := entity.GetRouteIdentifyMapToRoute(key)
	data := map[string]interface{}{
		"Route":  route,
		"Routes": routes,
	}

	var result bytes.Buffer
	if err := pkg.templ.ExecuteTemplate(&result, item.Template, data); err != nil {
		return err
	}

	resultFormatted := gohtml.FormatBytes(result.Bytes())

	path := filepath.Join(pkg.publicPath, item.Path)
	if err := ioutil.WriteFile(path, resultFormatted, 0644); err != nil {
		return err
	}

	return nil
}
