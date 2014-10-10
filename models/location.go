package models

type State struct {
	Id                     string `json:"id"`
	Name                   string `json:"name"`
	Capital                string `json:"capital"`
	Population             string `json:"population"`
	Area                   string `json:"area"`
	PopulationDensity      string `json:"population_density"`
	NumberOfMunicipalities string `json:"number_of_municipalities"`
}

type City struct {
	Name       string `json:"name"`
	IbgeCode   string `json:"ibge_code"`
	gentile    string `json:"gentile"`
	population string `json:"population"`
	area       string `json:"area"`
	density    string `json:"density"`
	pib        string `json:"pib"`
}
