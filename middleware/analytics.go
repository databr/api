package middleware

import (
	"time"

	"github.com/databr/api/config"
	"github.com/gin-gonic/gin"
	influxdb "github.com/influxdb/influxdb/client"
)

var influxdbC *influxdb.Client

func init() {
	if config.InfluxdbEnable != "true" {
		return
	}

	configInfluxdb := influxdb.ClientConfig{
		Host:     config.InfluxdbHost,
		Username: config.InfluxdbUser,
		Password: config.InfluxdbPassword,
		Database: config.InfluxdbDatabase,
		IsSecure: false,
	}

	var err error
	influxdbC, err = influxdb.NewClient(&configInfluxdb)
	checkErr(err)

	err = influxdbC.Ping()
	checkErr(err)
}

func Analytics() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Next()

		if config.InfluxdbEnable != "true" {
			return
		}
		if c.Request.URL.Path == "/status/pingdom" {
			return
		}

		cCopy := c.Copy()
		latency := time.Since(t)
		go func() {
			status := cCopy.Writer.Status()
			appId, _ := c.Get("app_id")

			s := []*influxdb.Series{{
				Name: "api_access",
				Columns: []string{
					"status", "latency", "value", "query", "app_id",
				},
				Points: [][]interface{}{
					{status, latency, cCopy.Request.URL.Path, cCopy.Request.URL.RawQuery, appId},
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
