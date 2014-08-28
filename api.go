package main

import (
	"os"

	"github.com/camarabook/camarabook-api/service"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(fixCorsMiddleware())

	scamarabook := service.ParliamentariansService{r}
	scamarabook.Run()

	r.Run(":" + os.Getenv("PORT"))
}

func fixCorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	}
}
