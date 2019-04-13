package main

import (
	"fmt"
	"os"
	// "strconv"
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
	case "warn":
		log.SetLevel(log.WarnLevel)
		fmt.Println("Logging Set to Warn")
	case "debug":
		log.SetLevel(log.DebugLevel)
		fmt.Println("Logging Set to Debug")
	default:
		log.SetLevel(log.WarnLevel)
		fmt.Println("Default")
	}

	// Output
	outputFormat, ok := os.LookupEnv("OUTPUT_FORMAT")

	switch outputFormat {
	case "statsd":

		// ToDo add StatsD support
		statsDServerURL, _ := os.LookupEnv("STATSD_SERVER_URL")
		fmt.Printf("Stats will output to StatsD Server at %s\n", statsDServerURL)

	default:
		fmt.Println("Stats will output to STDOUT")

	}

	// outputInterval, ok := os.LookupEnv("OUTPUT_INTERVAL")
	interval := 2

	// if ok {
	// 	interval, err := strconv.Atoi(outputInterval)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

	for {
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
