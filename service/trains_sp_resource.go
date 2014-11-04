package service

import (
	"github.com/databr/api/config"
	"github.com/databr/api/database"
	"github.com/databr/api/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

type TrainsSpResource struct {
	DB database.MongoDB
}

func (self TrainsSpResource) Lines(c *gin.Context) {
	var l []*models.Line

	err := self.DB.Find(bson.M{}, GetLimit(c.Request), 1, &l)

	if err != nil {
		c.JSON(500, gin.H{"error": "500", "message": err.Error()})
	} else {
		done := make(chan bool)
		go func() {
			self.setLinks(l)
			self.setStatus(l)
			done <- true
		}()
		<-done

		data := gin.H{
			"lines": l,
			"paging": pagination(
				"v1/sp/metro/lines",
				self.DB,
				GetLimit(c.Request),
				1,
				models.Line{},
				bson.M{},
			),
		}

		c.Render(200, DataRender{c.Request}, data)
	}
}

func (self TrainsSpResource) GetLine(c *gin.Context) {
	uri := c.Params.ByName("line_uri")
	var l models.Line

	err := self.DB.FindOne(bson.M{"id": uri}, &l)

	if err != nil {
		c.JSON(404, gin.H{"error": "404", "message": err.Error()})
	} else {

		done := make(chan bool)
		go func() {
			l.Links = self.getLink(l.Id, l.CannonicalUri)
			l.Status = self.getStatus(l.CannonicalUri)
			done <- true
		}()
		<-done

		c.Render(200, DataRender{c.Request}, gin.H{"line": l})
	}
}

func (self TrainsSpResource) GetLineStatuses(c *gin.Context) {
	uri := c.Params.ByName("line_uri")
	var s []*models.Status

	err := self.DB.Find(bson.M{"line_id": uri}, GetLimit(c.Request), 1, &s)

	if err != nil {
		c.JSON(404, gin.H{"error": "404", "message": err.Error()})
	} else {
		done := make(chan bool)
		go func() {
			self.setStatusLinks(s)
			done <- true
		}()
		<-done

		c.Render(200, DataRender{c.Request}, gin.H{"statuses": s})
	}
}

func (self TrainsSpResource) GetLineStatus(c *gin.Context) {
	uri := c.Params.ByName("line_uri")
	statusId := bson.ObjectIdHex(c.Params.ByName("status_id"))
	var status models.Status

	err := self.DB.FindOne(bson.M{"line_id": uri, "_id": statusId}, &status)
	if err != nil {
		c.JSON(404, gin.H{"error": "404", "message": err.Error()})
	} else {

		done := make(chan bool)
		go func() {
			status.Links = self.getStatusLink(uri, status.Id)
			done <- true
		}()
		<-done

		c.Render(200, DataRender{c.Request}, gin.H{"status": status})
	}
}

func (self TrainsSpResource) setLinks(l []*models.Line) {
	for i, _ := range l {
		l[i].Links = self.getLink(l[i].Id, l[i].CannonicalUri)
	}
}

func (_ TrainsSpResource) getLink(id, cannonical string) []models.Link {
	return []models.Link{
		{
			Url:  config.ApiRoot + "/v1/states/sp/transports/trains/lines/" + id,
			Note: "self",
		},
		{
			Url:  config.ApiRoot + "v1/states/sp/transports/trains/lines/" + cannonical + "/statuses",
			Note: "statuses",
		},
	}
}

func (self TrainsSpResource) setStatus(l []*models.Line) {
	for i, _ := range l {
		l[i].Status = self.getStatus(l[i].CannonicalUri)
	}
}

func (self TrainsSpResource) getStatus(id string) (status models.Status) {
	err := self.DB.FindOne(bson.M{"line_id": id}, &status)
	if err != nil {
		status.Links = self.getStatusLink(id, status.Id)
	}
	return status
}

func (self TrainsSpResource) setStatusLinks(statuses []*models.Status) {
	for i, _ := range statuses {
		statuses[i].Links = self.getStatusLink(statuses[i].LineId, statuses[i].Id)
	}
}

func (self TrainsSpResource) getStatusLink(lineId string, statusId bson.ObjectId) []models.Link {
	id := statusId.Hex()
	return []models.Link{
		{
			Url:  config.ApiRoot + "/v1/states/sp/transports/trains/lines/" + lineId + "/statuses/" + id,
			Note: "self",
		},
		{
			Url:  config.ApiRoot + "v1/states/sp/transports/trains/lines/" + lineId + "/statuses",
			Note: "statuses",
		},
		{
			Url:  config.ApiRoot + "v1/states/sp/transports/trains/lines/" + lineId,
			Note: "line",
		},
	}
}
