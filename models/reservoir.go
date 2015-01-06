package models

type Reservoir struct {
	Name string `json:"name"`
	Uri  string `json:"-"`
	Data []struct {
		Date    string `json:"date"`
		Percent string `json:"percent"`
	} `json:"data"`
}
