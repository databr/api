package service

import (
	"encoding/xml"
	"strconv"
	"time"

	"github.com/databr/api/database"
	"github.com/gin-gonic/gin"
)

type PingdomService struct {
	*gin.Engine
}

func (r PingdomService) Run() {
	r.GET("/status/pingdom", func(c *gin.Context) {
		n := time.Now()
		status := "OK"

		db := database.NewMongoDB()
		err := db.Ping()

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

}
