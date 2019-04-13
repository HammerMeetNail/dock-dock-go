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

func main() {

	// Docker
	maxDockerVersion := "1.39"
	cli := clients.Cli(maxDockerVersion)

	dockerVersion, ok := os.LookupEnv("DOCKER_VERSION")
	if ok {
		cli = clients.Cli(dockerVersion)
	}

	// // Logging
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
	outputFormat, ok := os.LookupEnv("OUTPUT_FORMAT")

	switch outputFormat {
	case "statsd":

		// ToDo add StatsD support
		statsDServerURL, _ := os.LookupEnv("STATSD_SERVER_URL")
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
				fmt.Println(name, size)
			}
		} else {
			fmt.Println("No volumes found")
		}

		time.Sleep(time.Duration(interval) * time.Second)
	}

}
