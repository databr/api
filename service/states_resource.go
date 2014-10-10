package service

import (
	"github.com/databr/api/database"
	"github.com/databr/api/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

type StatesResource struct {
	DB database.MongoDB
}

func (r *StatesResource) Index(c *gin.Context) {
	var s []*models.State
	search := bson.M{}

	err := r.DB.Find(search, PER_PAGE_LIMIT, 1, &s)

	if err != nil {
		c.JSON(500, gin.H{"error": "500", "message": err.Error()})
	} else {
		if len(s) == 0 {
			c.JSON(200, gin.H{"states": []string{}})
			return
		}

		data := gin.H{
			"states": s,
		}

		c.Render(200, DataRender{c.Request}, data)
	}
}

func (r *StatesResource) Get(c *gin.Context) {
	uri := c.Params.ByName("uri")

	var s models.State

	err := r.DB.FindOne(bson.M{"id": uri}, &s)

	if err != nil {
		c.JSON(404, gin.H{"error": "404", "message": err.Error()})
	} else {
		c.Render(200, DataRender{c.Request}, gin.H{"state": s})
	}
}
