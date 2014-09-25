package main

import (
	"encoding/xml"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/databr/api/middleware"
	"github.com/databr/api/models"
	"github.com/databr/api/service"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.CORS())
	r.Use(middleware.Analytics())
	r.Use(middleware.StatusPageIO())

	parliamentariansService := service.ParliamentariansService{r}
	parliamentariansService.Run()

	partiesService := service.PartiesService{r}
	partiesService.Run()

	r.GET("/status/pingdom", func(c *gin.Context) {
		n := time.Now()
		status := "OK"

		err := models.Ping()
		if err != nil {
			status = err.Error()
		}

		type Pingdom struct {
			XMLName      xml.Name `xml:"pingdom_http_custom_check"`
			Status       string   `xml:"status"`
			ResponseTime string   `xml:"response_time"`
		}

		latency := int(time.Since(n) / time.Millisecond)

		pingdom := &Pingdom{
			Status:       status,
			ResponseTime: strconv.Itoa(latency),
		}

		c.XML(200, pingdom)
	})

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
