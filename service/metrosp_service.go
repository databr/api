package service

import (
	"github.com/databr/api/database"
	"github.com/gin-gonic/gin"
)

type MetroSPService struct {
	*gin.Engine
}

func (cs MetroSPService) Run() {
	databaseDB := database.NewMongoDB()

	metrospResource := &MetroSPResource{DB: databaseDB}

	v1 := cs.Group("/v1")
	{
		v1.GET("/sp/metro/lines", metrospResource.Lines)
		v1.GET("/sp/metro/lines/:uri", metrospResource.GetLine)
		v1.GET("/sp/metro/lines/:uri/status", metrospResource.GetLineStatus)
	}
}
