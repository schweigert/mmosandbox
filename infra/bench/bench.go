package bench

import (
	"fmt"
	"strings"
	"time"

	"github.com/schweigert/mmosandbox/config"
	"github.com/schweigert/mmosandbox/lib/dont"
	"github.com/schweigert/mmosandbox/lib/metrics"
)

var metronome *metrics.Metronome

func init() {
	var err error
	metronome, err = metrics.Connect(config.Metric().Host(), config.Metric().Port())
	dont.Panic(err)
}

// Bench milliseconds to graphite
func Bench(tag string, fun func()) {
	start := time.Now()

	fun()

	elapsed := time.Since(start)
	metronome.Bip(bipStat(tag), bipMilliseconds(elapsed))
}

func bipMilliseconds(elapsed time.Duration) string {
	return fmt.Sprintf("%d", int64(elapsed/time.Millisecond))
}

func bipStat(tag string) string {
	serviceName := strings.ToLower(config.Service().Name())

	return fmt.Sprintf("bench.%s.%s", serviceName, tag)
}
