package database

import (
	"reflect"
	"strings"

	"github.com/databr/api/config"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoDB struct {
	Current *mgo.Database
}

func NewMongoDB(options ...string) MongoDB {
	var mongodb MongoDB

	// logger := log.New(os.Stdout, "LOG: ", log.Ldate|log.Ltime|log.Lshortfile)
	// mgo.SetDebug(false)
	// mgo.SetLogger(logger)

	session, err := mgo.Dial(config.MongoURL)
	checkErr(err)

	databaseName := config.MongoDatabaseName
	if len(options) > 0 {
		databaseName = options[0]
	}
	mongodb.Current = session.DB(databaseName)

	return mongodb
}

func (m MongoDB) Ping() error {
	return m.Current.Session.Ping()
}

func (d MongoDB) FindAll(results interface{}) error {
	return d.collection(results).Find(bson.M{}).All(results)
}

func (d MongoDB) FindOne(query, result interface{}) error {
	return d.collection(result).Find(query).Sort("-updatedat").One(result)
}

func (d MongoDB) Find(query interface{}, limit, page int, result interface{}) error {
	offset := limit * (page - 1)
	return d.collection(result).Find(query).Sort("-updatedat").Limit(limit).Skip(offset).All(result)
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

func (d MongoDB) FindAndGroupBy(by string, query interface{}, result interface{}, collection interface{}) error {
	return d.collection(collection).Pipe([]bson.M{
		{
			"$group": bson.M{
				"_id": by,
				d.collectionName(collection) + "s": bson.M{"$push": "$$ROOT"},
			},
		},
	}).All(result)
}

func (d MongoDB) Upsert(query, data interface{}, _type interface{}) (*mgo.ChangeInfo, error) {
	return d.collection(_type).Upsert(query, data)
}

func (d MongoDB) collection(t interface{}) *mgo.Collection {
	return d.Current.C(d.collectionName(t))
}

func (d MongoDB) collectionName(t interface{}) string {
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

	return strings.ToLower(elem.Name())
}
