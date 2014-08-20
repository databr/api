package main

import (
	"log"
	"os"

	"github.com/camarabook/camarabook-api/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

// database
func databaseMiddleware() gin.HandlerFunc {
	databaseDB := models.New()
	return func(c *gin.Context) {
		c.Set("database", databaseDB)
		c.Next()
	}
}

func getDB(c *gin.Context) models.Database {
	database, _ := c.Get("database")
	return database.(models.Database)
}

// cors
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	}
}

// Handlers
func getParliamentarian(c *gin.Context) {
	uri := c.Params.ByName("uri")
	DB := getDB(c)

	var p models.Parliamentarian

	err := DB.FindOne(bson.M{"id": uri}, &p)

	if err != nil {
		c.JSON(404, gin.H{"error": "404", "status": err.Error()})
	} else {
		c.JSON(200, gin.H{"parliamentarian": p})
	}
}

func getParliamentarianActivities(c *gin.Context) {
	uri := c.Params.ByName("uri")
	DB := getDB(c)
	activities := make([]models.Activity, 0)

	var q []models.Quota
	DB.Find(bson.M{"parliamentarian": uri}, &q)
	for _, item := range q {
		log.Println(item)
		activities = append(activities, item.ToActivity(DB))
	}

	c.JSON(200, gin.H{"activities": activities})
}

func main() {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(databaseMiddleware())
	r.Use(corsMiddleware())

	v1 := r.Group("/v1")
	{
		v1.GET("/parliamentarians", func(c *gin.Context) {
			c.String(200, "Hi!")
		})

		v1.GET("/parliamentarians/:uri", getParliamentarian)
		v1.GET("/parliamentarians/:uri/activities", getParliamentarianActivities)
	}

	r.Run(":" + os.Getenv("PORT"))
}
