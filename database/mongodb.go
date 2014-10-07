package database

import (
	"log"
	"reflect"
	"strings"

	"github.com/databr/api/config"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoDB struct {
	current *mgo.Database
}

func NewMongoDB() MongoDB {
	var mongodb MongoDB

	// logger := log.New(os.Stdout, "LOG: ", log.Ldate|log.Ltime|log.Lshortfile)
	// mgo.SetDebug(false)
	// mgo.SetLogger(logger)

	log.Println("Trying connect to", config.MongoURL, "on", config.MongoDatabaseName)
	session, err := mgo.Dial(config.MongoURL)
	checkErr(err)

	mongodb.current = session.DB(config.MongoDatabaseName)

	return mongodb
}

func (m MongoDB) Ping() error {
	return m.current.Session.Ping()
}

func (d MongoDB) FindAll(results interface{}) error {
	return d.collection(results).Find(bson.M{}).All(results)
}

func (d MongoDB) FindOne(query, result interface{}) error {
	return d.collection(result).Find(query).One(result)
}

func (d MongoDB) Find(query interface{}, limit, page int, result interface{}) error {
	offset := limit * (page - 1)
	return d.collection(result).Find(query).Sort("id").Limit(limit).Skip(offset).All(result)
}

func (d MongoDB) Count(resource interface{}, query bson.M) (int, error) {
	return d.collection(resource).Find(query).Count()
}

func (d MongoDB) Create(data interface{}) error {
	return d.collection(data).Insert(data)
}

func (d MongoDB) Update(query, data interface{}, _type interface{}) error {
	return d.collection(_type).Update(query, data)
}

func (d MongoDB) Upsert(query, data interface{}, _type interface{}) (*mgo.ChangeInfo, error) {
	return d.collection(_type).Upsert(query, data)
}

func (d MongoDB) collection(t interface{}) *mgo.Collection {
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
