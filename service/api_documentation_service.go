package service

import (
	"reflect"
	"strings"

	"github.com/databr/api/config"
	"github.com/databr/api/models"
	"github.com/gin-gonic/gin"
)

type ApiDocumentationService struct {
	*gin.Engine
}

func (a ApiDocumentationService) Run() {
	v1 := a.Group("/v1")
	{
		v1.GET("/doc", func(c *gin.Context) {
			doc := map[string]interface{}{}
			doc["swagger"] = 2.0
			doc["info"] = map[string]interface{}{
				"description": "DataBR é um conjunto de API para ajudar desenvolvedores, jornalistas, analistas e quem mais tiver interesse em trabalhar dados do governo brasileiro. Acreditamos que com nosso esforço na coleta e analise de dados, possibilitando a criação de aplicativos, jogos e visualizações, estamos ajudando para um Brasil melhor.",
				"version":     "1.0.0",
				"title":       "DataBR Console",
				"contact": map[string]string{
					"name": "contato@datanr.io",
				},
			}
			doc["host"] = strings.Replace(config.ApiRoot, "http://", "", -1)
			doc["basePath"] = "/v1"
			doc["schemes"] = []string{"http"}
			doc["consumes"] = []string{"application/json"}
			doc["paths"] = map[string]interface{}{
				"/parliamentarians": map[string]interface{}{
					"get": map[string]interface{}{
						"tags":        []string{"Parlamentares"},
						"summary":     "Retorna parlamentares das casas legislativas",
						"description": "Retorna parlamentares das casas legislativas, podendo ser filtrado por ID. Retornara um JSON com atributo paging, esse atributo ira conter next e/ou previous caso tenha resultados anteriores ou posteriores para o request, o valor de next e previous será sempre um link a ser seguido para buscar mais resultados.",
						"parameters": []map[string]interface{}{
							{
								"name":        "identifier",
								"in":          "query",
								"description": "Pode ser: ID usado Senado; os 3 IDs que a Câ¢mara Federal usa(numero de matricula, ID parlamentar, ID de cadastro); o ID usado pelo Transparencia Brasil; o CPF do parlamentar",
								"required":    false,
							}, {
								"name":        "page",
								"in":          "query",
								"description": "A paginação se dá atraves da query string page, sendo 1 a primeira pagina e a pagina padrão do request. Cada request retorna 100 registros.",
								"required":    false,
							},
						},
						"responses": map[string]interface{}{
							"200": map[string]interface{}{
								"description": "Sucesso",
								"schema": map[string]interface{}{
									"$ref": "#/definitions/ParliamentariansResponse",
								},
							},
							"500": map[string]interface{}{
								"description": "Erro interno",
							},
						},
					},
				},
				"/parliamentarians/{id}": map[string]interface{}{
					"get": map[string]interface{}{
						"tags":   []string{"Parlamentares"},
						"sumary": "Retorna dados de um parlamentar",
						"parameters": []map[string]interface{}{
							{
								"name":        "id",
								"in":          "path",
								"description": "id to deputado no databr.io, exemplo: tiririca",
								"required":    true,
							},
						},
						"responses": map[string]interface{}{
							"200": map[string]interface{}{
								"description": "Sucesso",
								"schema": map[string]interface{}{
									"$ref": "#/definitions/ParliamentarianResponse",
								},
							},
							"500": map[string]interface{}{
								"description": "Erro interno",
							},
							"404": map[string]interface{}{
								"description": "Parlamentar não encontrado",
							},
						},
					},
				},
				"/parties": map[string]interface{}{
					"get": map[string]interface{}{
						"tags":        []string{"Partidos"},
						"description": "Retorna dados de um Partido",
						"responses": map[string]interface{}{
							"200": map[string]interface{}{
								"description": "Sucesso",
								"schema": map[string]interface{}{
									"$ref": "#/definitions/PartyResponse",
								},
							},
							"500": map[string]interface{}{
								"description": "Erro interno",
							},
						},
					},
				},
				"/parties/{id}": map[string]interface{}{
					"get": map[string]interface{}{
						"tags":        []string{"Partidos"},
						"summary":     "Partidos politicos",
						"description": "Retorna dados d",
						"parameters": []map[string]interface{}{
							{
								"name":        "id",
								"in":          "path",
								"description": "ID do partido, exemplo: psdb",
								"required":    true,
							},
						},

						"responses": map[string]interface{}{
							"200": map[string]interface{}{
								"description": "Sucesso",
								"schema": map[string]interface{}{
									"$ref": "#/definitions/PartiesResponse",
								},
							},
							"500": map[string]interface{}{
								"description": "Erro interno",
							},
						},
					},
				},
				"/states/sp/transport/trains/lines": map[string]interface{}{
					"get": map[string]interface{}{
						"summary": "Test",
						"tags":    []string{"Transporte SP"},
					},
				},
				"/states/sp/transport/trains/lines/{uri}": map[string]interface{}{
					"get": map[string]interface{}{
						"summary": "Test",
						"tags":    []string{"Transporte SP"},
					},
				},
				"/states/sp/transport/trains/lines/{uri}/statuses": map[string]interface{}{
					"get": map[string]interface{}{
						"summary": "Test",
						"tags":    []string{"Transporte SP"},
					},
				},
			}
			doc["definitions"] = map[string]interface{}{
				"ParliamentariansResponse": map[string]interface{}{
					"properties": map[string]interface{}{
						"paging": map[string]interface{}{
							"$ref": "Pagging",
						},
						"parliamentarians": map[string]interface{}{
							"type": "array",
							"items": map[string]interface{}{
								"$ref": "Parliamentarian",
							},
						},
					},
				},
				"ParliamentarianResponse": map[string]interface{}{
					"properties": map[string]interface{}{
						"parliamentarian": map[string]interface{}{
							"$ref": "Parliamentarian",
						},
					},
				},
				"PartiesResponse": map[string]interface{}{
					"properties": map[string]interface{}{
						"paging": map[string]interface{}{
							"$ref": "Pagging",
						},
						"parties": map[string]interface{}{
							"$ref": "Party",
						},
					},
				},
				"PartyResponse": map[string]interface{}{
					"properties": map[string]interface{}{
						"party": map[string]interface{}{
							"$ref": "Party",
						},
					},
				},

				"Pagging": map[string]interface{}{
					"properties": map[string]interface{}{
						"next": map[string]interface{}{
							"type":    "string",
							"example": "http://api.databr.io/v1/parliamentarians/?page=3",
						},
						"prev": map[string]interface{}{
							"type":    "string",
							"example": "http://api.databr.io/v1/parliamentarians/?page=1",
						},
					},
				},
				"Parliamentarian": generateDefinition(models.Parliamentarian{}),
				"ContactDetail":   generateDefinition(models.ContactDetail{}),
				"Membership":      generateDefinition(models.Membership{}),
				"Source":          generateDefinition(models.Source{}),
				"OtherNames":      generateDefinition(models.OtherNames{}),
				"Party":           generateDefinition(models.Party{}),
				"Rel":             generateDefinition(models.Rel{}),
			}

			c.Render(200, DataRender{c.Request}, doc)
		})
	}
}

func generateDefinition(inter interface{}) map[string]map[string]map[string]interface{} {
	model := reflect.ValueOf(inter).Type()

	d := map[string]map[string]map[string]interface{}{
		"properties": map[string]map[string]interface{}{},
	}

	for i := 0; i < model.NumField(); i++ {
		field := model.Field(i)
		key := getKey(field)
		fType := getFieldType(field)

		switch fType.Kind() {
		case reflect.Struct:
			property := strings.ToLower(fType.Name())
			d["properties"][key] = propertyType(property)
		case reflect.Slice:
			ref := fType.Elem().Name()
			d["properties"][key] = propertyArrayRef(ref)
		default:
			property := fType.String()
			d["properties"][key] = propertyType(property)
		}
	}

	return d
}

func getKey(field reflect.StructField) string {
	return strings.Split(field.Tag.Get("json"), ",")[0]
}

func getFieldType(f reflect.StructField) reflect.Type {
	t := f.Type
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}

func propertyType(s string) map[string]interface{} {
	return map[string]interface{}{
		"type": s,
	}
}

func propertyArrayRef(s string) map[string]interface{} {
	return map[string]interface{}{
		"type":  "array",
		"items": propertyRef(s),
	}
}

func propertyRef(s string) map[string]interface{} {
	return map[string]interface{}{
		"$ref": s,
	}
}
