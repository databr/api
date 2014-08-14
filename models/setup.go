package models

import (
	"log"
	"os"
	"reflect"
	"regexp"
	"strings"

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

	logger := log.New(os.Stdout, "LOG: ", log.Ldate|log.Ltime|log.Lshortfile)

	mgo.SetDebug(false)
	mgo.SetLogger(logger)

	session, err := mgo.Dial(os.Getenv("MONGO_URL"))
	checkErr(err)

	database.current = session.DB(os.Getenv("DATABASE_NAME"))

	return database
}

func (d Database) FindAll(t interface{}) error {
	return d.collection(t).Find(bson.M{}).All(t)
}

func (d Database) FindOne(query, t interface{}) error {
	return d.collection(t).Find(query).One(t)
}

func (d Database) Create(data interface{}) error {
	return d.collection(data).Insert(data)
}

func (d Database) Update(query, update interface{}, t interface{}) error {
	return d.collection(t).Update(query, update)
}

func (d Database) Upsert(query, data interface{}) (*mgo.ChangeInfo, error) {
	return d.collection(data).Upsert(query, data)
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
