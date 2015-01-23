package service

import (
	"strconv"

	"github.com/databr/api/config"
	"github.com/databr/api/database"
	"github.com/databr/api/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

type ParliamentarianResource struct {
	DB database.MongoDB
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

	err := r.DB.Find(search, GetLimit(c.Request), page, "-updatedat", &p)

	if err != nil {
		c.JSON(500, gin.H{"error": "500", "message": err.Error()})
	} else {
		if len(p) == 0 {
			c.JSON(200, gin.H{"parliamentarians": []string{}})
			return
		}

		done := make(chan bool)

		go func() {
			r.setLinks(p)
			r.setMemberships(p)
			done <- true
		}()

		<-done
		data := gin.H{
			"parliamentarians": p,
			"paging": pagination(
				"v1/parliamentarians",
				r.DB,
				GetLimit(c.Request),
				page,
				models.Parliamentarian{},
				search,
			),
		}

		render := &DataRender{c.Request, data}
		c.Render(200, render)
	}
}

func (r *ParliamentarianResource) Get(c *gin.Context) {
	uri := c.Params.ByName("uri")

	var p models.Parliamentarian

	err := r.DB.FindOne(bson.M{"id": uri}, &p)

	done := make(chan bool)

	go func() {
		p.Links = r.getLink(p.Id)
		p.Memberships = r.getMemberships(p.Id)
		done <- true
	}()

	<-done

	if err != nil {
		c.JSON(404, gin.H{"error": "404", "message": err.Error()})
	} else {

		render := &DataRender{c.Request, gin.H{"parliamentarian": p}}
		c.Render(200, render)
	}
}

func (r ParliamentarianResource) setLinks(p []*models.Parliamentarian) {
	for i, _ := range p {
		p[i].Links = r.getLink(p[i].Id)
	}
}

func (r ParliamentarianResource) setMemberships(p []*models.Parliamentarian) {
	for i, _ := range p {
		p[i].Memberships = r.getMemberships(p[i].Id)
	}
}

func (_ ParliamentarianResource) getLink(id string) []models.Link {
	return []models.Link{
		{
			Url:  config.ApiRoot + "/v1/parliamentarians/" + id,
			Note: "self",
		},
	}
}

func (r ParliamentarianResource) getMemberships(id string) (memberships []models.Membership) {
	r.DB.Find(bson.M{"member.id": id}, 200, 1, "-updatedat", &memberships)
	return memberships
}
