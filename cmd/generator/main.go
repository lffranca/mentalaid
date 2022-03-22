package main

import (
	"github.com/lffranca/mentalaid/internal/entity"
	"github.com/lffranca/mentalaid/internal/service"
	"github.com/lffranca/mentalaid/internal/util"
	"html/template"
	"log"
	"os"
)

func main() {
	files, err := util.ReadFile(os.Getenv("TEMPLATE_PATH"))
	if err != nil {
		log.Panicln(err)
	}

	templ := template.Must(
		template.
			New("").
			Funcs(util.GetFuncMaps()).
			ParseFiles(files...))

	pageGenerator := service.NewPageGenerator(os.Getenv("PUBLIC_PATH"), templ)

	for key, item := range entity.GetRouteIdentifyMap() {
		if err := pageGenerator.Generate(key, item); err != nil {
			log.Println(err)
		}
	}
}

func init() {
	for _, envVar := range []string{
		"TEMPLATE_PATH",
		"PUBLIC_PATH",
	} {
		if _, ok := os.LookupEnv(envVar); !ok {
			log.Panicf("Required enviroment variable not set: %s\n", envVar)
		}
	}
}
