package main

import (
	"log"
	"net/http"
	"os"

	"github.com/databr/api/middleware"
	"github.com/databr/api/service"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.CORS())
	r.Use(middleware.Analytics())

	parliamentariansService := service.ParliamentariansService{r}
	parliamentariansService.Run()

	partiesService := service.PartiesService{r}
	partiesService.Run()

	r.GET("/", func(c *gin.Context) {
		http.Redirect(
			c.Writer,
			c.Request,
			"http://databr.io",
			302)
	})

	log.Println("Listening port", os.Getenv("PORT"))
	r.Run(":" + os.Getenv("PORT"))
}
