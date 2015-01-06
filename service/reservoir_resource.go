package service

import (
	"github.com/databr/api/database"
	"github.com/databr/api/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

type ReservoirSpResource struct {
	DB database.MongoDB
}

func (r *ReservoirSpResource) Index(c *gin.Context) {
	var rv []struct {
		Id         string             `bson:"_id" json:"granularity"`
		Reservoirs []models.Reservoir `json:"reservoirs"`
	}

	search := bson.M{}
	err := r.DB.FindAndGroupBy("granularity", search, &rv, models.Reservoir{})

	if err != nil {
		c.JSON(500, gin.H{"error": "500", "message": err.Error()})
	} else {
		data := make(map[string][]models.Reservoir, 0)

		for _, item := range rv {
			data[item.Id] = item.Reservoirs
		}

		c.Render(200, DataRender{c.Request}, data)
	}
}
