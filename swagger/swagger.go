package swagger

import (
	"reflect"
	"strings"
)

type Swagger struct {
	RefPrefix   []string                `json:"-"`
	Version     float64                 `json:"swagger"`
	Host        string                  `json:"host"`
	Info        Info                    `json:"info"`
	BasePath    string                  `json:"base_path"`
	Consumes    []string                `json:"consumes"`
	Schemes     []string                `json:"schemes"`
	Paths       []map[string]Path       `json:"paths"`
	Definitions []map[string]Definition `json:"definitions"`
}

type Info struct {
	Title       string  `json:"title"`
	Version     string  `json:"version"`
	Description string  `json:"description"`
	Contact     Contact `json:"contact"`
}

type Contact struct {
	Name string `json:"name"`
}

type Path struct {
	Get Request `json:"get"`
}

type Request struct {
	Tags        []string    `json:"tags"`
	Summary     string      `json:"summary"`
	Description string      `json:"description"`
	Parameters  []Parameter `json:"parameters"`
	Responses   Responses   `json:"responses"`
}

type Parameter struct {
	Name        string `json:"name"`
	In          string `json:"in"`
	Description string `json:"description"`
	Required    bool   `json:"required"`
}

type Responses struct {
	Ok          Response `json:"200,omitempty"`
	NotFound    Response `json:"404,omitempty"`
	ServerError Response `json:"500,omitempty"`
}

type Response struct {
	Schema      Schema `json:"schema,omitempty"`
	Description string `json:"description,omitempty"`
}

type Schema struct {
	Ref string `json:"$ref",omitempty`
}

type Definition struct {
	Properties map[string]interface{} `json:"properties"`
}

func New() *Swagger {
	return &Swagger{
		Version:     2.0,
		Paths:       make([]map[string]Path, 0),
		Definitions: make([]map[string]Definition, 0),
		RefPrefix:   make([]string, 0),
	}
}

func (s *Swagger) NewDefinition(name string, properties map[string]interface{}) {
	d := Definition{}
	d.Properties = properties

	o := map[string]Definition{}
	o[name] = d

	s.Definitions = append(s.Definitions, o)
}

func (s *Swagger) GenerateDefinition(c interface{}) {
	model := reflect.ValueOf(c).Type()

	properties := map[string]interface{}{}

	for i := 0; i < model.NumField(); i++ {
		field := model.Field(i)
		key := getKey(field)
		fType := getFieldType(field)

		switch fType.Kind() {
		case reflect.Struct:
			if m := strings.Split(fType.String(), "."); m[0] == "models" {
				properties[key] = propertyRef(fType.Name())
			} else {
				property := strings.ToLower(fType.Name())
				properties[key] = propertyType(property)
			}
		case reflect.Slice:
			if fType.Elem().Kind() == reflect.Ptr {
				fType = fType.Elem()
			}
			ref := fType.Elem().Name()
			properties[key] = propertyArrayRef(ref)
		default:
			property := fType.String()
			properties[key] = propertyType(property)
		}
	}

	s.NewDefinition(model.Name(), properties)
}

func (s *Swagger) NewGetPath(path string, r Request) {
	k := map[string]Path{}
	k[path] = Path{
		Get: r,
	}

	s.Paths = append(s.Paths, k)
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
	d := map[string]interface{}{}

	if s == "time" {
		s = "date-time"
	}

	if s == "bson.ObjectId" {
		s = "string"
	}

	if s != "string" {
		d["format"] = s
		s = "string"
	}

	d["type"] = s
	return d
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

func (s *Swagger) inRefPrefix(name string) bool {
	for _, n := range s.RefPrefix {
		if n == name {
			return true
		}
	}
	return false
}