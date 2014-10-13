package service

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/databr/api/config"
	"github.com/databr/api/database"
	"gopkg.in/mgo.v2/bson"
)

const (
	PER_PAGE_LIMIT = 100
)

func pagination(resourceURI string,
	database database.MongoDB,
	limit,
	currentPage int,
	resourceClass interface{},
	query bson.M,
) map[string]interface{} {
	total, _ := database.Count(resourceClass, query)

	pagination := map[string]interface{}{}

	if currentPage > 1 {
		pagination["previous"] = fmt.Sprintf("%s/%s/?page=%d", config.ApiRoot, resourceURI, currentPage-1)
	}

	if total > (limit * currentPage) {
		pagination["next"] = fmt.Sprintf("%s/%s/?page=%d", config.ApiRoot, resourceURI, currentPage+1)
	}

	return pagination
}

func GetLimit(request *http.Request) int {
	limitS := request.URL.Query().Get("limit")
	if limitS == "" {
		return PER_PAGE_LIMIT
	}

	limit, err := strconv.Atoi(limitS)

	if err != nil {
		return PER_PAGE_LIMIT

	}

	return limit
}
