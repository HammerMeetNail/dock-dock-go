package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/HammerMeetNail/dock-dock-go/pkg/clients"
	"github.com/HammerMeetNail/dock-dock-go/pkg/stats"
)

type data struct {
	dest string
	url  string
	name string
	size int64
}

func output(data data) {

	switch data.dest {
	case "statsd":
		fmt.Println("Sending to statsd!!!!")
	default:
		fmt.Println(data.name, data.size)
	}

}

func main() {

	// Docker
	maxDockerVersion := "1.39"
	cli := clients.Cli(maxDockerVersion)

	dockerVersion, ok := os.LookupEnv("DOCKER_VERSION")
	if ok {
		cli = clients.Cli(dockerVersion)
	}

	// Logging
	logLevel, ok := os.LookupEnv("LOG_LEVEL")

	switch logLevel {
	case "debug":
		log.SetLevel(log.DebugLevel)
		fmt.Println("Logging set to Debug")
	default:
		log.SetLevel(log.WarnLevel)
		fmt.Println("Logging set to Warn")
	}

	// Output Format
	outputType, ok := os.LookupEnv("OUTPUT_TYPE")
	data := data{dest: outputType}

	switch outputType {
	case "statsd":

		statsDServerURI, ok := os.LookupEnv("STATSD_SERVER_URI")
		if !ok {
			statsDServerURI = "http://localhost"
		}

		statsDServerPort, ok := os.LookupEnv("STATSD_SERVER_PORT")
		if !ok {
			statsDServerPort = "8125"
		}

		statsDServerURL := fmt.Sprintf("%s:%s", statsDServerURI, statsDServerPort)

		// ToDo add StatsD support
		fmt.Printf("Stats will output to StatsD Server at %s\n", statsDServerURL)

	default:
		fmt.Println("Stats will output to STDOUT")

	}

	// Output Interval
	outputInterval, ok := os.LookupEnv("OUTPUT_INTERVAL")
	var interval int

	if ok {
		var err error
		interval, err = strconv.Atoi(outputInterval)
		if err != nil {
			panic(err)
		}
	} else {
		interval = 2
	}

	// Output
	for {
		log.Info("Sending stats")

		volumes := stats.GetVolumeSize(cli)
		if len(volumes) > 0 {
			for name, size := range volumes {
				data.name = name
				data.size = size
				output(data)
			}
		}

		time.Sleep(time.Duration(interval) * time.Second)
	}

}
