package service

import (
	"strconv"

	"github.com/databr/api/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

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

	pageS := query.Get("page")
	if pageS == "" {
		pageS = "1"
	}
	page, _ := strconv.Atoi(pageS)

	err := r.DB.Find(search, PER_PAGE_LIMIT, page, &p)

	if err != nil {
		c.JSON(500, gin.H{"error": "500", "message": err.Error()})
	} else {
		if len(p) == 0 {
			c.JSON(200, gin.H{"parliamentarians": []string{}})
			return
		}
		setLinks(p)

		data := gin.H{
			"parliamentarians": p,
			"paging": pagination(
				"v1/parliamentarians",
				r.DB,
				PER_PAGE_LIMIT,
				page,
				models.Parliamentarian{},
				search,
			),
		}

		c.Render(200, DataRender{c.Request}, data)
	}
}

func (r *ParliamentarianResource) Get(c *gin.Context) {
	uri := c.Params.ByName("uri")

	var p models.Parliamentarian

	err := r.DB.FindOne(bson.M{"id": uri}, &p)

	if err != nil {
		c.JSON(404, gin.H{"error": "404", "message": err.Error()})
	} else {
		c.Render(200, DataRender{c.Request}, gin.H{"parliamentarian": p})
	}
}

func setLinks(p []*models.Parliamentarian) {
	for i, _ := range p {
		p[i].Links = getLink(p[i])
	}
}

func getLink(p *models.Parliamentarian) []models.Link {
	return []models.Link{
		{
			Url:  API_ROOT + "/v1/parliamentarians/" + p.Id,
			Note: "self",
		},
	}
}
