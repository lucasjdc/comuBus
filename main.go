package main

import (
  "github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Define a pasta onde estão os arquivos HTML
	r.LoadHTMLGlob("templates/*.html")

	r.Static("/static", "./static")

	// Rota principal
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "tela_1.html", gin.H{
			"title":"Página Inicial",
		})
	})

	r.GET("/comunidades", func(c *gin.Context) {
		c.HTML(200, "tela_2.html", gin.H{
			"title":"Comunidades",
		})
	})

	r.GET("/busca", func(c *gin.Context) {
		c.HTML(200, "tela_3.html", gin.H{
			"title":"Busca",
		})
	})

	r.Run()
}
