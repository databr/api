package config

import "os"

const (
	StatusPageIOApiBase = "https://api.statuspage.io/v1"
)

var (
	ApiRoot string
	Env     string
	Port    string
	Debug   string

	StatusPageIOApiKey   string
	StatusPageIOMetricID string
	StatusPageIOPageID   string
	StatusPageIOEnable   string

	InfluxdbHost     string
	InfluxdbUser     string
	InfluxdbPassword string
	InfluxdbDatabase string
	InfluxdbEnable   string

	MongoURL          string
	MongoDatabaseName string

	PrivateKey string
)

func init() {
	ApiRoot = env("API_ROOT", false)
	Env = env("ENV", false)
	Port = env("PORT", false)
	Debug = env("DEBUG", true)

	StatusPageIOEnable = env("STATUSPAGEIO_ENABLE", false)
	StatusPageIOApiKey = env("STATUSPAGEIO_API_KEY", StatusPageIOEnable != "true")
	StatusPageIOMetricID = env("STATUSPAGEIO_METRIC_ID", StatusPageIOEnable != "true")
	StatusPageIOPageID = env("STATUSPAGEIO_PAGE_ID", StatusPageIOEnable != "true")

	InfluxdbEnable = env("INFLUXDB_ENABLE", false)
	InfluxdbHost = env("INFLUXDB_HOST", InfluxdbEnable != "true")
	InfluxdbUser = env("INFLUXDB_USERNAME", InfluxdbEnable != "true")
	InfluxdbPassword = env("INFLUXDB_PASSWORD", InfluxdbEnable != "true")
	InfluxdbDatabase = env("INFLUXDB_DATABASE", InfluxdbEnable != "true")

	MongoURL = env("MONGO_URL", false)
	MongoDatabaseName = env("MONGO_DATABASE_NAME", false)

	PrivateKey = env("PRIVATE_KEY", false)
}

func env(s string, optional bool) string {
	v := os.Getenv(s)
	if v == "" && !optional {
		panic("VAR " + s + " empty, run:\n\nexport " + s + "='SOME VALUE'\n\n")
	}

	return v
}
