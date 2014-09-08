package models

import (
	"regexp"
	"strings"
	"time"

	"github.com/databr/go-popolo"
	. "github.com/fiam/gounidecode/unidecode"
)

// Parliamentarian

type Parliamentarian popolo.Person

type Party popolo.Organization

type Company popolo.Organization

type Quota struct {
	Company         string
	Date            time.Time
	Parliamentarian string
	Order           string
	Value           float64

	PassengerName string
	Route         string
	Ticket        string
}

// helper
func MakeUri(txt string) string {
	re := regexp.MustCompile(`\W`)
	uri := Unidecode(txt)
	uri = re.ReplaceAllString(uri, "")
	uri = strings.ToLower(uri)
	return uri
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
