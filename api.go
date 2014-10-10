package main

import (
	"log"
	"net/http"

	"github.com/databr/api/config"
	"github.com/databr/api/database"
	"github.com/databr/api/middleware"
	"github.com/databr/api/service"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.CORS())
	// r.Use(middleware.Authentication())
	r.Use(middleware.Analytics())
	r.Use(middleware.StatusPageIO())

	if config.Env == "production" {
		r.Use(middleware.NewRelic())
		middleware.InitNewrelicAgent(config.NewRelicLicense, config.NewRelicAppName, true)
	}

	databaseDB := database.NewMongoDB()

	parliamentarians := service.ParliamentariansService{r}
	parliamentarians.Run(databaseDB)

	parties := service.PartiesService{r}
	parties.Run(databaseDB)

	states := service.StatesService{r}
	states.Run(databaseDB)

	metrosp := service.MetroSPService{r}
	metrosp.Run(databaseDB)

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
