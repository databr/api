package main

import (
	"os"

	"github.com/camarabook/camarabook-api/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// database
func databaseMiddleware() gin.HandlerFunc {
	databaseDB := models.New()
	return func(c *gin.Context) {
		c.Set("database", databaseDB)
		c.Next()
	}
}

func getDB(c *gin.Context) gorm.DB {
	database, _ := c.Get("database")
	return database.(gorm.DB)
}

// cors
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	}
}

// entities

type ParliamentarianEntity struct {
	models.Parliamentarian
}

// Handlers
func getParliamentarian(c *gin.Context) {
	uri := c.Params.ByName("uri")
	DB := getDB(c)

	var p ParliamentarianEntity

	err := DB.Table("parliamentarians").Where("uri = ?", uri).Find(&p).Error

	if err != nil {
		c.JSON(404, gin.H{"error": "404", "status": err.Error()})
	} else {
		var about []models.ParliamentarianAbout
		DB.Where("parliamentarian_id = ?", p.RegisterId).Find(&about)
		p.About = about
		c.JSON(200, gin.H{"parliamentarian": p})
	}
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
	}

	r.Run(":" + os.Getenv("PORT"))
}
