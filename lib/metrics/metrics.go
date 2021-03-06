package metrics

import (
	"log"
	"strconv"

	"github.com/logrusorgru/aurora"
	"github.com/marpaia/graphite-golang"
	"github.com/schweigert/mmosandbox/lib/dont"
)

// Metronome to send float metrics
type Metronome struct {
	Conn *graphite.Graphite
	Chan chan graphite.Metric
}

// Connect constructor
func Connect(host string, port string) (*Metronome, error) {
	portAsInt, err := strconv.Atoi(port)

	if err == nil {
		conn, err := graphite.NewGraphite(host, portAsInt)
		return &Metronome{Conn: conn}, err
	}

	return nil, err
}

// Bip beacon
func (metronome *Metronome) Bip(stat string, value string) {
	err := metronome.Conn.SimpleSend(stat, value)
	log.Println(aurora.Red(stat).Bold(), "|>", aurora.Magenta(value))
	dont.Panic(err)
}
