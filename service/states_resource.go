package service

import (
	"strconv"

	"github.com/databr/api/config"
	"github.com/databr/api/database"
	"github.com/databr/api/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

type StatesResource struct {
	DB database.MongoDB
}

func (r *StatesResource) Index(c *gin.Context) {
	var s []*models.State
	search := bson.M{}

	err := r.DB.Find(search, PER_PAGE_LIMIT, 1, &s)

	if err != nil {
		c.JSON(500, gin.H{"error": "500", "message": err.Error()})
	} else {
		if len(s) == 0 {
			c.JSON(200, gin.H{"states": []string{}})
			return
		}

		for i, _ := range s {
			s[i].Capital = r.getCapital(s[i].CapitalId)
			s[i].Links = r.getStateLinks(*s[i])
		}

		data := gin.H{
			"states": s,
		}

		c.Render(200, DataRender{c.Request}, data)
	}
}

func (r *StatesResource) Get(c *gin.Context) {
	uri := c.Params.ByName("uri")

	var s models.State

	err := r.DB.FindOne(bson.M{"id": uri}, &s)

	if err != nil {
		c.JSON(404, gin.H{"error": "404", "message": err.Error()})
	} else {
		s.Capital = r.getCapital(s.CapitalId)
		s.Links = r.getStateLinks(s)
		c.Render(200, DataRender{c.Request}, gin.H{"state": s})
	}
}

func (s *StatesResource) Cities(c *gin.Context) {
	var ci []*models.City
	stateUri := c.Params.ByName("uri")
	query := c.Request.URL.Query()

	search := bson.M{"stateid": stateUri}

	pageS := query.Get("page")
	if pageS == "" {
		pageS = "1"
	}
	page, _ := strconv.Atoi(pageS)

	err := s.DB.Find(search, PER_PAGE_LIMIT, page, &ci)

	if err != nil {
		c.JSON(500, gin.H{"error": "500", "message": err.Error()})
	} else {
		if len(ci) == 0 {
			c.JSON(200, gin.H{"cities": []string{}})
			return
		}

		done := make(chan bool)

		go func() {
			s.setCitiesLinks(ci)
			done <- true
		}()

		<-done
		data := gin.H{
			"cities": ci,
			"paging": pagination(
				"v1/states/"+stateUri+"/cities",
				s.DB,
				PER_PAGE_LIMIT,
				page,
				models.City{},
				search,
			),
		}

		c.Render(200, DataRender{c.Request}, data)
	}
}

func (s *StatesResource) GetCity(c *gin.Context) {
	var city models.City
	stateUri := c.Params.ByName("uri")
	cityUri := c.Params.ByName("city_uri")

	q := bson.M{"stateid": stateUri, "id": cityUri}
	err := s.DB.FindOne(q, &city)

	if err != nil {
		c.JSON(404, gin.H{"error": "404", "message": err.Error()})
	} else {
		done := make(chan bool)
		go func() {
			city.Links = s.setCityLinks(city)
			done <- true
		}()
		<-done

		c.Render(200, DataRender{c.Request}, gin.H{"city": city})
	}
}

func (s *StatesResource) getStateLinks(state models.State) []models.Link {
	return []models.Link{
		{
			Url:  config.ApiRoot + "/v1/states/" + state.Id + "/cities/",
			Note: "cities",
		}, {
			Url:  config.ApiRoot + "/v1/states/" + state.Id,
			Note: "self",
		}, {
			Url:  config.ApiRoot + "/v1/states/" + state.Id + "/cities/" + state.CapitalId,
			Note: "capital",
		},
	}

}

func (s *StatesResource) setCitiesLinks(c []*models.City) {
	for i, _ := range c {
		c[i].Links = s.setCityLinks(*c[i])
	}
}

func (s *StatesResource) getCapital(capitalId string) (city models.City) {
	err := s.DB.FindOne(bson.M{"id": capitalId}, &city)
	if err == nil {
		city.Links = s.setCityLinks(city)
	}

	return city
}

func (s *StatesResource) setCityLinks(c models.City) []models.Link {
	return []models.Link{
		{
			Url:  config.ApiRoot + "/v1/states/" + c.StateId + "/cities/" + c.Id,
			Note: "self",
		}, {
			Url:  config.ApiRoot + "/v1/states/" + c.StateId,
			Note: "state",
		},
	}
}
