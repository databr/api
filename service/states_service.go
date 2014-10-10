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

	v1 := s.Group("/v1")
	{
		v1.GET("/states", statesResource.Index)
		v1.GET("/states/:uri", statesResource.Get)
	}
}
