package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasjdc/comuBus/routes"
)

func main() {
	r := gin.Default()

	// Define a pasta onde estão os arquivos HTML
	r.LoadHTMLGlob("templates/*.html")

	r.Static("/static", "./static")

	routes.HandleRequest(r)


	r.Run()
}
