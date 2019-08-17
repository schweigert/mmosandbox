package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	parser "github.com/docker/go-units"
	"github.com/schweigert/mmosandbox/config"
	"github.com/schweigert/mmosandbox/lib/dont"
	"github.com/schweigert/mmosandbox/lib/metrics"
)

// Public vars
var (
	metronome *metrics.Metronome
)

func init() {
	var err error
	metronome, err = metrics.Connect(config.Metric().Host(), config.Metric().Port())
	dont.Panic(err)
}

type dockerMetric struct {
	name             string
	cpuPerc          float32
	memPerc          float32
	netIn            int64
	netInUnformat    string
	netOut           int64
	netOutUnformat   string
	blockIn          int64
	blockInUnformat  string
	blockOut         int64
	blockOutUnformat string
	pids             int64
}

const statsCmd = `docker stats --no-stream --format "{{.Name}} {{.CPUPerc}} {{.MemPerc}} {{.NetIO}} {{.BlockIO}} {{.PIDs}}" --no-trunc`

func cmd(bash string) string {
	out, err := exec.Command(`/bin/bash`, `-c`, bash).Output()
	dont.Panic(err)

	return string(out)
}

func parse(element string) int64 {
	// element => 32Kib
	i, err := parser.FromHumanSize(element)
	dont.Panic(err)
	return i
}

func node() string {
	return os.Getenv("NODE")
}

func submit() {
	defer func() {
		_ = recover()
	}()

	lines := strings.Split(cmd(statsCmd), "\n")
	for _, line := range lines {
		metric := dockerMetric{}

		_, err := fmt.Sscanf(
			line,
			"%s %f%% %f%% %s / %s %s / %s %d",
			&metric.name,
			&metric.cpuPerc,
			&metric.memPerc,
			&metric.netInUnformat,
			&metric.netOutUnformat,
			&metric.blockInUnformat,
			&metric.blockOutUnformat,
			&metric.pids,
		)

		if err != nil {
			continue
		}

		metric.netIn = parse(metric.netInUnformat)
		metric.netOut = parse(metric.netOutUnformat)
		metric.blockIn = parse(metric.blockInUnformat)
		metric.blockOut = parse(metric.blockOutUnformat)

		metronome.Bip(fmt.Sprintf("%s.metrics.%s.cpu_perc", node(), metric.name), fmt.Sprintf("%f", metric.cpuPerc))
		metronome.Bip(fmt.Sprintf("%s.metrics.%s.mem_perc", node(), metric.name), fmt.Sprintf("%f", metric.memPerc))
		metronome.Bip(fmt.Sprintf("%s.metrics.%s.net_in", node(), metric.name), fmt.Sprintf("%d", metric.netIn))
		metronome.Bip(fmt.Sprintf("%s.metrics.%s.net_out", node(), metric.name), fmt.Sprintf("%d", metric.netOut))
		metronome.Bip(fmt.Sprintf("%s.metrics.%s.block_in", node(), metric.name), fmt.Sprintf("%d", metric.blockIn))
		metronome.Bip(fmt.Sprintf("%s.metrics.%s.block_out", node(), metric.name), fmt.Sprintf("%d", metric.blockOut))
		metronome.Bip(fmt.Sprintf("%s.metrics.%s.pids", node(), metric.name), fmt.Sprintf("%d", metric.pids))
	}
}

func main() {
	for {
		submit()
	}
}
