package routes

import (
	"github.com/gin-gonic/gin"
)

// eu coloca a função HandleRequest no controller ou deixo aqui?

func HandleRequest(r *gin.Engine) {
	addRoute := func(path string, template string, title string) {
		r.GET(path, func(c *gin.Context) {
			c.HTML(200, template, gin.H{
				"title": title,
			})
		})
	}

	addRoute("/", "tela_1.html", "Página Inicial")
	addRoute("/comunidades", "tela_2.html", "Comunidades")
	addRoute("/busca", "tela_3.html", "Busca")
	addRoute("/favoritos", "tela_4.html", "Favoritos")
	addRoute("/notificacoes", "tela_5.html", "Notificações")
	addRoute("/chat", "tela_6.html", "Chat")
	addRoute("/encaminhar", "tela_7.html", "Encaminhar Demanda")
	addRoute("/futuro", "tela_8.html", "Funcionalidades Futuras")
}
