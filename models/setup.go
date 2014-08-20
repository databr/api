package models

import (
	"log"
	"os"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/camarabook/go-popolo"
	. "github.com/fiam/gounidecode/unidecode"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Database struct {
	current *mgo.Database
}

func New() Database {
	var database Database

	// logger := log.New(os.Stdout, "LOG: ", log.Ldate|log.Ltime|log.Lshortfile)
	//mgo.SetDebug(false)
	//mgo.SetLogger(logger)

	session, err := mgo.Dial(os.Getenv("MONGO_URL"))
	checkErr(err)

	database.current = session.DB(os.Getenv("DATABASE_NAME"))

	return database
}

func (d Database) FindAll(results interface{}) error {
	return d.collection(results).Find(bson.M{}).All(results)
}

func (d Database) FindOne(query, result interface{}) error {
	return d.collection(result).Find(query).One(result)
}

func (d Database) Create(data interface{}) error {
	return d.collection(data).Insert(data)
}

func (d Database) Update(query, data interface{}, _type interface{}) error {
	return d.collection(_type).Update(query, data)
}

func (d Database) Upsert(query, data interface{}, _type interface{}) (*mgo.ChangeInfo, error) {
	return d.collection(_type).Upsert(query, data)
}

func (d Database) collection(t interface{}) *mgo.Collection {
	v := reflect.ValueOf(t)

	elem := v.Type()
	if elem.Kind() == reflect.Ptr {
		elem = elem.Elem()
	}

	if elem.Kind() == reflect.Slice {
		elem = elem.Elem()
	}

	collection := strings.ToLower(elem.Name())
	return d.current.C(collection)
}

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
