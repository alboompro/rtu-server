package server

import (
	"github.com/gin-gonic/gin"
)

// Server Criar o server da API
func Server() {
	router := gin.Default()

	// Redirecionamento para o site do projeto quando tentado
	// acessar direto na raiz
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong" })
	})

	router.Run(":80")
}

