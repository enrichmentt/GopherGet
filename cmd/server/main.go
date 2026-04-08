package main

import (
	"github.com/enrichmentt/GopherGet/internal/api"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/information", api.Information)

	r.RunTLS(":8443", "cert.pem", "key.pem")
}
