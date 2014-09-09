package main

import (
	"log"
	"os"

	"github.com/databr/api/service"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(fixCorsMiddleware())

	scamarabook := service.ParliamentariansService{r}
	scamarabook.Run()

	log.Println("Listening port", os.Getenv("PORT"))
	r.Run(":" + os.Getenv("PORT"))
}

func fixCorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	}
}
