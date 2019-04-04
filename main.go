package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"

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
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
		fmt.Println("Info")
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
		fmt.Println("Warn")
	default:
		logrus.SetLevel(logrus.WarnLevel)
		fmt.Println("Default")
	}

	// // Output
	// OUTPUT_FORMAT, output_format_ok := os.LookupEnv("OUTPUT_FORMAT")
	// TARGET_URL, target_url_ok := os.LookupEnv("TARGET")

	volumes := stats.GetVolumeSize(cli)
	for name, size := range volumes {
		fmt.Println(name, size)
	}
}
