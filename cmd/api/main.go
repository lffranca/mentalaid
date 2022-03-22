package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lffranca/mentalaid/internal/entity"
	"github.com/lffranca/mentalaid/internal/util"
	"log"
	"net/http"
	"os"
)

func main() {
	r := gin.Default()

	files, err := util.ReadFile(os.Getenv("TEMPLATE_PATH"))
	if err != nil {
		log.Panicln(err)
	}

	r.SetFuncMap(util.GetFuncMaps())
	r.LoadHTMLFiles(files...)

	r.Static("/assets", os.Getenv("ASSETS_PATH"))

	for key, item := range entity.GetRouteIdentifyMap() {
		if item.Index {
			r.GET("/", func(c *gin.Context) {
				c.Redirect(http.StatusPermanentRedirect, item.Path)
			})
		}

		r.GET(item.Path, func(keyParent entity.RouteIdentify, itemParent entity.Route) gin.HandlerFunc {
			route, routes := entity.GetRouteIdentifyMapToRoute(keyParent)
			return func(c *gin.Context) {
				c.HTML(http.StatusOK, itemParent.Template, gin.H{
					"Route":  route,
					"Routes": routes,
				})
			}
		}(key, item))
	}

	if err := r.Run(); err != nil {
		log.Panicln(err)
	}
}
