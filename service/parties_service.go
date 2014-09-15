package service

import (
	"github.com/databr/api/models"
	"github.com/gin-gonic/gin"
)

type PartiesService struct {
	*gin.Engine
}

func (cs *PartiesService) Run() {
	databaseDB := models.New()

	partyResource := &PartyResource{DB: databaseDB}

	v1 := cs.Group("/v1")
	{
		v1.GET("/parties", partyResource.Index)

		v1.GET("/parties/:uri", partyResource.Get)
	}
}
