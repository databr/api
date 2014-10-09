package main

import (
	"log"
	"net/http"

	"github.com/databr/api/config"
	"github.com/databr/api/middleware"
	"github.com/databr/api/service"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.NewRelic())
	r.Use(middleware.CORS())
	// r.Use(middleware.Authentication())
	r.Use(middleware.Analytics())
	r.Use(middleware.StatusPageIO())

	middleware.InitNewrelicAgent(config.NewRelicLicense, config.NewRelicAppName, true)

	parliamentarians := service.ParliamentariansService{r}
	parliamentarians.Run()

	partians := service.PartiesService{r}
	partians.Run()

	metrosp := service.MetroSPService{r}
	metrosp.Run()

	pingdom := service.PingdomService{r}
	pingdom.Run()

	r.GET("/", func(c *gin.Context) {
		http.Redirect(
			c.Writer,
			c.Request,
			"http://databr.io",
			302)
	})

	log.Println("Listening port", config.Port)
	r.Run(":" + config.Port)
}
