package handlers

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type Usuario struct {
	Nome string `json:"nome"`
	Email string `json:"email"`
	Senha string `json:"senha"`
	Cidade string `json:"cidade"`
	LinhasFav []string `json:"linhas_favoritas"`
}

var (
	usuarios []Usuario 		// slice para amazenar usuários
	usuariosMu sync.Mutex 	// mutex para acesso concorrente
)

func CadastrarUsuario(c *gin.Context) {
	var novo Usuario
	if err := c.ShouldBindJSON(&novo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Dados inválidos"})
		return
	}

	usuariosMu.Lock()
	usuarios = append(usuarios, novo)
	usuariosMu.Unlock()

	c.JSON(http.StatusOK, gin.H{"mensagem": "Usuário cadastrado com sucesso!"})
}