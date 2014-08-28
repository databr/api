package service

import (
	"github.com/camarabook/camarabook-api/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

type ParliamentarianResource struct {
	DB models.Database
}

func (pr *ParliamentarianResource) Get(c *gin.Context) {
	uri := c.Params.ByName("uri")

	var p models.Parliamentarian

	err := pr.DB.FindOne(bson.M{"id": uri}, &p)

	if err != nil {
		c.JSON(404, gin.H{"error": "404", "message": err.Error()})
	} else {
		c.JSON(200, gin.H{"parliamentarian": p})
	}
}
