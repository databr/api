package service_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/databr/api/database"
	"github.com/databr/api/models"
	"github.com/databr/api/service"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/*
Convert JSON data into a slice.
*/
func sliceFromJSON(data []byte) []interface{} {
	var result interface{}
	json.Unmarshal(data, &result)
	return result.([]interface{})
}

/*
Convert JSON data into a map.
*/
func mapFromJSON(data []byte) map[string]interface{} {
	var result interface{}
	json.Unmarshal(data, &result)
	return result.(map[string]interface{})
}

var _ = Describe("Service", func() {
	var databaseDB database.MongoDB
	var request *http.Request
	var recorder *httptest.ResponseRecorder
	var r *gin.Engine

	BeforeEach(func() {
		databaseDB = database.NewMongoDB("test")
		r = gin.Default()

		parliamentarians := service.ParliamentariansService{r}
		parliamentarians.Run(databaseDB)
		parties := service.PartiesService{r}
		parties.Run(databaseDB)

		recorder = httptest.NewRecorder()
	})

	AfterEach(func() {
		databaseDB.Current.DropDatabase()
	})

	Describe("GET /v1/parliamentarians", func() {
		BeforeEach(func() {
			request, _ = http.NewRequest("GET", "/v1/parliamentarians", nil)
		})

		Context("when no parliamentarians exist", func() {
			It("returns a status code of 200", func() {
				r.ServeHTTP(recorder, request)
				Expect(recorder.Code).To(Equal(200))
			})

			It("returns a empty body", func() {
				r.ServeHTTP(recorder, request)
				Expect(recorder.Body.String()).To(Equal("{\"parliamentarians\":[]}\n"))
			})
		})

		Context("when parliamentarians exist", func() {
			BeforeEach(func() {
				databaseDB.Create(models.Parliamentarian{Name: "Jose"})
				databaseDB.Create(models.Parliamentarian{Name: "Joao"})
			})

			It("returns a status code of 200", func() {
				r.ServeHTTP(recorder, request)
				Expect(recorder.Code).To(Equal(200))
			})

			It("returns those parliamentarians in the body", func() {
				r.ServeHTTP(recorder, request)

				parliamentariansJSON := mapFromJSON(recorder.Body.Bytes())["parliamentarians"].([]interface{})
				Expect(len(parliamentariansJSON)).To(Equal(2))

				parliamentarianJSON := parliamentariansJSON[0].(map[string]interface{})
				Expect(parliamentarianJSON["name"]).To(Equal("Jose"))
			})
		})
	})

	Describe("GET /v1/parliamentarians/:uri", func() {
		BeforeEach(func() {
			request, _ = http.NewRequest("GET", "/v1/parliamentarians/ze", nil)
		})

		Context("when the parliamentarian exist", func() {
			It("returns a status code of 200", func() {
				r.ServeHTTP(recorder, request)
				Expect(recorder.Code).To(Equal(404))
			})

			It("returns a empty body", func() {
				r.ServeHTTP(recorder, request)
				Expect(recorder.Body.String()).To(Equal("{\"error\":\"404\",\"message\":\"not found\"}\n"))
			})
		})

		Context("when the parliamentarian exist", func() {
			BeforeEach(func() {
				databaseDB.Create(models.Parliamentarian{Name: "Jose", Id: "ze"})
			})

			It("returns a status code of 200", func() {
				r.ServeHTTP(recorder, request)
				Expect(recorder.Code).To(Equal(200))
			})

			It("returns the parliamentarian in the body", func() {
				r.ServeHTTP(recorder, request)

				parliamentarianJSON := mapFromJSON(recorder.Body.Bytes())["parliamentarian"].(map[string]interface{})
				Expect(parliamentarianJSON["name"]).To(Equal("Jose"))
			})
		})
	})

	Describe("GET /v1/parties", func() {
		BeforeEach(func() {
			request, _ = http.NewRequest("GET", "/v1/parties", nil)
		})

		Context("when no parties exist", func() {
			It("returns a status code of 200", func() {
				r.ServeHTTP(recorder, request)
				Expect(recorder.Code).To(Equal(200))
			})

			It("returns a empty body", func() {
				r.ServeHTTP(recorder, request)
				Expect(recorder.Body.String()).To(Equal("{\"parties\":[]}\n"))
			})
		})

		Context("when parties exist", func() {
			BeforeEach(func() {
				databaseDB.Create(models.Party{Name: "PPM"})
				databaseDB.Create(models.Party{Name: "PPN"})
			})

			It("returns a status code of 200", func() {
				r.ServeHTTP(recorder, request)
				Expect(recorder.Code).To(Equal(200))
			})

			It("returns those parties in the body", func() {
				r.ServeHTTP(recorder, request)

				partiesJSON := mapFromJSON(recorder.Body.Bytes())["parties"].([]interface{})
				Expect(len(partiesJSON)).To(Equal(2))

				partieJSON := partiesJSON[0].(map[string]interface{})
				Expect(partieJSON["name"]).To(Equal("PPM"))
			})
		})
	})

	Describe("GET /v1/parties/:uri", func() {
		BeforeEach(func() {
			request, _ = http.NewRequest("GET", "/v1/parties/ppm", nil)
		})

		Context("when no parliamentarians exist", func() {
			It("returns a status code of 200", func() {
				r.ServeHTTP(recorder, request)
				Expect(recorder.Code).To(Equal(404))
			})

			It("returns a empty body", func() {
				r.ServeHTTP(recorder, request)
				Expect(recorder.Body.String()).To(Equal("{\"error\":\"404\",\"message\":\"not found\"}\n"))
			})
		})

		Context("when parliamentarians exist", func() {
			BeforeEach(func() {
				databaseDB.Create(models.Party{Name: "PPM", Id: "ppm"})
			})

			It("returns a status code of 200", func() {
				r.ServeHTTP(recorder, request)
				Expect(recorder.Code).To(Equal(200))
			})

			It("returns those signatures in the body", func() {
				r.ServeHTTP(recorder, request)

				parliamentarianJSON := mapFromJSON(recorder.Body.Bytes())["party"].(map[string]interface{})
				Expect(parliamentarianJSON["name"]).To(Equal("PPM"))
			})
		})
	})
})
