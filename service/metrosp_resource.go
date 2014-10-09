package service

import (
	"github.com/databr/api/config"
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
		done := make(chan bool)
		go func() {
			self.setLinks(l)
			done <- true
		}()
		<-done

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

func (self MetroSPResource) GetLine(c *gin.Context) {
	uri := c.Params.ByName("uri")
	var l models.Line

	err := self.DB.FindOne(bson.M{"id": uri}, &l)

	if err != nil {
		c.JSON(404, gin.H{"error": "404", "message": err.Error()})
	} else {

		done := make(chan bool)
		go func() {
			l.Links = self.getLink(l.Id)
			done <- true
		}()
		<-done

		c.Render(200, DataRender{c.Request}, gin.H{"line": l})
	}
}

func (self MetroSPResource) GetLineStatus(c *gin.Context) {
	uri := c.Params.ByName("uri")
	var s models.Status

	err := self.DB.FindOne(bson.M{"line_id": uri}, &s)

	if err != nil {
		c.JSON(404, gin.H{"error": "404", "message": err.Error()})
	} else {
		c.Render(200, DataRender{c.Request}, gin.H{"status": s})
	}
}

func (self MetroSPResource) setLinks(p []*models.Line) {
	for i, _ := range p {
		p[i].Links = self.getLink(p[i].Id)
	}
}

func (_ MetroSPResource) getLink(id string) []models.Link {
	return []models.Link{
		{
			Url:  config.ApiRoot + "/v1/sp/metro/lines/" + id,
			Note: "self",
		},
		{
			Url:  config.ApiRoot + "/v1/sp/metro/lines/" + id + "/status",
			Note: "status",
		},
	}
}
