package middleware

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	influxdb "github.com/influxdb/influxdb/client"
)

var influxdbC *influxdb.Client

func init() {

	config := influxdb.ClientConfig{
		Host:     os.Getenv("INFLUXDB_HOST"),
		Username: os.Getenv("INFLUXDB_USERNAME"),
		Password: os.Getenv("INFLUXDB_PASSWORD"),
		Database: os.Getenv("INFLUXDB_DATABASE"),
		IsSecure: false,
	}

	var err error
	influxdbC, err = influxdb.NewClient(&config)
	checkErr(err)

	err = influxdbC.Ping()
	checkErr(err)
}

func Analytics() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Next()
		c_c := c.Copy()
		latency := time.Since(t)
		go func() {
			status := c_c.Writer.Status()

			s := []*influxdb.Series{{
				Name: "api_access",
				Columns: []string{
					"status", "latency", "value", "query",
				},
				Points: [][]interface{}{
					{status, latency, c_c.Request.URL.Path, c_c.Request.URL.RawQuery},
				},
			}}

			time.Sleep(time.Second * 10)
			err := influxdbC.WriteSeries(s)
			checkErr(err)
		}()
	}
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
