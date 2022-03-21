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

	r.GET("/", func(c *gin.Context) {
		route, routes := entity.GetRouteIdentifyMapToRoute(entity.RouteIdentifyIndex)
		c.HTML(http.StatusOK, "index.tpl", gin.H{
			"Route":  route,
			"Routes": routes,
		})
	})

	r.GET("/index.html", func(c *gin.Context) {
		c.Redirect(http.StatusPermanentRedirect, "/")
	})

	r.GET("/timer.html", func(c *gin.Context) {
		route, routes := entity.GetRouteIdentifyMapToRoute(entity.RouteIdentifyTimer)
		c.HTML(http.StatusOK, "timer.tpl", gin.H{
			"Route":  route,
			"Routes": routes,
		})
	})

	r.GET("/play.html", func(c *gin.Context) {
		route, routes := entity.GetRouteIdentifyMapToRoute(entity.RouteIdentifyPlay)
		c.HTML(http.StatusOK, "play.tpl", gin.H{
			"Route":  route,
			"Routes": routes,
		})
	})

	if err := r.Run(); err != nil {
		log.Panicln(err)
	}
}
