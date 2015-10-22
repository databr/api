package service

import (
	"strconv"

	"github.com/databr/api/database"
	"github.com/databr/api/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

type PartyResource struct {
	DB database.MongoDB
}

func (r *PartyResource) Index(c *gin.Context) {
	var p []*models.Party
	search := bson.M{}
	query := c.Request.URL.Query()

	pageS := query.Get("page")
	if pageS == "" {
		pageS = "1"
	}
	page, _ := strconv.Atoi(pageS)

	err := r.DB.Find(search, GetLimit(c.Request), page, &p)

	if err != nil {
		c.JSON(500, gin.H{"error": "500", "message": err.Error()})
	} else {
		if len(p) == 0 {
			c.JSON(200, gin.H{"parties": []string{}})
			return
		}

		data := gin.H{
			"parties": p,
			"paging": pagination(
				"v1/parties",
				r.DB,
				GetLimit(c.Request),
				page,
				models.Party{},
				search,
			),
		}

		render := &DataRender{c.Request, data}
		c.Render(200, render)
	}
}

func (r *PartyResource) Get(c *gin.Context) {
	uri := c.Params.ByName("uri")

	var p models.Party

	err := r.DB.FindOne(bson.M{"id": uri}, &p)

	if err != nil {
		c.JSON(404, gin.H{"error": "404", "message": err.Error()})
	} else {
		render := &DataRender{c.Request, gin.H{"party": p}}
		c.Render(200, render)
	}
}
