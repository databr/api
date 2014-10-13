package models

type State struct {
	Id                     string   `json:"id"`
	Name                   string   `json:"name"`
	Capital                string   `json:"capital"`
	Population             float64  `json:"population"`
	Area                   float64  `json:"area"`
	PopulationDensity      float64  `json:"population_density"`
	NumberOfMunicipalities int      `json:"number_of_municipalities"`
	Sources                []Source `json:"sources"`
}

type City struct {
	Name       string   `json:"name"`
	IbgeCode   string   `json:"ibge_code"`
	gentile    string   `json:"gentile"`
	population string   `json:"population"`
	area       string   `json:"area"`
	density    string   `json:"density"`
	pib        string   `json:"pib"`
	Sources    []Source `json:"sources"`
}
