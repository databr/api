package models

import (
	"log"
	"os"
	"reflect"
	"strings"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Database struct {
	current *mgo.Database
}

func New() Database {
	var database Database

	// logger := log.New(os.Stdout, "LOG: ", log.Ldate|log.Ltime|log.Lshortfile)
	// mgo.SetDebug(false)
	// mgo.SetLogger(logger)

	log.Println("Trying connect to", os.Getenv("MONGO_URL"), "on", os.Getenv("DATABASE_NAME"))
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

func (d Database) Find(query interface{}, limit, page int, result interface{}) error {
	offset := limit * (page - 1)
	return d.collection(result).Find(query).Sort("id").Limit(limit).Skip(offset).All(result)
}

func (d Database) Count(resource interface{}, query bson.M) (int, error) {
	return d.collection(resource).Find(query).Count()
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

	if elem.Kind() == reflect.Ptr {
		elem = elem.Elem()
	}

	collection := strings.ToLower(elem.Name())
	return d.current.C(collection)
}
