package middleware

import (
	"bytes"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/databr/api/config"
	"github.com/gin-gonic/gin"
)

func StatusPageIO() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Next()
		latency := time.Since(t)

		go func() {
			if config.StatusPageIOEnable != "true" {
				return
			}
			requestURL := "" + config.StatusPageIOApiBase + "/pages/" + config.StatusPageIOPageID + "/metrics/" + config.StatusPageIOMetricID + "/data.json"

			timestamp := strconv.Itoa(int(time.Now().Unix()))
			value := int(latency / time.Millisecond)

			data := url.Values{}
			data.Set("data[timestamp]", timestamp)
			data.Set("data[value]", strconv.Itoa(value))

			r, err := http.NewRequest("POST", requestURL, bytes.NewBufferString(data.Encode()))
			if err != nil {
				return
			}
			r.Header.Add("Authorization", "OAuth "+config.StatusPageIOApiKey)
			r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

			resp, err := http.DefaultClient.Do(r)
			log.Println(resp, err)
		}()
	}
}
