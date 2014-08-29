package service

import (
	"os"

	"github.com/camarabook/camarabook-api/models"
	"github.com/camarabook/go-popolo"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

var API_ROOT string

func init() {
	API_ROOT = os.Getenv("API_ROOT")
}

type ParliamentarianResource struct {
	DB models.Database
}

func (r *ParliamentarianResource) Index(c *gin.Context) {
	var p []*models.Parliamentarian
	search := bson.M{}
	query := c.Request.URL.Query()
	identifier := query.Get("identifier")

	if identifier != "" {
		search["identifiers"] = bson.M{
			"$elemMatch": bson.M{
				"identifier": identifier,
			},
		}
	}

	err := r.DB.Find(search, &p)

	if err != nil {
		c.JSON(500, gin.H{"error": "500", "message": err.Error()})
	} else {
		if len(p) == 0 {
			p = make([]*models.Parliamentarian, 0)
		}
		setLinks(p)
		c.JSON(200, gin.H{"parliamentarians": p})
	}
}

func (r *ParliamentarianResource) Get(c *gin.Context) {
	uri := c.Params.ByName("uri")

	var p models.Parliamentarian

	err := r.DB.FindOne(bson.M{"id": uri}, &p)

	if err != nil {
		c.JSON(404, gin.H{"error": "404", "message": err.Error()})
	} else {
		c.JSON(200, gin.H{"parliamentarian": p})
	}
}

func setLinks(p []*models.Parliamentarian) {
	for i, _ := range p {
		p[i].Links = getLink(p[i])
	}
}

func getLink(p *models.Parliamentarian) []popolo.Link {
	return []popolo.Link{
		{
			Url:  toPtr(API_ROOT + "/v1/parliamentarians/" + *p.Id),
			Note: toPtr("self"),
		},
	}
}

func toPtr(s string) *string {
	return &s
}
