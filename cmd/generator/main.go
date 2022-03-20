package main

import (
	"github.com/gin-gonic/gin"
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

	r.LoadHTMLFiles(files...)

	r.Static("/assets", os.Getenv("ASSETS_PATH"))

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tpl", gin.H{})
	})

	if err := r.Run(); err != nil {
		log.Panicln(err)
	}
}
