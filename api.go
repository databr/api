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
	r.Use(middleware.StatusPageIO())

	if config.Env == "production" {
		r.Use(middleware.NewRelic())
		middleware.InitNewrelicAgent(config.NewRelicLicense, config.NewRelicAppName, true)
	}

	databaseDB := database.NewMongoDB()

	parliamentarians := service.ParliamentariansService{r}
	parties := service.PartiesService{r}
	states := service.StatesService{r}
	pingdom := service.PingdomService{r}
	doc := service.ApiDocumentationService{r}

	parliamentarians.Run(databaseDB)
	parties.Run(databaseDB)
	states.Run(databaseDB)
	pingdom.Run()
	doc.Run()

	r.OPTIONS("/*path", func(c *gin.Context) {
		c.Writer.Header().Set("Allow", "GET, OPTIONS")
		c.Abort(200)
	})

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
