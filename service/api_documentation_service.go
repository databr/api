package service

import (
	"strings"

	"github.com/databr/api/config"
	"github.com/databr/api/models"
	"github.com/databr/api/swagger"
	"github.com/gin-gonic/gin"
)

type ApiDocumentationService struct {
	*gin.Engine
}

const (
	cacheFile = "/tmp/documentation"
)

func (a ApiDocumentationService) Run() {
	v1 := a.Group("/v1")
	{
		v1.GET("/doc", func(c *gin.Context) {
			doc := generateDocumentation()
			c.Render(200, DataRender{c.Request}, doc)
		})
	}
}

func generateDocumentation() *swagger.Swagger {
	s := swagger.New()
	s.RefPrefix = []string{"models"}
	s.BasePath = "/v1"
	s.Host = strings.Replace(config.ApiRoot, "http://", "", -1)
	s.Schemes = []string{"http"}
	s.Consumes = []string{"application/json"}
	s.Info = swagger.Info{
		Title:       "DataBR Console",
		Version:     "1.0.0",
		Description: "DataBR é um conjunto de API para ajudar desenvolvedores, jornalistas, analistas e quem mais tiver interesse em trabalhar dados do governo brasileiro. Acreditamos que com nosso esforço na coleta e analise de dados, possibilitando a criação de aplicativos, jogos e visualizações, estamos ajudando para um Brasil melhor.",
		Contact: swagger.Contact{
			Name: "contato@databr.io",
		},
	}

	s.NewGetPath("/parliamentarians", swagger.Request{
		Tags:        []string{"Parlamentares"},
		Summary:     "Retorna parlamentares das casas legislativas",
		Description: "Retorna parlamentares das casas legislativas, podendo ser filtrado por ID. Retornara um JSON com atributo paging, esse atributo ira conter next e/ou previous caso tenha resultados anteriores ou posteriores para o request, o valor de next e previous será sempre um link a ser seguido para buscar mais resultados.",
		Parameters: []swagger.Parameter{
			{
				Name:        "identifier",
				In:          "query",
				Description: "Pode ser: ID usado Senado; os 3 IDs que a Câ¢mara Federal usa(numero de matricula, ID parlamentar, ID de cadastro); o ID usado pelo Transparencia Brasil; o CPF do parlamentar",
				Required:    false,
			}, {
				Name:        "page",
				In:          "query",
				Description: "A paginação se dá atraves da query string page, sendo 1 a primeira pagina e a pagina padrão do request. Cada request retorna 100 registros.",
				Required:    false,
			},
		},
		Responses: swagger.Responses{
			Ok: swagger.Response{
				Description: "Sucesso",
				Schema: swagger.Schema{
					Ref: "#/definitions/ParliamentariansResponse",
				},
			},
			ServerError: swagger.Response{
				Description: "Erro interno",
			},
		},
	})

	s.NewGetPath("/parliamentarians/{id}", swagger.Request{
		Tags:    []string{"Parlamentares"},
		Summary: "Retorna dados de um parlamentar",
		Parameters: []swagger.Parameter{
			{
				Name:        "id",
				In:          "path",
				Description: "id to deputado no databr.io, exemplo: tiririca",
				Required:    true,
			},
		},
		Responses: swagger.Responses{
			Ok: swagger.Response{
				Description: "Sucesso",
				Schema: swagger.Schema{
					Ref: "#/definitions/ParliamentarianResponse",
				},
			},
			ServerError: swagger.Response{
				Description: "Erro interno",
			},
			NotFound: swagger.Response{
				Description: "Parlamentar não encontrado",
			},
		},
	})

	s.NewGetPath("/parties", swagger.Request{
		Tags:        []string{"Partidos"},
		Description: "Retorna dados de um Partido",
		Responses: swagger.Responses{
			Ok: swagger.Response{
				Description: "Sucesso",
				Schema: swagger.Schema{
					Ref: "#/definitions/PartyResponse",
				},
			},
			ServerError: swagger.Response{
				Description: "Erro interno",
			},
		},
	})

	s.NewGetPath("/parties/{id}", swagger.Request{
		Tags:        []string{"Partidos"},
		Summary:     "Dados do Partido",
		Description: "Retorna dados do Partido",
		Parameters: []swagger.Parameter{
			{
				Name:        "id",
				In:          "path",
				Description: "ID do partido, exemplo: psdb",
				Required:    true,
			},
		},
		Responses: swagger.Responses{
			Ok: swagger.Response{
				Description: "Sucesso",
				Schema: swagger.Schema{
					Ref: "#/definitions/PartiesResponse",
				},
			},
			ServerError: swagger.Response{
				Description: "Erro interno",
			},
		},
	})

	s.NewGetPath("/states/sp/transports/trains/lines", swagger.Request{
		Summary: "Linhas de Trem e Metro de São Paulo",
		Tags:    []string{"Trens SP"},
		Responses: swagger.Responses{
			Ok: swagger.Response{
				Description: "Sucesso",
				Schema: swagger.Schema{
					Ref: "#/definitions/LinesResponse",
				},
			},
			ServerError: swagger.Response{
				Description: "Erro interno",
			},
		},
	})

	s.NewGetPath("/states/sp/transports/trains/lines/{uri}", swagger.Request{
		Summary: "Dados da linha solicitada",
		Tags:    []string{"Trens SP"},
		Parameters: []swagger.Parameter{
			{
				Name:        "uri",
				In:          "path",
				Description: "ID da Linha, exemplo: linha1azul",
				Required:    true,
			},
		},
		Responses: swagger.Responses{
			Ok: swagger.Response{
				Description: "Sucesso",
				Schema: swagger.Schema{
					Ref: "#/definitions/LineResponse",
				},
			},
			ServerError: swagger.Response{
				Description: "Erro interno",
			},
			NotFound: swagger.Response{
				Description: "Linha de Trem não encontrada",
			},
		},
	})

	s.NewGetPath("/states/sp/transports/trains/lines/{uri}/statuses", swagger.Request{
		Summary:     "Retorna Ultimos Status da Linha",
		Description: "Histórico da linha, um novo status é criado quando o status da linha muda, caso contratio apenas é atualizado o campo updated_at",
		Tags:        []string{"Trens SP"},
		Parameters: []swagger.Parameter{
			{
				Name:        "uri",
				In:          "path",
				Description: "ID da Linha, exemplo: linha1azul",
				Required:    true,
			},
		},
		Responses: swagger.Responses{
			Ok: swagger.Response{
				Description: "Sucesso",
				Schema: swagger.Schema{
					Ref: "#/definitions/StatusesResponse",
				},
			},
			ServerError: swagger.Response{
				Description: "Erro interno",
			},
			NotFound: swagger.Response{
				Description: "Linha de Trem não encontrada",
			},
		},
	})

	s.NewGetPath("/states/sp/transports/trains/lines/{uri}/statuses/{status_id}", swagger.Request{
		Summary: "Retorna Status solicitado",
		Tags:    []string{"Trens SP"},
		Parameters: []swagger.Parameter{
			{
				Name:        "uri",
				In:          "path",
				Description: "ID da Linha, exemplo: linha1azul",
				Required:    true,
			}, {
				Name:        "status_id",
				In:          "path",
				Description: "ID do Status",
				Required:    true,
			},
		},
		Responses: swagger.Responses{
			Ok: swagger.Response{
				Description: "Sucesso",
				Schema: swagger.Schema{
					Ref: "#/definitions/StatusResponse",
				},
			},
			ServerError: swagger.Response{
				Description: "Erro interno",
			},
			NotFound: swagger.Response{
				Description: "Linha de Trem ou Status não encontrada",
			},
		},
	})

	s.NewDefinition("ParliamentariansResponse", map[string]interface{}{
		"paging": map[string]interface{}{
			"$ref": "Pagging",
		},
		"parliamentarians": map[string]interface{}{
			"type": "array",
			"items": map[string]interface{}{
				"$ref": "Parliamentarian",
			},
		},
	})
	s.NewDefinition("ParliamentarianResponse", map[string]interface{}{
		"parliamentarian": map[string]interface{}{
			"$ref": "Parliamentarian",
		},
	})
	s.NewDefinition("PartiesResponse", map[string]interface{}{
		"paging": map[string]interface{}{
			"$ref": "Pagging",
		},
		"parties": map[string]interface{}{
			"type": "array",
			"items": map[string]interface{}{
				"$ref": "Party",
			},
		},
	})
	s.NewDefinition("PartyResponse", map[string]interface{}{
		"party": map[string]interface{}{
			"$ref": "Party",
		},
	})
	s.NewDefinition("LinesResponse", map[string]interface{}{
		"paging": map[string]interface{}{
			"$ref": "Pagging",
		},
		"lines": map[string]interface{}{
			"type": "array",
			"items": map[string]interface{}{
				"$ref": "Line",
			},
		},
	})
	s.NewDefinition("LineResponse", map[string]interface{}{
		"line": map[string]interface{}{
			"$ref": "Line",
		},
	})
	s.NewDefinition("StatusesResponse", map[string]interface{}{
		"paging": map[string]interface{}{
			"$ref": "Pagging",
		},
		"statuses": map[string]interface{}{
			"type": "array",
			"items": map[string]interface{}{
				"$ref": "Status",
			},
		},
	})
	s.NewDefinition("StatusResponse", map[string]interface{}{
		"status": map[string]interface{}{
			"$ref": "Status",
		},
	})
	s.NewDefinition("Pagging", map[string]interface{}{
		"next": map[string]interface{}{
			"type":    "string",
			"example": "http://api.databr.io/v1/parliamentarians/?page=3",
		},
		"prev": map[string]interface{}{
			"type":    "string",
			"example": "http://api.databr.io/v1/parliamentarians/?page=1",
		},
	})

	s.GenerateDefinition(models.Parliamentarian{})
	s.GenerateDefinition(models.ContactDetail{})
	s.GenerateDefinition(models.Membership{})
	s.GenerateDefinition(models.Source{})
	s.GenerateDefinition(models.OtherNames{})
	s.GenerateDefinition(models.Party{})
	s.GenerateDefinition(models.Rel{})
	s.GenerateDefinition(models.Identifier{})
	s.GenerateDefinition(models.Line{})
	s.GenerateDefinition(models.Link{})
	s.GenerateDefinition(models.Status{})

	return s
}
