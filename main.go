package main

import (
	"api_rest/api"
	"api_rest/config"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar .env")
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	r := gin.Default()

	r.POST("/decode", api.AuthMiddleware(cfg.SecretKey), api.DecodeJWTHandler)

	r.Run(cfg.ApiPort)
}
