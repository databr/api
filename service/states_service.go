package service

import (
	"github.com/databr/api/database"
	"github.com/gin-gonic/gin"
)

type StatesService struct {
	*gin.Engine
}

func (s StatesService) Run(databaseDB database.MongoDB) {
	statesResource := &StatesResource{DB: databaseDB}
	trainsSpResource := &TrainsSpResource{DB: databaseDB}

	v1 := s.Group("/v1")
	{

		v1.GET("/states", statesResource.Index)
		v1.GET("/states/:uri", statesResource.Get)
		v1.GET("/states/:uri/cities", statesResource.Cities)
		v1.GET("/states/:uri/cities/:city_uri", statesResource.GetCity)

		v1.GET("/states/:uri/transports/trains/lines", onlySP(), trainsSpResource.Lines)
		v1.GET("/states/:uri/transports/trains/lines/:line_uri", onlySP(), trainsSpResource.GetLine)
		v1.GET("/states/:uri/transports/trains/lines/:line_uri/statuses", onlySP(), trainsSpResource.GetLineStatuses)
		v1.GET("/states/:uri/transports/trains/lines/:line_uri/statuses/:status_id", onlySP(), trainsSpResource.GetLineStatus)
	}
}

func onlySP() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Params.ByName("uri") == "sp" {
			c.Next()
		} else {
			c.Abort(404)
		}
	}
}
