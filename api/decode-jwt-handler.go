package api

import (
	b64 "encoding/base64"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func DecodeJWTHandler(c *gin.Context) {
	var req struct {
		Token string `json:"token"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	parts := strings.Split(req.Token, ".")
	if len(parts) != 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token malformatado"})
	}

	payload64 := parts[1]

	payloadBytes, err := b64.RawURLEncoding.DecodeString(payload64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao decodificar o payload."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Decodificado": string(payloadBytes)})
}
