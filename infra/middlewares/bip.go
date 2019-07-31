package middlewares

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/krakenlab/ternary"
	"github.com/schweigert/mmosandbox/config"
	"github.com/schweigert/mmosandbox/lib/dont"
	"github.com/schweigert/mmosandbox/lib/metrics"
)

// Bip milliseconds to graphite
func Bip() gin.HandlerFunc {
	metronome, err := metrics.Connect(config.Metric().Host(), config.Metric().Port())

	dont.Panic(err)

	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		url := c.Request.URL.EscapedPath()
		elapsed := time.Since(start)

		metronome.Bip(bipStat(url), bipMilliseconds(elapsed))
	}
}

func bipMilliseconds(elapsed time.Duration) string {
	return fmt.Sprintf("%d", int64(elapsed/time.Millisecond))
}

func bipStat(url string) string {
	serviceName := strings.ToLower(config.Service().Name())
	url = strings.ReplaceAll(url, "/", "_")
	url = strings.TrimSuffix(url, "_")
	url = strings.TrimPrefix(url, "_")
	url = ternary.String(url == "", "root", url)

	return fmt.Sprintf("http.%s.%s", serviceName, url)
}
