package service

import (
	"github.com/camarabook/camarabook-api/models"
	"github.com/gin-gonic/gin"
)

type CamarabookService struct {
}

func (cs *CamarabookService) Run() *gin.Engine {
	databaseDB := models.New()

	parliamentarianResource := &ParliamentarianResource{DB: databaseDB}

	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(corsMiddleware())

	v1 := r.Group("/v1")
	{
		v1.GET("/parliamentarians", func(c *gin.Context) {
			c.String(200, "Hi!")
		})

		v1.GET("/parliamentarians/:uri", parliamentarianResource.Get)
		v1.GET("/parliamentarians/:uri/activities", parliamentarianResource.GetActivities)
	}

	return r
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	}
}
