package service

import (
	"log"
	"strconv"

	"github.com/databr/api/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

type PartyResource struct {
	DB models.Database
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

	err := r.DB.Find(search, PER_PAGE_LIMIT, page, &p)

	if err != nil {
		log.Println(err)
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
				PER_PAGE_LIMIT,
				page,
				models.Party{},
				search,
			),
		}

		c.Render(200, DataRender{c.Request}, data)
	}
}

func (r *PartyResource) Get(c *gin.Context) {
	uri := c.Params.ByName("uri")

	var p models.Party

	err := r.DB.FindOne(bson.M{"id": uri}, &p)

	if err != nil {
		c.JSON(404, gin.H{"error": "404", "message": err.Error()})
	} else {
		c.Render(200, DataRender{c.Request}, gin.H{"party": p})
	}
}