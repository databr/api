package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yvasiyarov/go-metrics"
	"github.com/yvasiyarov/gorelic"
)

var agent *gorelic.Agent

func NewRelic() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		agent.HTTPTimer.UpdateSince(startTime)
	}
}

func InitNewrelicAgent(license string, appname string, verbose bool) error {

	if license == "" {
		panic("Please specify NewRelic license")
	}

	agent = gorelic.NewAgent()
	agent.NewrelicLicense = license

	agent.HTTPTimer = metrics.NewTimer()
	agent.CollectHTTPStat = true
	agent.Verbose = verbose

	agent.NewrelicName = appname
	agent.Run()
	return nil
}
