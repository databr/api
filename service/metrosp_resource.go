package service

import (
	"github.com/databr/api/database"
	"github.com/databr/api/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

type MetroSPResource struct {
	DB database.MongoDB
}

func (self MetroSPResource) Lines(c *gin.Context) {
	var l []*models.Line

	err := self.DB.Find(bson.M{}, PER_PAGE_LIMIT, 1, &l)

	if err != nil {
		c.JSON(500, gin.H{"error": "500", "message": err.Error()})
	} else {
		data := gin.H{
			"lines": l,
			"paging": pagination(
				"v1/sp/metro/lines",
				self.DB,
				PER_PAGE_LIMIT,
				1,
				models.Line{},
				bson.M{},
			),
		}

		c.Render(200, DataRender{c.Request}, data)
	}
}

func (self MetroSPResource) GetLines(c *gin.Context) {
}
func (self MetroSPResource) GetLinesStatus(c *gin.Context) {
}
