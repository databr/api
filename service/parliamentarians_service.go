package service

import (
	"github.com/databr/api/database"
	"github.com/gin-gonic/gin"
)

type ParliamentariansService struct {
	*gin.Engine
}

func (cs *ParliamentariansService) Run() {
	databaseDB := database.NewMongoDB()

	parliamentarianResource := &ParliamentarianResource{DB: databaseDB}

	v1 := cs.Group("/v1")
	{
		v1.GET("/parliamentarians", parliamentarianResource.Index)

		v1.GET("/parliamentarians/:uri", parliamentarianResource.Get)
	}
}
