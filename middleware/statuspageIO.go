package middleware

import (
	"bytes"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	StatusPageIO_APIBase  = "https://api.statuspage.io/v1"
	StatusPageIO_APIKEY   string
	StatusPageIO_MetricID string
	StatusPageIO_PageID   string
)

func init() {
	StatusPageIO_APIKEY = os.Getenv("STATUSPAGEIO_API_KEY")
	StatusPageIO_MetricID = os.Getenv("STATUSPAGEIO_METRIC_ID")
	StatusPageIO_PageID = os.Getenv("STATUSPAGEIO_PAGE_ID")
}

func StatusPageIO() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Next()
		latency := time.Since(t)

		go func() {
			requestURL := "" + StatusPageIO_APIBase + "/pages/" + StatusPageIO_PageID + "/metrics/" + StatusPageIO_MetricID + "/data.json"

			timestamp := strconv.Itoa(int(time.Now().Unix()))
			value := int(latency / time.Millisecond)
			log.Println("V", strconv.Itoa(value))

			data := url.Values{}
			data.Set("data[timestamp]", timestamp)
			data.Set("data[value]", strconv.Itoa(value))

			log.Println(requestURL, data)
			r, err := http.NewRequest("POST", requestURL, bytes.NewBufferString(data.Encode()))
			if err != nil {
				return
			}
			r.Header.Add("Authorization", "OAuth "+StatusPageIO_APIKEY)
			r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

			resp, err := http.DefaultClient.Do(r)
			log.Println(resp, err)
		}()
	}
}
