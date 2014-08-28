package service

import (
	"github.com/camarabook/camarabook-api/models"
	"github.com/gin-gonic/gin"
)

type ParliamentariansService struct {
	*gin.Engine
}

func (cs *ParliamentariansService) Run() {
	databaseDB := models.New()

	parliamentarianResource := &ParliamentarianResource{DB: databaseDB}

	v1 := cs.Group("/v1")
	{
		v1.GET("/parliamentarians", func(c *gin.Context) {
			c.String(200, "Hi!")
		})

		v1.GET("/parliamentarians/:uri", parliamentarianResource.Get)
	}
}
