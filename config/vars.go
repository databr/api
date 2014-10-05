package config

import "os"

const (
	StatusPageIOApiBase = "https://api.statuspage.io/v1"
)

var (
	ApiRoot string
	Env     string
	Port    string

	StatusPageIOApiKey   string
	StatusPageIOMetricID string
	StatusPageIOPageID   string
	StatusPageIOEnable   string

	InfluxdbHost     string
	InfluxdbUser     string
	InfluxdbPassword string
	InfluxdbDatabase string

	MongoURL          string
	MongoDatabaseName string
)

func init() {
	ApiRoot = env("API_ROOT")
	Env = env("ENV")
	Port = env("PORT")

	StatusPageIOApiKey = os.Getenv("STATUSPAGEIO_API_KEY")
	StatusPageIOMetricID = os.Getenv("STATUSPAGEIO_METRIC_ID")
	StatusPageIOPageID = os.Getenv("STATUSPAGEIO_PAGE_ID")
	StatusPageIOEnable = env("STATUSPAGEIO_ENABLE")

	InfluxdbHost = env("INFLUXDB_HOST")
	InfluxdbUser = env("INFLUXDB_USERNAME")
	InfluxdbPassword = env("INFLUXDB_PASSWORD")
	InfluxdbDatabase = env("INFLUXDB_DATABASE")

	MongoURL = env("MONGO_URL")
	MongoDatabaseName = env("MONGO_DATABASE_NAME")
}

func env(s string) string {
	v := os.Getenv(s)
	if v == "" {
		panic("VAR " + s + " empty, run:\n\nexport " + s + "='SOME VALUE'\n\n")
	}

	return v
}
