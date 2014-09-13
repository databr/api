package service

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/databr/api/models"
	"github.com/databr/go-popolo"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

const (
	PER_PAGE_LIMIT = 100
)

var (
	API_ROOT string
	ENV      string
)

func init() {
	API_ROOT = os.Getenv("API_ROOT")
	ENV = os.Getenv("ENV")
}

func pagination(resourceURI string,
	database models.Database,
	limit,
	currentPage int,
	resourceClass interface{},
	query bson.M,
) map[string]interface{} {
	total, _ := database.Count(resourceClass, query)

	pagination := map[string]interface{}{}

	if currentPage > 1 {
		pagination["previous"] = fmt.Sprintf("%s/%s/?page=%d", API_ROOT, resourceURI, currentPage-1)
	}

	if total > (limit * currentPage) {
		pagination["next"] = fmt.Sprintf("%s/%s/?page=%d", API_ROOT, resourceURI, currentPage+1)
	}

	return pagination
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
		gzipJSON(c, 200, gin.H{
			"parliamentarians": p,
			"paging": pagination(
				"v1/parliamentarians",
				r.DB,
				PER_PAGE_LIMIT,
				page,
				models.Parliamentarian{},
				search,
			),
		})
	}
}

func (r *ParliamentarianResource) Get(c *gin.Context) {
	uri := c.Params.ByName("uri")

	var p models.Parliamentarian

	err := r.DB.FindOne(bson.M{"id": uri}, &p)

	if err != nil {
		c.JSON(404, gin.H{"error": "404", "message": err.Error()})
	} else {
		gzipJSON(c, 200, gin.H{"parliamentarian": p})
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

func gzipJSON(c *gin.Context, code int, data ...interface{}) {
	var writer io.Writer

	w := c.Writer
	r := c.Request

	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")

	if useGzip(r.Header) {
		gz := gzip.NewWriter(w)
		w.Header().Set("Content-Encoding", "gzip")
		defer gz.Close()
		writer = gz
	} else {
		writer = w
	}

	json.NewEncoder(writer).Encode(data[0])
}

func useGzip(h http.Header) bool {
	return ENV != "development" &&
		strings.Contains(h.Get("User-Agent"), "Mozilla") &&
		strings.Contains(h.Get("User-Agent"), "Opera") &&
		strings.Contains(h.Get("Accept-Encoding"), "gzip")
}
