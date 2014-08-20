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
		c.JSON(404, gin.H{"error": "404", "status": err.Error()})
	} else {
		c.JSON(200, gin.H{"parliamentarian": p})
	}
}

func (pr *ParliamentarianResource) GetActivities(c *gin.Context) {
	uri := c.Params.ByName("uri")
	activities := make([]models.Activity, 0)

	var q []models.Quota
	pr.DB.Find(bson.M{"parliamentarian": uri}, &q)
	for _, item := range q {
		activities = append(activities, item.ToActivity(pr.DB))
	}

	c.JSON(200, gin.H{"activities": activities})
}
