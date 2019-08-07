package middlewares

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/krakenlab/ternary"
	"github.com/schweigert/mmosandbox/config"
	"github.com/schweigert/mmosandbox/lib/bench"
)

// Bip milliseconds to graphite
func Bip() gin.HandlerFunc {
	return func(c *gin.Context) {
		url := c.Request.URL.EscapedPath()
		_ = bench.Bench(bipStat(url), func() error {
			c.Next()
			return nil
		})
	}
}

func bipStat(url string) string {
	serviceName := strings.ToLower(config.Service().Name())
	url = strings.ReplaceAll(url, "/", "_")
	url = strings.TrimSuffix(url, "_")
	url = strings.TrimPrefix(url, "_")
	url = ternary.String(url == "", "root", url)

	return fmt.Sprintf("http.%s.%s", serviceName, url)
}
