package models

import (
	"log"
	"os"
	"regexp"
	"strings"

	. "github.com/fiam/gounidecode/unidecode"
	"github.com/jinzhu/gorm"
	pq "github.com/lib/pq"
)

func New() gorm.DB {
	var database gorm.DB
	var err error

	databaseUrl, _ := pq.ParseURL(os.Getenv("DATABASE_URL"))
	log.Println(databaseUrl)
	database, err = gorm.Open("postgres", databaseUrl)

	if err != nil {
		log.Fatalln(err.Error())
	}

	database.LogMode(os.Getenv("DEBUG") == "true")
	database.AutoMigrate(Parliamentarian{})
	database.AutoMigrate(ParliamentarianAbout{})
	database.AutoMigrate(Party{})

	database.Model(Parliamentarian{}).AddUniqueIndex("x_parliamentarian_uri", "uri")
	database.Model(Parliamentarian{}).AddIndex("x_parliamentarian_gov_id", "gov_id")
	database.Model(Party{}).AddIndex("x_party_title", "title")

	return database
}

// Parliamentarian

type Parliamentarian struct {
	Id         int64                  `json:"-"`
	RegisterId int64                  `json:"register_id"`
	GovId      int64                  `json:"gov_id" sql:"not null;unique"`
	Name       string                 `json:"name" sql:"not null;unique"`
	RealName   string                 `json:"real_name"`
	State      string                 `json:"state"`
	Uri        string                 `json:"id" sql:"not null;unique"`
	Party      Party                  `json:"-"`
	About      []ParliamentarianAbout `json:"about"`
	PartyId    int64                  `json:"party_id" sql:"not null"`
	Gender     string                 `json:"gender"`
	Email      string                 `json:"email"`
	Phone      string                 `json:"phone"`
	ImageUrl   string                 `json:"image_url"`
}

func (p *Parliamentarian) AfterCreate(tx *gorm.DB) (err error) {
	tx.Model(p).Update("uri", makeUri(p.Name))
	return
}

type ParliamentarianAbout struct {
	Id                int64           `json:"-"`
	Parliamentarian   Parliamentarian `json:"-"`
	ParliamentarianId int64           `json:"parliamentarian_id"`
	SectionKey        string          `json:"section_key"`
	Title             string          `json:"title"`
	Body              string          `json:"body"`
}

// Party

type Party struct {
	Id    int64
	Title string `json:"title"`
}

// helper
func makeUri(txt string) string {
	re := regexp.MustCompile(`\W`)
	uri := Unidecode(txt)
	uri = re.ReplaceAllString(uri, "")
	uri = strings.ToLower(uri)
	return uri
}
