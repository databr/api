package models

type State struct {
	Id                     string   `json:"id"`
	Name                   string   `json:"name"`
	CapitalId              string   `json:"capital_id"`
	Capital                City     `json:"capital"`
	Population             float64  `json:"population"`
	Area                   float64  `json:"area"`
	PopulationDensity      float64  `json:"population_density"`
	NumberOfMunicipalities int      `json:"number_of_municipalities"`
	Sources                []Source `json:"sources"`
	Links                  []Link   `json:"links"`
}

type City struct {
	Id         string   `json:"id"`
	Name       string   `json:"name"`
	IbgeCode   int      `json:"ibge_code"`
	Gentile    string   `json:"gentile"`
	Population float64  `json:"population"`
	Area       float64  `json:"area"`
	Density    float64  `json:"density"`
	Pib        float64  `json:"pib"`
	StateId    string   `json:"state_id"`
	Sources    []Source `json:"sources"`
	Links      []Link   `json:"links"`
}
